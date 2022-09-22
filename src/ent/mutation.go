// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"aiisx.com/src/ent/githubevent"
	"aiisx.com/src/ent/predicate"
	"github.com/google/go-github/v47/github"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGithubEvent = "GithubEvent"
)

// GithubEventMutation represents an operation that mutates the GithubEvent nodes in the graph.
type GithubEventMutation struct {
	config
	op            Op
	typ           string
	id            *int
	event_id      *string
	event_type    *string
	created_at    *time.Time
	public        *bool
	actor_id      *int64
	addactor_id   *int64
	actor         **github.User
	repo_id       *int64
	addrepo_id    *int64
	repo          **github.Repository
	payload       *map[string]interface{}
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*GithubEvent, error)
	predicates    []predicate.GithubEvent
}

var _ ent.Mutation = (*GithubEventMutation)(nil)

// githubeventOption allows management of the mutation configuration using functional options.
type githubeventOption func(*GithubEventMutation)

// newGithubEventMutation creates new mutation for the GithubEvent entity.
func newGithubEventMutation(c config, op Op, opts ...githubeventOption) *GithubEventMutation {
	m := &GithubEventMutation{
		config:        c,
		op:            op,
		typ:           TypeGithubEvent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGithubEventID sets the ID field of the mutation.
func withGithubEventID(id int) githubeventOption {
	return func(m *GithubEventMutation) {
		var (
			err   error
			once  sync.Once
			value *GithubEvent
		)
		m.oldValue = func(ctx context.Context) (*GithubEvent, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().GithubEvent.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGithubEvent sets the old GithubEvent of the mutation.
func withGithubEvent(node *GithubEvent) githubeventOption {
	return func(m *GithubEventMutation) {
		m.oldValue = func(context.Context) (*GithubEvent, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GithubEventMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GithubEventMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GithubEventMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GithubEventMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().GithubEvent.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEventID sets the "event_id" field.
func (m *GithubEventMutation) SetEventID(s string) {
	m.event_id = &s
}

// EventID returns the value of the "event_id" field in the mutation.
func (m *GithubEventMutation) EventID() (r string, exists bool) {
	v := m.event_id
	if v == nil {
		return
	}
	return *v, true
}

// OldEventID returns the old "event_id" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldEventID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEventID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEventID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEventID: %w", err)
	}
	return oldValue.EventID, nil
}

// ResetEventID resets all changes to the "event_id" field.
func (m *GithubEventMutation) ResetEventID() {
	m.event_id = nil
}

// SetEventType sets the "event_type" field.
func (m *GithubEventMutation) SetEventType(s string) {
	m.event_type = &s
}

// EventType returns the value of the "event_type" field in the mutation.
func (m *GithubEventMutation) EventType() (r string, exists bool) {
	v := m.event_type
	if v == nil {
		return
	}
	return *v, true
}

// OldEventType returns the old "event_type" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldEventType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEventType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEventType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEventType: %w", err)
	}
	return oldValue.EventType, nil
}

// ResetEventType resets all changes to the "event_type" field.
func (m *GithubEventMutation) ResetEventType() {
	m.event_type = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *GithubEventMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *GithubEventMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *GithubEventMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetPublic sets the "public" field.
func (m *GithubEventMutation) SetPublic(b bool) {
	m.public = &b
}

// Public returns the value of the "public" field in the mutation.
func (m *GithubEventMutation) Public() (r bool, exists bool) {
	v := m.public
	if v == nil {
		return
	}
	return *v, true
}

// OldPublic returns the old "public" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldPublic(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPublic is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPublic requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPublic: %w", err)
	}
	return oldValue.Public, nil
}

// ResetPublic resets all changes to the "public" field.
func (m *GithubEventMutation) ResetPublic() {
	m.public = nil
}

// SetActorID sets the "actor_id" field.
func (m *GithubEventMutation) SetActorID(i int64) {
	m.actor_id = &i
	m.addactor_id = nil
}

// ActorID returns the value of the "actor_id" field in the mutation.
func (m *GithubEventMutation) ActorID() (r int64, exists bool) {
	v := m.actor_id
	if v == nil {
		return
	}
	return *v, true
}

// OldActorID returns the old "actor_id" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldActorID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActorID: %w", err)
	}
	return oldValue.ActorID, nil
}

// AddActorID adds i to the "actor_id" field.
func (m *GithubEventMutation) AddActorID(i int64) {
	if m.addactor_id != nil {
		*m.addactor_id += i
	} else {
		m.addactor_id = &i
	}
}

// AddedActorID returns the value that was added to the "actor_id" field in this mutation.
func (m *GithubEventMutation) AddedActorID() (r int64, exists bool) {
	v := m.addactor_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetActorID resets all changes to the "actor_id" field.
func (m *GithubEventMutation) ResetActorID() {
	m.actor_id = nil
	m.addactor_id = nil
}

// SetActor sets the "actor" field.
func (m *GithubEventMutation) SetActor(gi *github.User) {
	m.actor = &gi
}

// Actor returns the value of the "actor" field in the mutation.
func (m *GithubEventMutation) Actor() (r *github.User, exists bool) {
	v := m.actor
	if v == nil {
		return
	}
	return *v, true
}

// OldActor returns the old "actor" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldActor(ctx context.Context) (v *github.User, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActor is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActor requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActor: %w", err)
	}
	return oldValue.Actor, nil
}

// ResetActor resets all changes to the "actor" field.
func (m *GithubEventMutation) ResetActor() {
	m.actor = nil
}

// SetRepoID sets the "repo_id" field.
func (m *GithubEventMutation) SetRepoID(i int64) {
	m.repo_id = &i
	m.addrepo_id = nil
}

// RepoID returns the value of the "repo_id" field in the mutation.
func (m *GithubEventMutation) RepoID() (r int64, exists bool) {
	v := m.repo_id
	if v == nil {
		return
	}
	return *v, true
}

// OldRepoID returns the old "repo_id" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldRepoID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRepoID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRepoID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRepoID: %w", err)
	}
	return oldValue.RepoID, nil
}

// AddRepoID adds i to the "repo_id" field.
func (m *GithubEventMutation) AddRepoID(i int64) {
	if m.addrepo_id != nil {
		*m.addrepo_id += i
	} else {
		m.addrepo_id = &i
	}
}

// AddedRepoID returns the value that was added to the "repo_id" field in this mutation.
func (m *GithubEventMutation) AddedRepoID() (r int64, exists bool) {
	v := m.addrepo_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetRepoID resets all changes to the "repo_id" field.
func (m *GithubEventMutation) ResetRepoID() {
	m.repo_id = nil
	m.addrepo_id = nil
}

// SetRepo sets the "repo" field.
func (m *GithubEventMutation) SetRepo(gi *github.Repository) {
	m.repo = &gi
}

// Repo returns the value of the "repo" field in the mutation.
func (m *GithubEventMutation) Repo() (r *github.Repository, exists bool) {
	v := m.repo
	if v == nil {
		return
	}
	return *v, true
}

// OldRepo returns the old "repo" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldRepo(ctx context.Context) (v *github.Repository, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRepo is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRepo requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRepo: %w", err)
	}
	return oldValue.Repo, nil
}

// ResetRepo resets all changes to the "repo" field.
func (m *GithubEventMutation) ResetRepo() {
	m.repo = nil
}

// SetPayload sets the "payload" field.
func (m *GithubEventMutation) SetPayload(value map[string]interface{}) {
	m.payload = &value
}

// Payload returns the value of the "payload" field in the mutation.
func (m *GithubEventMutation) Payload() (r map[string]interface{}, exists bool) {
	v := m.payload
	if v == nil {
		return
	}
	return *v, true
}

// OldPayload returns the old "payload" field's value of the GithubEvent entity.
// If the GithubEvent object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GithubEventMutation) OldPayload(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPayload is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPayload requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPayload: %w", err)
	}
	return oldValue.Payload, nil
}

// ResetPayload resets all changes to the "payload" field.
func (m *GithubEventMutation) ResetPayload() {
	m.payload = nil
}

// Where appends a list predicates to the GithubEventMutation builder.
func (m *GithubEventMutation) Where(ps ...predicate.GithubEvent) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *GithubEventMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (GithubEvent).
func (m *GithubEventMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GithubEventMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.event_id != nil {
		fields = append(fields, githubevent.FieldEventID)
	}
	if m.event_type != nil {
		fields = append(fields, githubevent.FieldEventType)
	}
	if m.created_at != nil {
		fields = append(fields, githubevent.FieldCreatedAt)
	}
	if m.public != nil {
		fields = append(fields, githubevent.FieldPublic)
	}
	if m.actor_id != nil {
		fields = append(fields, githubevent.FieldActorID)
	}
	if m.actor != nil {
		fields = append(fields, githubevent.FieldActor)
	}
	if m.repo_id != nil {
		fields = append(fields, githubevent.FieldRepoID)
	}
	if m.repo != nil {
		fields = append(fields, githubevent.FieldRepo)
	}
	if m.payload != nil {
		fields = append(fields, githubevent.FieldPayload)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GithubEventMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case githubevent.FieldEventID:
		return m.EventID()
	case githubevent.FieldEventType:
		return m.EventType()
	case githubevent.FieldCreatedAt:
		return m.CreatedAt()
	case githubevent.FieldPublic:
		return m.Public()
	case githubevent.FieldActorID:
		return m.ActorID()
	case githubevent.FieldActor:
		return m.Actor()
	case githubevent.FieldRepoID:
		return m.RepoID()
	case githubevent.FieldRepo:
		return m.Repo()
	case githubevent.FieldPayload:
		return m.Payload()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GithubEventMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case githubevent.FieldEventID:
		return m.OldEventID(ctx)
	case githubevent.FieldEventType:
		return m.OldEventType(ctx)
	case githubevent.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case githubevent.FieldPublic:
		return m.OldPublic(ctx)
	case githubevent.FieldActorID:
		return m.OldActorID(ctx)
	case githubevent.FieldActor:
		return m.OldActor(ctx)
	case githubevent.FieldRepoID:
		return m.OldRepoID(ctx)
	case githubevent.FieldRepo:
		return m.OldRepo(ctx)
	case githubevent.FieldPayload:
		return m.OldPayload(ctx)
	}
	return nil, fmt.Errorf("unknown GithubEvent field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GithubEventMutation) SetField(name string, value ent.Value) error {
	switch name {
	case githubevent.FieldEventID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEventID(v)
		return nil
	case githubevent.FieldEventType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEventType(v)
		return nil
	case githubevent.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case githubevent.FieldPublic:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPublic(v)
		return nil
	case githubevent.FieldActorID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActorID(v)
		return nil
	case githubevent.FieldActor:
		v, ok := value.(*github.User)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActor(v)
		return nil
	case githubevent.FieldRepoID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRepoID(v)
		return nil
	case githubevent.FieldRepo:
		v, ok := value.(*github.Repository)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRepo(v)
		return nil
	case githubevent.FieldPayload:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPayload(v)
		return nil
	}
	return fmt.Errorf("unknown GithubEvent field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GithubEventMutation) AddedFields() []string {
	var fields []string
	if m.addactor_id != nil {
		fields = append(fields, githubevent.FieldActorID)
	}
	if m.addrepo_id != nil {
		fields = append(fields, githubevent.FieldRepoID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GithubEventMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case githubevent.FieldActorID:
		return m.AddedActorID()
	case githubevent.FieldRepoID:
		return m.AddedRepoID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GithubEventMutation) AddField(name string, value ent.Value) error {
	switch name {
	case githubevent.FieldActorID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddActorID(v)
		return nil
	case githubevent.FieldRepoID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRepoID(v)
		return nil
	}
	return fmt.Errorf("unknown GithubEvent numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GithubEventMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GithubEventMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GithubEventMutation) ClearField(name string) error {
	return fmt.Errorf("unknown GithubEvent nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GithubEventMutation) ResetField(name string) error {
	switch name {
	case githubevent.FieldEventID:
		m.ResetEventID()
		return nil
	case githubevent.FieldEventType:
		m.ResetEventType()
		return nil
	case githubevent.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case githubevent.FieldPublic:
		m.ResetPublic()
		return nil
	case githubevent.FieldActorID:
		m.ResetActorID()
		return nil
	case githubevent.FieldActor:
		m.ResetActor()
		return nil
	case githubevent.FieldRepoID:
		m.ResetRepoID()
		return nil
	case githubevent.FieldRepo:
		m.ResetRepo()
		return nil
	case githubevent.FieldPayload:
		m.ResetPayload()
		return nil
	}
	return fmt.Errorf("unknown GithubEvent field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GithubEventMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GithubEventMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GithubEventMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GithubEventMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GithubEventMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GithubEventMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GithubEventMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown GithubEvent unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GithubEventMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown GithubEvent edge %s", name)
}
