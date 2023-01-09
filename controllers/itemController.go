package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (idb *InDB) GetItem(c *gin.Context) {
	var item models.Item

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&item).Error
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
		"result": item,
		"count":  1,
	})

	c.Abort()
}

func (idb *InDB) GetItems(c *gin.Context) {
	var items []models.Item

	idb.DB.Scopes(utils.SearchItemKeyword(c.Request, "item_name"), utils.Paginate(c.Request)).Find(&items)

	if len(items) <= 0 {
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
		"result": items,
		"count":  len(items),
	})
	c.Abort()
}

func (idb *InDB) CreateItem(c *gin.Context) {
	var items []models.Item

	reqBody, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(reqBody, &items)

	for i := 0; i < len(items); i++ {
		items[i].Id = uuid.NewString()
	}

	err := idb.DB.Create(&items).Error

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
		"message": "Barang sukses di inputkan",
	})
	c.Abort()
}

func (idb *InDB) UpdateItem(c *gin.Context) {
	id := c.Query("id")
	var (
		item        models.Item
		updatedItem models.Item
	)

	err := idb.DB.First(&item, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data not found",
		})
		c.Abort()
		return
	}

	updatedItem.Item_Name = item.Item_Name
	updatedItem.Capital_Price = item.Capital_Price
	updatedItem.Selling_Price = item.Selling_Price
	updatedItem.Photo_Link = item.Photo_Link
	updatedItem.Sales_ID = item.Sales_ID

	err = idb.DB.Model(&item).Updates(updatedItem).Error
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

func (idb *InDB) DeleteItem(c *gin.Context) {
	var item models.Item

	id := c.Param("id")
	err := idb.DB.First(&item, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data tidak ditemukan",
		})
		c.Abort()
		return
	}

	err = idb.DB.Delete(&item).Error
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
