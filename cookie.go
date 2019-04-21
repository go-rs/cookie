/*!
 * go-rs/cookie
 * Copyright(c) 2019 Roshan Gade
 * MIT Licensed
 */
package cookie

import (
	"net/http"

	"github.com/go-rs/crypto"
	"github.com/go-rs/rest-api-framework"
)

var x crypto.AESGCM

type Cookie struct {
	ctx *rest.Context
}

func (c *Cookie) Set(cookie *http.Cookie) {
	http.SetCookie(c.ctx.Response, cookie)
}

func (c *Cookie) SetSigned(cookie *http.Cookie) error {
	var err error
	cookie.Value, err = x.Encrypt(cookie.Value)
	if err != nil {
		return err
	}
	http.SetCookie(c.ctx.Response, cookie)
	return err
}

func (c *Cookie) Get(name string) string {
	cookie, err := c.ctx.Request.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (c *Cookie) GetSigned(name string) string {
	cookie, err := c.ctx.Request.Cookie(name)
	if err != nil {
		return ""
	}
	decrypted, err := x.Decrypt(cookie.Value)
	if err != nil {
		return ""
	}
	return decrypted
}

func Load(cr crypto.AESGCM) rest.Handler {
	x = cr
	return func(ctx *rest.Context) {
		ctx.Set("Cookie", Cookie{ctx})
	}
}
