# Imagen base con Go
FROM golang:latest

ADD . /backend

WORKDIR /backend

# Instalar dependencias de Go
RUN go mod init backend
RUN go mod tidy
RUN go mod download

# Compilar la aplicacións
RUN go build -o main .

# Exponer el puerto en el que la aplicación escucha
EXPOSE 4000


# Comando para ejecutar la aplicación
CMD ["./main"]

