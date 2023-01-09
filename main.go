package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-go-dasar/configs"
	"project-go-dasar/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := configs.DBInit(os.Getenv("DSN"))
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// router.POST("/login", inDB.LoginUser)
	// router.GET("/person/:id", auth, inDB.GetPerson)
	// router.GET("/persons", auth, inDB.GetPersons)
	// router.POST("/person", auth, inDB.CreatePerson)
	// router.PUT("/person", auth, inDB.UpdatePerson)
	// router.DELETE("/person/:id", auth, inDB.DeletePerson)

	// Region User
	router.POST("/login", inDB.LoginUser)
	router.GET("/user/:id", auth, controllers.SARoleValidation, inDB.GetUser)
	router.GET("/users", auth, inDB.GetUsers)
	router.POST("/user", inDB.CreateUser)
	router.PUT("/user", auth, inDB.UpdateUser)
	router.DELETE("/user/:id", auth, inDB.DeleteUser)
	router.GET("/me", auth, inDB.Me)

	// Region Company
	router.GET("/company/:id", auth, inDB.GetCompany)
	router.GET("/companies", auth, inDB.GetCompanies)
	router.POST("/company", auth, inDB.CreateCompany)
	router.PUT("/company", auth, inDB.UpdateCompany)
	router.DELETE("/company/:id", auth, inDB.DeleteCompany)

	// Region Sales
	router.GET("/sales/:id", auth, inDB.GetSales)
	router.GET("/sales", auth, inDB.GetSalesList)
	router.POST("/sales", auth, inDB.CreateSales)
	router.PUT("/sales", auth, inDB.UpdateSales)
	router.DELETE("/sales/:id", auth, inDB.DeleteSales)

	// Region Item
	router.GET("/item/:id", auth, inDB.GetItem)
	router.GET("/items", auth, inDB.GetItems)
	router.POST("/item", auth, inDB.CreateItem)
	router.PUT("/item", auth, inDB.UpdateItem)
	router.DELETE("/item/:id", auth, inDB.DeleteItem)

	// Region Note
	router.GET("/note/:id", auth, inDB.GetNote)
	router.GET("/notes", auth, inDB.GetNotes)
	router.POST("/note", auth, inDB.CreateNote)
	router.PUT("/note", auth, inDB.UpdateNote)
	router.DELETE("/note/:id", auth, inDB.DeleteNote)

	// fmt.Println("Starting server at ", )
	router.Run(os.Getenv("PORT"))
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "Not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
