package schema

import (
	"regexp"

	"entgo.io/ent"
)

var reLabel = regexp.MustCompile(`^[a-z\d][a-z\d-]*$`)

type Label struct {
	ent.Schema
}
