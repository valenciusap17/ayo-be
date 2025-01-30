package router

import (
	"ayo/cmd/config"
	"ayo/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserEngine(
    engine *gin.Engine, 
    routes *config.VendorRoutes, 
    userSvc user.UserService,
) {
    engine.GET(routes.SignIn, func(ctx *gin.Context) {
        ctx.JSON(http.StatusCreated, "mantep")
    })

}