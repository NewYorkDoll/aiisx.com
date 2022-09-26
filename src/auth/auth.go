package auth

import (
	"encoding/hex"
	"net/http"
	"sync"

	"aiisx.com/src/database"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

var (
	// https时设置为true
	DefaultCookieSecure = false

	// cookie过期时间
	DefaulltCookieMaxAge = 30 * 86400

	gothInit sync.Once
)

type AuthHandler[Ident any, ID comparable] struct {
	Auth         database.AuthService[Ident, ID]
	Ident        *Ident
	ID           *ID
	router       http.Handler
	errorHandler func(w http.ResponseWriter, r *http.Request, err error) (ok bool)
}

func NewAuthHandler[Ident any, ID comparable](auth database.AuthService[Ident, ID], authKey, encryptKey string) *AuthHandler[Ident, ID] {
	authKeyBytes, err := hex.DecodeString(authKey)
	if err != nil {
		panic(err)
	}
	encryptKeyBytes, err := hex.DecodeString(encryptKey)
	if err != nil {
		panic(err)
	}

	gothInit.Do(func() {
		authStore := sessions.NewCookieStore(authKeyBytes, encryptKeyBytes)
		authStore.MaxAge(DefaulltCookieMaxAge)
		authStore.Options.Path = "/"
		authStore.Options.HttpOnly = true
		authStore.Options.Secure = DefaultCookieSecure
		gothic.Store = authStore
	})

	h := &AuthHandler[Ident, ID]{
		Auth:  auth,
		Ident: new(Ident),
		ID:    new(ID),
		errorHandler: func(w http.ResponseWriter, r *http.Request, err error) (ok bool) {
			return true
		},
	}
	return h
}
