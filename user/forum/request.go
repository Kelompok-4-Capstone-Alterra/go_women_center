package forum

type QueryParamRequest struct {
	IdUser     string
	Topic      string `query:"topic"`
	Created    string `query:"created"`
	Popular    string `query:"popular"`
	Categories int    `query:"categories"`
	MyForum    string `query:"myforum"`
	Page       int    `query:"page"`
	Offset     int    `query:"offset"`
	Limit      int    `query:"limit"`
}
