/**
 * Created by 93201 on 2017/7/21.
 */
package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client *http.Client

func init() {
	if client == nil {
		client = &http.Client{}
	}
}

func HttpDo(method string, add string, params url.Values) ([]byte, error) {
	if method == "" {
		method = http.MethodGet
	}
	req, e := http.NewRequest(method, add, strings.NewReader(params.Encode()))
	if e != nil {
		return nil, e
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 QQ/8.9.1.20437 Safari/537.36")
	resp, e := client.Do(req)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

}

func HttpGet(u string) ([]byte, error) {
	resp, e := client.Get(u)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
