# Laboratorio 1 - Sistemas Distribuidos 2024-2

## Grupo-06

- Felipe Marchant 202173643-3
- Felipe Muñoz 201973512-8

## Instrucciones

– make docker-logistica: Iniciará el código hecho en Docker para el sistema de logı́stica (Konzu).
– make docker-finanzas: Iniciará el código hecho en Docker para el sistema financiero (Raquis).
– make docker-caravanas: Iniciará el código hecho en Docker para los camiones.
– make docker-clientes: Iniciará el código hecho en Docker para los clientes.

## Consideraciones

El servidor RabbitMQ se encuentra en el mismo contenedor que el servidor de logística

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
