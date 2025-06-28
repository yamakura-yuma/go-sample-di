.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "  build            Goバイナリをビルド（swag initも実行）"
	@echo "  clean            Goのキャッシュとモジュールキャッシュをクリア"
	@echo "  start            Goアプリをローカルで起動 (8080番ポート)"
	@echo "  kill		      8080番ポートを使用しているプロセスを強制終了"
	@echo "  test             ユニットテストを実行"
	@echo "  swagger-ui       Swagger UI (http://localhost:8080/swagger/index.html) をブラウザで開く"
	@echo "  api-create-user  サンプルユーザー作成APIリクエスト (curl)"
	@echo "  api-get-user     サンプルユーザー取得APIリクエスト (curl)"
	@echo "  api-update-user  サンプルユーザー更新APIリクエスト (curl)"
	@echo "  api-delete-user  サンプルユーザー削除APIリクエスト (curl)"



.PHONY: build clean run clean start kill test swagger-ui
APP_NAME=app

build:
	@echo ""
	go mod tidy
	@echo ""
	swag init -g ./cmd/main/main.go --parseDependency --parseInternal --output ./docs || true
	@echo ""
	go build -o $(APP_NAME) ./cmd/main/main.go

clean:
	@echo ""
	go clean -cache -modcache -i -r
	@echo ""
	rm -rf ./docs

start:
	@echo ""
	go run ./cmd/main/main.go

kill:
	@fuser -k 8080/tcp || true

test:
	@echo ""
	go test ./internal/... \
		-v \
		-cover

swagger-ui:
	@echo 'Swagger UI: http://localhost:8080/swagger/index.html'
	@python3 -m webbrowser 'http://localhost:8080/swagger/index.html' || true



.PHONY: api-create-user api-get-user api-update-user api-delete-user

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
