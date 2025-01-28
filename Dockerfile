# Dockerfile
FROM golang:1.22.11 AS builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie os arquivos go.mod e go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante do código
COPY . .

# Compile o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o CEPrun api/cmd/main.go

# Etapa 2: Criar a imagem final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/CEPrun .
COPY .env ./

CMD ["./CEPrun"]