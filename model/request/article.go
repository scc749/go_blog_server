package request

type ArticleInfoByID struct {
	ID string `json:"id" form:"id" uri:"id" binding:"required"`
}

type ArticleSearch struct {
	Query    string `json:"query" form:"query" uri:"query" binding:"max=50"`
	Category string `json:"category" form:"category" uri:"category"`
	Tag      string `json:"tag" form:"tag" uri:"tag"`
	Sort     string `json:"sort" form:"sort" uri:"sort"`
	Order    string `json:"order" form:"order" uri:"order" binding:"required"`
	PageInfo
}

type ArticleLike struct {
	UserID    uint   `json:"-"`
	ArticleID string `json:"article_id" form:"article_id" binding:"required"`
}

type ArticleLikesList struct {
	UserID uint `json:"-"`
	PageInfo
}

type ArticleCreate struct {
	Cover    string   `json:"cover" binding:"required"`
	Title    string   `json:"title" binding:"required"`
	Category string   `json:"category" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
	Abstract string   `json:"abstract" binding:"required"`
	Content  string   `json:"content" binding:"required"`
}

type ArticleDelete struct {
	IDs []string `json:"ids" binding:"required"`
}

type ArticleUpdate struct {
	ID       string   `json:"id" binding:"required"`
	Cover    string   `json:"cover" binding:"required"`
	Title    string   `json:"title" binding:"required"`
	Category string   `json:"category" binding:"required"`
	Tags     []string `json:"tags" binding:"required"`
	Abstract string   `json:"abstract" binding:"required"`
	Content  string   `json:"content" binding:"required"`
}

type ArticleList struct {
	Title    *string `json:"title" form:"title"`
	Category *string `json:"category" form:"category"`
	Abstract *string `json:"abstract" form:"abstract"`
	PageInfo
}
