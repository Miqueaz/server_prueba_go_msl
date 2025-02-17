# Etapa de construcción
FROM golang:1.23.5 as builder

WORKDIR /app

# Copiar los módulos de Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Construir el binario
RUN go build -o server main.go

# Imagen final sin Go instalado
FROM debian:bookworm-slim

# Instalar curl en la imagen final
RUN apt-get update && apt-get install -y curl ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/server .

# Exponer el puerto
EXPOSE 50051

# Comando por defecto (sobreescrito en el compose)
CMD ["./server"]
