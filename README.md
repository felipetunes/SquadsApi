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

## API documentation
This API uses __Swagger__ to automatically generate API documentation with the OpenAPI 2.0 standard. You can access the API documentation via `/swagger/index.html` in your browser. The API documentation contains information about routes, parameters, return types, and request and response examples.

## Installation and Configuration
To install and configure this API locally, follow these steps:
1. Clone este repositório para o seu ambiente local.
2. Instale Go, se ainda não estiver instalado.
3. Instale o MySQL e configure uma instância do Amazon RDS.
4. Atualize as configurações do banco de dados no código para apontar para a sua instância do Amazon RDS.
5. Execute o servidor Go.

## Contributing

Contributions are welcome! Feel free to open an issue or send a pull request.

## License

This project is licensed under the terms of the MIT License.
