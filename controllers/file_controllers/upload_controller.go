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
		err := DownlinkFile(r, "uploadfile")
		err2 := DownlinkFile(r, "uploadfile2")
		c.JSON(200, gin.H{
			"Err":  err,
			"Err2": err2,
		})
	} else {
		c.JSON(200, "ok")
	}
}

func DownlinkFile(r *http.Request, fileUrl string) error {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(fileUrl)
	if err != nil {
		fmt.Println("FormFile Error:", err)
		return err
	}
	defer file.Close()
	// fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
	if err != nil {
		fmt.Println("Open File Error:", err)
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}
