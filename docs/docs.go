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
        "/songs/delete/{id}": {
            "delete": {
                "description": "Удаляет песню из библиотеки по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Song deleted",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/filter": {
            "get": {
                "description": "Получает песни по заданным фильтрам с пагинацией",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название группы",
                        "name": "group_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата выхода песни",
                        "name": "release_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Количество элементов на странице",
                        "name": "limit_num",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы",
                        "name": "page_num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Filtered songs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/songtext/{id}": {
            "get": {
                "description": "Получает текст песни с разбивкой на куплеты по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Couplets of the song",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/update/{id}": {
            "patch": {
                "description": "Обновляет данные песни по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Обновленные данные песни",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Song not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "group_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}