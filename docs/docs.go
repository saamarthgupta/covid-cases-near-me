// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/covidCases": {
            "get": {
                "description": "Get the active covid cases in state and Country.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Covid Cases"
                ],
                "summary": "Get the active covid cases in state and Country.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Active Covid cases in State and Country",
                        "schema": {
                            "$ref": "#/definitions/model.CovidCasesResponse"
                        }
                    },
                    "404": {
                        "description": "Unable to Fetch Data for Given Latitude and Longitudes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "fetch Covid Case Data from API and store in DB.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Covid Cases"
                ],
                "summary": "fetch Covid Case Data from API and store in DB.",
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Resource Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CovidCasesInLocation": {
            "type": "object",
            "properties": {
                "activeCases": {
                    "type": "integer"
                },
                "lastUpdated": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                }
            }
        },
        "model.CovidCasesResponse": {
            "type": "object",
            "properties": {
                "country": {
                    "$ref": "#/definitions/model.CovidCasesInLocation"
                },
                "state": {
                    "$ref": "#/definitions/model.CovidCasesInLocation"
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
	Version:     "1.0",
	Host:        "covidCases.swagger.io",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a covid case tracking server.",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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