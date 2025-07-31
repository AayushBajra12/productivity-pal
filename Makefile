build:
	docker build -t productivity-pal .

run:
	@echo "Starting application"
	docker build -t productivity-pal .

	#docker run --name propal productivity-pal
	docker-compose up --build

# Delete all containers
clean-containers:
	docker rm -f $$(docker ps -aq) || true

# Delete all images
clean-images:
	docker rmi -f $$(docker images -q) || true

# Delete all volumes
clean-volumes:
	docker volume rm $$(docker volume ls -q) || true

# Prune everything (stopped containers, networks, build cache, etc.)
prune-all:
	docker system prune -af --volumes

# Clean all: containers, images, volumes
clean-all: clean-containers clean-images clean-volumes

# Hard reset using prune
reset-docker: