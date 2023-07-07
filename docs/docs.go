// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "登录请求参数",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/refresh_token": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "RefreshToken",
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40005": {
                        "description": "认证失败",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserApi"
                ],
                "summary": "User Register",
                "parameters": [
                    {
                        "description": "注册请求参数",
                        "name": "registerRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40002": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    },
                    "40003": {
                        "description": "系统错误",
                        "schema": {
                            "$ref": "#/definitions/r.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "r.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "properties": {
                "userAccount": {
                    "type": "string"
                },
                "userPassword": {
                    "type": "string"
                }
            }
        },
        "requests.RegisterRequest": {
            "type": "object",
            "properties": {
                "checkPassword": {
                    "type": "string"
                },
                "userAccount": {
                    "type": "string"
                },
                "userPassword": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "BI Pro API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
