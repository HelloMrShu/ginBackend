package apis

import (
	. "financial/models"
	"log"
	"net/http"
	"strconv"

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

//SectorSaveAPI 保存sector
func SectorSaveAPI(c *gin.Context) {
	name := c.PostForm("name")
	intro := c.PostForm("intro")
	s := Sector{Name: name, Intro: intro}

	s.SectorSave()

}

//SectorDeleteAPI 保存sector
func SectorDeleteAPI(c *gin.Context) {
	strID := c.PostForm("id")
	id, _ := strconv.Atoi(strID)
	s := Sector{ID: id}
	log.Println(id)

	s.SectorDelete()

}
