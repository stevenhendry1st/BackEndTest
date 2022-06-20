package app

import (
	"backend-test/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()	

	//ENDPOINT CREATE
	router.POST("/article", func(c *gin.Context) {
		var requestData model.Article

		errBind := c.ShouldBindJSON(&requestData)
		if errBind != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "input invalid",
			})
			return
		}

		if requestData.Title == "" || requestData.Content == "" || requestData.Category == "" || requestData.Status == "" {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "make sure to input all the fields",
			})
			return
		}
		
		requestData.CreatedDate = time.Now()
		requestData.UpdatedDate = time.Now()

		if err := Db.Create(&requestData).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "storing data failed",
			})
			return			
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "create article success",
			"data": requestData,
		})
	})	

	//ENDPOINT GET BY ID
	router.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		idConverted, _ := strconv.Atoi(id)

		var article model.Article

		if err := Db.Where("id = ?", idConverted).Find(&article).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "get article by id failed",
			})
			return			
		}

		if article.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "article not found",
				"data": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "get article by id success",
			"data": article,
		})
	})

	//ENDPOINT UPDATE BY ID
	router.PUT("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		idConverted, _ := strconv.Atoi(id)

		var article model.Article
		var articleCheck model.Article
		
		errBind := c.ShouldBindJSON(&article)
		if errBind != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "input invalid",
			})
			return
		}
		
		if article.Title == "" || article.Content == "" || article.Category == "" || article.Status == "" {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "make sure to input all the fields",
			})
			return
		}
		
		//CHECK KE DB DATA EXISTED OR NOT
		if err := Db.Where("id = ?", idConverted).Find(&articleCheck).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "get article by id failed",
			})
			return			
		}

		if articleCheck.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "article not found",
				"data": nil,
			})
			return
		}
		//

		//UPDATE DATA TO DB BY ID
		if err := Db.Model(&model.Article{}).Where("id = ?", idConverted).Updates(map[string]interface{}{"title": article.Title, "content": article.Content, "category": article.Category, "status": article.Status, "updated_date": time.Now()}).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "update article by id failed",
			})
			return			
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "update article by id success",
			"data": article,
		})
	})

	//ENDPOINT DELETE
	router.DELETE("/article/:id", func(c *gin.Context) {
		id := c.Param("id")

		idConverted, _ := strconv.Atoi(id)

		var articleCheck model.Article
		
		//CHECK KE DB DATA EXISTED OR NOT
		if err := Db.Where("id = ?", idConverted).Find(&articleCheck).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "get article by id failed",
			})
			return			
		}

		if articleCheck.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "article not found",
				"data": nil,
			})
			return
		}
		//

		//DELETE DATA TO DB BY ID
		if err := Db.Model(&model.Article{}).Where("id = ?", idConverted).Delete(&articleCheck).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "delete article by id failed",
			})
			return			
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "delete article by id success",
			"data": nil,
		})
	})


	router.GET("/articles/:limit/:offset", func(c *gin.Context) {
		limit := c.Param("limit")
		offset := c.Param("offset")

		limitConverted, _ := strconv.Atoi(limit)
		offsetConverted, _ := strconv.Atoi(offset)

		var articles []model.Article

		if err := Db.Limit(limitConverted).Offset(offsetConverted).Find(&articles).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "get data failed",
			})
			return			
		}

		
		c.JSON(http.StatusOK, gin.H{
			"message": "get article success",
			"data": articles,
		})
	})

	



	router.Run(":8000")
}