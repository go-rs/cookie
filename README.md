# Cookie
Cookie helper for REST API framework

## How to use?
Need go-rs/crypto package for cookie encryption & decryption.

````
var c crypto.AESGCM
_ = c.Config("0123456789abcdefg", "0123456789abcd0123456789")

api.Use(cookie.Load(c))

// Authenticator
api.Use(func(ctx *rest.Context) {
    helper := ctx.Get("Cookie").(cookie.Cookie) // this is cookie helper object
    fmt.Println(c.Get("authtoken"))
})
````

## Supported methods 
- Get(cookieName string) > return string
- Set(cookie *http.Cookie) > return error if any
- GetSigned(cookieName string) > return string
- SetSigned(cookie *http.Cookie) > return error if any
