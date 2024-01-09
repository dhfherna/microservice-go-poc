# Dockerfile
FROM mcr.microsoft.com/devcontainers/universal:2

# Install Go
RUN curl -O https://dl.google.com/go/go1.18.linux-amd64.tar.gz \
    && tar -xvf go1.18.linux-amd64.tar.gz \
    && mv go /usr/local \
    && rm go1.18.linux-amd64.tar.gz

# Set environment variables for Go
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Instala Postgres
RUN apt-get update && apt-get install -y postgresql postgresql-contrib

# Configura Postgres para que se inicie automáticamente
RUN service postgresql start && \
    sudo -u postgres psql -c "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" && \
    sudo -u postgres createdb -O docker docker

# Expone el puerto de Postgres
EXPOSE 5432
