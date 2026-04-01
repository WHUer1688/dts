package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"travel/internal/config"

	"github.com/gin-gonic/gin"
)

func HttpHello(c *gin.Context) {
	// Function name should be Capital uppercase
	// for other packages to call
	c.JSON(http.StatusOK, gin.H{
		"message": "hello http",
	})
}

// 处理`POST photos`请求
// TODO：未更新数据库，仅保存了文件
func HttpPostPhotos(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "接受文件失败: " + err.Error(),
		})
		return
	}
	
	saveDir := config.GlobalConfig.SFHD
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "服务器创建目录失败",
		})
		return
	}
	
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	savePath := filepath.Join(saveDir, newFileName)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "保存文件到服务器失败",
		})
		return
	}

	// 读取配置
	host := config.GlobalConfig.Host
	port := config.GlobalConfig.Port

	fileUrl := fmt.Sprintf("%s:%d/photos/%s", host, port, newFileName)
	c.JSON(http.StatusOK, gin.H{
		"remote_url": fileUrl,
	})
}
