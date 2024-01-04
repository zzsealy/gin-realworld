package users

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct{
	Code		int		`json:"code" example:"200"`
	Message     string  `json:"message" example:"请求成功"`
}

type FailResponse struct {
	Code		int		`json:"code" example:"400"`
	Message     string  `json:"message" example:"请求失败"`
}


func UserRegister(router *gin.RouterGroup) {
	router.POST("/register", UsersRegisterView)
}

// PingExample godoc
// @Summary 用户注册
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Param username query string true "用户名，必须唯一"
// @Param password query string true "用户密码"
// @Param email query string true "用户邮箱"
// @Success 200 {object} SuccessResponse "成功时返回的数据结构"
// @Failure 400 {object} FailResponse "请求参数错误"
// @Router /api/users/register [get]
func UsersRegisterView(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	fmt.Println(userModelValidator)

	successResponse := SuccessResponse{Code:200, Message: "注册成功"}
	c.JSON(http.StatusOK, successResponse)
	
}