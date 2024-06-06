
<h1 align="center">Smart News Fetcher</h1>

<p align='center'> 
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/>  
    <img src="https://img.shields.io/badge/JWT-F2F4F9?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=black"/>
    <img src="https://img.shields.io/badge/GoLand-0f0f0f?&style=for-the-badge&logo=goland&logoColor=white"/>
</p>  


<p align="center">
  <img src="https://github.com/matheusvidal21/smart-news-fetcher/assets/102569695/1b1ec142-e71c-4c2e-b054-909fab4a97db" alt="Logo" height="250">
</p>


O Smart News Fetcher é um agregador de notícias que coleta e organiza artigos de várias fontes. Ele permite aos usuários adicionar fontes de notícias, buscar artigos, se inscrever em newsletters e gerenciar seu conteúdo de forma eficiente.

## Índice
- 📖 [Introdução](#-introducao)
- 📁 [Estrutura de pacotes](#-estrutura-de-pacotes)
- 💻 [Tecnologias utilizadas](#-tecnologias-utilizadas)
- 🚀 [Uso](#-uso)
- 🔧 [Instalação](#-instalacao)
- 👤 [Autor](#-autor)

# 📖 Introdução

O Smart News Fetcher é uma solução completa para agregar e organizar notícias de diversas fontes. Ele é projetado para facilitar a coleta, processamento e exibição de artigos de notícias de diferentes sites, permitindo que os usuários mantenham-se atualizados com suas fontes preferidas em um único lugar.

Esta aplicação robusta utiliza diversas tecnologias modernas para garantir alta performance, segurança e facilidade de manutenção. Com uma API RESTful bem definida, o Smart News Fetcher possibilita a integração com outros sistemas e a criação de diversas funcionalidades, como autenticação de usuários, gerenciamento de fontes de notícias, manipulação de artigos e envio de newsletters.

Principais recursos:
- Adição e gerenciamento de fontes de notícias.
- Busca e exibição de artigos.
- Inscrição e desinscrição em newsletters.
- Carregamento periódico de feed de fontes.
- Autenticação JWT.
- Envio de emails.

# 📁 Estrutura de pacotes
```
smart-news-fetcher/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.go
├── internal/
│   ├── auth/
│   │   └── jwt_service.go
│   ├── di/
│   │   └── wire.go
│   ├── email/
│   │   └── email_service.go
│   ├── fetcher/
│   │   └── fetcher.go
│   ├── infra/
│   │   ├── database/
│   │   │   └── article_db.go
│   │   │   └── source_db.go
│   │   │   └── user_db.go
│   │   └── handler/
│   │       └── article_handler.go
│   │       └── source_handler.go
│   │       └── user_handler.go
│   │   └── service/
│   │       └── article_service.go
│   │       └── source_service.go
│   │       └── user_service.go
│   ├── interfaces/
│   │   └── article.go
│   │   └── email.go
│   │   └── fetcher.go
│   │   └── source.go
│   │   └── user.go
│   ├── middleware/
│   │   └── middleware.go
├── pkg/
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
├── sql/
│   └── migrations/
│       └── init.sql
├── Dockerfile
├── docker-compose.yaml
├── go.mod
├── go.sum
└── Makefile
````
- `cmd/server`: Contém o ponto de entrada principal da aplicação.
- `configs`: Gerencia as configurações da aplicação.
internal:
- `auth`: Implementa o serviço de autenticação JWT.
- `di`: Gerencia a injeção de dependências utilizando Wire.
- `email`: Implementa o serviço de envio de emails.
- `fetcher`: Contém a lógica para coletar e analisar feeds RSS.
- `database`: Implementa os repositórios para interação com o banco de dados.
- `handler`: Contém os manipuladores HTTP para os endpoints da API.
- `interfaces`: Define interfaces utilizadas em toda a aplicação.
- `middleware`: Implementa middlewares, como o de autenticação.
- `service`: Contém a lógica de negócios dos serviços (artigo, fonte e usuário).
pkg:
- `logger`: Configura e gerencia a criação de logs.
- `utils`: Funções utilitárias utilizadas na aplicação.
- `sql/migrations`: Scripts de migração para o banco de dados.
- `Dockerfile`: Define o ambiente Docker para a aplicação.
- `docker-compose.yaml`: Configuração do Docker Compose para orquestrar os serviços.
- `Makefile`: Contém comandos úteis migrações.

# 🚀 Uso
A aplicação fornece uma API RESTful para interagir com o agregador de notícias. Aqui estão alguns exemplos de endpoints disponíveis:
### Autenticação
- Login de usuário
```
POST /users/login
Content-Type: application/json

{
  "email": "test@gmail.com",
  "password": "password"
}
```
### Artigos
- Obter todos os artigos com paginação e ordenação
```
GET /articles?page=2&limit=1&sort=desc
Content-Type: application/json
Authorization: Bearer {token}
```
### Fontes
- Criar uma nova fonte
```
POST /sources
Content-Type: application/json
Authorization: Bearer {token}

{
  "name": "Folha de São Paulo - Esportes",
  "url": "https://feeds.folha.uol.com.br/esporte/rss091.xml",
  "user_id": 1,
  "update_interval": 1
}
```
- Carregar Feed de uma Fonte
```
GET /sources/load_feed/1
Content-Type: application/json
Authorization: Bearer {token}
```
- Inscrição no Newsletter
``` 
GET /sources/subscribe/1
Content-Type: application/json
Authorization: Bearer {token}
```
### Usuários
- Deletar um usuário
```
DELETE /users/find_by_email/test@gmail.com
Content-Type: application/json
```

# 💻 Tecnologias utilizadas
- `Go`: Linguagem de programação utilizada para desenvolver a aplicação.
- `Gofeed`: Biblioteca Go para coletar e analisar feeds RSS.
- `Gomail`: Biblioteca Go para serviços de emails.
- `Gin`: Framework web escrito em Go, utilizado para criar APIs RESTful.
- `Wire`: Utilizado para injeção de dependências na aplicação.
- `Migrate`: Ferramenta de migração de banco de dados.
- `JWT (JSON Web Tokens)`: Utilizado para autenticação e autorização.
- `MySQL`: Sistema de gerenciamento de banco de dados relacional.
- `Docker`: Utilizado para containerização da aplicação, facilitando a implantação e escalabilidade.

# 🔧 Instalação 
### Pré-requisitos
- Docker
### Passos:
1. Clone o repositório:
```
git clone https://github.com/matheusvidal21/smart-news-fetcher.git
cd smart-news-fetcher
```
2. Configure o arquivo docker-compose.yaml com as informações necessárias, especialmente as linhas de SMTP para o serviço de Gmail:
```
- DB_DRIVER=mysql
- DB_HOST=db
- DB_PORT=3306
- DB_USER=root
- DB_PASSWORD=root
- DB_SOURCE=root:root@tcp(db:3306)/news_aggregator
- DB_NAME=news_aggregator
- WEB_SERVER_PORT=:8080
- JWT_SECRET_KEY=your_secret_key
- JWT_EXPIRATION_MINUTES=60
- SMTP_HOST=smtp.gmail.com
- SMTP_PORT=587
- SMTP_USER=seu-email@gmail.com
- SMTP_PASSWORD=sua-senha
- SMTP_FROM_EMAIL=seu-email@gmail.com
```
3. Construa e inicie os containers Docker:
```
docker-compose up --build
```

Alguns exemplos de fontes RSS que é possível adicionar ao Smart News Fetcher:
- TechCrunch
<br> URL: https://techcrunch.com/feed/
<br> Descrição: Últimas notícias sobre startups, tecnologia e venture capital.

- BBC News - World
<br> URL: http://feeds.bbci.co.uk/news/world/rss.xml
<br> Descrição: Notícias globais e atualizações da BBC.

- CNN - Top Stories
<br> URL: http://rss.cnn.com/rss/edition.rss
<br> Descrição: Principais notícias e histórias do CNN.

# 👤 Autor

| [<img src="https://avatars.githubusercontent.com/u/102569695?s=400&u=f20bbb53cc46ec2bae01f8d60a28492bfdccbdd5&v=4" width=115><br><sub>Matheus Vidal</sub>](https://github.com/matheusvidal21) |
| :---: |

