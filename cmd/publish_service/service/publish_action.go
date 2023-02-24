package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"tiktok/cmd/publish_service/dal/db"
	"tiktok/cmd/publish_service/global"
	"tiktok/pkg/utils"
	"tiktok/pkg/utils/upload"
)

func PublishAction(ctx context.Context, userID int64, data []byte, title string) error {
	fmt.Println(title)
	reader := bytes.NewReader(data)
	// 上传文件
	err, video := UploadFile(title, reader, "0", int64(len(data))) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("上传文件失败")
		return err
	}
	video.UserID = userID
	err = db.CreateVideo(ctx, video)

	if err != nil {
		global.LOG.Error(err.Error())
		err = errors.New("创建视频失败")
		return err
	}
	return nil
}

// UploadFile 上传文件
func UploadFile(title string, VReader io.Reader, noSave string, size int64) (err error, file db.Video) {
	oss := upload.NewOss(global.Config, *global.LOG)
	VideoFilename := uuid.New().String() + ".mp4"
	url, _, err := oss.UploadFile(VReader, VideoFilename, size)
	if err != nil {
		panic(err)
	}
	reader, ImgSize := utils.ExampleReadFrameAsJpeg(url, 1)
	//img, err := imaging.Decode(reader)
	//if err != nil {
	//	return err, file
	//}
	// 保存封面图片
	//pathJPG := global.TiK_CONFIG.Local.Path + "/pathjpg/" + key[:len(key)-4] + ".jpg"
	//fmt.Println
	//imgUrl, _, err := oss.UploadImg(reader)
	imgFilename := uuid.New().String() + ".jpg"
	imgUrl, _, err := oss.UploadFile(reader, imgFilename, ImgSize)
	fmt.Println(imgUrl)
	//if err = imaging.Save(img, pathJPG); err != nil {
	//	return err, file
	//}
	if noSave == "0" {
		//s := strings.Split(header.Filename, ".")
		f := db.Video{
			Title:         title,
			PlayUrl:       url,
			CoverUrl:      imgUrl,
			FavoriteCount: 0,
			CommentCount:  0,
			Name:          VideoFilename,
			//Tag:           s[len(s)-1],
			//Key:           key,
		}
		return err, f
	}
	return
}
