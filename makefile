run: 
	go build -o ./build/app ./cmd/main.go && ./build/app

ps_run: 
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=Qwerty1! -e POSTGRES_DB=turtle -d postgres:latest
	docker exec -it postgres bash
	psql -h localhost -U admin turtle

ps_stop: 
	docker stop postgres