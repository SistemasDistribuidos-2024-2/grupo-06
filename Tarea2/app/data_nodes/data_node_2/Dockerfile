FROM golang:latest

WORKDIR /app/data_nodes
#Copio todo lo que esta en el directorio actual
COPY . .


#Ejecuto el siguiente comando. Nos aseguramos que todas las dependencias esten alineadas
RUN go mod tidy
#Construyo el binario
RUN go build -o data_node_2

EXPOSE 50053

#Ejecuto el build creado
CMD ["./data_node_2"]