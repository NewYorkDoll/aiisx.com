// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package githubevent

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the githubevent type in the database.
	Label = "github_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldEventType holds the string denoting the event_type field in the database.
	FieldEventType = "event_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldActorID holds the string denoting the actor_id field in the database.
	FieldActorID = "actor_id"
	// FieldActor holds the string denoting the actor field in the database.
	FieldActor = "actor"
	// FieldRepoID holds the string denoting the repo_id field in the database.
	FieldRepoID = "repo_id"
	// FieldRepo holds the string denoting the repo field in the database.
	FieldRepo = "repo"
	// FieldPayload holds the string denoting the payload field in the database.
	FieldPayload = "payload"
	// Table holds the table name of the githubevent in the database.
	Table = "github_events"
)

// Columns holds all SQL columns for githubevent fields.
var Columns = []string{
	FieldID,
	FieldEventID,
	FieldEventType,
	FieldCreatedAt,
	FieldPublic,
	FieldActorID,
	FieldActor,
	FieldRepoID,
	FieldRepo,
	FieldPayload,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "aiisx.com/src/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// EventTypeValidator is a validator for the "event_type" field. It is called by the builders before save.
	EventTypeValidator func(string) error
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
	// RepoIDValidator is a validator for the "repo_id" field. It is called by the builders before save.
	RepoIDValidator func(int64) error
)
