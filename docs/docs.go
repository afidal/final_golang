// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://developers.ctd.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "API Support",
            "url": "https://developers.ctd.com.ar/support"
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
        "/odontologos": {
            "post": {
                "description": "Crea un nuevo odontólogo",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "POST odontologo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Odontologo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Odontologo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/odontologos/:id": {
            "get": {
                "description": "Obtiene un odontólogo por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "GET odontologo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Odontologo Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualiza un odontologo por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "PUT odontologo by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Odontologo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Odontologo"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Odontologo Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Elimina un odontólogo por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "DELETE odontologo by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Odontologo Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Actualizar parcialmente un odontólogo por su ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologo"
                ],
                "summary": "PATCH odontologo by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Odontologo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Odontologo"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Odontologo Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Odontologo": {
            "type": "object",
            "required": [
                "apellido",
                "matricula",
                "nombre"
            ],
            "properties": {
                "apellido": {
                    "type": "string"
                },
                "matricula": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                }
            }
        },
        "web.errorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "web.response": {
            "type": "object",
            "properties": {
                "data": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API consultorio odontológico",
	Description:      "Esta API permite realizar operaciones CRUD sobre la base de datos del consultorio, que contiene registros de odontólogos, pacientes y turnos",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
