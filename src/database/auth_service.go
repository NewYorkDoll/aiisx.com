package database

import (
	"context"
	"errors"
	"strconv"

	"aiisx.com/src/ent"
	"aiisx.com/src/ent/user"
	"ariga.io/entcache"
	"github.com/markbates/goth"
)

func NewAuthService(db *ent.Client, admin int) *authService {
	return &authService{
		db:    db,
		admin: admin,
	}
}

type authService struct {
	db    *ent.Client
	admin int
}

type AuthService[Ident ent.User, ID comparable] interface {
	Get(context.Context, ID) (*ent.User, error)
	Set(context.Context, *goth.User) (ID, error)
	Roles(context.Context, ID) ([]string, error)
}

var _ AuthService[ent.User, int] = (*authService)(nil)

func (s *authService) Get(ctx context.Context, id int) (*ent.User, error) {
	return s.db.User.Get(ctx, id)
}

func (s *authService) Set(ctx context.Context, guser *goth.User) (id int, err error) {
	uid, err := strconv.Atoi(guser.UserID)
	if err != nil {
		return 0, err
	}

	if uid != s.admin {
		return 0, errors.New("access denied")
	}

	q := s.db.User.Create().
		SetUserID(uid).
		SetName(guser.Name).
		SetLogin(guser.NickName).
		SetEmail(guser.Email).
		SetAvatarURL(guser.AvatarURL).
		SetLocation(guser.Location).
		SetBio(guser.Description)

	if data, ok := guser.RawData["html_url"]; ok {
		if url, ok := data.(string); ok {
			q = q.SetHTMLURL(url)
		}
	}

	return q.OnConflictColumns(user.FieldUserID).Ignore().UpdateNewValues().ID(entcache.Evict(ctx))
}

func (s *authService) Roles(ctx context.Context, id int) ([]string, error) {
	if ok, _ := s.db.User.Query().Where(user.IDEQ(id)).Exist(ctx); ok {
		return []string{"admin"}, nil
	}
	return nil, nil
}
