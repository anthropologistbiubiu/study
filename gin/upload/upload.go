package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// gin接收 前端的文件
func UploadFileControl(c *gin.Context) {
	logrus.Infof("[UploadFileControl] user_id =%d", 1)
	// GIN框架获取前端上传文件
	// 这里的参数其实是上传文件时的字段名，也就是上面图片中的file，如果前端是自己定义的其他字段名，需要替换下
	uploadFile, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "获取文件信息失败!" + err.Error(),
		})
	}
	if uploadFile != nil { // 记得及时关闭文件，避免内存泄漏
		defer uploadFile.Close()
	}

	// 读取上传文件的内容
	// 其实这里直接读取所有不太合理,如果文件过大时会占用很多内存，可以考虑使用缓冲区读取
	fileContent, err := ioutil.ReadAll(uploadFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "读取文件内容失败!" + err.Error(),
		})
	}

	// 接受到文件后的处理方式，可以保存到本地，也可以上传到文件对象，看自己了
	/*fileId, err := oss.GetOssClient().UploadFile(userId, fileHeader.Filename, fileContent)
	if err != nil {
		logrus.Errorf("[UploadFile] error,user_id = %d,err_msg= %s", userId, err.Error())
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "上传文件失败!请稍后再试",
		})
	}*/

	// 这里向前端返回下上传成功的信息
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "",
	})
}
