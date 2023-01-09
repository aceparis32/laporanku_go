package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SARoleValidation(c *gin.Context) {
	currentUser := ReadClaims(c.Request.Header.Get("Authorization"))
	if strings.ToLower(currentUser.Role) != "sa" {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "Forbidden",
		})
		c.Abort()
		return
	}
}

func EmployeeRoleValidation(c *gin.Context) {
	currentUser := ReadClaims(c.Request.Header.Get("Authorization"))
	if strings.ToLower(currentUser.Role) != "employee" {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": "Forbidden",
		})
		c.Abort()
		return
	}
}
