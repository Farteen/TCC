package travel

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TravelRouterRegister(r *gin.RouterGroup) {
	r.POST("/travelcost", TravelCostCreate)
}

func TravelCostCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "create travel success"})
}
