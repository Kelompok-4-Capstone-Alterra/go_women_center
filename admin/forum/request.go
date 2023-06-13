package forum

type GetAllRequest struct {
	Topic      string `query:"topic"`
	Created    string `query:"created"`
	Popular    string `query:"popular"`
	Categories int    `query:"categories"`
	Page       int    `query:"page"`
	Offset     int    `query:"offset"`
	Limit      int    `query:"limit"`
}
