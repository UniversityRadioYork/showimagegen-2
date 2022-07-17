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

// CtxKey is a `string` type for keys in the context
type CtxKey string

const (
	// CtxKeyTimeoutSeconds is the key for an integer which is the number
	// of seconds given for a request to timeout
	CtxKeyTimeoutSeconds CtxKey = "timeoutSeconds"

	// CtxKeyXSRFToken is the context key for the MyRadio XSRF token in
	// the session
	CtxKeyXSRFToken CtxKey = "xsrfToken"

	// CtxKeyMyRadioUsername is the key for the username to authenticate
	// the MyRadio session with.
	CtxKeyMyRadioUsername CtxKey = "myRadioUsername"

	// CtxKeyMyRadioPassword is the key for the password to authenticate
	// the MyRadio session with.
	CtxKeyMyRadioPassword CtxKey = "myRadioPassword"
)

// LoginSession provides a login session for MyRadio after authenticating
// with a MyRadio username and password
type LoginSession struct {
	client    *http.Client
	xsrfToken string
	timeout   time.Duration
}

// Close will log out the MyRadio session
func (e *LoginSession) Close() {
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

// CreateMyRadioLoginSession gets an XSRF token from MyRadio, and will
// then log in to MyRadio, returning the LoginSession
func CreateMyRadioLoginSession(ctx context.Context) (*LoginSession, error) {
	myr := LoginSession{
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

func (e *LoginSession) login(ctx context.Context) error {
	ctxUser := ctx.Value(CtxKeyMyRadioUsername)
	username, ok := ctxUser.(string)
	if !ok {
		return fmt.Errorf("%v can't be used as a string myradio username", ctxUser)
	}

	ctxPass := ctx.Value(CtxKeyMyRadioPassword)
	password, ok := ctxPass.(string)
	if !ok {
		return fmt.Errorf("%v can't be used as a string myradio password", ctxPass)
	}

	form := url.Values{
		"myradio_login-user":         []string{username},
		"myradio_login-password":     []string{password},
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
