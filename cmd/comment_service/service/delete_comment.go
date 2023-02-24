package service

import (
	"context"
	"gorm.io/gorm"
	"tiktok/cmd/comment_service/dal/db"
)

func DeleteComment(ctx context.Context, commentID int64) error {
	commentModel := &db.Comment{
		Model: gorm.Model{
			ID: uint(commentID),
		},
	}
	return db.DeleteComment(ctx, commentModel)
}
