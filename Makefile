include .env

server:
	gow run cmd/blog/main.go -e .env

migrate:
	goose -dir 'app/db/migrations' mysql ${DATABASE_URL} up

migrations:
	goose -dir app/db/migrations create $(name) sql

rollback:
	goose -dir 'app/db/migrations' mysql ${DATABASE_URL} down