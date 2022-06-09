package form

import (
	"github.com/ernestngugi/go-blog/app/model"
	"gopkg.in/guregu/null.v4"
)

type Blog struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBlog struct {
	Title       null.String      `json:"title"`
	Description null.String      `json:"description"`
	Status      model.BlogStatus `json:"status"`
}
