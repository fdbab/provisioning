# Verwenden Sie ein offizielles Golang-Image als Basis
FROM golang:1.20.0-alpine

# Definieren Sie den Arbeitsverzeichnis
WORKDIR /app

# Kopieren Sie den Quellcode in den Container
COPY . .

# Kompilieren Sie das Go-Programm
RUN go build -o app

# Verwenden Sie ein offizielles Alpine Linux-Image als Basis
FROM alpine:latest

# Kopieren Sie das kompilierte Go-Programm in den Container
COPY --from=0 /app/app /usr/local/bin/app

# Führen Sie das Go-Programm im Container aus
CMD ["app"]
