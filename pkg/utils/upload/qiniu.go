package upload

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"io"
	"tiktok/pkg/utils"
	"time"
)

type Qiniu struct {
	LOG    *zap.Logger
	config *utils.Config
}

func (q *Qiniu) UploadFile(f io.Reader, Filename string, size int64) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: q.config.Viper.GetString("Qiniu.Bucket")}
	mac := qbox.NewMac(q.config.Viper.GetString("Qiniu.AccessKey"), q.config.Viper.GetString("Qiniu.SecretKey"))
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig(q.config)
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, size, &putExtra)
	if putErr != nil {
		q.LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return "http://" + q.config.Viper.GetString("Qiniu.ImgPath") + "/" + ret.Key, ret.Key, nil
}

//func (*Qiniu) DeleteFile(key string) error {
//	mac := qbox.NewMac(global.TiK_CONFIG.Qiniu.AccessKey, global.TiK_CONFIG.Qiniu.SecretKey)
//	cfg := qiniuConfig()
//	bucketManager := storage.NewBucketManager(mac, cfg)
//	if err := bucketManager.Delete(global.TiK_CONFIG.Qiniu.Bucket, key); err != nil {
//		global.TiK_LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
//		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
//	}
//	return nil
//}

func qiniuConfig(config *utils.Config) *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      config.Viper.GetBool("Qiniu.UseHTTPS"),
		UseCdnDomains: config.Viper.GetBool("Qiniu.UseCdnDomains"),
	}
	switch config.Viper.GetString("Qiniu.Zone") { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
