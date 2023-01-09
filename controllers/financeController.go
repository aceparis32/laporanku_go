package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project-go-dasar/models"
	"project-go-dasar/utils"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetNote(c *gin.Context) {
	var note models.Note

	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&note).Error

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
		"result": note,
		"count":  1,
	})

	c.Abort()
}

func (idb *InDB) GetNotes(c *gin.Context) {
	var notes []models.Note

	idb.DB.Scopes(utils.Paginate(c.Request)).Find(&notes)

	if len(notes) <= 0 {
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
		"result": notes,
		"count":  len(notes),
	})
	c.Abort()
}

func (idb *InDB) CreateNote(c *gin.Context) {
	var notes []models.Note

	reqBody, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(reqBody, &notes)

	err := idb.DB.Create(&notes).Error

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
		"message": "Nota sukses di inputkan",
	})
	c.Abort()
}

func (idb *InDB) UpdateNote(c *gin.Context) {
	id := c.Query("id")
	var (
		note        models.Note
		updatedNote models.Note
	)

	err := idb.DB.First(&note, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data tidak ditemukan",
		})
		c.Abort()
		return
	}

	updatedNote.NoteNumber = note.NoteNumber
	updatedNote.Income = note.Income
	updatedNote.Outcome = note.Outcome
	updatedNote.Input_Date = note.Input_Date
	updatedNote.Description = note.Description

	err = idb.DB.Model(&note).Updates(updatedNote).Error
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
		"message": "Nota berhasil di update",
	})
	c.Abort()
}

func (idb *InDB) DeleteNote(c *gin.Context) {
	var note models.Note

	id := c.Param("id")
	err := idb.DB.First(&note, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data tidak ditemukan",
		})
		c.Abort()
		return
	}

	err = idb.DB.Delete(&note).Error
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
		"message": "Data nota berhasil dihapus",
	})
	c.Abort()
}
