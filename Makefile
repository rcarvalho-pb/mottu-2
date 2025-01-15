DB_PATH=/Users/ramon/Projects/go/mottu/database/db.db
USER_BINARY=userApp
USER_SERVICE_PORT=12345
create-migration:
	@goose -dir ./db_migration/files/migrations/ create ${name} sql
config-db:
	@cd ./db_migration/ && export DB_PATH=${DB_PATH} && go run main.go
user-service:
	@echo building user service...
	@cd ./user_service/ && go build -o app/${USER_BINARY} ./cmd
	@echo building done, starting user service on port ${USER_SERVICE_PORT}
	@cd ./user_service/ && export DB_PATH=${DB_PATH} USER_SERVICE_PORT=${USER_SERVICE_PORT} && app/${USER_BINARY} &
	@echo user service started

stop:
	@echo Stoping services
	pkill -SIGTERM -f "${USER_BINARY}"
	@echo services stoped
