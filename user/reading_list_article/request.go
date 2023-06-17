package readingListArticle

type CreateRequest struct {
	ID            string `json:"id"`
	ArticleId     string `json:"article_id"`
	ReadingListId string `json:"reading_list_id"`
	UserId        string `json:"user_id"`
}
