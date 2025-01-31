package router

import (
	"ayo/cmd/config"
	"ayo/internal/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserEngine(
    engine *gin.Engine, 
    userRoutes *config.UserRoutes, 
    accSvc account.AccountService,
) {
    engine.POST(userRoutes.SignUp, func(c *gin.Context) {
        spec := account.AuthenticationSpec{}
        if err := c.ShouldBindJSON(&spec); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        user, err := accSvc.Register(c, spec);
        if err != nil {
            c.JSON(err.Code, gin.H{
                "error": err.Message,
            })
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "user": user,
            "message": "User sign up succeed",
        })
    })

    engine.POST(userRoutes.SignIn, func(c *gin.Context) {
        spec := account.AuthenticationSpec{}
        if err := c.ShouldBindJSON(&spec); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        err := accSvc.Login(c, spec);
        if err != nil {
            c.JSON(err.Code, gin.H{
                "error": err.Message,
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "User sign in succeed",
        })
    })

}