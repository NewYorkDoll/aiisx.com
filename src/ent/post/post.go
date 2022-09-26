// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package post

import (
	"time"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldContentHTML holds the string denoting the content_html field in the database.
	FieldContentHTML = "content_html"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldPublishedAt holds the string denoting the published_at field in the database.
	FieldPublishedAt = "published_at"
	// FieldViewCount holds the string denoting the view_count field in the database.
	FieldViewCount = "view_count"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeLabels holds the string denoting the labels edge name in mutations.
	EdgeLabels = "labels"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "posts"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_posts"
	// LabelsTable is the table that holds the labels relation/edge. The primary key declared below.
	LabelsTable = "label_posts"
	// LabelsInverseTable is the table name for the Label entity.
	// It exists in this package in order to avoid circular dependency with the "label" package.
	LabelsInverseTable = "labels"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldSlug,
	FieldTitle,
	FieldContent,
	FieldContentHTML,
	FieldSummary,
	FieldPublishedAt,
	FieldViewCount,
	FieldPublic,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_posts",
}

var (
	// LabelsPrimaryKey and LabelsColumn2 are the table columns denoting the
	// primary key for the labels relation (M2M).
	LabelsPrimaryKey = []string{"label_id", "post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// ContentHTMLValidator is a validator for the "content_html" field. It is called by the builders before save.
	ContentHTMLValidator func(string) error
	// SummaryValidator is a validator for the "summary" field. It is called by the builders before save.
	SummaryValidator func(string) error
	// DefaultPublishedAt holds the default value on creation for the "published_at" field.
	DefaultPublishedAt func() time.Time
	// DefaultViewCount holds the default value on creation for the "view_count" field.
	DefaultViewCount int
	// ViewCountValidator is a validator for the "view_count" field. It is called by the builders before save.
	ViewCountValidator func(int) error
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
)
