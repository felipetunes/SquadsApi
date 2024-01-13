# Squads API

## Descrição

Este repositório contém uma API desenvolvida em *Go* que fornece informações sobre times de futebol ao redor do mundo e muitos jogadores que fazem parte desses times. A API utiliza o framework Echo para roteamento e um banco de dados *MySQL* hospedado no ####Amazon RDS####.

## Rotas da API

A API possui as seguintes rotas:

### Rotas dos Times

- `GET /team/getall`: Retorna todos os times.
- `GET /team/insert`: Insere um novo time.
- `GET /team/update`: Atualiza um time existente.
- `GET /team/delete`: Exclui um time.
- `GET /team/getbyid/:id`: Retorna um time pelo seu ID.
- `GET /team/getbyname/:name`: Retorna um time pelo seu nome.
- `GET /team/getbycountry/:country`: Retorna todos os times de um determinado país.

### Rotas dos Jogadores

- `GET /player/getall`: Retorna todos os jogadores.
- `GET /player/getbyidteam/:idteam`: Retorna todos os jogadores de um determinado time.

## Como usar

Para usar esta API, você pode enviar uma solicitação HTTP para a rota desejada. Por exemplo, para obter todos os times, você pode enviar uma solicitação GET para `/team/getall`.

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
