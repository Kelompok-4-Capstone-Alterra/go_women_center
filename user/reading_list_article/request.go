package readingListArticle

type CreateRequest struct {
	ID            string `json:"id"`
	ArticleId     string `json:"article_id" form:"article_id"`
	ReadingListId string `json:"reading_list_id" form:"reading_list_id"`
	UserId        string `json:"user_id" form:"user_id"`
}
