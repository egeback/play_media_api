// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://xxxx.xxx.xx",
            "email": "support@egeback.se"
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
        "/ping": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "summary": "ping common",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shows": {
            "get": {
                "description": "get shows",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shows"
                ],
                "summary": "List shows",
                "parameters": [
                    {
                        "type": "string",
                        "format": "bool",
                        "description": "pretty print show",
                        "name": "prettyPrint",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "bool",
                        "description": "show seasons",
                        "name": "showSeasons",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Show"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/shows/{slug}": {
            "get": {
                "description": "get show by slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shows"
                ],
                "summary": "Show an show",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Show ID",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "bool",
                        "description": "pretty print show",
                        "name": "prettyPrint",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Show"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                },
                "more_info": {
                    "type": "string",
                    "example": "http://"
                }
            }
        },
        "models.Show": {
            "type": "object",
            "properties": {
                "api_url": {
                    "type": "string",
                    "example": "http://adad.ad/se"
                },
                "decription": {
                    "type": "string",
                    "example": "Show about x"
                },
                "genre": {
                    "type": "string",
                    "example": "2019-12-22"
                },
                "imageUrl": {
                    "type": "string",
                    "example": "http://adad.ad/se"
                },
                "name": {
                    "type": "string",
                    "example": "Show Name"
                },
                "page_url": {
                    "type": "string",
                    "example": "http://adad.ad/se"
                },
                "seasons": {
                    "type": "array",
                    "items": {
                        "type": "Season"
                    }
                },
                "service": {
                    "type": "string"
                },
                "slug": {
                    "type": "string",
                    "example": "show_name"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2019-12-22"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.5",
	Host:        "",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "Play API",
	Description: "API including SVT and TV4 Play",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
