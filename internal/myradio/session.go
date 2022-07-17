/**
URY Show Image Generator 2

Author: Michael Grace <michael.grace@ury.org.uk>
*/

package myradio

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CtxKey string

const (
	CtxKeyTimeoutSeconds CtxKey = "timeoutSeconds"
	CtxKeyXSRFToken      CtxKey = "xsrfToken"
)

type MyRadioLoginEnvironment struct {
	client    *http.Client
	xsrfToken string
	timeout   time.Duration
}

func (e *MyRadioLoginEnvironment) Close() {
	ctx, cnl := context.WithTimeout(context.Background(), e.timeout)
	defer cnl()

	form := url.Values{
		"myradio_logout-__xsrftoken": []string{e.xsrfToken},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://ury.org.uk/myradio/MyRadio/logout", strings.NewReader(form.Encode()))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	e.client.Do(req)
}

func CreateMyRadioLoginEnvironment(ctx context.Context) (*MyRadioLoginEnvironment, error) {
	myr := MyRadioLoginEnvironment{
		client: http.DefaultClient,
	}

	myr.client.Jar = cookieJar{
		cookies: make(map[string][]*http.Cookie),
	}

	timeoutSeconds, ok := ctx.Value(CtxKeyTimeoutSeconds).(int)
	if !ok {
		timeoutSeconds = 5
	}

	myr.timeout = time.Duration(timeoutSeconds) * time.Second

	ctx, cnl := context.WithTimeout(ctx, time.Duration(2)*myr.timeout)
	defer cnl()

	var err error
	myr.xsrfToken, err = myr.getXSRFTokenFromMyRadio(ctx)
	if err != nil {
		return nil, err
	}

	if err := myr.login(ctx); err != nil {
		return nil, err
	}

	return &myr, nil
}

func (e *MyRadioLoginEnvironment) login(ctx context.Context) error {
	form := url.Values{
		"myradio_login-user":         []string{"***"},
		"myradio_login-password":     []string{"***"},
		"myradio_login-next":         []string{"/myradio"},
		"myradio_login-__xsrf-token": []string{e.xsrfToken},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://ury.org.uk/myradio/MyRadio/login", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := e.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("not OK: %v", res.Status)
	}

	return nil

}
