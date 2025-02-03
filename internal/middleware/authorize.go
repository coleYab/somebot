package middleware

import (
	"fmt"
	"net/http"
	"portfolio/config"
	"portfolio/types"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func AuthorizeUser(next func(w http.ResponseWriter, r *http.Request), st *types.UserStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// make sure it is created witin past 5 hours
		data := "sfjsdkfhjsdkhfjkshfjkdhfjh"
		token := config.Envs.BotToken
		expIn := 5 * time.Hour

		fmt.Println(initdata.Validate(data, token, expIn))

		next(w, r)
	}
}
