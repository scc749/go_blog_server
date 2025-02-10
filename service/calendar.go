package service

import (
	"encoding/json"
	"server/global"
	"server/model/other"
	"server/utils"
	"time"
)

type CalendarService struct {
}

func (calendarService *CalendarService) GetCalendarByDate(dateStr string) (other.Calendar, error) {
	result, err := global.Redis.Get("calendar-" + dateStr).Result()
	if err != nil {
		calendar, err := utils.GetCalendar(dateStr)
		if err != nil {
			return other.Calendar{}, err
		}
		data, err := json.Marshal(calendar)
		if err != nil {
			return other.Calendar{}, err
		}
		if err := global.Redis.Set("calendar-"+dateStr, data, time.Hour*24).Err(); err != nil {
			return other.Calendar{}, err
		}
		return calendar, nil
	}
	var calendar other.Calendar
	if err := json.Unmarshal([]byte(result), &calendar); err != nil {
		return other.Calendar{}, err
	}
	return calendar, nil
}
