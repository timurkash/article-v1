package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"

	"gitlab.com/mcsolutions/find-psy/back/common/consts"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "article"},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(consts.Id, uuid.UUID{}).Unique(),
		field.String(consts.Email),
		field.String(consts.UserName),
		field.String(consts.Title),
		field.String(consts.Latinized),
		field.Text(consts.Html),
		field.Text(consts.ShortText),
		field.String(consts.Lang).MaxLen(2),
		field.String(consts.Tags).Optional().Nillable(),
		field.String(consts.B17Account).Optional().Nillable(),
		field.String(consts.Source).Unique().Optional().Nillable(),
		field.Bool(consts.IsPublished).Default(false),
		field.Time(consts.CDate).Default(time.Now),
		field.Time(consts.UDate).Optional().Nillable(),
		field.Time(consts.PubDate).Optional().Nillable(),
		field.Bool("is_processed").Default(false),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(consts.PubDate),
		index.Fields(consts.Email),
		index.Fields(consts.Tags),
	}
}
