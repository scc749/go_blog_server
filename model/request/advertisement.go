package request

type AdvertisementCreate struct {
	AdImage string `json:"ad_image" binding:"required"`
	Link    string `json:"link" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type AdvertisementDelete struct {
	IDs []uint `json:"ids" binding:"required"`
}

type AdvertisementUpdate struct {
	ID      uint   `json:"id" binding:"required"`
	Link    string `json:"link" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type AdvertisementList struct {
	Title   *string `json:"title" form:"title"`
	Content *string `json:"content" form:"content"`
	PageInfo
}
