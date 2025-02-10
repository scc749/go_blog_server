package response

import "server/model/database"

type AdvertisementInfo struct {
	List  []database.Advertisement `json:"list"`
	Total int64                    `json:"total"`
}
