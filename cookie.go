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

var algo crypto.AESGCM

type Cookie struct {
	ctx *rest.Context
}

func (c *Cookie) Set(cookie *http.Cookie) (err error) {
	http.SetCookie(c.ctx.Response, cookie)
	return
}

func (c *Cookie) SetSigned(cookie *http.Cookie) (err error) {
	cookie.Value, err = algo.Encrypt(cookie.Value)
	if err != nil {
		return
	}
	http.SetCookie(c.ctx.Response, cookie)
	return
}

func (c *Cookie) Get(name string) (val string) {
	cookie, err := c.ctx.Request.Cookie(name)
	if err != nil {
		return
	}
	return cookie.Value
}

func (c *Cookie) GetSigned(name string) (val string) {
	cookie, err := c.ctx.Request.Cookie(name)
	if err != nil {
		return
	}
	val, err = algo.Decrypt(cookie.Value)
	if err != nil {
		return
	}
	return
}

func Load(x crypto.AESGCM) rest.Handler {
	algo = x
	return func(ctx *rest.Context) {
		ctx.Set("Cookie", Cookie{ctx})
	}
}
