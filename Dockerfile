# Utiliza a imagem oficial do Go para construção da aplicação
FROM golang:1.20-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos go.mod e go.sum e faz download das dependências
COPY go.mod go.sum ./
RUN go mod download

# Instala a ferramenta migrate
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copia o restante do código fonte
COPY . .

# Compila a aplicação
RUN go build -o /smart-news-fetcher ./cmd/server

# Utiliza uma imagem minimalista para execução da aplicação
FROM alpine:latest

# Instala as dependências necessárias
RUN apk --no-cache add ca-certificates bash mysql-client

# Define o diretório de trabalho
WORKDIR /root/

# Cria o diretório de logs
RUN mkdir -p /root/logs

# Copia o binário da aplicação do estágio de construção
COPY --from=builder /smart-news-fetcher .

# Copia o arquivo de configuração
COPY --from=builder /app/cmd/server/.env .

# Copia os scripts wait-for-it e start.sh
COPY wait-for-it.sh start.sh .

# Copia o binário da ferramenta migrate do estágio de construção
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Permite a execução dos scripts
RUN chmod +x wait-for-it.sh start.sh

# Copia os scripts de migração
COPY --from=builder /app/sql /root/sql

# Exponha a porta que o servidor irá rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./start.sh"]
