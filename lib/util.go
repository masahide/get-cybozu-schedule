package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ClientID     string
	ClientSecret string
	Url          string
}

// 設定ファイルを読み込む
func Parse(filename string) (config Config, err error) {
	var c Config
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("error: readFile %v", err)
		return
	}
	err = json.Unmarshal(jsonString, &c)
	if err != nil {
		err = fmt.Errorf("error: json.Unmarshal %v", err)
		return
	}
	return c, nil
}
