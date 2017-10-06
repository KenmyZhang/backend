package controllers

import (

)

type FileController struct {
	MainController	
}

// @Title login
// @Description login by username and password
// @Success 200 {object} models.ZDTProduct.ProductList
// @Param   category_id     query   int false       "category id"
// @Param   brand_id    query   int false       "brand id"
// @Param   query   query   string  false       "query of search"
// @Param   segment query   string  false       "segment"
// @Param   sort    query   string  false       "sort option"
// @Param   dir     query   string  false       "direction asc or desc"
// @Param   offset  query   int     false       "offset"
// @Param   limit   query   int     false       "count limit"
// @Param   price           query   float       false       "price"
// @Param   special_price   query   bool        false       "whether this is special price"
// @Param   size            query   string      false       "size filter"
// @Param   color           query   string      false       "color filter"
// @Param   format          query   bool        false       "choose return format"
// @Failure 400 no enough input
// @Failure 500 get products common error
// @router /upload [post]
func (c *FileController) UploadFile() {
    f, h, err := c.GetFile("uploadname")
    if err != nil {
    	c.Ctx.WriteString("failed")
    	return
    }
    defer f.Close()
    c.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
    c.Ctx.WriteString("success")
}
