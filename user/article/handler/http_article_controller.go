package handler

import (
	"net/http"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/article/usecase"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	ArticleUsecase usecase.ArticleUsecase
}

func NewArticleHandler(ArticleUsecase usecase.ArticleUsecase) *articleHandler {
	return &articleHandler{ArticleUsecase: ArticleUsecase}
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

func (h *articleHandler) GetAllComment(c echo.Context) error {

	getAllCommentReq := article.GetAllCommentRequest{}

	c.Bind(&getAllCommentReq)

	if err := isRequestValid(&getAllCommentReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	page, offset, limit := helper.GetPaginateData(getAllCommentReq.Page, getAllCommentReq.Limit, "mobile")

	comments, totalData, err := h.ArticleUsecase.GetAllComment(getAllCommentReq.ArticleID, offset, limit)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
		case article.ErrArticleNotFound:
			status = http.StatusNotFound
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseData(err.Error(), http.StatusInternalServerError, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success get all article comment", http.StatusOK, echo.Map{
		"comments":      comments,
		"current_pages": page,
		"total_pages":   totalPages,
	}))
}

func (h *articleHandler) CreateComment(c echo.Context) error {

	var commentReq article.CreateCommentRequest

	var user = c.Get("user").(*helper.JwtCustomUserClaims)

	c.Bind(&commentReq)

	if err := isRequestValid(&commentReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	commentReq.UserID = user.ID

	err := h.ArticleUsecase.CreateComment(commentReq)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
		case article.ErrArticleNotFound:
			status = http.StatusNotFound
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success create comment", http.StatusOK, nil))
}

func (h *articleHandler) DeleteComment(c echo.Context) error {
	var commentReq article.DeleteCommentRequest

	var user = c.Get("user").(*helper.JwtCustomUserClaims)
	c.Bind(&commentReq)

	if err := isRequestValid(&commentReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseData(err.Error(), http.StatusBadRequest, nil))
	}

	commentReq.UserID = user.ID

	err := h.ArticleUsecase.DeleteComment(commentReq.UserID, commentReq.ArticleID, commentReq.CommentID)

	if err != nil {
		status := http.StatusInternalServerError

		switch err {
		case article.ErrArticleNotFound:
			status = http.StatusNotFound
		}

		return c.JSON(status, helper.ResponseData(err.Error(), status, nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseData("success delete comment", http.StatusOK, nil))
}
