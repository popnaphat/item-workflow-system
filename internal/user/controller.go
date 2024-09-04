package user

import (
	"fmt"
	"net/http"
	"task-api/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB, secret string) Controller {
	return Controller{
		Service: NewService(db, secret),
	}
}

func (controller Controller) Login(ctx *gin.Context) {
	var request model.RequestLogin

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID, token, err := controller.Service.Login(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Set the userID in a separate cookie
	ctx.SetCookie(
		"userID",
		fmt.Sprintf("%v", userID),
		int(60*time.Second), // set a realistic expiration time
		"/",
		"localhost", // replace with your domain
		false,       // secure: set to true in production (for HTTPS)
		true,        // http-only
	)
	// Set token cookie
	ctx.SetCookie(
		"token",
		fmt.Sprintf("Bearer %v", token),
		int(60*time.Second),
		"/",
		"localhost",
		false,
		true, // http-only cookie flag
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login succeed",
	})
}
