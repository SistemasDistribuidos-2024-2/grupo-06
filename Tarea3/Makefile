# Comando para crear e inciar el servicio de servidores
docker-servidor1:
	@docker-compose ./app/docker-compose.yml run hextech1
docker-servidor2:
	@docker-compose ./app/docker-compose.yml run hextech2
docker-servidor3:
	@docker-compose ./app/docker-compose.yml run hextech3

# Comando para crear e inciar  el servicio de jayce
docker-jayce:
	@docker-compose -f ./app/docker-compose.yml run jayce
# Comando para crear e inciar el servicio de Broker
docker-broker:
	@docker-compose -f ./app/docker-compose.yml up broker

# Comando para crear e inciar un servicio de Supervisor
docker-supervisor:
	@docker-compose -f ./app/docker-compose.yml run supervisor
