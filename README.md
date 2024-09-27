# Laboratorio 1 - Sistemas Distribuidos 2024-2

## Grupo-06

- Felipe Marchant 202173643-3
- Felipe Muñoz 201973512-8

## Instrucciones
2-Para crear las imagenes y construir el contenedor:
Desde la carpeta grupo-06 escribir en la consola:
    1- make docker-logistica
    2- make docker-caravanas
    3- make docker-clientes
    4- make docker-finanzas

3-Para iniciar los contenedores
    1- docker run app_logistica
    2- docker run app_caravanas
    3- docker run app_facciones
    4- docker run app_finanzas

## Consideraciones

1-Se hicieron modifaciones solo a lo que se respecta a el manejor de los puertos:
    Rabbit-Falto exponer el puerto( y se hizo una modificación con lo que respecta a la conexion del puerto en logistica.go 
)
    505001(en la primera entrega falto exponerlo en docker-compose)

2-Se cambio el nombre del Makefile MAKEFILE ->makefile


3-El servidor RabbitMQ se encuentra en el mismo contenedor que el servidor de logística

dist021--->Logistica, dist022--->Caravanas, dist023--->Facciones/clientes, dist024---->Finanzas


Consideramos que todos los pedidos se harán desde un archivo "input.txt" como se planteó en el foro del aula.
Ejemplo de input que consideramos:

0001,Ostronitas,Choripán,200,Escolta1,Destino1,123
0002,Normal,Anticucho,450,Escolta2,Destino5,777
0003,Prioritario,Pan,100,Escolta1,Destino2,465
0004,Prioritario,Coca-Cola,75,Escolta2,Destino1,768
...

Además, para que todo finalice y se muestre el saldo final de finanzas, se usará el input "0,Normal,100,final,final,0" como último pedido en el archivo de inputs. Por lo tanto, un ejemplo de inputs que funcionan será:

0001,Ostronitas,Choripán,200,Escolta1,Destino1,123
0002,Normal,Anticucho,450,Escolta2,Destino5,777
0003,Prioritario,Pan,100,Escolta1,Destino2,465
0004,Prioritario,Coca-Cola,75,Escolta2,Destino1,768
0,Normal,100,final,final,0
