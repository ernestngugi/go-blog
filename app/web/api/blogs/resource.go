package blogs

import (
	"net/http"
	"strconv"

	"github.com/ernestngugi/go-blog/app/controller"
	"github.com/ernestngugi/go-blog/app/db"
	"github.com/ernestngugi/go-blog/app/form"
	"github.com/gin-gonic/gin"
)

func createBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		var form form.Blog
		err := c.BindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "invalid form"})
			return
		}

		ctx := c.Request.Context()

		blog, err := blogController.CreateBlog(ctx, dB, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, blog)
	}
}

func getBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		blogID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "invalid id"})
			return
		}

		ctx := c.Request.Context()

		blog, err := blogController.BlogByID(ctx, dB, blogID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
			return
		}

		c.JSON(http.StatusOK, blog)
	}
}

func updateBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		var blog form.UpdateBlog
		err := c.BindJSON(&blog)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "invalid form"})
			return
		}

		blogID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "invalid id"})
			return
		}

		ctx := c.Request.Context()

		err = blogController.UpdateBlog(ctx, dB, blogID, &blog)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err})
			return
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func deleteBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		blogID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": "invalid id"})
			return
		}

		ctx := c.Request.Context()

		err = blogController.DeleteBlog(ctx, dB, blogID)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, "OK")
	}
}

func allBlogs(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		blogList, err := blogController.AllBlogs(ctx, dB, "active")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, blogList)
	}
}
