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
        "/admin/customers/all": {
            "get": {
                "description": "Truy xuất danh sách tất cả khách hàng được phân trang",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Liệt kê tất cả khách hàng",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number, defaults to 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size, defaults to 10",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/admin/customers/pending": {
            "get": {
                "description": "Truy xuất danh sách khách hàng được phân trang với trạng thái \"PENDING\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Liệt kê khách hàng đang \"PENDING\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number, defaults to 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size, defaults to 10",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/admin/customers/{id}": {
            "patch": {
                "description": "Cập nhật trạng thái của một khách hàng dựa trên ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
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
                        "name": "approval-status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/admin/delivery-persons/all": {
            "get": {
                "description": "Truy xuất danh sách được phân trang của tất cả những người giao hàng",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Liệt kê tất cả những người giao hàng",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number, defaults to 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size, defaults to 10",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/admin/delivery-persons/pending": {
            "get": {
                "description": "Truy xuất danh sách được phân trang của những người giao hàng với trạng thái \"PENDING\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Liệt kê những người giao hàng đang \"PENDING\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number, defaults to 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page size, defaults to 10",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/admin/delivery-persons/{id}": {
            "patch": {
                "description": "Cập nhật trạng thái của một người giao hàng dựa trên ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Cập nhật trạng thái của người giao hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DeliveryPerson ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Trạng thái mới của người giao hàng (accept, deny)",
                        "name": "approval-status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/customers/{id}": {
            "get": {
                "description": "Truy xuất thông tin chi tiết về một khách hàng cụ thể theo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Nhận thông tin khách hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Cập nhật chi tiết của một khách hàng cụ thể theo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Cập nhật thông tin khách hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer Update Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdateCustomerReq"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Thực hiện xóa khách hàng theo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Xóa tài khoản khách hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/delivery-persons/{id}": {
            "get": {
                "description": "Truy xuất thông tin chi tiết về người giao hàng cụ thể theo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DeliveryPersons"
                ],
                "summary": "Lấy thông tin người giao hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delivery Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Cập nhật chi tiết người giao hàng cụ thể theo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DeliveryPersons"
                ],
                "summary": "Cập nhật thông tin người giao hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delivery Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Delivery Person Update Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdateDeliveryPersonReq"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Thực hiện xóa người giao hàng theo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DeliveryPersons"
                ],
                "summary": "Xóa tài khoản người giao hàng",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delivery Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "Login for different roles (admin, customer, delivery_person, accounting)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login",
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
                    "Authentication"
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
                "description": "Register for different roles (admin, accounting, customer, delivery_person) account_type customer (prepaid, postpaid)\nExample payload:\n\n{\n\"account_type\": \"prepaid\",\n\"current_address\": \"Shop Address of Customer\",\n\"date_of_birth\": \"23/10/2002\",\n\"full_name\": \"Nguyen Cong Manh\",\n\"gender\": \"Nam\",\n\"identification_number\": \"052202014579\",\n\"latitude\": 37.7749,\n\"longtitude\": 122.4194,\n\"nationality\": \"VN\",\n\"password\": \"strongpassword123\",\n\"phone\": \"0977683511\",\n\"place_of_origin\": \"Hoài Sơn, Thị xã Hoài Nhơn, Bình Định\",\n\"place_of_residence\": \"Thôn Phú Nông, Hoài Sơn, Hoài Nhơn, Bình Định\",\n\"role\": \"customer\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
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
                "description": "Đổi mật khẩu người dùng bằng số điện thoại",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Đổi mật khẩu",
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
        }
    },
    "definitions": {
        "req.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "phone",
                "role"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
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
                "current_address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "identification_number": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longtitude": {
                    "type": "number"
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
                }
            }
        },
        "req.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "new_password",
                "phone",
                "role"
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
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "req.UpdateCustomerReq": {
            "type": "object",
            "properties": {
                "account_type": {
                    "type": "string"
                },
                "current_address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "identification_number": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longtitude": {
                    "type": "number"
                },
                "nationality": {
                    "type": "string"
                },
                "place_of_origin": {
                    "type": "string"
                },
                "place_of_residence": {
                    "type": "string"
                }
            }
        },
        "req.UpdateDeliveryPersonReq": {
            "type": "object",
            "properties": {
                "current_address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "identification_number": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "place_of_origin": {
                    "type": "string"
                },
                "place_of_residence": {
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
