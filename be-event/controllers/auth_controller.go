package controllers

import (
	"be-event/models"
	"be-event/request"
	"be-event/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req request.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Convert request sang model.User
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := c.service.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Đăng ký thành công",
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	token, err := c.service.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Đăng nhập thành công",
		"token":   token,
	})
}

func (c *AuthController) Logout(ctx *gin.Context) {
	// Lấy token từ header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token không được cung cấp"})
		return
	}

	// Tách "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Định dạng token không hợp lệ"})
		return
	}

	token := parts[1]

	// Gọi service logout
	if err := c.service.Logout(token); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Logout thất bại"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout thành công"})
}
