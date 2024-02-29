package service

import (
	"fmt"
	"ginchat/utils"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	writer := c.Writer
	request := c.Request
	// file 是文件的内容，head 是文件的元信息
	file, head, err := request.FormFile("file")
	if err != nil {
		utils.RespFail(writer, err.Error())
	}
	// 默认文件后缀为 ".png"。
	suffix := ".png"
	fileName := head.Filename
	// 获取文件后缀
	tem := strings.Split(fileName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	// 包含当前时间戳和随机数，以确保唯一
	newfileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	// 新文件保存的路径
	dstFile, err := os.Create("./frontend/src/assets/upload/" + newfileName)
	if err != nil {
		utils.RespFail(writer, err.Error())
	}
	// 将文件内容拷贝到新文件中
	_, err = io.Copy(dstFile, file)
	if err != nil {
		utils.RespFail(writer, err.Error())
	}
	// 上传成功后的文件访问 URL
	url := "/src/assets/upload/" + newfileName
	utils.RespOK(writer, "upload success", url)
}
