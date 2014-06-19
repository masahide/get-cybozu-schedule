package lib

import (
	"encoding/xml"
	"io"
)

type XMLexclusive struct {
	XMLName xml.Name `xml:"exclusive_datetime"`
	Start   string   `xml:"start,attr"`
	End     string   `xml:"end,attr"`
}
type XMLrepeat_condition struct {
	XMLName   xml.Name `xml:"condition"`
	Type      string   `xml:"type,attr"`
	Day       string   `xml:"day,attr"`
	Week      string   `xml:"week,attr"`
	StartDate string   `xml:"start_date,attr"`
	EndDate   string   `xml:"end_date,attr"`
	StartTime string   `xml:"start_time,attr"`
	EndTime   string   `xml:"end_time,attr"`
}
type XMLrepeat struct {
	XMLName   xml.Name             `xml:"repeat_info"`
	Condition *XMLrepeat_condition `xml:"condition"`
	Exclusive []*XMLexclusive      `xml:"exclusive_datetimes>exclusive_datetime"`
}

type XMLdatetime struct {
	XMLName xml.Name `xml:"datetime"`
	Start   string   `xml:"start,attr"`
	End     string   `xml:"end,attr"`
}

type XMLfacility struct {
	XMLName xml.Name `xml:"facility"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

type XMLuser struct {
	XMLName xml.Name `xml:"user"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}
type XMLmember struct {
	XMLName  xml.Name     `xml:"member"`
	User     *XMLuser     `xml:"user"`
	Facility *XMLfacility `xml:"facility"`
}
type XMLschedule_event struct {
	XMLName    xml.Name     `xml:"schedule_event"`
	Id         int          `xml:"id,attr"`
	EventType  string       `xml:"event_type,attr"`
	PublicType string       `xml:"public_type,attr"`
	Plan       string       `xml:"plan,attr"`
	Detail     string       `xml:"detail,attr"`
	TimeZone   string       `xml:"timezone,attr"`
	Version    int          `xml:"version,attr"`
	Allday     bool         `xml:"allday,attr"`
	StartOnly  bool         `xml:"start_only,attr"`
	Datetime   *XMLdatetime `xml:"when>datetime"`
	Repeat     *XMLrepeat   `xml:"repeat_info"`
	Member     []*XMLmember `xml:"members>member"`
}

type XMLSoap struct {
	XMLName        xml.Name             `xml:"Envelope"`
	Schedule_event []*XMLschedule_event `xml:"Body>ScheduleGetEventsResponse>returns>schedule_event"`
}

func ReadSoap(reader io.Reader) ([]*XMLschedule_event, error) {
	xmlSoap := &XMLSoap{}
	decoder := xml.NewDecoder(reader)

	if err := decoder.Decode(xmlSoap); err != nil {
		return nil, err
	}

	return xmlSoap.Schedule_event, nil
}
