package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"singo/model"
	"singo/serializer"
)

// 获取轮播数据接口
func GainHomeItemList(c *gin.Context) {

	c.JSON(200, serializer.Response{
		Code: 0,
		Data:  SelectAllItem(),
		Msg: "",
	})
}

type Project struct {
	gorm.Model
	Id         int           `gorm:"column:id"`
	Logo        string        `gorm:"column:logo"`
	Item_name   string        `gorm:"column:item_name"`
	Intro       string        `gorm:"column:intro"`
	Price   float64 		`gorm:"column:price"`
	Sn    string           `gorm:"column:sn"`
	City_id   int           `gorm:"column:city_id"`
}


func SelectAllItem() ([]Project) {
	//执行查询语句
	//rows, err := model.DB.Raw("SELECT * from home_item").Row()
	rows, err := model.DB.Raw("SELECT * from home_item").Rows()
	if err != nil{
		fmt.Println("查询出错了")
	}
	var projects []Project
	//循环读取结果
	for rows.Next(){
		var project Project
		//将每一行的结果都赋值到一个project对象中
		err := rows.Scan(&project.Id, &project.Logo, &project.Item_name,&project.Intro, &project.Price, &project.Sn,&project.City_id)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		projects = append(projects, project)
	}
	return projects
}
