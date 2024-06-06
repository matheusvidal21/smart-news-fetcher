createmigration:
	docker-compose run --rm app migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	docker-compose run --rm app migrate -path=sql/migrations -database "mysql://root:root@tcp(db:3306)/news_aggregator" -verbose up

migratedown:
	docker-compose run --rm app migrate -path=sql/migrations -database "mysql://root:root@tcp(db:3306)/news_aggregator" -verbose down

.PHONY: migrate migratedown createmigration
