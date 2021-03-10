package api


import (
	"github.com/gin-gonic/gin"
	"singo/serializer"
)

// 获取轮播数据接口
func GainHomeMiddle(c *gin.Context) {
	balance := [7]string{"erwererwrewerwrewerwe", "dhsfjkhjfdsjhkfhjkdshjkdfs","dhsfjkhjfbvdsjhkfhjkdshjkdfs"}
	c.JSON(200, serializer.Response{
		Code: 0,
		Data:  balance,
		Msg: "主页中间数据",
	})
}