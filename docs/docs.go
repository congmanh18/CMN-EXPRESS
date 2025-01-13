// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://sharkytech.vercel.app/terms",
        "contact": {
            "name": "Lucas",
            "url": "https://sharkytech.vercel.app",
            "email": "nguyenmanh180102@gmail.com"
        },
        "license": {
            "name": "MIT License",
            "url": "https://opensource.org/licenses/MIT"
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
                    "Sprint1"
                ],
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.LoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/refresh-token": {
            "post": {
                "description": "Refresh Access Token using a valid Refresh Token from headers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "summary": "Refresh Access Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The refresh token identifier stored in headers",
                        "name": "refresh_token_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/register": {
            "post": {
                "description": "Register for different roles (admin, accounting, customer, delivery_person) account_type customer (prepaid, postpaid)\nExample customer payload:\n{\n\"phone\": \"0990123456\",\n\"password\": \"wrightdaniel\",\n\"last_name\": \"Wright\",\n\"middle_name\": \"Thomas\",\n\"first_name\": \"Daniel\",\n\"specific_address\": \"789 Tran Hung Dao\",\n\"ward\": \"Duong Dong\",\n\"district\": \"Duong Dong\",\n\"city\": \"Phu Quoc\",\n\"identification_number\": \"321987654\",\n\"gender\": \"male\",\n\"date_of_birth\": \"1994-09-09\",\n\"nationality\": \"Vietnamese\",\n\"place_of_origin\": \"Phu Quoc\",\n\"place_of_residence\": \"Phu Quoc\",\n\"latitude\": 10.2899,\n\"longtitude\": 103.9840,\n\"account_type\": \"postpaid\",\n\"role\": \"customer\"\n}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register Request Example",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.RegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/reset-password": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "parameters": [
                    {
                        "description": "Reset Password Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/search": {
            "get": {
                "description": "Search a list of users (customers and delivery persons) with optional filters by status and role, including pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "summary": "Search paginated users",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size (default is 10)",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by customer status (e.g., pending, verified, blocked, active, inactive) Filter by delivery_person (e.g., on_duty, off_duty)",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search by user name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search by user phone",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by user role (e.g., customer, delivery_person)",
                        "name": "role",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users": {
            "get": {
                "description": "Get a list of users (customers and delivery persons) with optional filters by status and role, including pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "summary": "Fetch paginated users",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size (default is 10)",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by customer status (e.g., pending, verified, blocked, active, inactive) Filter by delivery_person (e.g., on_duty, off_duty)",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by user role (e.g., customer, delivery_person)",
                        "name": "role",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Fetch the details of a user based on their ID. The user can be either a customer or a delivery person.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "summary": "Get user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "Cập nhật trạng thái của một khách hàng dựa trên ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sprint1"
                ],
                "summary": "Cập nhật trạng thái khách hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Trạng thái mới của khách hàng (accept, deny)",
                        "name": "approval_status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "req.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "req.RegisterRequest": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "account_type": {
                    "description": "2. Field for customer",
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "identification_number": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longtitude": {
                    "type": "number"
                },
                "middle_name": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "description": "1. Default field",
                    "type": "string"
                },
                "place_of_origin": {
                    "type": "string"
                },
                "place_of_residence": {
                    "type": "string"
                },
                "role": {
                    "description": "3. Field for delivery person",
                    "type": "string"
                },
                "specific_address": {
                    "type": "string"
                },
                "ward": {
                    "type": "string"
                }
            }
        },
        "req.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "new_password",
                "phone"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "new_password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "203.145.47.225",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CMN Express API Documentation",
	Description:      "Welcome to the CMN Express API documentation. This server facilitates seamless integration for managing logistics, tracking, and delivery services efficiently.\nExplore the endpoints to unlock robust features designed to enhance your operational workflow.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
