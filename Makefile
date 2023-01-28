createdb:
	docker exec -it pg_container createdb --username=ayoub --owner=ayoub youtube

dropdb:
		docker exec -it pg_continer dropdb youtube

postgres:
	  docker run --name  pg_container -p 5432:5432  -e POSTGRES_USER=ayoub -e POSTGRES_PASSWORD=myPwd -d postgres:alpine

migrateup:
	 migrate -path db/migration -database "postgresql://ayoub:myPwd@localhost:5432/youtube?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://ayoub:myPwd@localhost:5432/youtube?sslmode=disable" -verbose down

#sqlc:
#	/snap/bin/sqlc generate

test:
	go test -v -cover ./...

.PHONY:createdb dropdb postgres migrateup migratedown  test
