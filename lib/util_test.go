package lib

import "testing"

func TestParse(t *testing.T) {

	_, err := Parse("test")
	if err.Error() != "error: readFile open test: no such file or directory" {
		t.Errorf("Error Parse: %#v", err)
		return
	}
	config, err := Parse("test1.json")
	if err != nil {
		t.Errorf("Error Parse: %#v", err)
		return
	}
	if config.Installed.ClientID != "hogeid" {
		t.Errorf("Error Parse: %#v", config.Installed.ClientID)
		return
	}
	_, err = Parse("test2.json")
	if err.Error() != `error: json.Unmarshal invalid character 'b' looking for beginning of object key string` {
		t.Errorf("Error Parse: %#v", err)
		return
	}

}
