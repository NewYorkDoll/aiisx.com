// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GithubEventsColumns holds the columns for the "github_events" table.
	GithubEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_id", Type: field.TypeString, Unique: true},
		{Name: "event_type", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "public", Type: field.TypeBool, Default: false},
		{Name: "actor_id", Type: field.TypeInt64},
		{Name: "actor", Type: field.TypeJSON},
		{Name: "repo_id", Type: field.TypeInt64},
		{Name: "repo", Type: field.TypeJSON},
		{Name: "payload", Type: field.TypeJSON},
	}
	// GithubEventsTable holds the schema information for the "github_events" table.
	GithubEventsTable = &schema.Table{
		Name:       "github_events",
		Columns:    GithubEventsColumns,
		PrimaryKey: []*schema.Column{GithubEventsColumns[0]},
	}
	// GithubRepositoriesColumns holds the columns for the "github_repositories" table.
	GithubRepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "repo_id", Type: field.TypeInt64, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "full_name", Type: field.TypeString},
		{Name: "owner_login", Type: field.TypeString},
		{Name: "owner", Type: field.TypeJSON},
		{Name: "public", Type: field.TypeBool, Default: false},
		{Name: "html_url", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "fork", Type: field.TypeBool, Default: false},
		{Name: "homepage", Type: field.TypeString, Nullable: true},
		{Name: "star_count", Type: field.TypeInt, Default: 0},
		{Name: "default_branch", Type: field.TypeString},
		{Name: "is_template", Type: field.TypeBool, Default: false},
		{Name: "has_issues", Type: field.TypeBool, Default: true},
		{Name: "archived", Type: field.TypeBool, Default: false},
		{Name: "pushed_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "license", Type: field.TypeJSON, Nullable: true},
	}
	// GithubRepositoriesTable holds the schema information for the "github_repositories" table.
	GithubRepositoriesTable = &schema.Table{
		Name:       "github_repositories",
		Columns:    GithubRepositoriesColumns,
		PrimaryKey: []*schema.Column{GithubRepositoriesColumns[0]},
	}
	// LabelsColumns holds the columns for the "labels" table.
	LabelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// LabelsTable holds the schema information for the "labels" table.
	LabelsTable = &schema.Table{
		Name:       "labels",
		Columns:    LabelsColumns,
		PrimaryKey: []*schema.Column{LabelsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GithubEventsTable,
		GithubRepositoriesTable,
		LabelsTable,
	}
)

func init() {
}
