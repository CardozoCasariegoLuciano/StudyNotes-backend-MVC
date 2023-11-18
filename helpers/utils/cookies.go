package utils

import (
	"net/http"
	"time"
)

var CookieName = "Authorization"

func DeleteCookie() *http.Cookie {
	cookie := http.Cookie{
		Name:    CookieName,
		Path:    BasePath,
		Expires: time.Now().AddDate(0, 0, -1),
	}

	return &cookie
}

func CreateCookie(t string) *http.Cookie {
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    t,
		Path:     BasePath,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
	}

	return &cookie
}
