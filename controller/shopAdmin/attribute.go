package shopAdmin

import (
	"login_register_demo/config"
	"login_register_demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAttributeCategory(c *gin.Context) {
	shopId, ok := c.GetQuery("shop_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code":    http.StatusBadRequest,
			"msg ":    "请求参数错误",
		})
	}

	res := make([]model.GoodsAttributeCategory, 0)
	err := config.Engine.Table("goods_attribute_category").Where("shop_id = ?", shopId).Find(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"code":    http.StatusInternalServerError,
			"msg ":    "内部查询错误",
		})
	}

	c.JSON(http.StatusOK, res)
}