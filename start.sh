#!/bin/bash

# Espera até que o MySQL esteja disponível
./wait-for-it.sh db 3306 -- echo "MySQL is up"

# Executa as migrações
echo "Running migrations"
/usr/local/bin/migrate -source "file:///root/sql/migrations" -database "mysql://root:root@tcp(db:3306)/news_aggregator" up

# Verifica se o comando de migração foi bem-sucedido
if [ $? -ne 0 ]; then
  echo "Migration failed"
  exit 1
fi

# Inicia a aplicação
echo "Starting application"
./smart-news-fetcher
