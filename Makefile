migration-create:
	migrate create -ext sql -dir migrations/sql $(name)

migration-up:
	migrate -path migrations/sql -verbose -database "${DATABASE_URI}" up

migration-down:
	migrate -path migrations/sql -verbose -database "${DATABASE_URI}" down

run-db:
	docker run --name wepress -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres:latest

stop-db:
	docker stop wepress
	docker rm wepress -v

run-redis:
	docker run --name wepress-redis -p 6379:6379 -d redis:latest

stop-redis:
	docker stop wepress-redis
	docker rm wepress-redis -v

run-api:
	go run ./cmd/api/

build-api:
	docker build --target=exporter -t export-wepress-core --output=./dist .

watch-api:
	air -c .air.toml

test:
	go test -v -cover -benchmem ./...

mock:
	mockery --all

setup:
	go mod tidy
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

docs:
	swag i --dir ./cmd/api/,\
	./modules/,\
	./pkg/wrapper,\
	./pkg/contexts

git-hooks:
	echo "Installing hooks..." && \
	rm -rf .git/hooks/pre-commit && \
	ln -s ../../tools/scripts/pre-commit.sh .git/hooks/pre-commit && \
	chmod +x .git/hooks/pre-commit && \
	echo "Done!"

routes:
	go run ./tools/routes/

gen-proto:
	protoc -I. \
    --go_out=. \
    --go-grpc_out=. \
    modules/printer/proto/*.proto

.PHONY: routes run-api run-db build-api migration-create docs gen-proto