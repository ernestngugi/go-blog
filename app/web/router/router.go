package router

import (
	"net/http"

	"github.com/ernestngugi/go-blog/app/controller"
	"github.com/ernestngugi/go-blog/app/db"
	"github.com/ernestngugi/go-blog/app/repository"
	"github.com/ernestngugi/go-blog/app/web/api/blogs"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func BuildRouter(
	dB db.DB,
) *Router {

	router := gin.Default()

	v1Router := router.Group("/v1")

	blogRepository := repository.NewBlogRepository()

	blogController := controller.NewBlogRepository(blogRepository)

	blogs.Endpoints(v1Router, dB, blogController)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error_message": "Resource not found"})
	})
	
	return &Router{
		router,
	}
}