# Comando para iniciar el servicio de Logística (Konzu)
docker-logistica:
	@docker-compose -f ./app/docker-compose.yml up logistica

# Comando para iniciar el servicio de Finanzas (Raquis)
docker-finanzas:
	@docker-compose -f ./app/docker-compose.yml up finanzas

# Comando para iniciar el servicio de Caravanas
docker-caravanas:
	@docker-compose -f ./app/docker-compose.yml up caravanas

# Comando para iniciar el servicio de Clientes( Clientes  Se usa indistintamente con facciones)
docker-clientes:
	@docker-compose -f ./app/docker-compose.yml up facciones
