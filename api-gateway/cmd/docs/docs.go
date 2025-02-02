// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "description": "Returns a JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.LoginOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "405": {
                        "description": "Method Not Allowed"
                    }
                }
            }
        },
        "/api/v1/auth/reg": {
            "post": {
                "description": "Returns a signed up user response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reg"
                ],
                "summary": "Register API",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.RegisterOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "405": {
                        "description": "Method Not Allowed"
                    }
                }
            }
        },
        "/api/v1/auth/valid": {
            "post": {
                "description": "Check if JWT is still valid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "valid"
                ],
                "summary": "JWT validation API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ValidOkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "405": {
                        "description": "Method Not Allowed"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.LoginOkResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "main.RegisterOkResponse": {
            "type": "object"
        },
        "main.ValidOkResponse": {
            "type": "object",
            "properties": {
                "valid": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1/auth",
	Schemes:          []string{},
	Title:            "SSO-service API",
	Description:      "SSO-service for DronWallet",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
