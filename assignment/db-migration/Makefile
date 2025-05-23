
rollbackCount ?= 1

run_liquibase_in_go:
	go run main.go

run_migrations:
	liquibase update

rollback_migrations:
	liquibase rollbackCount $(rollbackCount)

# run_migrations_in_docker:
# 	docker --version
# 	docker info
# 	docker rm -f liquibase-runner 2>/dev/null
# 	docker network inspect liquibase-net >/dev/null 2>&1 || docker network create liquibase-net
# 	docker run --network liquibase-net --name liquibase-runner -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=testpass -e POSTGRES_DB=testdb -p 5432:5432 -d postgres
# 	# docker exec -it liquibase-runner psql -h localhost -U testuser -d testdb
# 	docker run liquibase-runner --rm \
# 		--network liquibase-net \
# 		-v ./db/changelog/changelog.xml:/liquibase/changelog \
# 		liquibase/liquibase \
# 		--url=jdbc:postgresql://host.docker.internal:5432/testdb \
# 		--changeLogFile=changelog.xml \
# 		--username=testuser \
# 		--password=testpass \
# 		update
