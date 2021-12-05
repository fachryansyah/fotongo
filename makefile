COMPONENT=fotongo

serve:
	go run main.go
build:
	go build main.go
run-prod:
	COMMAND="make prisma-push && make prisma-gen && make serve" APP_ENV="prod" docker-compose up -d --force-recreate
run-local:
	COMMAND="make prisma-push && make prisma-gen && make serve" APP_ENV="local" docker-compose up
gen-swagger:
	COMMAND="go install github.com/swaggo/swag/cmd/swag@latest && swag init" docker-compose up
test:
	docker container rm $(COMPONENT)_db_test_1
	docker volume rm $(COMPONENT)_$(COMPONENT)_data_test
	COMMAND="make prisma-push && make prisma-gen && go test ./app/modules/$(MODULE) -v" docker-compose up --force-recreate --abort-on-container-exit
test-all:
	docker container rm $(COMPONENT)_db_test_1
	docker volume rm $(COMPONENT)_$(COMPONENT)_data_test
	COMMAND="make migrate-up-test && go test ./app/modules... -v" docker-compose up --force-recreate --abort-on-container-exit
prisma-gen:
	go run github.com/prisma/prisma-client-go generate --schema=./infrastructure/services/prisma/schema.prisma
prisma-push:
	go run github.com/prisma/prisma-client-go db push --schema=./infrastructure/services/prisma/schema.prisma
prisma-pull:
	go run github.com/prisma/prisma-client-go db pull --schema=./infrastructure/services/prisma/schema.prisma
prisma-format:
	go run github.com/prisma/prisma-client-go format --schema=./infrastructure/services/prisma/schema.prisma
prisma-migrate-dev-init:
	go run github.com/prisma/prisma-client-go migrate dev --name init --schema=./infrastructure/services/prisma/schema.prisma
prisma-migrate-new:
	go run github.com/prisma/prisma-client-go migrate dev --name $(NAME) --schema=./infrastructure/services/prisma/schema.prisma
prisma-studio:
	go run github.com/prisma/prisma-client-go studio --schema=./infrastructure/services/prisma/schema.prisma
gen-swagger:
	COMMAND="go install github.com/swaggo/swag/cmd/swag@latest && swag init" docker-compose up
gen-wire:
	go run github.com/google/wire/cmd/wire