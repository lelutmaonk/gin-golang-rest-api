package productcontroller

import (
	"encoding/json"
	"go_tutorial/rest_api_gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []model.Product

	model.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {

	var product model.Product

	product_id := c.Param("product_id")

	if err := model.DB.First(&product, product_id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Create(c *gin.Context) {

	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	model.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {

	var product model.Product

	product_id := c.Param("product_id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if model.DB.Model(&product).Where("product_id = ?", product_id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update product"})

}

func Delete(c *gin.Context) {

	var product model.Product

	var input struct {
		ProductID json.Number `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	product_id, _ := input.ProductID.Int64()

	if model.DB.Delete(&product, product_id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete product"})

}
