migration-create:
	migrate create -ext sql -dir migrations/sql $(name)

migration-up:
	migrate -path migrations/sql -verbose -database "${DATABASE_URL}" up

migration-down:
	migrate -path migrations/sql -verbose -database "${DATABASE_URL}" down

run-db:
	docker compose up postgres

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

.PHONY: routes run-api run-db build-api migration-create docs