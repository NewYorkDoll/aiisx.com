package gh

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/apex/log"
	"github.com/google/go-github/v47/github"
)

// / 获取用户信息的轮询间隔
const userInterval = 30 * time.Minute

var User atomic.Pointer[github.User]

func UserRunner(ctx context.Context, username string) error {
	logger := log.FromContext(ctx).WithField("runner", "github_user")
	getUser(ctx, logger, username)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(userInterval):
			getUser(ctx, logger, username)
		}
	}
}
func getUser(ctx context.Context, logger log.Interface, username string) {
	user, _, err := Client.Users.Get(ctx, username)
	if err != nil {
		logger.WithError(err).Error("failed to get user")
		return
	}
	logger.WithField("user", user.GetLogin()).Info("got user")
	User.Store(user)
}
