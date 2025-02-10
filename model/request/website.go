package request

type WebsiteCarouselOperation struct {
	Url string `json:"url" binding:"required"`
}
