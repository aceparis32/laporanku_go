package controllers

import (
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (idb *InDB) GetCompany(c *gin.Context) {
	var company models.Company

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&company).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
			"count":   0,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": company,
		"count":   1,
	})
	c.Abort()
}

func (idb *InDB) GetCompanies(c *gin.Context) {
	var companies []models.Company
	idb.DB.Scopes(utils.SearchCompanyKeyword(c.Request), utils.Paginate(c.Request)).Find(&companies)

	if len(companies) <= 0 {
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
		"result": companies,
		"count":  len(companies),
	})
	c.Abort()
}

func (idb *InDB) CreateCompany(c *gin.Context) {
	var company models.Company

	name := c.PostForm("name")
	address := c.PostForm("address")
	phoneNumber := c.PostForm("phone_number")
	bankName := c.PostForm("bank_account_name")
	bankNumber := c.PostForm("bank_account_number")

	company.Id = uuid.New().String()
	company.Name = name
	company.Address = address
	company.Phone_Number = phoneNumber
	company.Bank_Account_Name = bankName
	company.Bank_Account_Number = bankNumber

	bankNameValid := Contains(utils.GetBankNamesConstant(), company.Bank_Account_Name)
	if !bankNameValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Nama bank tidak ditemukan",
		})
		c.Abort()
		return
	}

	err := idb.DB.Create(&company).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data created",
	})
	c.Abort()
}

func (idb *InDB) UpdateCompany(c *gin.Context) {
	id := c.Query("id")

	var (
		company        models.Company
		updatedCompany models.Company
	)

	err := idb.DB.First(&company, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data tidak ditemukan",
		})
		c.Abort()
		return
	}

	updatedCompany.Name = c.PostForm("name")
	updatedCompany.Address = c.PostForm("address")
	updatedCompany.Phone_Number = c.PostForm("phone_number")
	updatedCompany.Bank_Account_Name = c.PostForm("bank_account_name")
	updatedCompany.Bank_Account_Number = c.PostForm("bank_account_number")

	bankNameValid := Contains(utils.GetBankNamesConstant(), updatedCompany.Bank_Account_Name)
	if !bankNameValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Nama bank tidak ditemukan",
		})
		c.Abort()
		return
	}

	err = idb.DB.Model(&company).Updates(updatedCompany).Error
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
		"message": "Company updated",
	})
	c.Abort()
}

func (idb *InDB) DeleteCompany(c *gin.Context) {
	id := c.Query("id")

	var company models.Company

	err := idb.DB.First(&company, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	err = idb.DB.Delete(&company).Error

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
		"message": "Company data deleted",
	})
	c.Abort()
}
