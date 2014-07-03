package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"code.google.com/p/goauth2/oauth"
)

func TestRedirectHandler(t *testing.T) {
	redirect := NewRedirect(make(chan RedirectResult, 1))
	ts := httptest.NewServer(http.HandlerFunc(redirect.GetCode))
	defer redirect.Stop()
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

type WrapRedirect struct {
	*Redirect
}

func NewWrapRedirect(result chan RedirectResult) *WrapRedirect {
	return &WrapRedirect{NewRedirect(result)}
}
func (this WrapRedirect) Handler(w http.ResponseWriter, r *http.Request) {
	this.Result <- RedirectResult{Code: "111"}
}

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
	code, err := getAuthCode(config, LocalServerConfig{20343, 1, "test"})
	if err == nil {
		t.Errorf("Error getAuthCode: %#v", err)
		return
	}
	if err.Error() != "リダイレクト待ち時間がタイムアウトしました" {
		t.Errorf("Error getAuthCode: %#v", err)
		return
	}
	if code != "" {
		t.Errorf("getAuthCode: %#v", code)
		return
	}
}

func TestServer(t *testing.T) {
	redirect := NewRedirect(make(chan RedirectResult, 1))
	go redirect.Server(2000)
	<-redirect.ServerStart
	redirect.Stop()
	var result RedirectResult
	result = <-redirect.Result
	if result.Err != nil {
		t.Errorf("Error Server: %#v", result)
	}
}

func TestServerError(t *testing.T) {
	redirect := NewRedirect(make(chan RedirectResult, 1))
	go redirect.Server(-1)
	var result RedirectResult
	result = <-redirect.Result
	if result.Err == nil {
		t.Errorf("Error Server: %#v", result.Err)
	}
}
