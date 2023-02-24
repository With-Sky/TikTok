package upload

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gopkg.in/dutchcoders/goftp.v1"
	"io"
	"tiktok/pkg/utils"
)

type FtpUtils struct {
	LOG    *zap.Logger
	config *utils.Config
}

func (f *FtpUtils) DeleteFile(key string) error {
	//global.TiK_FTP.Dele()
	panic("implement me")
}

func (f *FtpUtils) UploadFile(file io.Reader, Filename string, size int64) (string, string, error) {
	var err error
	var ftp *goftp.FTP

	if ftp, err = goftp.Connect(f.config.Viper.GetString("Ftp.Address") + ":" + f.config.Viper.GetString("Ftp.Port")); err != nil {
		return "", "", err
		f.LOG.Error("连接失败" + err.Error())
	}

	if err = ftp.Login(f.config.Viper.GetString("Ftp.Account"), f.config.Viper.GetString("Ftp.Password")); err != nil {
		return "", "", err
		f.LOG.Error("登录失败" + err.Error())
		defer ftp.Close()
	}
	f.LOG.Info("登录成功")
	var curpath string
	if curpath, err = ftp.Pwd(); err != nil {
		panic(err)
	}
	//
	fmt.Printf("Current path: %s", curpath)
	//Ofile, err := file.Open()
	//if err != nil {
	//	return "", "", err
	//	//global.TiK_LOG.Error("打开错误" + err.Error())
	//	defer ftp.Close()
	//}
	path := uuid.New().String()
	if err = ftp.Mkd(path); err != nil {
		return "", "", err
		f.LOG.Error("创建文件夹失败")
		defer ftp.Close()
	}
	//Filename:=uuid.New().String()+".mp4"
	if err = ftp.Stor(path+"/"+Filename, file); err != nil {
		return "", "", err
		f.LOG.Error("上传错误" + err.Error())
		fmt.Println(err.Error())
		defer ftp.Close()
	}
	f.LOG.Info("上传成功，路径：" + path + "/" + Filename)
	defer ftp.Close()
	return "http://" + f.config.Viper.GetString("Ftp.Address") + ":" + f.config.Viper.GetString("Ftp.GetPort") + "/" + path + "/" + Filename, Filename, err
}
func (f *FtpUtils) UploadImg(file io.Reader) (string, string, error) {
	var err error
	var ftp *goftp.FTP

	if ftp, err = goftp.Connect(f.config.Viper.GetString("Ftp.Address") + ":" + f.config.Viper.GetString("Ftp.Port")); err != nil {
		return "", "", err
		f.LOG.Error("连接失败" + err.Error())
	}
	f.LOG.Info("连接成功")
	if err = ftp.Login(f.config.Viper.GetString("Ftp.Account"), f.config.Viper.GetString("Ftp.Password")); err != nil {
		return "", "", err
		f.LOG.Error("登录失败" + err.Error())
		defer ftp.Close()
	}
	f.LOG.Info("登录成功")
	var curpath string
	if curpath, err = ftp.Pwd(); err != nil {
		panic(err)
	}
	//
	fmt.Printf("Current path: %s", curpath)
	if err != nil {
		return "", "", err
		f.LOG.Error("打开错误" + err.Error())
		defer ftp.Close()
	}
	path := uuid.New().String()
	if err = ftp.Mkd(path); err != nil {
		return "", "", err
		f.LOG.Error("创建文件夹失败")
		defer ftp.Close()
	}
	imgName := uuid.New().String() + ".jpg"
	if err = ftp.Stor(path+"/"+imgName, file); err != nil {
		return "", "", err
		f.LOG.Error("上传错误" + err.Error())
		fmt.Println(err.Error())
		defer ftp.Close()
	}
	f.LOG.Info("上传成功，路径：" + path + "/" + imgName)
	defer ftp.Close()
	return "http://" + f.config.Viper.GetString("Ftp.Address") + ":" + f.config.Viper.GetString("Ftp.GetPort") + "/" + path + "/" + imgName, "", err
}
