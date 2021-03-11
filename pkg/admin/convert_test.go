package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ziru/openmock"
	model "github.com/ziru/openmock/swagger_gen/models"
)

func TestSwaggerToOpenmockMock(t *testing.T) {
	// swagger := model.Mock{}
	key := "ping"
	method := "GET"
	path := "/ping"
	statusCode := int64(200)
	swagger := model.Mock{
		Kind: "Behavior",
		Key:  key,
		Expect: &model.Expect{
			HTTP: &model.ExpectHTTP{
				Method: method,
				Path:   path,
			},
		},
		Actions: []*model.ActionDispatcher{{
			ReplyHTTP: &model.ActionReplyHTTP{
				Body:       "OK",
				StatusCode: &statusCode,
			},
		}},
	}

	omMock := openmock.Mock{
		Kind: "Behavior",
		Key:  key,
		Expect: openmock.Expect{
			HTTP: openmock.ExpectHTTP{
				Method: method,
				Path:   path,
			},
		},
		Actions: []openmock.ActionDispatcher{{
			ActionReplyHTTP: openmock.ActionReplyHTTP{
				Body:       "OK",
				StatusCode: int(statusCode),
			},
		}},
	}

	t.Run("test converting a swagger mock", func(t *testing.T) {
		om := swaggerToOpenmockMock(swagger)
		assert.Equal(t, omMock, om)
	})

	t.Run("test converting swagger mocks array", func(t *testing.T) {
		swaggerMocks := model.Mocks{
			&swagger,
		}
		omMocks := swaggerToOpenmockMockArray(swaggerMocks)
		assert.Equal(t, []*openmock.Mock{&omMock}, omMocks)
	})
}
