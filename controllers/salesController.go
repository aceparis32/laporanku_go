package controllers

import (
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (idb *InDB) GetSales(c *gin.Context) {
	var sales models.Sales

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&sales).Error
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
		"result": sales,
		"count":  1,
	})

	c.Abort()
}

func (idb *InDB) GetSalesList(c *gin.Context) {
	var salesList []models.Sales

	idb.DB.Scopes(utils.SearchSalesKeyword(c.Request), utils.Paginate(c.Request)).Find(&salesList)

	if len(salesList) <= 0 {
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
		"result": salesList,
		"count":  len(salesList),
	})
	c.Abort()
}

func (idb *InDB) CreateSales(c *gin.Context) {
	var sales models.Sales
	var company models.Company

	sales.Id = uuid.New().String()
	sales.Company_ID = c.PostForm("company_id")
	sales.Name = c.PostForm("name")
	sales.Phone_Number = c.PostForm("phone_number")

	idb.DB.Where("id = ?", sales.Company_ID).Limit(1).Find(&company)
	if company.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"result": "Data 'Perusahaan' tidak valid / tidak ditemukan",
		})
		c.Abort()
		return
	}

	err := idb.DB.Create(&sales).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"result": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Sales created",
	})
	c.Abort()
}

func (idb *InDB) UpdateSales(c *gin.Context) {
	id := c.Query("id")
	var (
		sales        models.Sales
		updatedSales models.Sales
	)

	err := idb.DB.First(&sales, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		c.Abort()
		return
	}

	updatedSales.Company_ID = sales.Company_ID
	updatedSales.Name = sales.Name
	updatedSales.Phone_Number = sales.Phone_Number

	err = idb.DB.Model(&sales).Updates(updatedSales).Error
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
		"message": "Sales updated",
	})
	c.Abort()
}

func (idb *InDB) DeleteSales(c *gin.Context) {
	var sales models.Sales

	id := c.Param("id")
	err := idb.DB.First(&sales, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		c.Abort()
		return
	}

	err = idb.DB.Delete(&sales).Error
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
		"message": "Data deleted successfully",
	})
	c.Abort()
}
