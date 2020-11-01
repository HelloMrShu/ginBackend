package apis

import (
	. "financial/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SectorListAPI 获取sector列表
func SectorListAPI(c *gin.Context) {

	var st Sector
	sectors, err := st.SectorList()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sectors,
		"msg":  "success",
	})
}
