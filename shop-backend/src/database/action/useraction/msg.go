package useraction

import (
	"github.com/SongZihuan/cat-shop/shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop/shop-backend/src/model"
)

func SendMsgByUser(msg string, user *model.User) error {
	return internal.DB().Create(model.NewMsg(user.ID, msg)).Error
}
