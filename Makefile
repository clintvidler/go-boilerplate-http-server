logs:
	docker-compose logs -f main

run:
	docker-compose up -d

run -b:
	docker-compose up --build -d

stop:
	docker-compose stop
