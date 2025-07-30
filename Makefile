build:
	docker build -t productivity-pal .

run:
	@echo "Starting application"
	#docker run --name propal productivity-pal
	docker-compose up --build