// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package ent

import (
	"time"
)

// CreateLabelInput represents a mutation input for creating labels.
type CreateLabelInput struct {
	CreateTime          *time.Time
	UpdateTime          *time.Time
	Name                string
	PostIDs             []int
	GithubRepositoryIDs []int
}

// Mutate applies the CreateLabelInput on the LabelMutation builder.
func (i *CreateLabelInput) Mutate(m *LabelMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetName(i.Name)
	if v := i.PostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.GithubRepositoryIDs; len(v) > 0 {
		m.AddGithubRepositoryIDs(v...)
	}
}

// SetInput applies the change-set in the CreateLabelInput on the LabelCreate builder.
func (c *LabelCreate) SetInput(i CreateLabelInput) *LabelCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateLabelInput represents a mutation input for updating labels.
type UpdateLabelInput struct {
	UpdateTime                *time.Time
	Name                      *string
	AddPostIDs                []int
	RemovePostIDs             []int
	AddGithubRepositoryIDs    []int
	RemoveGithubRepositoryIDs []int
}

// Mutate applies the UpdateLabelInput on the LabelMutation builder.
func (i *UpdateLabelInput) Mutate(m *LabelMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddPostIDs; len(v) > 0 {
		m.AddPostIDs(v...)
	}
	if v := i.RemovePostIDs; len(v) > 0 {
		m.RemovePostIDs(v...)
	}
	if v := i.AddGithubRepositoryIDs; len(v) > 0 {
		m.AddGithubRepositoryIDs(v...)
	}
	if v := i.RemoveGithubRepositoryIDs; len(v) > 0 {
		m.RemoveGithubRepositoryIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateLabelInput on the LabelUpdate builder.
func (c *LabelUpdate) SetInput(i UpdateLabelInput) *LabelUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateLabelInput on the LabelUpdateOne builder.
func (c *LabelUpdateOne) SetInput(i UpdateLabelInput) *LabelUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePostInput represents a mutation input for creating posts.
type CreatePostInput struct {
	CreateTime  *time.Time
	UpdateTime  *time.Time
	Slug        string
	Title       string
	Content     string
	PublishedAt *time.Time
	Public      *bool
	LabelIDs    []int
}

// Mutate applies the CreatePostInput on the PostMutation builder.
func (i *CreatePostInput) Mutate(m *PostMutation) {
	if v := i.CreateTime; v != nil {
		m.SetCreateTime(*v)
	}
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	m.SetSlug(i.Slug)
	m.SetTitle(i.Title)
	m.SetContent(i.Content)
	if v := i.PublishedAt; v != nil {
		m.SetPublishedAt(*v)
	}
	if v := i.Public; v != nil {
		m.SetPublic(*v)
	}
	if v := i.LabelIDs; len(v) > 0 {
		m.AddLabelIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePostInput on the PostCreate builder.
func (c *PostCreate) SetInput(i CreatePostInput) *PostCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePostInput represents a mutation input for updating posts.
type UpdatePostInput struct {
	UpdateTime     *time.Time
	Slug           *string
	Title          *string
	Content        *string
	PublishedAt    *time.Time
	Public         *bool
	AddLabelIDs    []int
	RemoveLabelIDs []int
}

// Mutate applies the UpdatePostInput on the PostMutation builder.
func (i *UpdatePostInput) Mutate(m *PostMutation) {
	if v := i.UpdateTime; v != nil {
		m.SetUpdateTime(*v)
	}
	if v := i.Slug; v != nil {
		m.SetSlug(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Content; v != nil {
		m.SetContent(*v)
	}
	if v := i.PublishedAt; v != nil {
		m.SetPublishedAt(*v)
	}
	if v := i.Public; v != nil {
		m.SetPublic(*v)
	}
	if v := i.AddLabelIDs; len(v) > 0 {
		m.AddLabelIDs(v...)
	}
	if v := i.RemoveLabelIDs; len(v) > 0 {
		m.RemoveLabelIDs(v...)
	}
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdate builder.
func (c *PostUpdate) SetInput(i UpdatePostInput) *PostUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePostInput on the PostUpdateOne builder.
func (c *PostUpdateOne) SetInput(i UpdatePostInput) *PostUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
