package dal

import "tiktok/cmd/comment_service/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
