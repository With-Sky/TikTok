package dal

import "tiktok/cmd/message_service/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
