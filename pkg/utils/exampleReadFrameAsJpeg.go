package utils

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
)

/*
#cgo LDFLAGS: -llibavformat  -llibavcodec -llibavutil -llibavdevice -llibavfilter -llibswresample -llibswscale
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavutil/opt.h>
#include <libavdevice/avdevice.h>
static const AVStream *go_av_streams_get(const AVStream **streams,unsigned int n)
{
    return streams[n];
}
*/

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) (io.Reader, int64) {
	buf := bytes.NewBuffer(nil)
	fmt.Println(inFileName)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//buf.Len()
	return buf, int64(buf.Len())
}
