package service

import (
	"context"
	"tiktok/cmd/favorite_service/dal/db"
)

func FavoriteAction(ctx context.Context, userID int64, videoID int64) error {
	favoriteModel := &db.UserFavorite{
		VideoID: videoID,
		UserID:  userID,
	}
	return db.CreateFavorite(ctx, favoriteModel)
}
