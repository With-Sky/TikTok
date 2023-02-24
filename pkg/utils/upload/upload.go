package upload

import (
	"go.uber.org/zap"
	"io"
	"tiktok/pkg/utils"
)

// OSS 对象存储接口
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
type OSS interface {
	UploadFile(file io.Reader, Filename string, size int64) (string, string, error)
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss(config utils.Config, logger zap.Logger) OSS {
	switch config.Viper.GetString("System.OssType") {
	//case "local":
	//	return &Local{}
	case "qiniu":
		return &Qiniu{
			LOG:    &logger,
			config: &config,
		}
	//case "tencent-cos":
	//	return &TencentCOS{}
	//case "aliyun-oss":
	//	return &AliyunOSS{}
	//case "huawei-obs":
	//	return HuaWeiObs
	//case "aws-s3":
	//	return &AwsS3{}
	//case "minio":
	//	return &Minio{
	//		config: &config,
	//	}
	case "ftp":
		return &FtpUtils{
			LOG:    &logger,
			config: &config,
		}
	default:
		return &Qiniu{
			LOG:    &logger,
			config: &config,
		}
		//return &Local{}
	}
}
