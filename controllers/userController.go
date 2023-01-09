package controllers

import (
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (idb *InDB) LoginUser(c *gin.Context) {
	var user models.User

	username := c.PostForm("username")
	password := c.PostForm("password")

	// hashPwd, err := HashPassword(password)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"message": "Password tidak dapat di hash",
	// 	})
	// 	c.Abort()
	// 	return
	// }

	err := idb.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Username atau password salah",
		})
		c.Abort()
		return
	}

	if !CheckPasswordHash(password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Username atau password salah",
		})
		c.Abort()
		return
	}

	claims := &jwt.MapClaims{
		"username": user.Username,
		"name":     user.Name,
		"role":     user.Role,
	}

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (idb *InDB) CreateUser(c *gin.Context) {
	var user models.User
	idb.DB.Where("username = ? AND role = ?", c.PostForm("username"), c.PostForm("role")).Limit(1).Find(&user)
	if user.Username != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Username sudah pernah dibuat",
		})
		c.Abort()
		return
	}

	user.Username = c.PostForm("username")
	hashPwd, err := HashPassword(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Password tidak dapat di hash",
		})
		c.Abort()
		return
	}
	user.Id = uuid.New().String()
	user.Password = hashPwd
	user.Is_Active = true
	user.Name = c.PostForm("name")
	user.Phone_Number = c.PostForm("phone_number")
	user.Role = c.PostForm("role")
	idb.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User berhasil dibuat",
	})
}

func (idb *InDB) GetUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User tidak ditemukan",
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"user":   user,
		})
	}
}

func (idb *InDB) GetUsers(c *gin.Context) {
	var users []models.User

	idb.DB.Scopes(utils.Paginate(c.Request)).Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"result": nil,
			"count":  0,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": users,
		"count":  len(users),
	})

	c.Abort()
}

func (idb *InDB) DeleteUser(c *gin.Context) {
	var user models.User

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User tidak ditemukan",
		})
		c.Abort()
		return
	}

	err = idb.DB.Delete(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Gagal menghapus data",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User berhasil dihapus",
	})
	c.Abort()
}

func (idb *InDB) UpdateUser(c *gin.Context) {
	var (
		user        models.User
		updatedUser models.User
	)

	id := c.Query("id")
	name := c.PostForm("name")
	phoneNumber := c.PostForm("phone_number")

	err := idb.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data tidak ditemukan",
		})
		c.Abort()
		return
	}

	updatedUser.Name = name
	updatedUser.Phone_Number = phoneNumber

	err = idb.DB.Model(&user).Updates(updatedUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Gagal menghapus data",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User berhasil diupdate",
	})
	c.Abort()
}

func (idb *InDB) Me(c *gin.Context) {
	data := ReadClaims(c.Request.Header.Get("Authorization"))

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": data,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ReadClaims(tokenString string) models.Claims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {

	}

	claims := token.Claims.(jwt.MapClaims)

	user := &models.Claims{
		Username: claims["username"].(string),
		Name:     claims["name"].(string),
		Role:     claims["role"].(string),
	}

	return *user
}
