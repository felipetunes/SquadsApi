# Squads API

## Descrição

Este repositório contém uma API desenvolvida em ___'Go'___ que fornece informações sobre times de futebol ao redor do mundo e muitos jogadores que fazem parte desses times. A API utiliza o framework Echo para roteamento e um banco de dados ___'MySQL'___ hospedado no ___'Amazon RDS'___.

![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![GoLand](https://img.shields.io/badge/GoLand-0f0f0f?&style=for-the-badge&logo=goland&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

# Índice

1. Prefácio
2. Descrição
3. Rotas da API
   - Rotas dos Times
   - Rotas dos Jogadores
4. Como usar
5. Documentação da API
6. Instalação e Configuração
7. Contribuindo
8. Licença

## Rotas da API

A API possui as seguintes rotas:

Rotas dos Times
GET /api/v1/team/getall: Retorna todos os times.
POST /api/v1/team/insert: Insere um novo time.
PUT /api/v1/team/update: Atualiza um time existente.
DELETE /api/v1/team/delete: Exclui um time.
GET /api/v1/team/getbyid/:id: Retorna um time pelo seu ID.
GET /api/v1/team/getbyname/:name: Retorna um time pelo seu nome.
GET /api/v1/team/getbycountry/:country: Retorna todos os times de um determinado país.

Rotas dos Jogadores
GET /api/v1/player/getall: Retorna todos os jogadores.
GET /api/v1/player/getbyidteam/:idteam: Retorna todos os jogadores de um determinado time.

## Como usar

Para usar esta API, você pode enviar uma solicitação HTTP para a rota desejada. Por exemplo, para obter todos os times, você pode enviar uma solicitação GET para `/team/getall`.

## Documentação da API
Esta API utiliza o Swagger para gerar automaticamente a documentação da API com o padrão OpenAPI 2.0. Você pode acessar a documentação da API pelo caminho /swagger/index.html no seu navegador. A documentação da API contém informações sobre as rotas, os parâmetros, os tipos de retorno e os exemplos de solicitação e resposta.

## Instalação e Configuração 
Para instalar e configurar esta API localmente, siga estas etapas:

1. Clone este repositório para o seu ambiente local.
2. Instale Go, se ainda não estiver instalado.
3. Instale o MySQL e configure uma instância do Amazon RDS.
4. Atualize as configurações do banco de dados no código para apontar para a sua instância do Amazon RDS.
5. Execute o servidor Go.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob os termos da licença MIT.
