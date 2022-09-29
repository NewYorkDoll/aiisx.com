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
		{Name: "label_github_repositories", Type: field.TypeInt, Nullable: true},
	}
	// GithubRepositoriesTable holds the schema information for the "github_repositories" table.
	GithubRepositoriesTable = &schema.Table{
		Name:       "github_repositories",
		Columns:    GithubRepositoriesColumns,
		PrimaryKey: []*schema.Column{GithubRepositoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "github_repositories_labels_github_repositories",
				Columns:    []*schema.Column{GithubRepositoriesColumns[20]},
				RefColumns: []*schema.Column{LabelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LabelsColumns holds the columns for the "labels" table.
	LabelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// LabelsTable holds the schema information for the "labels" table.
	LabelsTable = &schema.Table{
		Name:       "labels",
		Columns:    LabelsColumns,
		PrimaryKey: []*schema.Column{LabelsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString, Size: 100},
		{Name: "content", Type: field.TypeString, Size: 15000, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "content_html", Type: field.TypeString, Size: 50000, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "summary", Type: field.TypeString, Size: 10000, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "view_count", Type: field.TypeInt, Default: 0},
		{Name: "public", Type: field.TypeBool, Default: false},
		{Name: "user_posts", Type: field.TypeInt},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Unique: true},
		{Name: "login", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 400},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true, Size: 2048},
		{Name: "html_url", Type: field.TypeString, Nullable: true, Size: 2048},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 320},
		{Name: "location", Type: field.TypeString, Nullable: true, Size: 400},
		{Name: "bio", Type: field.TypeString, Nullable: true, Size: 400},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// LabelPostsColumns holds the columns for the "label_posts" table.
	LabelPostsColumns = []*schema.Column{
		{Name: "label_id", Type: field.TypeInt},
		{Name: "post_id", Type: field.TypeInt},
	}
	// LabelPostsTable holds the schema information for the "label_posts" table.
	LabelPostsTable = &schema.Table{
		Name:       "label_posts",
		Columns:    LabelPostsColumns,
		PrimaryKey: []*schema.Column{LabelPostsColumns[0], LabelPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "label_posts_label_id",
				Columns:    []*schema.Column{LabelPostsColumns[0]},
				RefColumns: []*schema.Column{LabelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "label_posts_post_id",
				Columns:    []*schema.Column{LabelPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GithubEventsTable,
		GithubRepositoriesTable,
		LabelsTable,
		PostsTable,
		UsersTable,
		LabelPostsTable,
	}
)

func init() {
	GithubRepositoriesTable.ForeignKeys[0].RefTable = LabelsTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	LabelPostsTable.ForeignKeys[0].RefTable = LabelsTable
	LabelPostsTable.ForeignKeys[1].RefTable = PostsTable
}
