// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticleColumns holds the columns for the "article" table.
	ArticleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "latinized", Type: field.TypeString},
		{Name: "html", Type: field.TypeString, Size: 2147483647},
		{Name: "short_text", Type: field.TypeString, Size: 2147483647},
		{Name: "lang", Type: field.TypeString, Size: 2},
		{Name: "tags", Type: field.TypeString, Nullable: true},
		{Name: "b17_account", Type: field.TypeString, Nullable: true},
		{Name: "source", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "is_published", Type: field.TypeBool, Default: false},
		{Name: "cdate", Type: field.TypeTime},
		{Name: "udate", Type: field.TypeTime, Nullable: true},
		{Name: "pub_date", Type: field.TypeTime, Nullable: true},
		{Name: "is_processed", Type: field.TypeBool, Default: false},
	}
	// ArticleTable holds the schema information for the "article" table.
	ArticleTable = &schema.Table{
		Name:       "article",
		Columns:    ArticleColumns,
		PrimaryKey: []*schema.Column{ArticleColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "article_pub_date",
				Unique:  false,
				Columns: []*schema.Column{ArticleColumns[14]},
			},
			{
				Name:    "article_email",
				Unique:  false,
				Columns: []*schema.Column{ArticleColumns[1]},
			},
			{
				Name:    "article_tags",
				Unique:  false,
				Columns: []*schema.Column{ArticleColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticleTable,
	}
)

func init() {
	ArticleTable.Annotation = &entsql.Annotation{
		Table: "article",
	}
}
