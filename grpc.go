package openmock

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

const (
	grpcPayloadLen = 1
	grpcSizeLen    = 4
	grpcHeaderLen  = grpcPayloadLen + grpcSizeLen
)

var jsonMarshaler = jsonpb.Marshaler{
	OrigName:     true,
	EmitDefaults: true,
	EnumsAsInts:  false,
}

var jsonUnmarshaler = jsonpb.Unmarshaler{}

// length-prefixed message, see https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md
func msgHeader(data []byte) (hdr []byte, payload []byte) {
	hdr = make([]byte, grpcHeaderLen)

	hdr[0] = byte(0)

	// Write length of payload into buf
	binary.BigEndian.PutUint32(hdr[grpcPayloadLen:], uint32(len(data)))
	return hdr, data
}

// GRPCService is a map of service_name => GRPCRequestResponsePair
type GRPCService map[string]GRPCRequestResponsePair

// GRPCRequestResponsePair is a pair of proto.Message to define
// the message schema of request and response of a method
type GRPCRequestResponsePair struct {
	Request  proto.Message
	Response proto.Message
}

func (om *OpenMock) parseServiceMethod(ec echo.Context) (service string, method string) {
	tokens := strings.Split(ec.Path(), "/")
	return tokens[1], tokens[2]
}

func (om *OpenMock) findRegisteredRpcPair(ec echo.Context) (GRPCRequestResponsePair, error) {
	var result GRPCRequestResponsePair
	if om.GRPCServiceMap == nil {
		return result, fmt.Errorf("empty GRPCServiceMap")
	}

	service, method := om.parseServiceMethod(ec)

	if _, ok := om.GRPCServiceMap[service]; !ok {
		return result, fmt.Errorf("invalid service in GRPCServiceMap. %s", service)
	}

	if _, ok := om.GRPCServiceMap[service][method]; !ok {
		return result, fmt.Errorf("invalid method in GRPCServiceMap[%s]. %+v", service, method)
	}

	return om.GRPCServiceMap[service][method], nil
}

func (om *OpenMock) convertJSONToH2Response(ec echo.Context, resJSON string) (header []byte, data []byte, err error) {
	pair, err := om.findRegisteredRpcPair(ec)
	if err != nil {
		return nil, nil, err
	}
	res := pair.Response
	buf := bytes.NewBufferString(resJSON)
	err = jsonUnmarshaler.Unmarshal(buf, res)
	if err != nil {
		return nil, nil, err
	}
	b, err := proto.Marshal(res)
	if err != nil {
		return nil, nil, err
	}

	header, data = msgHeader(b)
	return header, data, nil
}

func (om *OpenMock) convertMsgToJSON(ec echo.Context, msg proto.Message, body []byte) (string, error) {
	if len(body) < grpcHeaderLen {
		return "", fmt.Errorf("invalid grpc body length. length: %d", len(body))
	}
	// first grpcHeaderLen bytes are compression and size information
	err := proto.Unmarshal(body[grpcHeaderLen:], msg)
	if err != nil {
		return "", err
	}

	return jsonMarshaler.MarshalToString(msg)
}

// convertRequestBodyToJSON is how we support JSONPath to take values from GRPC requests and include them in responses
func (om *OpenMock) convertRequestBodyToJSON(ec echo.Context, body []byte) (string, error) {
	pair, err := om.findRegisteredRpcPair(ec)
	if err != nil {
		return "", err
	}
	return om.convertMsgToJSON(ec, pair.Request, body)
}

func (om *OpenMock) convertResponseBodyToJSON(ec echo.Context, body []byte) (string, error) {
	pair, err := om.findRegisteredRpcPair(ec)
	if err != nil {
		return "", err
	}
	return om.convertMsgToJSON(ec, pair.Response, body)
}

func (om *OpenMock) prepareGRPCEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(ec echo.Context, reqBody, resBody []byte) {
		reqBodyStr, reqErr := om.convertRequestBodyToJSON(ec, reqBody)
		var resBodyStr string
		var resErr error
		if ec.Response().Status == 200 {
			resBodyStr, resErr = om.convertResponseBodyToJSON(ec, resBody)
		} else {
			resBodyStr = string(resBody)
		}
		logrus.WithFields(logrus.Fields{
			"grpc_path":    ec.Path(),
			"grpc_method":  ec.Request().Method,
			"grpc_host":    ec.Request().Host,
			"grpc_status":  ec.Response().Status,
			"grpc_req":     reqBodyStr,
			"grpc_req_err": reqErr,
			"grpc_res":     resBodyStr,
			"grpc_res_err": resErr,
		}).Info()
	}))
	if om.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowCredentials: true,
			AllowHeaders:     []string{"*"},
			AllowMethods:     []string{"*"},
		}))
	}
	mocks := om.repo.GRPCMocks
	for h, ms := range mocks {
		func(h ExpectGRPC, ms MocksArray) {
			e.Match(
				[]string{"POST"},
				fmt.Sprintf("/%s/%s", h.Service, h.Method),
				func(ec echo.Context) error {
					body, _ := ioutil.ReadAll(ec.Request().Body)
					JSONRequestBody, err := om.convertRequestBodyToJSON(ec, body)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"grpc_path":   ec.Path(),
							"grpc_method": ec.Request().Method,
							"grpc_host":   ec.Request().Host,
						}).WithError(err).Error("failed to convert gRPC request body to JSON")
						return err
					}

					c := Context{
						GRPCContext: ec,
						GRPCHeader:  ec.Request().Header,
						GRPCPayload: JSONRequestBody,
						GRPCMethod:  h.Method,
						GRPCService: h.Service,
						om:          om,
					}

					if !ms.DoActions(c) {
						return errors.New("no mock matched")
					}
					return nil
				},
			)
		}(h, ms)
	}
	return e
}

func (om *OpenMock) startGRPC() {
	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	e := om.prepareGRPCEcho()
	e.Logger.Fatal(e.StartH2CServer(fmt.Sprintf("%s:%d", om.GRPCHost, om.GRPCPort), s))
}
