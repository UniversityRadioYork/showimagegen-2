/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package myradio

import (
	"net/http"
	"net/url"
)

type cookieJar struct {
	cookies map[string][]*http.Cookie
}

func (j cookieJar) Cookies(u *url.URL) []*http.Cookie {
	return j.cookies[u.Host]
}

func (j cookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.cookies[u.Host] = cookies
}
