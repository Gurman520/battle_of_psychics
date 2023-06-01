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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Service The battle of psychics",
    "title": "battle",
    "version": "1.0.0"
  },
  "paths": {
    "/Conceive": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "server"
        ],
        "operationId": "conceive",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Response%D0%A1onceive"
              }
            }
          },
          "401": {
            "description": "incorrect api key auth"
          },
          "404": {
            "description": "a session with such a token was not found"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/Game": {
      "get": {
        "security": [
          {}
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "text/html; charset=utf-8"
        ],
        "tags": [
          "server"
        ],
        "operationId": "game",
        "responses": {
          "200": {
            "description": "success and returns some html text",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/Result": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "server"
        ],
        "operationId": "result",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ResponseResult"
              }
            }
          },
          "401": {
            "description": "incorrect api key auth"
          },
          "404": {
            "description": "a session with such a token was not found"
          },
          "406": {
            "description": "cannot call a method without first calling /Conceive"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/Session": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "server"
        ],
        "operationId": "getSession",
        "responses": {
          "200": {
            "description": "Success create session",
            "schema": {
              "$ref": "#/definitions/ResponseSession"
            }
          },
          "400": {
            "description": "error create session"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/healthCheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ResponseResult": {
      "type": "object",
      "properties": {
        "hypothesis": {
          "type": "integer"
        },
        "reliability": {
          "type": "integer",
          "x-nullable": true
        }
      }
    },
    "ResponseSession": {
      "type": "object",
      "properties": {
        "jwt": {
          "type": "string"
        }
      }
    },
    "ResponseСonceive": {
      "type": "object",
      "properties": {
        "hypothesis": {
          "type": "integer"
        }
      }
    },
    "Result": {
      "type": "object",
      "properties": {
        "number": {
          "type": "integer"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  },
  "security": [
    {
      "key": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Service The battle of psychics",
    "title": "battle",
    "version": "1.0.0"
  },
  "paths": {
    "/Conceive": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "server"
        ],
        "operationId": "conceive",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Response%D0%A1onceive"
              }
            }
          },
          "401": {
            "description": "incorrect api key auth"
          },
          "404": {
            "description": "a session with such a token was not found"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/Game": {
      "get": {
        "security": [
          {}
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "text/html; charset=utf-8"
        ],
        "tags": [
          "server"
        ],
        "operationId": "game",
        "responses": {
          "200": {
            "description": "success and returns some html text",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/Result": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "server"
        ],
        "operationId": "result",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ResponseResult"
              }
            }
          },
          "401": {
            "description": "incorrect api key auth"
          },
          "404": {
            "description": "a session with such a token was not found"
          },
          "406": {
            "description": "cannot call a method without first calling /Conceive"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/Session": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "server"
        ],
        "operationId": "getSession",
        "responses": {
          "200": {
            "description": "Success create session",
            "schema": {
              "$ref": "#/definitions/ResponseSession"
            }
          },
          "400": {
            "description": "error create session"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/healthCheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ResponseResult": {
      "type": "object",
      "properties": {
        "hypothesis": {
          "type": "integer"
        },
        "reliability": {
          "type": "integer",
          "x-nullable": true
        }
      }
    },
    "ResponseSession": {
      "type": "object",
      "properties": {
        "jwt": {
          "type": "string"
        }
      }
    },
    "ResponseСonceive": {
      "type": "object",
      "properties": {
        "hypothesis": {
          "type": "integer"
        }
      }
    },
    "Result": {
      "type": "object",
      "properties": {
        "number": {
          "type": "integer"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  },
  "security": [
    {
      "key": []
    }
  ]
}`))
}
