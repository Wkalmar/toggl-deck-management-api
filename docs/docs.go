// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/create-deck": {
            "post": {
                "description": "Creates deck that can be either shuffled or unshuffled. It can accept the list of exact cards which can be shuffled or unshuffled as well. In case no cards provided it returns a deck with 52 cards.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Creates new deck.",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "indicates whether deck is shuffled",
                        "name": "shuffled",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "description": "array of card codes i.e. 8C,AS,7D",
                        "name": "cards",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ClosedDeckDTO"
                        }
                    }
                }
            }
        },
        "/draw-cards": {
            "put": {
                "description": "Removes given number of cards from a deck and returns them as a response.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Draws cards from a deck.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the deck",
                        "name": "deck_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "number of cards to draw",
                        "name": "count",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.CardDTO"
                            }
                        }
                    }
                }
            }
        },
        "/open-deck": {
            "get": {
                "description": "Returns a deck with all of its cars revealed.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Opens deck.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the deck",
                        "name": "deck_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.OpenDeckDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.CardDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "suit": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api.ClosedDeckDTO": {
            "type": "object",
            "properties": {
                "deck_id": {
                    "type": "string"
                },
                "remaining": {
                    "type": "integer"
                },
                "shuffled": {
                    "type": "boolean"
                }
            }
        },
        "api.OpenDeckDTO": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.CardDTO"
                    }
                },
                "deck_id": {
                    "type": "string"
                },
                "remaining": {
                    "type": "integer"
                },
                "shuffled": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Deck Management API",
	Description:      "This is a sample server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
