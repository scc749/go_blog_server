package request

type FriendLinkCreate struct {
	Logo        string `json:"logo" binding:"required"`
	Link        string `json:"link" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type FriendLinkDelete struct {
	IDs []uint `json:"ids" binding:"required"`
}

type FriendLinkUpdate struct {
	ID          uint   `json:"id" binding:"required"`
	Link        string `json:"link" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type FriendLinkList struct {
	Name        *string `json:"name" form:"name"`
	Description *string `json:"description" form:"description"`
	PageInfo
}
