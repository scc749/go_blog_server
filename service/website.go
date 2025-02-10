package service

import (
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/model/other"
	"server/model/request"
	"server/utils"
)

type WebsiteService struct {
}

func (websiteService *WebsiteService) WebsiteCarousel() []string {
	var urls []string
	global.DB.Model(&database.Image{}).Where("category = ?", appTypes.Carousel).Pluck("url", &urls)
	return urls
}

func (websiteService *WebsiteService) WebsiteNews(sourceStr string) (other.HotSearchData, error) {
	hotSearchData, err := ServiceGroupApp.HotSearchService.GetHotSearchDataBySource(sourceStr)
	if err != nil {
		return other.HotSearchData{}, err
	}
	return hotSearchData, nil
}

func (websiteService *WebsiteService) WebsiteCalendar(dateStr string) (other.Calendar, error) {
	calendar, err := ServiceGroupApp.CalendarService.GetCalendarByDate(dateStr)
	if err != nil {
		return other.Calendar{}, err
	}
	return calendar, nil
}

func (websiteService *WebsiteService) WebsiteFooterLink() []database.FooterLink {
	var footerLinks []database.FooterLink
	global.DB.Find(&footerLinks)
	return footerLinks
}

func (websiteService *WebsiteService) WebsiteAddCarousel(req request.WebsiteCarouselOperation) error {
	return utils.ChangeImagesCategory(global.DB, []string{req.Url}, appTypes.Carousel)
}

func (websiteService *WebsiteService) WebsiteCancelCarousel(req request.WebsiteCarouselOperation) error {
	return utils.InitImagesCategory(global.DB, []string{req.Url})
}

func (websiteService *WebsiteService) WebsiteCreateFooterLink(req database.FooterLink) error {
	return global.DB.Save(&req).Error
}

func (websiteService *WebsiteService) WebsiteDeleteFooterLink(req database.FooterLink) error {
	return global.DB.Delete(&req).Error
}
