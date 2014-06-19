package lib

import (
	"bytes"
	"testing"
)

func TestReadSoap(t *testing.T) {
	var xmlSoapBody []*XMLschedule_event

	testXMLbuffer := bytes.NewBufferString(testXML)
	xmlSoapBody, err := ReadSoap(testXMLbuffer)
	if err != nil {
		t.Fail()
	}
	//pretty.Printf("xmlSoapBody:%# v\n", xmlSoapBody)
	if 111111 != xmlSoapBody[1].Id {
		t.Errorf("111111 != xmlSoapBody[1].Id")
	}
	if 1490467551 != xmlSoapBody[1].Version {
		t.Errorf("1490467551 != xmlSoapBody[1].Version")
	}
	if "week" != xmlSoapBody[1].Repeat.Condition.Type {
		t.Errorf(`"week" != xmlSoapBody[1].Repeat.Condition.Type`)
	}
	if "部屋1" != xmlSoapBody[0].Member[1].Facility.Name {
		t.Errorf(`"部屋1" != xmlSoapBody[1].Facility.Name`)
	}
	if "2014-01-23T00:00:00+09:00" != xmlSoapBody[1].Repeat.Exclusive[0].Start {
		t.Errorf(`"2014-01-23T00:00:00+09:00" != xmlSoapBody[1].Repeat.Exclusive[0].Start`)
	}
}

func TestBadXMLReadSoap(t *testing.T) {
	testXMLbuffer := bytes.NewBufferString("hoge")
	_, err := ReadSoap(testXMLbuffer)
	if "EOF" != err.Error() {
		t.Errorf("%s", err)
	}
}

const testXML = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope
 xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
 xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns:xsd="http://www.w3.org/2001/XMLSchema"
 xmlns:schedule="http://wsdl.cybozu.co.jp/schedule/2008">
 <soap:Header><vendor>Cybozu</vendor><product>Garoon</product><product_type>1</product_type><version>3.7.2</version><apiversion>1.3.0</apiversion></soap:Header>
 <soap:Body>
  <schedule:ScheduleGetEventsResponse>
   <returns>
    <schedule_event id="22222"
     event_type="repeat" 
     public_type="public" 
     plan="打合せ" 
     detail="詳細１" 
     version="122326785"
     timezone="Asia/Tokyo"
     allday="false" 
     start_only="false"
     >
     <members xmlns="http://schemas.cybozu.co.jp/schedule/2008">
      <member>
       <user id="111" name="名前1" order="0"/>
      </member>
      <member>
       <facility id="222" name="部屋1" order="1"/>
      </member>
     </members>
     <repeat_info xmlns="http://schemas.cybozu.co.jp/schedule/2008">
      <condition type="weekday" day="1" 
       week="3" start_date="2014-01-01" end_date="2014-01-31"
       start_time="10:30:00" end_time="13:00:00"/>
      <exclusive_datetimes>
      </exclusive_datetimes>
     </repeat_info>
     </schedule_event>   
	 <schedule_event id="111111"
     event_type="repeat" 
     public_type="public" 
     detail="詳細2" 
     version="1490467551"
     timezone="Asia/Tokyo"
     allday="false" 
     start_only="false"
     >
     <members xmlns="http://schemas.cybozu.co.jp/schedule/2008">
      <member>
       <user id="33" name="名前1" order="0"/>
      </member>
      <member>
       <user id="333" name="名前2" order="1"/>
      </member>
     </members>
     <repeat_info xmlns="http://schemas.cybozu.co.jp/schedule/2008">
      <condition type="week" day="26" 
       week="4" start_date="2014-01-01" end_date="2014-01-31"
       start_time="16:00:00" end_time="19:00:00"/>
      <exclusive_datetimes>
       <exclusive_datetime start="2014-01-23T00:00:00+09:00" end="2014-01-24T00:00:00+09:00" />
       <exclusive_datetime start="2014-01-16T00:00:00+09:00" end="2014-01-17T00:00:00+09:00" />
      </exclusive_datetimes>
     </repeat_info>
     </schedule_event> </returns>
  </schedule:ScheduleGetEventsResponse>
 </soap:Body>
</soap:Envelope>
`
