package controllers

import (
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetCustomer(c *gin.Context) {
	var customer models.Customer

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"result": err.Error(),
			"count":  0,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": customer,
		"count":  1,
	})

	c.Abort()
}

func (idb *InDB) GetCustomers(c *gin.Context) {
	var customers []models.Customer

	idb.DB.Scopes(utils.Paginate(c.Request)).Find(&customers)

	if len(customers) <= 0 {
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
		"result": customers,
		"count":  len(customers),
	})
	c.Abort()
}

func (idb *InDB) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	customer.Customer_Name = c.PostForm("customer_name")
	customer.Phone_Number = c.PostForm("phone_number")
	customer.Email = c.PostForm("email")

	idb.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Customer berhasil dibuat",
	})
	c.Abort()
}

func (idb *InDB) UpdateCustomer(c *gin.Context) {
	id := c.Query("id")
	var (
		customer        models.Customer
		updatedCustomer models.Customer
	)

	err := idb.DB.First(&customer, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		c.Abort()
		return
	}

	updatedCustomer.Customer_Name = customer.Customer_Name
	updatedCustomer.Phone_Number = customer.Phone_Number
	updatedCustomer.Email = customer.Email

	err = idb.DB.Model(&customer).Updates(updatedCustomer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data customer berhasil di update",
	})
	c.Abort()
}

func (idb *InDB) DeleteCustomer(c *gin.Context) {
	var customer models.Customer

	id := c.Param("id")
	err := idb.DB.First(&customer, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		c.Abort()
		return
	}

	err = idb.DB.Delete(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data customer berhasil dihapus",
	})
	c.Abort()
}
