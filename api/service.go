package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"singo/model"
	"singo/serializer"
	"strconv"

	//"singo/service"
)

// 获取轮播数据接口
func GainServiceItems(c *gin.Context) {

	//var service = CityIdService
		c.JSON(200, serializer.Response{
			Code: 0,
			Data: _SelectAllItem(c),
			Msg: "",
		})
}

type CityIdService struct {
	city_id int `form:"city_id" json:"city_id" binding:"required,min=1,max=30"`
}

type ServiceProject struct {
	gorm.Model
	Id         int           `gorm:"column:id"`
	Logo        string        `gorm:"column:logo"`
	Item_name   string        `gorm:"column:item_name"`
	Intro       string        `gorm:"column:intro"`
	Price   float64 		`gorm:"column:price"`
	Sn    string           `gorm:"column:sn"`
	City_id   int           `gorm:"column:city_id"`
}


func _SelectAllItem(c *gin.Context) ([]Project) {
	//执行查询语句
	//rows, err := model.DB.Raw("SELECT * from home_item").Row()
	fmt.Println(c.Request.Body)
	fmt.Println("fgfgdgdfdggfd")
	rows, err := model.DB.Raw("SELECT * from home_item").Rows()
	if err != nil{
		fmt.Println("查询出错了")
	}
	var projects []Project
	//循环读取结果
	for rows.Next(){
		var project Project
		//将每一行的结果都赋值到一个project对象中
		err := rows.Scan(&project.Id, &project.Logo, &project.Item_name,&project.Intro, &project.Price, &project.Sn, &project.City_id)
		if err != nil {
			fmt.Println("rows fail")
		}
		m, err := strconv.Atoi(c.PostForm("city_id"))
		if m==project.City_id {
			projects = append(projects, project)
		}
	}
	return projects
}
