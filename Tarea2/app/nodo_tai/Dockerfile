FROM golang:latest

WORKDIR /app
#Copio todo lo que esta en el directorio actual
COPY . .


#Ejecuto el siguiente comando. Nos aseguramos que todas las dependencias esten alineadas
RUN go mod tidy
#Construyo el binario
RUN go build -o nodo_tai

EXPOSE 50059

#Ejecuto el build creado
CMD ["./nodo_tai"]