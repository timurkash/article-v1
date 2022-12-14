// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/article"
)

// Article is the model entity for the Article schema.
type Article struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Latinized holds the value of the "latinized" field.
	Latinized string `json:"latinized,omitempty"`
	// HTML holds the value of the "html" field.
	HTML string `json:"html,omitempty"`
	// ShortText holds the value of the "short_text" field.
	ShortText string `json:"short_text,omitempty"`
	// Lang holds the value of the "lang" field.
	Lang string `json:"lang,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags *string `json:"tags,omitempty"`
	// B17Account holds the value of the "b17_account" field.
	B17Account *string `json:"b17_account,omitempty"`
	// Source holds the value of the "source" field.
	Source *string `json:"source,omitempty"`
	// IsPublished holds the value of the "is_published" field.
	IsPublished bool `json:"is_published,omitempty"`
	// Cdate holds the value of the "cdate" field.
	Cdate time.Time `json:"cdate,omitempty"`
	// Udate holds the value of the "udate" field.
	Udate *time.Time `json:"udate,omitempty"`
	// PubDate holds the value of the "pub_date" field.
	PubDate *time.Time `json:"pub_date,omitempty"`
	// IsProcessed holds the value of the "is_processed" field.
	IsProcessed bool `json:"is_processed,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Article) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case article.FieldIsPublished, article.FieldIsProcessed:
			values[i] = new(sql.NullBool)
		case article.FieldEmail, article.FieldUserName, article.FieldTitle, article.FieldLatinized, article.FieldHTML, article.FieldShortText, article.FieldLang, article.FieldTags, article.FieldB17Account, article.FieldSource:
			values[i] = new(sql.NullString)
		case article.FieldCdate, article.FieldUdate, article.FieldPubDate:
			values[i] = new(sql.NullTime)
		case article.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Article", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Article fields.
func (a *Article) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case article.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case article.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case article.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				a.UserName = value.String
			}
		case article.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case article.FieldLatinized:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latinized", values[i])
			} else if value.Valid {
				a.Latinized = value.String
			}
		case article.FieldHTML:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field html", values[i])
			} else if value.Valid {
				a.HTML = value.String
			}
		case article.FieldShortText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short_text", values[i])
			} else if value.Valid {
				a.ShortText = value.String
			}
		case article.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				a.Lang = value.String
			}
		case article.FieldTags:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value.Valid {
				a.Tags = new(string)
				*a.Tags = value.String
			}
		case article.FieldB17Account:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field b17_account", values[i])
			} else if value.Valid {
				a.B17Account = new(string)
				*a.B17Account = value.String
			}
		case article.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				a.Source = new(string)
				*a.Source = value.String
			}
		case article.FieldIsPublished:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_published", values[i])
			} else if value.Valid {
				a.IsPublished = value.Bool
			}
		case article.FieldCdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field cdate", values[i])
			} else if value.Valid {
				a.Cdate = value.Time
			}
		case article.FieldUdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field udate", values[i])
			} else if value.Valid {
				a.Udate = new(time.Time)
				*a.Udate = value.Time
			}
		case article.FieldPubDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field pub_date", values[i])
			} else if value.Valid {
				a.PubDate = new(time.Time)
				*a.PubDate = value.Time
			}
		case article.FieldIsProcessed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_processed", values[i])
			} else if value.Valid {
				a.IsProcessed = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Article.
// Note that you need to call Article.Unwrap() before calling this method if this Article
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Article) Update() *ArticleUpdateOne {
	return (&ArticleClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Article entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Article) Unwrap() *Article {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Article is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Article) String() string {
	var builder strings.Builder
	builder.WriteString("Article(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", email=")
	builder.WriteString(a.Email)
	builder.WriteString(", user_name=")
	builder.WriteString(a.UserName)
	builder.WriteString(", title=")
	builder.WriteString(a.Title)
	builder.WriteString(", latinized=")
	builder.WriteString(a.Latinized)
	builder.WriteString(", html=")
	builder.WriteString(a.HTML)
	builder.WriteString(", short_text=")
	builder.WriteString(a.ShortText)
	builder.WriteString(", lang=")
	builder.WriteString(a.Lang)
	if v := a.Tags; v != nil {
		builder.WriteString(", tags=")
		builder.WriteString(*v)
	}
	if v := a.B17Account; v != nil {
		builder.WriteString(", b17_account=")
		builder.WriteString(*v)
	}
	if v := a.Source; v != nil {
		builder.WriteString(", source=")
		builder.WriteString(*v)
	}
	builder.WriteString(", is_published=")
	builder.WriteString(fmt.Sprintf("%v", a.IsPublished))
	builder.WriteString(", cdate=")
	builder.WriteString(a.Cdate.Format(time.ANSIC))
	if v := a.Udate; v != nil {
		builder.WriteString(", udate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := a.PubDate; v != nil {
		builder.WriteString(", pub_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", is_processed=")
	builder.WriteString(fmt.Sprintf("%v", a.IsProcessed))
	builder.WriteByte(')')
	return builder.String()
}

// Articles is a parsable slice of Article.
type Articles []*Article

func (a Articles) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
