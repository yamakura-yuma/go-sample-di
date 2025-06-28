package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 認証ロジックをここに実装
		// 例: トークンの検証、ユーザーの認証など

		// 認証が成功した場合、次のハンドラーを呼び出す
		next.ServeHTTP(w, r)
	})
}
