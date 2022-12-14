// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/article"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/predicate"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetEmail sets the "email" field.
func (au *ArticleUpdate) SetEmail(s string) *ArticleUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetUserName sets the "user_name" field.
func (au *ArticleUpdate) SetUserName(s string) *ArticleUpdate {
	au.mutation.SetUserName(s)
	return au
}

// SetTitle sets the "title" field.
func (au *ArticleUpdate) SetTitle(s string) *ArticleUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetLatinized sets the "latinized" field.
func (au *ArticleUpdate) SetLatinized(s string) *ArticleUpdate {
	au.mutation.SetLatinized(s)
	return au
}

// SetHTML sets the "html" field.
func (au *ArticleUpdate) SetHTML(s string) *ArticleUpdate {
	au.mutation.SetHTML(s)
	return au
}

// SetShortText sets the "short_text" field.
func (au *ArticleUpdate) SetShortText(s string) *ArticleUpdate {
	au.mutation.SetShortText(s)
	return au
}

// SetLang sets the "lang" field.
func (au *ArticleUpdate) SetLang(s string) *ArticleUpdate {
	au.mutation.SetLang(s)
	return au
}

// SetTags sets the "tags" field.
func (au *ArticleUpdate) SetTags(s string) *ArticleUpdate {
	au.mutation.SetTags(s)
	return au
}

// SetNillableTags sets the "tags" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTags(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTags(*s)
	}
	return au
}

// ClearTags clears the value of the "tags" field.
func (au *ArticleUpdate) ClearTags() *ArticleUpdate {
	au.mutation.ClearTags()
	return au
}

// SetB17Account sets the "b17_account" field.
func (au *ArticleUpdate) SetB17Account(s string) *ArticleUpdate {
	au.mutation.SetB17Account(s)
	return au
}

// SetNillableB17Account sets the "b17_account" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableB17Account(s *string) *ArticleUpdate {
	if s != nil {
		au.SetB17Account(*s)
	}
	return au
}

// ClearB17Account clears the value of the "b17_account" field.
func (au *ArticleUpdate) ClearB17Account() *ArticleUpdate {
	au.mutation.ClearB17Account()
	return au
}

// SetSource sets the "source" field.
func (au *ArticleUpdate) SetSource(s string) *ArticleUpdate {
	au.mutation.SetSource(s)
	return au
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableSource(s *string) *ArticleUpdate {
	if s != nil {
		au.SetSource(*s)
	}
	return au
}

// ClearSource clears the value of the "source" field.
func (au *ArticleUpdate) ClearSource() *ArticleUpdate {
	au.mutation.ClearSource()
	return au
}

// SetIsPublished sets the "is_published" field.
func (au *ArticleUpdate) SetIsPublished(b bool) *ArticleUpdate {
	au.mutation.SetIsPublished(b)
	return au
}

// SetNillableIsPublished sets the "is_published" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableIsPublished(b *bool) *ArticleUpdate {
	if b != nil {
		au.SetIsPublished(*b)
	}
	return au
}

// SetCdate sets the "cdate" field.
func (au *ArticleUpdate) SetCdate(t time.Time) *ArticleUpdate {
	au.mutation.SetCdate(t)
	return au
}

// SetNillableCdate sets the "cdate" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCdate(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetCdate(*t)
	}
	return au
}

// SetUdate sets the "udate" field.
func (au *ArticleUpdate) SetUdate(t time.Time) *ArticleUpdate {
	au.mutation.SetUdate(t)
	return au
}

// SetNillableUdate sets the "udate" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableUdate(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetUdate(*t)
	}
	return au
}

// ClearUdate clears the value of the "udate" field.
func (au *ArticleUpdate) ClearUdate() *ArticleUpdate {
	au.mutation.ClearUdate()
	return au
}

// SetPubDate sets the "pub_date" field.
func (au *ArticleUpdate) SetPubDate(t time.Time) *ArticleUpdate {
	au.mutation.SetPubDate(t)
	return au
}

// SetNillablePubDate sets the "pub_date" field if the given value is not nil.
func (au *ArticleUpdate) SetNillablePubDate(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetPubDate(*t)
	}
	return au
}

// ClearPubDate clears the value of the "pub_date" field.
func (au *ArticleUpdate) ClearPubDate() *ArticleUpdate {
	au.mutation.ClearPubDate()
	return au
}

// SetIsProcessed sets the "is_processed" field.
func (au *ArticleUpdate) SetIsProcessed(b bool) *ArticleUpdate {
	au.mutation.SetIsProcessed(b)
	return au
}

// SetNillableIsProcessed sets the "is_processed" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableIsProcessed(b *bool) *ArticleUpdate {
	if b != nil {
		au.SetIsProcessed(*b)
	}
	return au
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArticleUpdate) check() error {
	if v, ok := au.mutation.Lang(); ok {
		if err := article.LangValidator(v); err != nil {
			return &ValidationError{Name: "lang", err: fmt.Errorf(`ent: validator failed for field "Article.lang": %w`, err)}
		}
	}
	return nil
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   article.Table,
			Columns: article.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: article.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldEmail,
		})
	}
	if value, ok := au.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldUserName,
		})
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTitle,
		})
	}
	if value, ok := au.mutation.Latinized(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldLatinized,
		})
	}
	if value, ok := au.mutation.HTML(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldHTML,
		})
	}
	if value, ok := au.mutation.ShortText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldShortText,
		})
	}
	if value, ok := au.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldLang,
		})
	}
	if value, ok := au.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTags,
		})
	}
	if au.mutation.TagsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: article.FieldTags,
		})
	}
	if value, ok := au.mutation.B17Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldB17Account,
		})
	}
	if au.mutation.B17AccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: article.FieldB17Account,
		})
	}
	if value, ok := au.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldSource,
		})
	}
	if au.mutation.SourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: article.FieldSource,
		})
	}
	if value, ok := au.mutation.IsPublished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: article.FieldIsPublished,
		})
	}
	if value, ok := au.mutation.Cdate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldCdate,
		})
	}
	if value, ok := au.mutation.Udate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldUdate,
		})
	}
	if au.mutation.UdateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: article.FieldUdate,
		})
	}
	if value, ok := au.mutation.PubDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldPubDate,
		})
	}
	if au.mutation.PubDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: article.FieldPubDate,
		})
	}
	if value, ok := au.mutation.IsProcessed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: article.FieldIsProcessed,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetEmail sets the "email" field.
func (auo *ArticleUpdateOne) SetEmail(s string) *ArticleUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetUserName sets the "user_name" field.
func (auo *ArticleUpdateOne) SetUserName(s string) *ArticleUpdateOne {
	auo.mutation.SetUserName(s)
	return auo
}

// SetTitle sets the "title" field.
func (auo *ArticleUpdateOne) SetTitle(s string) *ArticleUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetLatinized sets the "latinized" field.
func (auo *ArticleUpdateOne) SetLatinized(s string) *ArticleUpdateOne {
	auo.mutation.SetLatinized(s)
	return auo
}

// SetHTML sets the "html" field.
func (auo *ArticleUpdateOne) SetHTML(s string) *ArticleUpdateOne {
	auo.mutation.SetHTML(s)
	return auo
}

// SetShortText sets the "short_text" field.
func (auo *ArticleUpdateOne) SetShortText(s string) *ArticleUpdateOne {
	auo.mutation.SetShortText(s)
	return auo
}

// SetLang sets the "lang" field.
func (auo *ArticleUpdateOne) SetLang(s string) *ArticleUpdateOne {
	auo.mutation.SetLang(s)
	return auo
}

// SetTags sets the "tags" field.
func (auo *ArticleUpdateOne) SetTags(s string) *ArticleUpdateOne {
	auo.mutation.SetTags(s)
	return auo
}

// SetNillableTags sets the "tags" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTags(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTags(*s)
	}
	return auo
}

// ClearTags clears the value of the "tags" field.
func (auo *ArticleUpdateOne) ClearTags() *ArticleUpdateOne {
	auo.mutation.ClearTags()
	return auo
}

// SetB17Account sets the "b17_account" field.
func (auo *ArticleUpdateOne) SetB17Account(s string) *ArticleUpdateOne {
	auo.mutation.SetB17Account(s)
	return auo
}

// SetNillableB17Account sets the "b17_account" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableB17Account(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetB17Account(*s)
	}
	return auo
}

// ClearB17Account clears the value of the "b17_account" field.
func (auo *ArticleUpdateOne) ClearB17Account() *ArticleUpdateOne {
	auo.mutation.ClearB17Account()
	return auo
}

// SetSource sets the "source" field.
func (auo *ArticleUpdateOne) SetSource(s string) *ArticleUpdateOne {
	auo.mutation.SetSource(s)
	return auo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableSource(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetSource(*s)
	}
	return auo
}

// ClearSource clears the value of the "source" field.
func (auo *ArticleUpdateOne) ClearSource() *ArticleUpdateOne {
	auo.mutation.ClearSource()
	return auo
}

// SetIsPublished sets the "is_published" field.
func (auo *ArticleUpdateOne) SetIsPublished(b bool) *ArticleUpdateOne {
	auo.mutation.SetIsPublished(b)
	return auo
}

// SetNillableIsPublished sets the "is_published" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableIsPublished(b *bool) *ArticleUpdateOne {
	if b != nil {
		auo.SetIsPublished(*b)
	}
	return auo
}

// SetCdate sets the "cdate" field.
func (auo *ArticleUpdateOne) SetCdate(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetCdate(t)
	return auo
}

// SetNillableCdate sets the "cdate" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCdate(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetCdate(*t)
	}
	return auo
}

// SetUdate sets the "udate" field.
func (auo *ArticleUpdateOne) SetUdate(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetUdate(t)
	return auo
}

// SetNillableUdate sets the "udate" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableUdate(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetUdate(*t)
	}
	return auo
}

// ClearUdate clears the value of the "udate" field.
func (auo *ArticleUpdateOne) ClearUdate() *ArticleUpdateOne {
	auo.mutation.ClearUdate()
	return auo
}

// SetPubDate sets the "pub_date" field.
func (auo *ArticleUpdateOne) SetPubDate(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetPubDate(t)
	return auo
}

// SetNillablePubDate sets the "pub_date" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillablePubDate(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetPubDate(*t)
	}
	return auo
}

// ClearPubDate clears the value of the "pub_date" field.
func (auo *ArticleUpdateOne) ClearPubDate() *ArticleUpdateOne {
	auo.mutation.ClearPubDate()
	return auo
}

// SetIsProcessed sets the "is_processed" field.
func (auo *ArticleUpdateOne) SetIsProcessed(b bool) *ArticleUpdateOne {
	auo.mutation.SetIsProcessed(b)
	return auo
}

// SetNillableIsProcessed sets the "is_processed" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableIsProcessed(b *bool) *ArticleUpdateOne {
	if b != nil {
		auo.SetIsProcessed(*b)
	}
	return auo
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	var (
		err  error
		node *Article
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArticleUpdateOne) check() error {
	if v, ok := auo.mutation.Lang(); ok {
		if err := article.LangValidator(v); err != nil {
			return &ValidationError{Name: "lang", err: fmt.Errorf(`ent: validator failed for field "Article.lang": %w`, err)}
		}
	}
	return nil
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   article.Table,
			Columns: article.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: article.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldEmail,
		})
	}
	if value, ok := auo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldUserName,
		})
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTitle,
		})
	}
	if value, ok := auo.mutation.Latinized(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldLatinized,
		})
	}
	if value, ok := auo.mutation.HTML(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldHTML,
		})
	}
	if value, ok := auo.mutation.ShortText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldShortText,
		})
	}
	if value, ok := auo.mutation.Lang(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldLang,
		})
	}
	if value, ok := auo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldTags,
		})
	}
	if auo.mutation.TagsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: article.FieldTags,
		})
	}
	if value, ok := auo.mutation.B17Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldB17Account,
		})
	}
	if auo.mutation.B17AccountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: article.FieldB17Account,
		})
	}
	if value, ok := auo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: article.FieldSource,
		})
	}
	if auo.mutation.SourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: article.FieldSource,
		})
	}
	if value, ok := auo.mutation.IsPublished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: article.FieldIsPublished,
		})
	}
	if value, ok := auo.mutation.Cdate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldCdate,
		})
	}
	if value, ok := auo.mutation.Udate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldUdate,
		})
	}
	if auo.mutation.UdateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: article.FieldUdate,
		})
	}
	if value, ok := auo.mutation.PubDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: article.FieldPubDate,
		})
	}
	if auo.mutation.PubDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: article.FieldPubDate,
		})
	}
	if value, ok := auo.mutation.IsProcessed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: article.FieldIsProcessed,
		})
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
