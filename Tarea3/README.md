# Laboratorio 3 - Sistemas Distribuidos 2024-2

## Grupo-06

- Felipe Marchant 202173643-3
- Felipe Muñoz 201973512-8

## Instrucciones

Para crear las imágenes, construir los contenedores y correr todo:
Desde la carpeta grupo-06 escribir en la consola en el siguiente orden:

1. (en dist021) make docker-servidor1
2. (en dist022) make docker-servidor2
3. (en dist023) make docker-servidor3
4. (en dist024) make docker-broker
5. (en dist021) make docker-supervisor
6. (en dist023) make docker-supervisor
7. (en dist022) make docker-jayce

## Consideraciones

1. En las máquinas virtuales que contienen un servidor y un supervisor o un servidor y a Jayce, solo se muestran los logs de supervisor o Jayce. Para ver los logs de los servidores se debe ejecutar el comando docker logs <contenedor> en las respectivas máquinas, donde <contenedor> es el nombre del contenedor.

2. Se asume que la asignación de servidores hecha por el broker es al azar, tanto para los supervisores como para Jayce.

3. Dockerización:
   - Las entidades Servidor 1 y Supervisor 1 están dockerizadas en la MV1(dist021).
   - Las entidades Servidor 2 y Jayce están dockerizadas en la MV2 (dist022).
   - Las entidades Servidor 3 y Supervisor 2 están dockerizadas en la MV3 (dist023).
   - La entidad Broker está dockerizada en la MV4 (dist024).
