package apis

import (
	. "financial/models"
	. "financial/utils"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//SectorListAPI 获取sector列表
func SectorListAPI(c *gin.Context) {

	var st Sector
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	sectors, err := st.SectorList(page, pageSize)
	if err != nil {
		log.Fatalln(err)
	}

	total := st.SectorCount()

	var paginate Paginate
	paginate.Current = page
	paginate.PageSize = pageSize
	paginate.Pages = int(math.Ceil(float64(total / pageSize)))
	paginate.Total = total

	c.JSON(http.StatusOK, gin.H{
		"data":       sectors,
		"msg":        "success",
		"code":       200,
		"pagination": paginate,
	})
}

//SectorSaveAPI 保存sector
func SectorSaveAPI(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	intro := c.PostForm("intro")
	s := Sector{ID: id, Name: name, Intro: intro}

	err := s.SectorSave()
	code, msg := APIResponse(err)
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"code": code,
	})
}

//SectorDeleteAPI 保存sector
func SectorDeleteAPI(c *gin.Context) {
	strID := c.PostForm("id")
	id, _ := strconv.Atoi(strID)
	s := Sector{ID: id}
	err := s.SectorDelete()

	code, msg := APIResponse(err)

	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"code": code,
	})
}
