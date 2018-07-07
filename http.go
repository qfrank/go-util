package goutil

import (
	"net/http"
	"io/ioutil"
)

type ResponseCallback func(b []byte, r *http.Response)

func GetOrPanic(url string, callback ResponseCallback) {
	resp, err := http.Get(url)
	PanicIfErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	PanicIfErr(err)
	callback(body, resp)
}
