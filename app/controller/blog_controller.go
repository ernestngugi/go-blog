package controller

import (
	"context"
	"errors"

	"github.com/ernestngugi/go-blog/app/db"
	"github.com/ernestngugi/go-blog/app/form"
	"github.com/ernestngugi/go-blog/app/model"
	"github.com/ernestngugi/go-blog/app/repository"
)

type (
	BlogController interface {
		CreateBlog(ctx context.Context, dB db.DB, blog *form.Blog) (*model.Blog, error)
		BlogByID(ctx context.Context, dB db.DB, blogID int64) (*model.Blog, error)
		AllBlogs(ctx context.Context, dB db.DB, filter string) ([]*model.Blog, error)
		DeleteBlog(ctx context.Context, dB db.DB, blogID int64) error
		UpdateBlog(ctx context.Context, dB db.DB, blogID int64, form *form.UpdateBlog) error
	}

	AppBlogController struct {
		blogRepository repository.BlogRepository
	}
)

func NewBlogRepository(
	blogRepository repository.BlogRepository,
) *AppBlogController {
	return &AppBlogController{
		blogRepository: blogRepository,
	}
}

func (c *AppBlogController) DeleteBlog(
	ctx context.Context,
	dB db.DB,
	blogID int64,
) error {

	blog, err := c.blogRepository.BlogByID(ctx, dB, blogID)
	if err != nil {
		return err
	}

	return c.blogRepository.DeleteBlog(ctx, dB, blog.ID)
}

func (c *AppBlogController) AllBlogs(
	ctx context.Context,
	dB db.DB,
	status string,
) ([]*model.Blog, error) {

	if !model.BlogStatus(status).IsValid() {
		return []*model.Blog{}, errors.New("invalid status")
	}

	return c.blogRepository.ListBlogs(ctx, dB, status)
}

func (c *AppBlogController) UpdateBlog(
	ctx context.Context,
	dB db.DB,
	blogID int64,
	form *form.UpdateBlog,
) error {

	blog, err := c.blogRepository.BlogByID(ctx, dB, blogID)
	if err != nil {
		return err
	}

	if form.Title.Valid {
		blog.Title = form.Title.String
	}

	if form.Description.Valid {
		blog.Description = form.Description.String
	}

	if form.Status.IsValid() {
		blog.Status = model.BlogStatus(form.Status.String())
	}

	return c.blogRepository.Save(ctx, dB, blog)
}

func (c *AppBlogController) BlogByID(
	ctx context.Context,
	dB db.DB,
	blogID int64,
) (*model.Blog, error) {
	return c.blogRepository.BlogByID(ctx, dB, blogID)
}

func (c *AppBlogController) CreateBlog(
	ctx context.Context,
	dB db.DB,
	form *form.Blog,
) (*model.Blog, error) {

	blog := &model.Blog{
		Title:       form.Title,
		Description: form.Description,
		Status:      model.BlogInactive,
	}

	err := c.blogRepository.Save(ctx, dB, blog)
	if err != nil {
		return &model.Blog{}, err
	}

	return blog, nil
}
