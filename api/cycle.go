package api

import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
)

// 获取轮播数据接口
func GainCycleData(c *gin.Context) {
	balance := [2]string{"erwererwrewerwrewerwe", "dhsfjkhjfdsjhkfhjkdshjkdfs"}
	c.JSON(200, serializer.Response{
		Code: 0,
		Data:  balance,
		Msg: "djksdjks",
	})
}