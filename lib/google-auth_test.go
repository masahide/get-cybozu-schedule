package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"code.google.com/p/goauth2/oauth"
)

func TestRedirectHandler(t *testing.T) {
	redirect := NewRedirect(make(chan RedirectResult, 1))
	ts := httptest.NewServer(http.HandlerFunc(redirect.Handler))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?code=111")
	if err != nil {
		t.Errorf("unexpected: %#v", err)
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
	var result RedirectResult
	result = <-redirect.Result
	//fmt.Printf("result:%#v", result)
	if "111" != result.Code {
		t.Errorf("111 != result.Code: %#v", result.Code)
		return
	}

}

/*
type WrapRedirect struct {
	*Redirect
}

func NewWrapRedirect(result chan RedirectResult) *WrapRedirect {
	return &WrapRedirect{NewRedirect(result)}
}
func (this WrapRedirect) Handler(w http.ResponseWriter, r *http.Request) {
	this.Result <- RedirectResult{Code: "111"}
}
*/

func TestGetAuthCode(t *testing.T) {
	config := &oauth.Config{
		ClientId:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scope:        "",
		AuthURL:      "",
		TokenURL:     "",
		TokenCache:   oauth.CacheFile("cache.json"),
	}
	code, err := getAuthCode(config, LocalServerConfig{0, 0, "test"})
	if err != nil {
		t.Errorf("Error getAuthCode: %#v", err)
		return
	}
	if "200" != code {
		t.Errorf("Error getAuthCode 200 != code :%#v", code)
		return
	}

}

func TestServer(t *testing.T) {
	redirect := NewRedirect(make(chan RedirectResult, 1))
	go redirect.Server(-1)
	<-redirect.ServerStart
	//var result RedirectResult
	//result = <-redirect.Result
	//fmt.Printf("result:%#v", result)
}
