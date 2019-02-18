// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-02-19 01:24:04.148658 +0300 MSK m=+0.087481222

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample Payment service.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/screwyprof/s"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "paths": {
        "/accounts": {
            "get": {
                "description": "Retrieves available accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Retrieves available accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.AvailableAccount"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Open a new account with optional balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Open a new account",
                "parameters": [
                    {
                        "description": "Open account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.OpenAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/response.ShortAccountInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    }
                }
            }
        },
        "/accounts/{number}": {
            "get": {
                "description": "Show account info by number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Show an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "account number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/response.AccountInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    }
                }
            }
        },
        "/accounts/{number}/transfer": {
            "post": {
                "description": "Transfer money from an account to another account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Transfer money from an account to another account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "account number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transfer money",
                        "name": "transfer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.TransferMoney"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/responder.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "request.OpenAccount": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 77700
                },
                "currency": {
                    "type": "string",
                    "example": "USD"
                },
                "number": {
                    "type": "string",
                    "example": "ACC777"
                }
            }
        },
        "request.TransferMoney": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 10000
                },
                "currency": {
                    "type": "string",
                    "example": "USD"
                },
                "to": {
                    "type": "string",
                    "example": "ACC555"
                }
            }
        },
        "responder.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "response.AccountInfo": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string",
                    "example": "$100.00"
                },
                "ledgers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Ledger"
                    }
                },
                "number": {
                    "type": "string",
                    "example": "ACC777"
                }
            }
        },
        "response.AvailableAccount": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string",
                    "example": "ACC777"
                }
            }
        },
        "response.Ledger": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string",
                    "example": "Transfer from AK777, $100"
                }
            }
        },
        "response.ShortAccountInfo": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string",
                    "example": "$100.00"
                },
                "number": {
                    "type": "string",
                    "example": "ACC777"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}