# Squads API

## Description

This repository contains an API developed in ___'Go'___ that provides information about football teams around the world and many players who are part of these teams. The API uses the Echo framework for routing and a ___'MySQL'___ database hosted on ___'Amazon RDS'___.

![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![GoLand](https://img.shields.io/badge/GoLand-0f0f0f?&style=for-the-badge&logo=goland&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

## Index

1. Preface
2. Description
3. API Routes
   - Team Routes
   - Player Routes
4. How to use
5. API Documentation
6. Installation and Configuration
7. Contributing
8. License

# API Routes

The API has the following routes:

## Team Routes

- `GET /api/v1/team/getall`: Returns all teams.
- `POST /api/v1/team/insert`: Inserts a new team.
- `PUT /api/v1/team/update`: Updates an existing team.
- `DELETE /api/v1/team/delete`: Deletes a team.
- `GET /api/v1/team/getbyid/:id`: Returns a team by its ID.
- `GET /api/v1/team/getbyname/:name`: Returns a team by its name.
- `GET /api/v1/team/getbycountry/:country`: Returns all teams from a given country.

To use this API, you can send an HTTP request to the desired route. For example, to get all teams, you can send a GET request to `/team/getall`.

## Player Routes

- `GET /api/v1/player/getall`: Returns all players.
- `POST /api/v1/player/insert`: Inserts a new player.
- `PUT /api/v1/player/update`: Updates an existing player.
- `DELETE /api/v1/player/delete`: Deletes a player.
- `GET /api/v1/player/getbyid/:id`: Returns a player by its ID.
- `GET /api/v1/player/getbyname/:name`: Returns a player by its name.
- `GET /api/v1/player/getbycountry/:country`: Returns all players from a given country.

To use this API, you can send an HTTP request to the desired route. For example, to get all players, you can send a GET request to `/player/getall`.

## API documentation
This API uses __Swagger__ to automatically generate API documentation with the OpenAPI 2.0 standard. You can access the API documentation via `/swagger/index.html` in your browser. The API documentation contains information about routes, parameters, return types, and request and response examples.

## Installation and Configuration
To install and configure this API locally, follow these steps:
1. Clone this repository to your local environment.
2. Install Go if it is not already installed.
3. Install MySQL and configure an Amazon RDS instance.
4. Update the database settings in your code to point to your Amazon RDS instance.
5. Run the Go server on port `8080`.

## Contributing

Contributions are welcome! Feel free to open an issue or send a pull request.

## License

This project is licensed under the terms of the MIT License.
