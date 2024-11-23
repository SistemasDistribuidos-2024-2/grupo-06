# Comando para crear e inciar el servicio de servidores
docker-servidor1:
	@docker-compose -f ./app/docker-compose.yml up hextech1
docker-servidor2:
	@docker-compose -f ./app/docker-compose.yml up hextech2
docker-servidor3:
	@docker-compose -f ./app/docker-compose.yml up hextech3

# Comando para crear e inciar  el servicio de jayce
docker-jayce:
	@docker-compose -f ./app/docker-compose.yml up jayce

# Comando para crear e inciar el servicio de Broker
docker-broker:
	@docker-compose -f ./app/docker-compose.yml up broker

# Comando para crear e inciar el servicio de Supervisores
docker-supervisores:
	@docker-compose -f ./app/docker-compose.yml up supervisores


# Comando para iniciar el servicio de servidores
start-servidor1:
    @docker start hextech1
start-servidor2:
    @docker start hextech2
start-servidor3:
    @docker start hextech3

# Comando para iniciar el servicio de jayce
start-jayce:
    @docker start jayce

# Comando para iniciar el servicio de Broker
start-broker:
    @docker start broker

# Comando para iniciar el servicio de Supervisores
start-supervisores:
    @docker start supervisores