package file_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "log"
	"io"
	"net/http"
	"os"
)

func UploadSingleFile(c *gin.Context) {
	file, err := c.FormFile("index.tmpl")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully\n", file.Filename))
}

func UploadTest(c *gin.Context) {
	fmt.Println("---------------- request method", c.Request.Method, "----------------")
	r := c.Request
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("FormFile Error:", err)
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println("Open File Error:", err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.JSON(200, handler.Header)
	} else {
		c.JSON(200, "ok")
	}
}
