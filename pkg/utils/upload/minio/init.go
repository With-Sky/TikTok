package minio

//var (
//	minioClient          *minio.Client
//	Config               = ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
//	MinioEndpoint        = Config.Viper.GetString("minio.Endpoint")
//	MinioAccessKeyId     = Config.Viper.GetString("minio.AccessKeyId")
//	MinioSecretAccessKey = Config.Viper.GetString("minio.SecretAccessKey")
//	MinioUseSSL          = Config.Viper.GetBool("minio.UseSSL")
//	MinioVideoBucketName = Config.Viper.GetString("minio.VideoBucketName")
//)
//
//// Minio 对象存储初始化
//func init() {
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
//	if err := CreateBucket(MinioVideoBucketName); err != nil {
//		klog.Errorf("minio client init failed: %v", err)
//	}
//}
