package blogs

import (
	"github.com/ernestngugi/go-blog/app/controller"
	"github.com/ernestngugi/go-blog/app/db"
	"github.com/gin-gonic/gin"
)

func Endpoints(
	r *gin.RouterGroup,
	dB db.DB,
	blogController controller.BlogController,
) {
	r.POST("/blog", createBlog(dB, blogController))
	r.GET("/blog/:id", getBlog(dB, blogController))
	r.GET("/blogs", allBlogs(dB, blogController))
	r.PUT("/blog/:id", updateBlog(dB, blogController))
	r.DELETE("/blog:id", deleteBlog(dB, blogController))
}
