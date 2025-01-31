package router

import (
	"ayo/cmd/config"
	"ayo/internal/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAccountEngine(
    engine *gin.Engine, 
    accountRoutes *config.AccountRoutes, 
    accSvc account.AccountService,
) {
    engine.POST(accountRoutes.SignUp, func(c *gin.Context) {
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

    engine.POST(accountRoutes.SignIn, func(c *gin.Context) {
        spec := account.AuthenticationSpec{}
        if err := c.ShouldBindJSON(&spec); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        token, err := accSvc.Login(c, spec);
        if err != nil {
            c.JSON(err.Code, gin.H{
                "error": err.Message,
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "token": token,
            "message": "User sign in succeed",
        })
    })

}