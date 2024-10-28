# Laboratorio 2 - Sistemas Distribuidos 2024-2

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

1. Se proporcionó un archivo "DIGIMONS.txt" a cada servidor regional (isla file, continente server y continente folder).
2. A cada archivo de texto .txt ("DIGIMONS.txt" e "input.txt") se les proporcionó datos, lo cuales se pueden cambiar pero siguiendo el formato de cada archivo (especificado en las instrucciones del laboratorio).
3. El orden de ejecución es el siguiente: - Primero se deben ejecutar ambos data nodes (data node 1 y data node 2). - Luego se debe ejecutar el Primary node. - Seguido de eso, se deben ejecutar los servidores regionales en cualquier orden (isla file, continente server y continente folder). - Después, se puede ejecutar Diaboromon. - Finalmente se puede ejecutar el Nodo Tai.
   Este orden va en base a no generar errores como que los servidores regionales no encuentren a Primary Node (lo cual provocaría que se deban volver a ejecutar para que se comuniquen), que Primary node no encuentre a los Data Nodes (lo cual haría que los Data Nodes no guarden su información en sus archivos .txt), que el Nodo Tai no encuentre a Diaboromon (por lo que no se ejecutaría la pelea), o que Nodo Tai no encuentre al Primary Node (lo que haría que no se pudieran cargar los datos al mismo).
