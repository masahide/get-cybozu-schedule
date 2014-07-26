package lib

import (
	"net/http"

	"code.google.com/p/google-api-go-client/calendar/v3"
)

type Calendar interface {
	Get(string) (*calendar.Event, error)
	Insert(*calendar.Event) (*calendar.Event, error)
	Update(string, *calendar.Event) (*calendar.Event, error)
}

type GoogleCalendar struct {
	calendarId string
	svc        *calendar.Service
}

func NewGoogleCalendar(client *http.Client, calendarId string) *GoogleCalendar {
	this := &GoogleCalendar{}
	this.svc, _ = calendar.New(client)
	this.calendarId = calendarId
	return this
}
func (this *GoogleCalendar) Get(id string) (*calendar.Event, error) {
	return this.svc.Events.Get(this.calendarId, id).Do()
}
func (this *GoogleCalendar) Insert(event *calendar.Event) (*calendar.Event, error) {
	return this.svc.Events.Insert(this.calendarId, event).Do()
}
func (this *GoogleCalendar) Update(id string, event *calendar.Event) (*calendar.Event, error) {
	return this.svc.Events.Update(this.calendarId, id, event).Do()
}
func (this *GoogleCalendar) List() (*calendar.Events, error) {
	return this.svc.Events.List(this.calendarId).Do()
}
func (this *GoogleCalendar) ClendarList() (*calendar.CalendarList, error) {
	return this.svc.CalendarList.List().Do()
}

func (this *GoogleCalendar) InsertEvent() (*calendar.Event, error) {
	event := calendar.Event{
		Id:      "123456abcdef",
		Summary: "test test",
		Start: &calendar.EventDateTime{
			DateTime: `2014-07-15T12:30:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: `2014-07-15T13:00:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		Recurrence: []string{
			"RRULE:FREQ=WEEKLY;UNTIL=20140801T000000Z",
			"EXDATE:20140722T123000",
		},
	}
	return this.Insert(&event)
}
func (this *GoogleCalendar) UpdateEvent() (*calendar.Event, error) {
	event := calendar.Event{
		Summary: "test test",
		Start: &calendar.EventDateTime{
			DateTime: `2014-07-15T12:30:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: `2014-07-15T13:00:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		Recurrence: []string{
			"RRULE:FREQ=DAILY;BYDAY=MO,TU,WE,TH,FR;UNTIL=20140801T000000Z",
			"EXDATE:20140716T123000",
			"EXDATE:20140717T123000",
		},
	}
	return this.Update("123456abcdef", &event)
}
func (this *GoogleCalendar) InsertTestWeeklyEvent() (*calendar.Event, error) {
	event := calendar.Event{
		Summary: "test test",
		Start: &calendar.EventDateTime{
			DateTime: `2014-07-15T12:30:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: `2014-07-15T13:00:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		Recurrence: []string{
			"RRULE:FREQ=WEEKLY;UNTIL=20140801T000000Z",
			"EXDATE:20140722T123000",
		},
	}
	return this.Insert(&event)
}

func (this *GoogleCalendar) InsertTestWeekDayEvent() (*calendar.Event, error) {
	event := calendar.Event{
		Summary: "test test",
		Start: &calendar.EventDateTime{
			DateTime: `2014-07-15T12:30:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: `2014-07-15T13:00:00+09:00`,
			TimeZone: "Asia/Tokyo",
		},
		Recurrence: []string{
			"RRULE:FREQ=DAILY;BYDAY=MO,TU,WE,TH,FR;UNTIL=20140801T000000Z",
			"EXDATE:20140716T123000",
			"EXDATE:20140717T123000",
		},
	}
	return this.Insert(&event)
}
