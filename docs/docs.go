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
        "/customer": {
            "get": {
                "description": "Overview all customer's list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Overview customer list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/customer.Customer"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Create customer",
                "parameters": [
                    {
                        "description": "CustomerPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CustomerPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "customer.Customer",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\\"Invalid CPF\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/customer/{cpf}": {
            "get": {
                "description": "Overview a customer by cpf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Overview customer by cpf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer cpf",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "404": {
                        "description": "Customer not found"
                    }
                }
            }
        },
        "/customer/{id}": {
            "put": {
                "description": "Update a customer by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Update customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the customer",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CustomerPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CustomerPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "customer.Customer",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\\"Invalid CPF\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a customer by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Delete customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123456789",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid identifier informed"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CustomerPayload": {
            "type": "object",
            "required": [
                "cpf",
                "email",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 120,
                    "minimum": 18
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                }
            }
        },
        "customer.Customer": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Tech Challenge Customer Food API",
	Description:      "Fast Food API for FIAP Tech course",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
