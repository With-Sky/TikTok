package dal

import "tiktok/cmd/favorite_service/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
