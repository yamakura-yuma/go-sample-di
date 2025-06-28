.PHONY: help build run clean start kill-port api-create-user api-get-user api-update-user api-delete-user test test-html swagger-ui
kill-port:
	@fuser -k 8080/tcp || true

help:
	@echo "Usage: make [target]"
	@echo "  build            Goバイナリをビルド（swag initも実行）"
	@echo "  start            Goアプリをローカルで起動 (8080番ポート)"
	@echo "  swagger-ui       Swagger UI (http://localhost:8080/swagger/index.html) をブラウザで開く"
	@echo "  api-create-user  サンプルユーザー作成APIリクエスト (curl)"
	@echo "  api-get-user     サンプルユーザー取得APIリクエスト (curl)"
	@echo "  api-update-user  サンプルユーザー更新APIリクエスト (curl)"
	@echo "  api-delete-user  サンプルユーザー削除APIリクエスト (curl)"
	@echo "  test             ユニットテストを実行"

APP_NAME=app

build:
	go mod tidy
	cd cmd && swag init --parseDependency --parseInternal --output ../docs || true
	go build -o $(APP_NAME) ./cmd/main.go

start:
	go run ./cmd/main.go

swagger-ui:
	@echo 'Swagger UI: http://localhost:8080/swagger/index.html'
	@python3 -m webbrowser 'http://localhost:8080/swagger/index.html' || true

test:
	go test ./internal/... \
		-v \
		-cover
api-create-user:
	@echo 'POST /users'
	@curl -X POST -H "Content-Type: application/json" -d '{"id":"1","name":"テスト","email":"test@example.com"}' http://localhost:8080/users

api-get-user:
	@echo 'GET /users/1'
	@curl http://localhost:8080/users/1

api-update-user:
	@echo 'PUT /users/1'
	@curl -X PUT -H "Content-Type: application/json" -d '{"name":"テスト更新","email":"test2@example.com"}' http://localhost:8080/users/1

api-delete-user:
	@echo 'DELETE /users/1'
	@curl -X DELETE http://localhost:8080/users/1
