package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/article/usecase"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	ArticleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(ArticleUsecase usecase.ArticleUsecase) *articleHandler {
	return &articleHandler{ArticleUsecase: ArticleUsecase}
}

func (h *articleHandler) Create(c echo.Context) error {

	articleReq := article.CreateRequest{}
	imgInput, _ := c.FormFile("image")
	articleReq.Image = imgInput
	c.Bind(&articleReq)

	if err := isRequestValid(articleReq); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.ArticleUsecase.Create(articleReq, imgInput)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success create article", http.StatusOK, nil))

}

func (h *articleHandler) GetAll(c echo.Context) error {

	page, _ := helper.StringToInt(c.QueryParam("page"))
	limit, _ := helper.StringToInt(c.QueryParam("limit"))

	page, offset, limit := helper.GetPaginateData(page, limit)
	search := c.QueryParam("search")
	articles, totalPages, err := h.ArticleUsecase.GetAll(search, offset, limit)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	if page > totalPages {
		return c.JSON(
			http.StatusNotFound,
			helper.ResponseData(article.ErrPageNotFound.Error(), http.StatusBadRequest, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all article", http.StatusOK, echo.Map{
		"articles":      articles,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (h *articleHandler) GetById(c echo.Context) error {

	var id article.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	article, err := h.ArticleUsecase.GetById(id.ID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get article by id", http.StatusOK, article))
}

func (h *articleHandler) Update(c echo.Context) error {

	var articleReq article.UpdateRequest
	imgInput, _ := c.FormFile("image")
	articleReq.Image = imgInput
	c.Bind(&articleReq)

	if err := isRequestValid(articleReq); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.ArticleUsecase.Update(articleReq, imgInput)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success update article", http.StatusOK, nil))
}

func (h *articleHandler) Delete(c echo.Context) error {

	var id article.IdRequest

	c.Bind(&id)

	if err := isRequestValid(id); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	err := h.ArticleUsecase.Delete(id.ID)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			helper.ResponseData(err.Error(), http.StatusBadRequest, nil),
		)
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success delete article", http.StatusOK, nil))

}
