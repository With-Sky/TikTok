package main

import (
	"context"
	"errors"
	"fmt"
	"tiktok/cmd/feed_service/global"
	"tiktok/cmd/feed_service/pack"
	"tiktok/cmd/feed_service/service"
	feed "tiktok/kitex_gen/feed"
	"tiktok/pkg/utils"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedReq) (*feed.FeedRes, error) {
	fmt.Println(req)
	var resp = new(feed.FeedRes)
	videoList, nextTime, err := service.GetFeed(ctx, req.LatestTime, req.Token)
	fmt.Println("视频列表", videoList)
	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("视频流失败")
		resp.BaseResp = pack.BuildBaseResp(err, 400)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(err, 0)
	resp.VideoList = videoList
	nextTimeInt, err := utils.TimeToInt64(nextTime)
	resp.NextTime = nextTimeInt
	return resp, nil
}
