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
	sectors, err := st.SectorList()
	if err != nil {
		log.Fatalln(err)
	}

	total := len(sectors)

	page := c.DefaultQuery("page", "1")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var paginate Paginate
	paginate.Current, _ = strconv.Atoi(page)
	paginate.PageSize = pageSize
	paginate.Pages = int(math.Ceil(float64(total / pageSize)))
	paginate.Total = total

	c.JSON(http.StatusOK, gin.H{
		"data":       sectors,
		"msg":        "success",
		"code":       0,
		"pagination": paginate,
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
	s.SectorDelete()
}
