package utils

import (
	"context"
	"portfolio/types"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func GetUserFromContext(ctx context.Context) *initdata.InitData {
	userData := ctx.Value(types.UserKey("user"))
	if uInitData, ok := userData.(initdata.InitData); ok {
		return &uInitData
	}

	return nil
}
