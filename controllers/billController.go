package controllers

import (
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetBill(c *gin.Context) {
	var bill models.Bill

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&bill).Error
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
		"result": bill,
		"count":  1,
	})

	c.Abort()
}

func (idb *InDB) GetBills(c *gin.Context) {
	var bills []models.Bill

	idb.DB.Scopes(utils.Paginate(c.Request)).Find(&bills)

	if len(bills) <= 0 {
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
		"result": bills,
		"count":  len(bills),
	})
	c.Abort()
}

func (idb *InDB) CreateBill(c *gin.Context) {
	var bill models.Bill
	dateString := "2022-07-06"

	salesId, err := strconv.ParseUint(c.PostForm("sales_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Wrong input",
		})
		c.Abort()
		return
	}

	bill.Sales_ID = uint(salesId)
	bill.Bill_Number = c.PostForm("bill_number")
	bill.Bill_Total, err = strconv.Atoi(c.PostForm("bill_total"))
	bill.Input_Date, err = time.Parse(dateString, c.PostForm("input_date"))
	bill.Due_Date, err = time.Parse(dateString, c.PostForm("due_date"))

	idb.DB.Create(&bill)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Input tagihan sukses",
	})
	c.Abort()
}
