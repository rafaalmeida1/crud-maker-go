# Dockerfile.dev
FROM golang:1.18

# Instalar fresh
RUN go install github.com/gravityblast/fresh@latest

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o código-fonte
COPY . .

# Instalar dependências
RUN go mod tidy

# Rodar fresh para assistir mudanças
CMD ["fresh"]
