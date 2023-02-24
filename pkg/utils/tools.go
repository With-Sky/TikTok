package utils

import (
	"io"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"time"
)

// NowDate 获取当前日期
func NowDate() string {
	return time.Unix(time.Now().Unix(), 0).Format("01-02")
}

// NowTime 获取当前时间
func NowTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

// NowUnix 获取当前时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// TimeToFormatData Time获取时间戳日期
func TimeToFormatData(timeStamp time.Time) string {
	return timeStamp.Format("01-02")
}
func TimeToInt64(timeStamp time.Time) (int64, error) {

	int64, err := strconv.ParseInt(timeStamp.Format("01-02"), 10, 64)
	if err != nil {
		return 0, err
	}
	return int64, nil
}

// UnixToFormatData 获取时间戳日期
func UnixToFormatData(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format("01-02")
}

// UnixToFormatTime 获取时间戳时间
func UnixToFormatTime(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
}

// VerifyMobileFormat 校验手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// BoolIntoByte bool转[]bytez
func BoolIntoByte(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// ByteIntoBool byte转bool
func ByteIntoBool(b []byte) bool {
	if b[0] == 1 {
		return true
	} else {
		return false
	}
}

// ByteIntoInt byte转int
func ByteIntoInt(b []byte) uint8 {
	if b[0] == 1 {
		return 1
	} else {
		return 0
	}
}

// BoolIntoInt bool转int
func BoolIntoInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

// StrBoolIntoByte 用于转化前端字符串布尔值为[]byte
func StrBoolIntoByte(s string) []byte {
	if s == "true" {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// StrGenderIntoByte 用于转化前端字符串布尔值为[]byte
func StrGenderIntoByte(s string) []byte {
	if s == "男" {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// ByteEnabledToString 启用未启用状态布尔转字符串
func ByteEnabledToString(b []byte) string {
	if b[0] == 0 {
		return `禁用`
	} else {
		return `启用`
	}
}

// SaveFile 保存文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	err := os.MkdirAll("./static/uploadfile/", os.ModePerm)
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
