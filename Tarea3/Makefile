# Comando para crear e inciar el servicio de servidores
docker-servidor1:
	@docker-compose -f ./app/docker-compose.yml run -d hextech1
docker-servidor2:
	@docker-compose -f ./app/docker-compose.yml run -d hextech2
docker-servidor3:
	@docker-compose -f ./app/docker-compose.yml run -d hextech3
# Comando para crear e inciar  el servicio de jayce
docker-jayce:
	@docker-compose -f ./app/docker-compose.yml run jayce
# Comando para crear e inciar el servicio de Broker
docker-broker:
	@docker-compose -f ./app/docker-compose.yml run broker
# Comando para crear e inciar un servicio de Supervisor
docker-supervisor:
	@docker-compose -f ./app/docker-compose.yml run supervisor
# Comando para iniciar el contenedor suspicious_leakey
docker-run-servidor1:
	@sudo docker run -d -p 50051:50051 -e SERVER_ID=1 -v $(pwd)/logs/hextech1:/app/logs -v $(pwd)/mercancias/hextech1:/app/mercancias 10e2eb97c1e9
docker-run-servidor2:
	@sudo docker run -d -p 50052:50052 -e SERVER_ID=2 \
	-v $(pwd)/logs/hextech2:/app/logs \
	-v $(pwd)/mercancias/hextech2:/app/mercancias \
	547e37b93036
docker-run-servidor3:
	@sudo docker run -d -p 50053:50053 -e SERVER_ID=3 \
	-v $(pwd)/logs/hextech3:/app/logs \
	-v $(pwd)/mercancias/hextech3:/app/mercancias \
	9879ba137678