package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func oauthSetup() {
	err := godotenv.Load(".env")
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	gomniauth.SetSecurityKey("セキュリティキー")
	gomniauth.WithProviders(
		// facebook.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/facebook"),
		// github.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/github"),
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/callback/google"),
	)
}
