package lib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"text/template"
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

type APIParameters struct {
	Username, Password, Action, Parameters string
}

func GetResponse(url string, parameters APIParameters) (result string, err error) {

	t := template.Must(template.New("SOAPRequest").Parse(SOAPRequest))

	body := bytes.NewBufferString("")
	err = t.Execute(body, parameters)
	//if err != nil {
	//	err = fmt.Errorf("template.Execute: %#v", err)
	//	return
	//}
	response, err := http.Post(url, "text/xml; charset=utf-8", body)
	if err != nil {
		err = fmt.Errorf("http.Post: %#v", err)
		return
	}
	b, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	//if err != nil {
	//	err = fmt.Errorf("ioutil.ReadAll: %#v", err)
	//	return
	//}
	result = string(b)
	return

}

const SOAPRequest = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope"
 xmlns:xsd="http://www.w3.org/2001/XMLSchema"
 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/"
 xmlns:base_services="http://wsdl.cybozu.co.jp/base/2008">
 <SOAP-ENV:Header>
  <Action SOAP-ENV:mustUnderstand="1" xmlns="http://schemas.xmlsoap.org/ws/2003/03/addressing">{{.Action}}</Action>
  <Security xmlns:wsu="http://schemas.xmlsoap.org/ws/2002/07/utility"
   SOAP-ENV:mustUnderstand="1"
   xmlns="http://schemas.xmlsoap.org/ws/2002/12/secext">
   <UsernameToken wsu:Id="id"><Username>{{.Username}}</Username><Password>{{.Password}}</Password></UsernameToken>
  </Security>
  <Timestamp SOAP-ENV:mustUnderstand="1" Id="id"
   xmlns="http://schemas.xmlsoap.org/ws/2002/07/utility">
   <Created>2037-08-12T14:45:00Z</Created>
   <Expires>2037-08-12T14:45:00Z</Expires>
  </Timestamp>
  <Locale>jp</Locale>
  </SOAP-ENV:Header><SOAP-ENV:Body>
  <{{.Action}}>
  {{.Parameters}}
 </{{.Action}}>
</SOAP-ENV:Body></SOAP-ENV:Envelope>
`
