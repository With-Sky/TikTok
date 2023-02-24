package dal

import "tiktok/cmd/publish_service/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
