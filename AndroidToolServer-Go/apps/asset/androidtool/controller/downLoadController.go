package controller

import "github.com/gin-gonic/gin"

var DownloadController *download

type download struct {
}

func init() {
	DownloadController = &download{}
}
func (d *download) DownApk(ctx *gin.Context) {
	ctx.Header("Content-Disposition", "attachment; filename="+"t.txt") // 用来指定下载下来的文件名
	ctx.File("D:\\Go\\test.txt")
	/*	ctx.Header("Content-Type", "application/octet-stream")
		disposition := ""
		ctx.Header("Content-Disposition",)
		ctx.Header("Content-Transfer-Encoding", "binary")*/
}
