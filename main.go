package main

import (
	"assignment/docs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	// swagger configuration
	// docs.SwaggerInfo.Host = "8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		person := v1.Group("/person")
		{
			person.GET("/get", Get)
			person.POST("/post", Post)
			person.POST("/delete", Delete)
			person.POST("/update", Update)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:4000/swagger/doc.json")))
	r.Run(":4000")
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

var global *[]Person

// @BasePath /api/v1
// Get godoc
// @Summary get
// @Schemes
// @Description Get
// @Tags person
// @Accept json
// @Produce json
// @Router /person [get]
func Get(c *gin.Context) {
	if global == nil || len(*global) == 0 {
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    nil,
			"message": "No Data Available",
		}
		c.JSON(http.StatusOK, result)
	} else {
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    global,
			"message": "Successfully Get Data",
		}
		c.JSON(http.StatusOK, result)
	}

}

// @BasePath /api/v1
// Post godoc
// @Summary Post
// @Schemes
// @Description Post
// @Tags person
// @Accept json
// @Produce json
// @Param   Body body   Person  true  "payload"
// @Router /person [post]
func Post(c *gin.Context) {
	var (
		input       Person
		todos, data []Person
	)

	c.BindJSON(&input)
	if global == nil {
		input.ID = 1
		todos = append(todos, input)
		global = &todos
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    global,
			"message": "Successfully Posted Data",
		}
		c.JSON(http.StatusOK, result)
	} else {
		data = *global
		input.ID = len(*global) + 1
		data = append(data, input)
		global = &data
		result := map[string]interface{}{
			"status":  http.StatusOK,
			"data":    global,
			"message": "Successfully Posted Data",
		}
		c.JSON(http.StatusOK, result)
	}
}

// @BasePath /api/v1
// Delete godoc
// @Summary delete
// @Schemes
// @Description Delete
// @Tags person
// @Accept json
// @Produce json
// @Param   id  query   string  true  "id"
// @Param   Body body   Person  true  "payload"
// @Router /person/:id [delete]
func Delete(c *gin.Context) {
	var data []Person
	data = *global

	id, _ := strconv.Atoi(c.Param("id"))

	for index, todo := range data {
		if todo.ID == id {
			data = append(data[:index], data[index+1:]...)
			global = &data
			result := map[string]interface{}{
				"status":  http.StatusOK,
				"data":    global,
				"message": "Successfully Deleted Data",
			}
			c.JSON(http.StatusOK, result)
		}
	}

}

// @BasePath /api/v1
// Update godoc
// @Summary update
// @Schemes
// @Description Update
// @Tags person
// @Accept json
// @Produce json
// @Param   id  query   string  true  "id"
// @Param   Body body   Person  true  "payload"
// @Router /person/:id [put]
func Update(c *gin.Context) {
	var (
		data  []Person
		input Person
	)
	data = *global

	id, _ := strconv.Atoi(c.Param("id"))

	if data == nil || len(data) == 0 {
		result := map[string]interface{}{
			"status":  http.StatusBadRequest,
			"data":    global,
			"message": "Invalid Request",
		}
		c.JSON(http.StatusOK, result)
	}

	c.BindJSON(&input)
	for index, curData := range data {
		if curData.ID == id {
			data[index].FirstName = input.FirstName
			data[index].LastName = input.LastName
			global = &data
			result := map[string]interface{}{
				"status":  http.StatusOK,
				"data":    global,
				"message": "Successfully Updated Data",
			}
			c.JSON(http.StatusOK, result)
		}
	}
}
