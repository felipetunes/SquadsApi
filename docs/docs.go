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
        "/api/v1/live/getall": {
            "get": {
                "description": "Get all matches",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Live"
                ],
                "summary": "Get all matches",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Live"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/live/getallbyidteam/{id}": {
            "get": {
                "description": "Get all matches by team id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Live"
                ],
                "summary": "Get all matches by team id",
                "parameters": [
                    {
                        "type": "integer",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Live"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/live/getalltoday": {
            "get": {
                "description": "Get all matches today",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Live"
                ],
                "summary": "Get all matches today",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Live"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/player/delete/{id}": {
            "delete": {
                "description": "Delete a player by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Delete a player by ID",
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
                            "type": "string"
                        }
                    }
                }
            }
        },
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
                        "description": "IdTeam",
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
                "description": "Get players by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Get players by name",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Player"
                            }
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
                        "description": "Id Team",
                        "name": "idteam",
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
                        "description": "Birth",
                        "name": "birth",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Height",
                        "name": "height",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Position",
                        "name": "position",
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
        "/api/v1/player/update": {
            "put": {
                "description": "Update a player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Players"
                ],
                "summary": "Update a player",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID Player",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Player Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Id Team",
                        "name": "idteam",
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
                        "description": "Birth",
                        "name": "birth",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Height",
                        "name": "height",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Position",
                        "name": "position",
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
        "/api/v1/team/delete/{id}": {
            "delete": {
                "description": "Delete a team by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Delete a team by ID",
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
                            "type": "string"
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
                "description": "Get teams by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Get teams by name",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Team"
                            }
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
                "parameters": [
                    {
                        "type": "string",
                        "description": "Team Name",
                        "name": "name",
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
                        "description": "Color1",
                        "name": "color1",
                        "in": "query",
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
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID Team",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Team Name",
                        "name": "name",
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
                        "description": "Color1",
                        "name": "color1",
                        "in": "query",
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
        }
    },
    "definitions": {
        "structs.Live": {
            "type": "object",
            "properties": {
                "championship": {
                    "type": "string"
                },
                "datematch": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "stadium": {
                    "type": "string"
                },
                "teamid1": {
                    "type": "integer"
                },
                "teamid2": {
                    "type": "integer"
                },
                "teampoints1": {
                    "type": "integer"
                },
                "teampoints2": {
                    "type": "integer"
                }
            }
        },
        "structs.Player": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "height": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "idteam": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "position": {
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
                "color1": {
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
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Squads API",
	Description:      "This API allows users to manage football teams and players. Users can enter new teams and players, update existing information, and view details.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
