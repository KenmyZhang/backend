package controllers

import (
	"net/http"
	"github.com/KenmyZhang/mini-chat/models"
)

type FileController struct {
	MainController	
}

// @Title upload
// @Description upload file
// @Success 200 upload success 
// @Param   formData   string  true   "the path of upload file"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /upload [post]
func (c *FileController) UploadFile() {
    f, h, err := c.GetFile("uploadname")
    if err != nil {
		models.SetInvalidParam(c.Ctx, "uploadname", http.StatusBadRequest)	
    	return
    }
    defer f.Close()
    c.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
    c.Ctx.WriteString("success")
}
