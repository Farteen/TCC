package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Farteen/TCC/user"
	"github.com/Farteen/TCC/travel"
)

func main() {
	r := gin.Default()
	user.UserRouterRegister(r.Group("/user"))
	travel.TravelRouterRegister(r.Group("/travel"))
	r.Run(":3000")
}