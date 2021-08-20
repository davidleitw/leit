package calendar

import (
	"leit/pkg/internal"
	"time"
)

type CalendarDay struct {
	date    time.Time
	events  *CalendarEvent
	records *Record
}

func (day *CalendarDay) Date() time.Time {
	return day.date
}

func ParseCalendarDayCreationByFullStr(inputDateStr string) error {
	internal.VaildDateStr(inputDateStr)

	return nil
}

func CreateNewCalendarDay(year, month, day int) *CalendarDay {
	// check specific day is exist or not.

	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &CalendarDay{
		date:    d,
		events:  new(CalendarEvent),
		records: new(Record),
	}
}

func (Cal *CalendarDay) CreateNewCalendarEvent() {

}
