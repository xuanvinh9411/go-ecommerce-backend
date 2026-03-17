package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuanvinh9411/go-ecommerce-backend/internal/service"
	"github.com/xuanvinh9411/go-ecommerce-backend/pkg/response"
)


type UserController struct {
 userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func(uc *UserController) GetUser(c *gin.Context)  {
	  c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"user": "xuanvinh9411",
	})
}

//query user by query param
func(uc *UserController) GetUserQuery(c *gin.Context)  {
	uid := c.Query("uid")
	old := c.Query("old")
	allParam := c.Request.URL.Query()
	fmt.Printf("allParam: ",  allParam)
	if uid == "123" {
		response.SuccessResponse(c, response.ErrCodeSuccess, gin.H{
			"msg": "userid of user: " + uid + ", old: " + old + " is valid",
		})
	} else {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, "Create user failed")
	}

}

func(uc *UserController) CreateUser(c *gin.Context)  {
	uid := c.Param("uid")
	if uid != "123" {
		response.SuccessResponse(c, response.ErrCodeSuccess, gin.H{
			"uid": uid,
			"name": "xuanvinh9411",
		})
	} else {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, "Create user failed")
	}

}