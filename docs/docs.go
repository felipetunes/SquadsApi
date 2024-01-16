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
        "/api/v1/player/getall": {
            "get": {
                "description": "Get all players",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get all players",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Player"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/player/getbycountry/{country}": {
            "get": {
                "description": "Get players by country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get players by country",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Country Name",
                        "name": "country",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Player"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/player/getbyid/{id}": {
            "get": {
                "description": "Get a player by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get a player by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Player"
                        }
                    }
                }
            }
        },
        "/api/v1/player/getbyidteam/{idteam}": {
            "get": {
                "description": "Get players by team ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get players by team ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "idteam",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Player"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/player/getbyname/{name}": {
            "get": {
                "description": "Get a player by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get a player by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Player Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Player"
                        }
                    }
                }
            }
        },
        "/api/v1/player/insert": {
            "post": {
                "description": "Insert a player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Insert a player",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Player Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "team",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "City",
                        "name": "city",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Country",
                        "name": "country",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Age",
                        "name": "age",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Player"
                        }
                    }
                }
            }
        },
        "/api/v1/team/getall": {
            "get": {
                "description": "Get all teams",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get all teams",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Team"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/team/getbycountry/{country}": {
            "get": {
                "description": "Get teams by country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get teams by country",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Country Name",
                        "name": "country",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Team"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/team/getbyid/{id}": {
            "get": {
                "description": "Get a team by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get a team by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Team"
                        }
                    }
                }
            }
        },
        "/api/v1/team/getbyname/{name}": {
            "get": {
                "description": "Get a team by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get a team by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Team"
                        }
                    }
                }
            }
        },
        "/api/v1/team/insert": {
            "post": {
                "description": "Insert a team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Insert a team",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Team"
                        }
                    }
                }
            }
        },
        "/api/v1/team/update": {
            "put": {
                "description": "Update a team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Update a team",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Team"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.Player": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "idteam": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "structs.Team": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
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
