// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "MIT",
            "url": "http://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/forgot-password": {
            "post": {
                "description": "Request To Reset Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Forgot Password",
                "parameters": [
                    {
                        "description": "Forgot Password Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ForgotPasswordRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Post Login",
                "parameters": [
                    {
                        "description": "Login Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AuthResponseBody"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "Set-Cookie": {
                                "type": "string",
                                "description": "token=jwt-token; Path=/; Secure"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get User Profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get User Profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserResponseBody"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    }
                }
            }
        },
        "/auth/verify": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Verify Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verify Token",
                "parameters": [
                    {
                        "description": "Token to verify",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VerifyTokenRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    }
                }
            }
        },
        "/demo": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get demo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "demo"
                ],
                "summary": "Get demo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domains.Demo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/documents": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Save Document",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "documents"
                ],
                "summary": "Save Document",
                "parameters": [
                    {
                        "description": "Upload Document Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UploadDocumentRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    }
                }
            }
        },
        "/documents/generate-presigned-url": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Generate Presigned URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "documents"
                ],
                "summary": "Generate Presigned URL",
                "parameters": [
                    {
                        "description": "Presigned URL Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PresignedURLRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.PresignedURLResponseBody"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.FailResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domains.Demo": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.AuthResponseBody": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.ForgotPasswordRequestBody": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@email.com"
                }
            }
        },
        "dto.LoginRequestBody": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@email.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 8,
                    "example": "password"
                }
            }
        },
        "dto.MetaDataBody": {
            "type": "object",
            "properties": {
                "extension": {
                    "type": "string"
                },
                "mime_type": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "dto.PresignedURLRequestBody": {
            "type": "object",
            "required": [
                "action",
                "name"
            ],
            "properties": {
                "action": {
                    "type": "string",
                    "example": "upload"
                },
                "name": {
                    "type": "string",
                    "example": "example.pdf"
                }
            }
        },
        "dto.PresignedURLResponseBody": {
            "type": "object",
            "required": [
                "object_key",
                "url"
            ],
            "properties": {
                "object_key": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.UploadDocumentRequestBody": {
            "type": "object",
            "required": [
                "key",
                "metadata"
            ],
            "properties": {
                "key": {
                    "type": "string",
                    "example": "example.pdf"
                },
                "metadata": {
                    "$ref": "#/definitions/dto.MetaDataBody"
                }
            }
        },
        "dto.UserResponseBody": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "pubkey": {
                    "type": "string"
                }
            }
        },
        "dto.VerifyTokenRequestBody": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string",
                    "example": "token"
                }
            }
        },
        "wrapper.FailResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Example message"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "wrapper.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "metadata": {},
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "description": "Enter the token with the ` + "`" + `Bearer ` + "`" + ` prefix, e.g. ` + "`" + `Bearer jwt_token_string` + "`" + `.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "WePress API",
	Description:      "This is the API Document for WePress",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
