package goutil

import (
	"net/http"
	"io/ioutil"
)

type HttpResponseCallback func(b []byte, r *http.Response)

func MakeHttpGet(url string, callback HttpResponseCallback) {
	resp, err := http.Get(url)
	PanicIfErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	PanicIfErr(err)
	callback(body, resp)
}
