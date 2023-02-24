package upload

//
//import (
//	"context"
//	"github.com/minio/minio-go/v7/pkg/credentials"
//	"io"
//	"net/url"
//	"strings"
//	"tiktok/pkg/utils"
//	"time"
//
//	"github.com/cloudwego/kitex/pkg/klog"
//	"github.com/minio/minio-go/v7"
//)
//
//type Minio struct {
//	config *utils.Config
//}
//
//var (
//	minioClient *minio.Client
//)
//
////	Config               = ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
////	MinioEndpoint        = .Viper.GetString("minio.Endpoint")
////	MinioAccessKeyId     = Config.Viper.GetString("minio.AccessKeyId")
////	MinioSecretAccessKey = Config.Viper.GetString("minio.SecretAccessKey")
////	MinioUseSSL          = Config.Viper.GetBool("minio.UseSSL")
////	MinioVideoBucketName = Config.Viper.GetString("minio.VideoBucketName")
////)
//
//// Minio 对象存储初始化
//func (m *Minio) init() {
//	//var minioClient          *minio.Client
//	//Config               := ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
//	MinioEndpoint := m.config.Viper.GetString("minio.Endpoint")
//	MinioAccessKeyId := m.config.Viper.GetString("minio.AccessKeyId")
//	MinioSecretAccessKey := m.config.Viper.GetString("minio.SecretAccessKey")
//	MinioUseSSL := m.config.Viper.GetBool("minio.UseSSL")
//	MinioVideoBucketName := m.config.Viper.GetString("minio.VideoBucketName")
//	client, err := minio.New(MinioEndpoint, &minio.Options{
//		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
//		Secure: MinioUseSSL,
//	})
//	if err != nil {
//		klog.Errorf("minio client init failed: %v", err)
//	}
//	// fmt.Println(client)
//	klog.Debug("minio client init successfully")
//	minioClient = client
//	if err := m.CreateBucket(MinioVideoBucketName); err != nil {
//		klog.Errorf("minio client init failed: %v", err)
//	}
//}
//
//// CreateBucket 创建桶
//func (m *Minio) CreateBucket(bucketName string) error {
//	if len(bucketName) <= 0 {
//		klog.Error("bucketName invalid")
//	}
//
//	location := "beijing"
//	ctx := context.Background()
//
//	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
//	if err != nil {
//		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
//		if errBucketExists == nil && exists {
//			klog.Debugf("bucket %s already exists", bucketName)
//			return nil
//		} else {
//			return err
//		}
//	} else {
//		klog.Infof("bucket %s create successfully", bucketName)
//	}
//	return nil
//}
//
//// UploadLocalFile 上传本地文件（提供文件路径）至 minio
//func UploadLocalFile(bucketName string, objectName string, filePath string, contentType string) (int64, error) {
//	ctx := context.Background()
//	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
//		ContentType: contentType,
//	})
//	if err != nil {
//		klog.Errorf("localfile upload failed, %s", err)
//		return 0, err
//	}
//	klog.Infof("upload %s of size %d successfully", objectName, info.Size)
//	return info.Size, nil
//}
//
//// UploadFile 上传文件（提供reader）至 minio
//func (m *Minio) UploadFile(reader io.Reader, Filename string, objectsize int64) (string, string, error) {
//	m.init()
//	ctx := context.Background()
//	n, err := minioClient.PutObject(ctx, m.config.Viper.GetString("minio.VideoBucketName"), Filename, reader, objectsize, minio.PutObjectOptions{
//		ContentType: "application/octet-stream",
//	})
//	if err != nil {
//		klog.Errorf("upload %s of size %d failed, %s", m.config.Viper.GetString("minio.VideoBucketName"), objectsize, err)
//		return "", "", err
//	}
//	klog.Infof("upload %s of bytes %d successfully", Filename, n.Size)
//	url, err := m.GetFileUrl(m.config.Viper.GetString("minio.VideoBucketName"), Filename, 0)
//	playUrl := strings.Split(url.String(), "?")[0]
//	return playUrl, "", nil
//}
//
//// GetFileUrl 从 minio 获取文件Url
//func (m *Minio) GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
//	ctx := context.Background()
//	reqParams := make(url.Values)
//	if expires <= 0 {
//		expires = time.Second * 60 * 60 * 24
//	}
//	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
//	if err != nil {
//		klog.Errorf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
//		return nil, err
//	}
//	// TODO: url可能要做截取
//	return presignedUrl, nil
//}
