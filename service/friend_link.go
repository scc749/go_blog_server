package service

import (
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/model/other"
	"server/model/request"
	"server/utils"

	"gorm.io/gorm"
)

type FriendLinkService struct {
}

func (friendLinkService *FriendLinkService) FriendLinkInfo() (friendLinks []database.FriendLink, total int64, err error) {
	err = global.DB.Model(&database.FriendLink{}).Count(&total).Find(&friendLinks).Error
	if err != nil {
		return nil, 0, err
	}
	return friendLinks, total, nil
}

func (friendLinkService *FriendLinkService) FriendLinkCreate(req request.FriendLinkCreate) error {
	friendLinkToCreate := database.FriendLink{
		Logo:        req.Logo,
		Link:        req.Link,
		Name:        req.Name,
		Description: req.Description,
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := utils.ChangeImagesCategory(tx, []string{friendLinkToCreate.Logo}, appTypes.Logo); err != nil {
			return err
		}

		return tx.Create(&friendLinkToCreate).Error
	})
}

func (friendLinkService *FriendLinkService) FriendLinkDelete(req request.FriendLinkDelete) error {
	if len(req.IDs) == 0 {
		return nil
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range req.IDs {
			var friendLinkToDelete database.FriendLink
			if err := tx.Take(&friendLinkToDelete, id).Error; err != nil {
				return err
			}
			if err := utils.InitImagesCategory(tx, []string{friendLinkToDelete.Logo}); err != nil {
				return err
			}
			if err := tx.Delete(&friendLinkToDelete).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (friendLinkService *FriendLinkService) FriendLinkUpdate(req request.FriendLinkUpdate) error {
	updates := struct {
		Link        string `json:"link"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Link:        req.Link,
		Name:        req.Name,
		Description: req.Description,
	}
	return global.DB.Take(&database.FriendLink{}, req.ID).Updates(updates).Error
}

func (friendLinkService *FriendLinkService) FriendLinkList(info request.FriendLinkList) (list interface{}, total int64, err error) {
	db := global.DB

	if info.Name != nil {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}

	if info.Description != nil {
		db = db.Where("description LIKE ?", "%"+*info.Description+"%")
	}

	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
	}

	return utils.MySQLPagination(&database.FriendLink{}, option)
}
