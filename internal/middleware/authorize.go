package middleware

import (
	"log"
	"net/http"
	"portfolio/config"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure it is created witin past 5 hours
		initData := r.Header.Get("AUTH")

		log.Println("Logging the init data sent with the request: ", initData)
		log.Println("Logging the request headers: ", r.Header)

		token := config.Envs.BotToken
		expIn := 5 * time.Hour

		log.Println(initdata.Validate(initData, token, expIn))

		next.ServeHTTP(w, r)
	})
}
