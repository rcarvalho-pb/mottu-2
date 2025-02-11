DB_PATH=/Users/ramon/Projects/go/mottu/database/db.db
IMAGES_DIRECTORY=/Users/ramon/Projects/go/mottu/database/files
USER_BINARY=userApp
USER_SERVICE_PORT=12346
BROKER_BINARY=brokerApp
BROKER_SERVICE_PORT=12345
AUTH_BINARY=brokerApp
AUTH_SERVICE_PORT=12347
run: user-service broker-service
user-service:
	@echo building user service...
	@cd ./user_service/ && go build -o app/${USER_BINARY} ./cmd
	@echo building done, starting user service on port ${USER_SERVICE_PORT}
	@cd ./user_service/ && export DB_PATH=${DB_PATH} IMAGES_DIRECTORY=${IMAGES_DIRECTORY} USER_SERVICE_PORT=${USER_SERVICE_PORT} && app/${USER_BINARY} &
	@echo user service started
auth-service:
	@echo building auth service...
	@cd ./broker_service/ && go build -o app/${AUTH_BINARY} ./cmd
	@echo building done, starting broker service on port ${AUTH_SERVICE_PORT}
	@cd ./broker_service/ && export USER_SERVICE_PORT=${USER_SERVICE_PORT} AUTH_SERVICE_PORT=${AUTH_SERVICE_PORT} && app/${AUTH_BINARY} &
	@echo auth service started
broker-service:
	@echo building broker service...
	@cd ./broker_service/ && go build -o app/${BROKER_BINARY} ./cmd
	@echo building done, starting broker service on port ${BROKER_SERVICE_PORT}
	@cd ./broker_service/ && export BROKER_SERVICE_PORT=${BROKER_SERVICE_PORT} USER_SERVICE_PORT=${USER_SERVICE_PORT} && app/${BROKER_BINARY} &
	@echo broker service started
restart: stop run
stop:
	@echo Stoping services
	pkill -SIGTERM -f "${USER_BINARY}"
	pkill -SIGTERM -f "${BROKER_BINARY}"
	@echo services stoped
create-migration:
	@goose -dir ./db_migration/files/migrations/ create ${name} sql
config-db:
	@cd ./db_migration/ && export DB_PATH=${DB_PATH} && go run main.go
