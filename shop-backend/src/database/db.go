package database

import "github.com/SongZihuan/cat-shop/shop-backend/src/database/action"

func ConnectToMySQL() error {
	return action.ConnectToMySQL()
}

func CloseMySQL() {
	action.CloseMySQL()
}

func IsReady() bool {
	return action.IsReady()
}
