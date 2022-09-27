package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"aiisx.com/src/database"
	"aiisx.com/src/ent"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

type contextKey string

var (
	// https时设置为true
	DefaultCookieSecure = false

	// cookie过期时间
	DefaulltCookieMaxAge = 30 * 86400

	gothInit sync.Once

	User atomic.Pointer[ent.User]
)

const (
	contextDebug       contextKey = "debug"
	contextAuth        contextKey = "auth"
	contextAuthID      contextKey = "auth_id"
	contextAuthRoles   contextKey = "auth_roles"
	contextNextURL     contextKey = "next_url"
	contextSkipNextURL contextKey = "skip_next_url"
	contextIP          contextKey = "ip"

	authSessionKey = "_auth"
	nextSessionKey = "_next"
)

type AuthHandler[Ident ent.User, ID comparable] struct {
	Auth         database.AuthService[Ident, ID]
	Ident        *Ident
	ID           *ID
	errorHandler func(w http.ResponseWriter, r *http.Request, err error) (ok bool)
}

func NewAuthHandler[Ident ent.User, ID comparable](auth database.AuthService[Ident, ID], authKey, encryptKey []byte) *AuthHandler[Ident, ID] {

	gothInit.Do(func() {
		authStore := sessions.NewCookieStore(authKey, encryptKey)
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

func (h *AuthHandler[Ident, ID]) AddToContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := User.Load()
		if u != nil {
			ctx.Next()
			return
		}
		r := ctx.Request.Context()
		id := h.getIDFromSession(ctx.Request)
		if id == nil {
			ctx.Next()
			return
		}

		ident, err := h.Auth.Get(r, *id)
		if err != nil {
			log.FromContext(r).WithError(err).WithField("user_id", *id).Warn("failed to get ident from session (but id set)")
			ctx.Next()
			return
		}
		User.Store(ident)
		ctx.Next()
	}
}

func (h *AuthHandler[Ident, ID]) Logout() {
	User.Store(nil)
}
func (h *AuthHandler[Ident, ID]) getIDFromSession(r *http.Request) *ID {
	key, _ := gothic.GetFromSession(authSessionKey, r)
	if key == "" {
		return nil
	}

	id, err := h.convertID(key)
	if err != nil {
		return nil
	}
	return &id
}

func (h *AuthHandler[Ident, ID]) convertID(in string) (ID, error) {
	var v any
	var err error

	switch any(h.ID).(type) {
	case *string:
		v = in
	case *int:
		v, err = strconv.Atoi(in)
	case *int64:
		v, err = strconv.ParseInt(in, 10, 64)
	case *float64:
		v, err = strconv.ParseFloat(in, 64)
	case *uint:
		v, err = strconv.ParseUint(in, 10, 64)
	case *uint16:
		v, err = strconv.ParseUint(in, 10, 16)
	case *uint32:
		v, err = strconv.ParseUint(in, 10, 32)
	case *uint64:
		v, err = strconv.ParseUint(in, 10, 64)
	default:
		panic("unsupported ID type")
	}
	if err != nil {
		return *new(ID), err
	}
	return v.(ID), nil
}

func IdentFromContext[Ident any](ctx context.Context) (auth *Ident) {
	auth, _ = ctx.Value(contextAuth).(*Ident)
	return auth
}

func (h *AuthHandler[Ident, ID]) RolesFromContext(ctx context.Context) (roles AuthRoles) {
	return RolesFromContext(ctx)
}

type AuthRoles []string

func RolesFromContext(ctx context.Context) (roles AuthRoles) {
	roles, _ = ctx.Value(contextAuthRoles).([]string)
	return AuthRoles(roles)
}

func (r AuthRoles) Has(role string) bool {
	if len(r) == 0 {
		return false
	}

	for _, r := range r {
		if strings.EqualFold(r, role) {
			return true
		}
	}

	return false
}
