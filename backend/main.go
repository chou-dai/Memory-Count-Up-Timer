package main

import (
	"gin_backend/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default() // ここからCorsの設定
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいオリジン
		AllowOrigins: []string{
			"http://localhost", // 今回はフロントエンドアプリケーションのみを指定
		},
		// 許可したいHTTPヘッダー
		AllowHeaders: []string{
			"Cookie",     // 過去にSet-Cookieヘッダーでブラウザに保存したクッキーをクライアント側からサーバ側へ送信することを許可するヘッダー
			"Set-Cookie", // サーバ側からクライアント側にクッキーを送信することを許可するヘッダー
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding", // クライアント側がサポートしている圧縮（エンコーディング）方式をサーバ側に伝えるためのヘッダー
			"Authorization",
			"Access-Control-Allow-Credentials", // Cookie、認証ヘッダー、または TLS クライアント証明書といった資格情報をクライアント側に公開することを許可するヘッダー
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Methods",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// routerをcontrollerに渡してルーティング
	controller.Router(router)

	// ポートの設定
	router.Run(":8080")
}
