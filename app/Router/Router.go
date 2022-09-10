package Router

import (
	"main/app/Business/Handle"
	e "main/app/ErrorHandle"

	"github.com/gin-gonic/gin"
)

func RouterStart() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//根目錄
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error_code": "",
			"error_text": "",
			"result":     "hi",
		})
	})
	api := r.Group("/api")
	{
		api.GET("/todo", e.ErrorWrapper(Handle.GetTodo))
		api.POST("/todo", e.ErrorWrapper(Handle.GetTodo))
		api.DELETE("/todo", e.ErrorWrapper(Handle.GetTodo))
	}

	r.Run(":8083")
}
