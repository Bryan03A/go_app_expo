# Usamos la imagen oficial de Go basada en Alpine
FROM golang:1.23.3-alpine

# Seteamos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos el go.mod y go.sum (si existen)
COPY go.mod ./

# Instalamos las dependencias con go mod
RUN go mod tidy

# Copiamos el código fuente y los archivos estáticos
COPY . .

# Exponemos el puerto 5005
EXPOSE 5005

# Comando para ejecutar la aplicación
CMD ["go", "run", "app.go"]