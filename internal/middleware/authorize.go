package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"portfolio/cmd/web/components"
	"portfolio/config"
	"portfolio/types"
	"time"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure it is created witin past 5 hours
		// initData := r.Header.Get("Auth")
		// use some mock initData
		initData := "user=%7B%22id%22%3A6813536683%2C%22first_name%22%3A%22Still%20makes%20no%20sense%20YoungBoy%20never%20broke%20again%22%2C%22last_name%22%3A%22%22%2C%22username%22%3A%22nbayeabsira%22%2C%22language_code%22%3A%22en%22%2C%22allows_write_to_pm%22%3Atrue%2C%22photo_url%22%3A%22https%3A%5C%2F%5C%2Ft.me%5C%2Fi%5C%2Fuserpic%5C%2F320%5C%2FO02ACpNz22hiMlIgGdZt0D2jZJzwSvWBK59Xi2riDBaGyZjNTYXkwkv1ae4CUP-v.svg%22%7D&chat_instance=-6644236584099621849&chat_type=sender&auth_date=1738572589&signature=I644H25bjMBzskrH-m3rNdtLPB-vtY7QTwkJL68gFmmxySbaFM3nfssXDc7FFeP1WoX-V9SIIiccj59DKiGiDg&hash=e4b54c83d1c8477329f917c342d7baf850607c2a3f6e8f98fbdcc56ec3b61721"
		// log.Println("Logging the init data sent with the request: ", initData)
		// log.Println("Logging the request headers: ", r.Header)

		// find the user by id otherwise return the register form to the user
		token := config.Envs.BotToken
		expIn := 1500 * time.Hour

		if err := initdata.Validate(initData, token, expIn); err != nil {
			log.Println(err)
			e := components.ErrorPage(fmt.Sprint(http.StatusUnauthorized), "UnAuthorized", err.Error())
			e.Render(r.Context(), w)
			return
		}

		dt, err := initdata.Parse(initData)
		if err != nil {
			log.Println(err)
			e := components.ErrorPage(fmt.Sprint(http.StatusInternalServerError), "Internal server error", err.Error())
			e.Render(r.Context(), w)
			return
		}

		log.Println(dt)
		// set the user data
		ctx := context.WithValue(r.Context(), types.UserKey("user"), dt)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
