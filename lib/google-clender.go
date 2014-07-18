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
	return this.svc.Events.Update(id, this.calendarId, event).Do()
}
func (this *GoogleCalendar) List() (*calendar.Events, error) {
	return this.svc.Events.List(this.calendarId).Do()
}
func (this *GoogleCalendar) ClendarList() (*calendar.CalendarList, error) {
	return this.svc.CalendarList.List().Do()
}
