// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/yaml",
    "application/json"
  ],
  "produces": [
    "application/yaml",
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "OpenMock is a Go service that can mock services in integration tests, staging environment, or anywhere.  The goal is to simplify the process of writing mocks in various channels.  Currently it supports four channels: HTTP Kafka AMQP (e.g. RabbitMQ) GRPC The admin API allows you to manipulate the mock behaviour provided by openmock, live.  The base path for the admin API is \"/api/v1\".\n",
    "title": "OpenMock",
    "version": "0.2.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/health": {
      "get": {
        "description": "Check if Flagr is healthy",
        "tags": [
          "health"
        ],
        "operationId": "getHealth",
        "responses": {
          "200": {
            "description": "status of health check",
            "schema": {
              "$ref": "#/definitions/health"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/template_sets/{setKey}": {
      "post": {
        "description": "creates / overrides the template set with setKey to the body contents",
        "tags": [
          "template_set"
        ],
        "operationId": "postTemplateSet",
        "parameters": [
          {
            "description": "mocks to add",
            "name": "mocks",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          {
            "type": "string",
            "description": "set key to create",
            "name": "setKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns the successfully posted templates",
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          "400": {
            "description": "if incoming templates were invalid",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "deletes specified template set",
        "tags": [
          "template_set"
        ],
        "operationId": "deleteTemplateSet",
        "parameters": [
          {
            "type": "string",
            "description": "set key to delete",
            "name": "setKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "when successfully deleted"
          },
          "default": {
            "description": "server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/templates": {
      "get": {
        "description": "Get all templates in the loaded model",
        "tags": [
          "template"
        ],
        "operationId": "getTemplates",
        "responses": {
          "200": {
            "description": "all the templates",
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Modify templates in the base model by POSTing new ones",
        "tags": [
          "template"
        ],
        "operationId": "postTemplates",
        "parameters": [
          {
            "description": "mocks to add",
            "name": "mocks",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the successfully posted templates",
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          "400": {
            "description": "if incoming templates were invalid",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Deletes any templates that have been added to the admin endpoint (without a setKey)",
        "tags": [
          "template"
        ],
        "operationId": "deleteTemplates",
        "responses": {
          "204": {
            "description": "when successfully deleted"
          },
          "default": {
            "description": "server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/templates/{templateKey}": {
      "delete": {
        "description": "delete a specific template that has been added by key",
        "tags": [
          "template"
        ],
        "operationId": "deleteTemplate",
        "parameters": [
          {
            "type": "string",
            "description": "key of template to delete",
            "name": "templateKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "when successfully deleted"
          },
          "400": {
            "description": "invalid key",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ActionDispatcher": {
      "type": "object",
      "properties": {
        "order": {
          "description": "used to explicitly order the actions run when a behavior triggers",
          "type": "integer",
          "format": "int64",
          "default": 0
        },
        "publish_amqp": {
          "$ref": "#/definitions/ActionPublishAMQP"
        },
        "publish_kafka": {
          "$ref": "#/definitions/ActionPublishKafka"
        },
        "redis": {
          "$ref": "#/definitions/ActionRedis"
        },
        "reply_grpc": {
          "$ref": "#/definitions/ActionReplyGRPC"
        },
        "reply_http": {
          "$ref": "#/definitions/ActionReplyHTTP"
        },
        "send_http": {
          "$ref": "#/definitions/ActionSendHTTP"
        },
        "sleep": {
          "$ref": "#/definitions/ActionSleep"
        }
      }
    },
    "ActionPublishAMQP": {
      "description": "publish a message on AMQP if this behaviors condition is met",
      "type": "object",
      "properties": {
        "exchange": {
          "description": "AMQP exchange name",
          "type": "string"
        },
        "payload": {
          "description": "string payload to send on AMQP",
          "type": "string"
        },
        "payload_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load body from",
          "type": "string"
        },
        "routing_key": {
          "description": "AMQP routing key",
          "type": "string"
        }
      }
    },
    "ActionPublishKafka": {
      "description": "publish a message on kafka",
      "type": "object",
      "properties": {
        "payload": {
          "description": "string payload to send on AMQP",
          "type": "string"
        },
        "payload_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load body from",
          "type": "string"
        },
        "topic": {
          "description": "kafka topic to publish on",
          "type": "string"
        }
      }
    },
    "ActionRedis": {
      "description": "a list of redis commands to run when the",
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "ActionReplyGRPC": {
      "description": "reply to incoming GRPC that triggered this behavior with a response",
      "type": "object",
      "properties": {
        "payload": {
          "description": "string payload to send via GRPC, this should be a json string that maps to the protobuf response object",
          "type": "string"
        },
        "payload_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load payload from, this should be a json string that maps to the protobuf response object",
          "type": "string"
        }
      }
    },
    "ActionReplyHTTP": {
      "description": "reply to incoming HTTP that triggered this behavior with a response",
      "type": "object",
      "properties": {
        "body": {
          "description": "Text body to send over HTTP, can use templating",
          "type": "string"
        },
        "body_from_file": {
          "description": "file name (relative to working directory of OM) to load HTTP body from",
          "type": "string"
        },
        "headers": {
          "description": "map of string to string specifying HTTP headers to attach to our message",
          "type": "object",
          "additionalProperties": true
        },
        "status_code": {
          "description": "HTTP status code to reply with",
          "type": "integer",
          "format": "int64",
          "default": 200
        }
      }
    },
    "ActionSendHTTP": {
      "description": "Send a HTTP message as an action",
      "type": "object",
      "properties": {
        "body": {
          "description": "Text body to send over HTTP, can use templating",
          "type": "string"
        },
        "body_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load HTTP body from",
          "type": "string"
        },
        "headers": {
          "description": "map of string to string specifying HTTP headers to attach to our message",
          "type": "object",
          "additionalProperties": true
        },
        "method": {
          "description": "HTTP method to use for the send",
          "type": "string",
          "enum": [
            "POST",
            "GET",
            "DELETE",
            "PUT",
            "OPTIONS",
            "HEAD"
          ]
        },
        "url": {
          "description": "The URL to send HTTP to",
          "type": "string"
        }
      }
    },
    "ActionSleep": {
      "description": "pause the action thread for a time",
      "type": "object",
      "properties": {
        "duration": {
          "description": "time to wait in seconds; e.g. '1s'",
          "type": "string"
        }
      }
    },
    "Expect": {
      "type": "object",
      "properties": {
        "amqp": {
          "$ref": "#/definitions/ExpectAMQP"
        },
        "condition": {
          "description": "a go template that determines if this behavior triggers",
          "type": "string"
        },
        "grpc": {
          "$ref": "#/definitions/ExpectGRPC"
        },
        "http": {
          "$ref": "#/definitions/ExpectHTTP"
        },
        "kafka": {
          "$ref": "#/definitions/ExpectKafka"
        }
      }
    },
    "ExpectAMQP": {
      "type": "object",
      "properties": {
        "exchange": {
          "description": "TODO",
          "type": "string"
        },
        "queue": {
          "description": "TODO",
          "type": "string"
        },
        "routing_key": {
          "description": "TODO",
          "type": "string"
        }
      }
    },
    "ExpectGRPC": {
      "type": "object",
      "properties": {
        "method": {
          "description": "GRPC method to expect to trigger this behavior",
          "type": "string"
        },
        "service": {
          "description": "GRPC service to expect to trigger this behavior",
          "type": "string"
        }
      }
    },
    "ExpectHTTP": {
      "type": "object",
      "properties": {
        "method": {
          "description": "HTTP method to expect to trigger this behavior",
          "type": "string",
          "enum": [
            "POST",
            "GET",
            "DELETE",
            "PUT",
            "OPTIONS",
            "HEAD"
          ]
        },
        "path": {
          "description": "HTTP path to expect to trigger this behavior",
          "type": "string"
        }
      }
    },
    "ExpectKafka": {
      "type": "object",
      "properties": {
        "topic": {
          "description": "kafka topic to listen on",
          "type": "string"
        }
      }
    },
    "Mock": {
      "type": "object",
      "properties": {
        "actions": {
          "description": "for behaviors, the actions this mock would do when the expect is met",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionDispatcher"
          }
        },
        "expect": {
          "$ref": "#/definitions/Expect"
        },
        "extend": {
          "description": "for behaviors, makes this behavior extend a specified AbstractBehavior",
          "type": "string"
        },
        "key": {
          "description": "Unique key for the item in OM's model",
          "type": "string",
          "pattern": "[\\w_\\-\\.]+"
        },
        "kind": {
          "description": "The type of item this is. possible types are: Behavior - creates a new mock behavior  AbstractBehavior - allows behaviors to use common features from this item Template - used in template language rendering to do fancy stuff\n",
          "type": "string",
          "enum": [
            "Behavior",
            "AbstractBehavior",
            "Template"
          ]
        },
        "template": {
          "description": "a go template to be embedded in other templates",
          "type": "string"
        },
        "values": {
          "description": "Arbitrary values that can be used in go templates rendered by this item",
          "type": "object",
          "additionalProperties": true
        }
      }
    },
    "Mocks": {
      "description": "collection of mocks",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Mock"
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "health": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "manipulating 'mocks' in the model",
      "name": "template"
    },
    {
      "description": "manipulating sets of 'mocks' in the model",
      "name": "template_set"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "application/yaml"
  ],
  "produces": [
    "application/json",
    "application/yaml"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "OpenMock is a Go service that can mock services in integration tests, staging environment, or anywhere.  The goal is to simplify the process of writing mocks in various channels.  Currently it supports four channels: HTTP Kafka AMQP (e.g. RabbitMQ) GRPC The admin API allows you to manipulate the mock behaviour provided by openmock, live.  The base path for the admin API is \"/api/v1\".\n",
    "title": "OpenMock",
    "version": "0.2.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/health": {
      "get": {
        "description": "Check if Flagr is healthy",
        "tags": [
          "health"
        ],
        "operationId": "getHealth",
        "responses": {
          "200": {
            "description": "status of health check",
            "schema": {
              "$ref": "#/definitions/health"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/template_sets/{setKey}": {
      "post": {
        "description": "creates / overrides the template set with setKey to the body contents",
        "tags": [
          "template_set"
        ],
        "operationId": "postTemplateSet",
        "parameters": [
          {
            "description": "mocks to add",
            "name": "mocks",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          {
            "type": "string",
            "description": "set key to create",
            "name": "setKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns the successfully posted templates",
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          "400": {
            "description": "if incoming templates were invalid",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "deletes specified template set",
        "tags": [
          "template_set"
        ],
        "operationId": "deleteTemplateSet",
        "parameters": [
          {
            "type": "string",
            "description": "set key to delete",
            "name": "setKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "when successfully deleted"
          },
          "default": {
            "description": "server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/templates": {
      "get": {
        "description": "Get all templates in the loaded model",
        "tags": [
          "template"
        ],
        "operationId": "getTemplates",
        "responses": {
          "200": {
            "description": "all the templates",
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "Modify templates in the base model by POSTing new ones",
        "tags": [
          "template"
        ],
        "operationId": "postTemplates",
        "parameters": [
          {
            "description": "mocks to add",
            "name": "mocks",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the successfully posted templates",
            "schema": {
              "$ref": "#/definitions/Mocks"
            }
          },
          "400": {
            "description": "if incoming templates were invalid",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Deletes any templates that have been added to the admin endpoint (without a setKey)",
        "tags": [
          "template"
        ],
        "operationId": "deleteTemplates",
        "responses": {
          "204": {
            "description": "when successfully deleted"
          },
          "default": {
            "description": "server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/templates/{templateKey}": {
      "delete": {
        "description": "delete a specific template that has been added by key",
        "tags": [
          "template"
        ],
        "operationId": "deleteTemplate",
        "parameters": [
          {
            "type": "string",
            "description": "key of template to delete",
            "name": "templateKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "when successfully deleted"
          },
          "400": {
            "description": "invalid key",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "server error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ActionDispatcher": {
      "type": "object",
      "properties": {
        "order": {
          "description": "used to explicitly order the actions run when a behavior triggers",
          "type": "integer",
          "format": "int64",
          "default": 0
        },
        "publish_amqp": {
          "$ref": "#/definitions/ActionPublishAMQP"
        },
        "publish_kafka": {
          "$ref": "#/definitions/ActionPublishKafka"
        },
        "redis": {
          "$ref": "#/definitions/ActionRedis"
        },
        "reply_grpc": {
          "$ref": "#/definitions/ActionReplyGRPC"
        },
        "reply_http": {
          "$ref": "#/definitions/ActionReplyHTTP"
        },
        "send_http": {
          "$ref": "#/definitions/ActionSendHTTP"
        },
        "sleep": {
          "$ref": "#/definitions/ActionSleep"
        }
      }
    },
    "ActionPublishAMQP": {
      "description": "publish a message on AMQP if this behaviors condition is met",
      "type": "object",
      "properties": {
        "exchange": {
          "description": "AMQP exchange name",
          "type": "string"
        },
        "payload": {
          "description": "string payload to send on AMQP",
          "type": "string"
        },
        "payload_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load body from",
          "type": "string"
        },
        "routing_key": {
          "description": "AMQP routing key",
          "type": "string"
        }
      }
    },
    "ActionPublishKafka": {
      "description": "publish a message on kafka",
      "type": "object",
      "properties": {
        "payload": {
          "description": "string payload to send on AMQP",
          "type": "string"
        },
        "payload_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load body from",
          "type": "string"
        },
        "topic": {
          "description": "kafka topic to publish on",
          "type": "string"
        }
      }
    },
    "ActionRedis": {
      "description": "a list of redis commands to run when the",
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "ActionReplyGRPC": {
      "description": "reply to incoming GRPC that triggered this behavior with a response",
      "type": "object",
      "properties": {
        "payload": {
          "description": "string payload to send via GRPC, this should be a json string that maps to the protobuf response object",
          "type": "string"
        },
        "payload_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load payload from, this should be a json string that maps to the protobuf response object",
          "type": "string"
        }
      }
    },
    "ActionReplyHTTP": {
      "description": "reply to incoming HTTP that triggered this behavior with a response",
      "type": "object",
      "properties": {
        "body": {
          "description": "Text body to send over HTTP, can use templating",
          "type": "string"
        },
        "body_from_file": {
          "description": "file name (relative to working directory of OM) to load HTTP body from",
          "type": "string"
        },
        "headers": {
          "description": "map of string to string specifying HTTP headers to attach to our message",
          "type": "object",
          "additionalProperties": true
        },
        "status_code": {
          "description": "HTTP status code to reply with",
          "type": "integer",
          "format": "int64",
          "default": 200
        }
      }
    },
    "ActionSendHTTP": {
      "description": "Send a HTTP message as an action",
      "type": "object",
      "properties": {
        "body": {
          "description": "Text body to send over HTTP, can use templating",
          "type": "string"
        },
        "body_from_file": {
          "description": "file path (relative to OPENMOCK_TEMPLATES_DIR of OM) to load HTTP body from",
          "type": "string"
        },
        "headers": {
          "description": "map of string to string specifying HTTP headers to attach to our message",
          "type": "object",
          "additionalProperties": true
        },
        "method": {
          "description": "HTTP method to use for the send",
          "type": "string",
          "enum": [
            "POST",
            "GET",
            "DELETE",
            "PUT",
            "OPTIONS",
            "HEAD"
          ]
        },
        "url": {
          "description": "The URL to send HTTP to",
          "type": "string"
        }
      }
    },
    "ActionSleep": {
      "description": "pause the action thread for a time",
      "type": "object",
      "properties": {
        "duration": {
          "description": "time to wait in seconds; e.g. '1s'",
          "type": "string"
        }
      }
    },
    "Expect": {
      "type": "object",
      "properties": {
        "amqp": {
          "$ref": "#/definitions/ExpectAMQP"
        },
        "condition": {
          "description": "a go template that determines if this behavior triggers",
          "type": "string"
        },
        "grpc": {
          "$ref": "#/definitions/ExpectGRPC"
        },
        "http": {
          "$ref": "#/definitions/ExpectHTTP"
        },
        "kafka": {
          "$ref": "#/definitions/ExpectKafka"
        }
      }
    },
    "ExpectAMQP": {
      "type": "object",
      "properties": {
        "exchange": {
          "description": "TODO",
          "type": "string"
        },
        "queue": {
          "description": "TODO",
          "type": "string"
        },
        "routing_key": {
          "description": "TODO",
          "type": "string"
        }
      }
    },
    "ExpectGRPC": {
      "type": "object",
      "properties": {
        "method": {
          "description": "GRPC method to expect to trigger this behavior",
          "type": "string"
        },
        "service": {
          "description": "GRPC service to expect to trigger this behavior",
          "type": "string"
        }
      }
    },
    "ExpectHTTP": {
      "type": "object",
      "properties": {
        "method": {
          "description": "HTTP method to expect to trigger this behavior",
          "type": "string",
          "enum": [
            "POST",
            "GET",
            "DELETE",
            "PUT",
            "OPTIONS",
            "HEAD"
          ]
        },
        "path": {
          "description": "HTTP path to expect to trigger this behavior",
          "type": "string"
        }
      }
    },
    "ExpectKafka": {
      "type": "object",
      "properties": {
        "topic": {
          "description": "kafka topic to listen on",
          "type": "string"
        }
      }
    },
    "Mock": {
      "type": "object",
      "properties": {
        "actions": {
          "description": "for behaviors, the actions this mock would do when the expect is met",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ActionDispatcher"
          }
        },
        "expect": {
          "$ref": "#/definitions/Expect"
        },
        "extend": {
          "description": "for behaviors, makes this behavior extend a specified AbstractBehavior",
          "type": "string"
        },
        "key": {
          "description": "Unique key for the item in OM's model",
          "type": "string",
          "pattern": "[\\w_\\-\\.]+"
        },
        "kind": {
          "description": "The type of item this is. possible types are: Behavior - creates a new mock behavior  AbstractBehavior - allows behaviors to use common features from this item Template - used in template language rendering to do fancy stuff\n",
          "type": "string",
          "enum": [
            "Behavior",
            "AbstractBehavior",
            "Template"
          ]
        },
        "template": {
          "description": "a go template to be embedded in other templates",
          "type": "string"
        },
        "values": {
          "description": "Arbitrary values that can be used in go templates rendered by this item",
          "type": "object",
          "additionalProperties": true
        }
      }
    },
    "Mocks": {
      "description": "collection of mocks",
      "type": "array",
      "items": {
        "$ref": "#/definitions/Mock"
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "health": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "manipulating 'mocks' in the model",
      "name": "template"
    },
    {
      "description": "manipulating sets of 'mocks' in the model",
      "name": "template_set"
    }
  ]
}`))
}
