package components

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func LogToFile()  {
	logfile, err := os.OpenFile("./runtime/app-info.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("Could not create log file")
		fmt.Println(err)
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
}
