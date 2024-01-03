# Dockerfile
FROM golang:1.21.5

# Instala las herramientas de Go necesarias
RUN go get github.com/go-delve/delve/cmd/dlv

# Instala Postgres
RUN apt-get update && apt-get install -y postgresql postgresql-contrib

# Configura Postgres para que se inicie automáticamente
RUN service postgresql start && \
    sudo -u postgres psql -c "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" && \
    sudo -u postgres createdb -O docker docker

# Expone el puerto de Postgres
EXPOSE 5432
