package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-jwt-demo/pkg/jwt"
	"go-jwt-demo/user/repository"
)

type Service interface {
	Login(c *gin.Context)
	GetUserInfo(c *gin.Context)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Login(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}

	user, err := s.repo.FindUserByUsername(req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", req.Name),
		})
		return
	}

	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	token, err := jwt.GenerateToken(*user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (s *service) GetUserInfo(c *gin.Context) {
	id, _, ok := getSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	user, err := s.repo.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func getSession(c *gin.Context) (uint, string, bool) {
	id, ok := c.Get("id")
	if !ok {
		return 0, "", false
	}

	username, ok := c.Get("name")
	if !ok {
		return 0, "", false
	}

	return id.(uint), username.(string), true
}
