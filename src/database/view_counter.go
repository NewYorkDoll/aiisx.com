package database

import (
	"context"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"aiisx.com/src/ent"
	"aiisx.com/src/ent/privacy"
	"github.com/apex/log"
)

type contextKey string

var (
	postViewMu            = &sync.Mutex{}
	postView              = map[int]map[string]time.Time{}
	contextIP  contextKey = "ip"
)

func PostViewCounter(ctx context.Context, post *ent.Post) {
	ip := GetContextIP(ctx)
	if ip == nil {
		return
	}

	nctx, cancel := context.WithTimeout(
		privacy.DecisionContext(context.Background(), privacy.Allow),
		time.Second*5,
	)
	defer cancel()

	postViewMu.Lock()
	if _, ok := postView[post.ID]; !ok {
		postView[post.ID] = map[string]time.Time{}
	}

	if t, ok := postView[post.ID][ip.String()]; ok {
		if time.Since(t) < 5*time.Minute {
			postViewMu.Unlock()
			return
		}
	}
	postView[post.ID][ip.String()] = time.Now()
	postViewMu.Unlock()

	if _, err := post.Update().AddViewCount(1).Save(nctx); err != nil {
		log.FromContext(ctx).WithError(err).WithField("post_id", post.ID).
			Error("failed to update post view count")
	}
}

func GetContextIP(ctx context.Context) net.IP {
	if ip, ok := ctx.Value(contextIP).(net.IP); ok {
		return ip
	}

	return nil
}
func UseContextIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(
			context.WithValue(
				r.Context(),
				contextIP,
				parseIP(sanitizeIP(r.RemoteAddr)),
			),
		))
	})
}

func parseIP(ip string) net.IP {
	parsedIP := net.ParseIP(strings.TrimSpace(ip))

	if parsedIP != nil {
		if v4 := parsedIP.To4(); v4 != nil {
			return v4
		}
	}

	return parsedIP
}
func sanitizeIP(input string) (ip string) {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(input))
	if err != nil || ip == "" {
		ip = input
	}
	return ip
}
