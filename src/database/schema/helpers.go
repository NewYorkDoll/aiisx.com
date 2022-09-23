package schema

import (
	"context"

	"aiisx.com/src/ent/privacy"
	"entgo.io/ent/entql"
)

func AllowPublicOnly() privacy.QueryRule {
	type PublicFilter interface {
		WherePublic(p entql.BoolP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		public, ok := f.(PublicFilter)
		if !ok {
			return privacy.Denyf("missing public field in filter")
		}

		public.WherePublic(entql.BoolEQ(true))
		return privacy.Skip
	})
}
