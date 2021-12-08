package tool

import (
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
	body, err := Get("https://raw.githubusercontent.com/http-live-streaming/github.com/q-meet/m3u8Tool/develop/README.md")
	if err != nil {
		t.Error(err)
	}
	defer body.Close()
	_, err = ioutil.ReadAll(body)
	if err != nil {
		t.Error(err)
	}
}
