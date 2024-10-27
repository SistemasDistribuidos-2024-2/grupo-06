# Comando para iniciar los servicios de la maquina 1: isla_file y  data_node_1
docker-MV1:
	@docker-compose -f ./app/docker-compose.yml up -d image_isla_file
	@docker-compose -f ./app/docker-compose.yml up -d image_data_node_1

docker-MV2:
	@docker-compose -f ./app/docker-compose.yml up -d image_continente_folder
	@docker-compose -f ./app/docker-compose.yml up -d image_diaboromon
docker-MV3:
	@docker-compose -f ./app/docker-compose.yml up -d image_continente_server
	@docker-compose -f ./app/docker-compose.yml up -d image_data_node_2
docker-MV4:
	@docker-compose -f ./app/docker-compose.yml up -d image_primary_node
	@docker-compose -f ./app/docker-compose.yml up -d image_nodo_tai
docker-PR:
	@docker-compose -f ./app/docker-compose.yml up -d image_primary_node
	@docker-compose -f ./app/docker-compose.yml up -d image_isla_file

