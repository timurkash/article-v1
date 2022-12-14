// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/article"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/predicate"

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
	TypeArticle = "Article"
)

// ArticleMutation represents an operation that mutates the Article nodes in the graph.
type ArticleMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	email         *string
	user_name     *string
	title         *string
	latinized     *string
	html          *string
	short_text    *string
	lang          *string
	tags          *string
	b17_account   *string
	source        *string
	is_published  *bool
	cdate         *time.Time
	udate         *time.Time
	pub_date      *time.Time
	is_processed  *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Article, error)
	predicates    []predicate.Article
}

var _ ent.Mutation = (*ArticleMutation)(nil)

// articleOption allows management of the mutation configuration using functional options.
type articleOption func(*ArticleMutation)

// newArticleMutation creates new mutation for the Article entity.
func newArticleMutation(c config, op Op, opts ...articleOption) *ArticleMutation {
	m := &ArticleMutation{
		config:        c,
		op:            op,
		typ:           TypeArticle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArticleID sets the ID field of the mutation.
func withArticleID(id uuid.UUID) articleOption {
	return func(m *ArticleMutation) {
		var (
			err   error
			once  sync.Once
			value *Article
		)
		m.oldValue = func(ctx context.Context) (*Article, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Article.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArticle sets the old Article of the mutation.
func withArticle(node *Article) articleOption {
	return func(m *ArticleMutation) {
		m.oldValue = func(context.Context) (*Article, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArticleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArticleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Article entities.
func (m *ArticleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArticleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArticleMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Article.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetEmail sets the "email" field.
func (m *ArticleMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *ArticleMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *ArticleMutation) ResetEmail() {
	m.email = nil
}

// SetUserName sets the "user_name" field.
func (m *ArticleMutation) SetUserName(s string) {
	m.user_name = &s
}

// UserName returns the value of the "user_name" field in the mutation.
func (m *ArticleMutation) UserName() (r string, exists bool) {
	v := m.user_name
	if v == nil {
		return
	}
	return *v, true
}

// OldUserName returns the old "user_name" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldUserName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserName: %w", err)
	}
	return oldValue.UserName, nil
}

// ResetUserName resets all changes to the "user_name" field.
func (m *ArticleMutation) ResetUserName() {
	m.user_name = nil
}

// SetTitle sets the "title" field.
func (m *ArticleMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *ArticleMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *ArticleMutation) ResetTitle() {
	m.title = nil
}

// SetLatinized sets the "latinized" field.
func (m *ArticleMutation) SetLatinized(s string) {
	m.latinized = &s
}

// Latinized returns the value of the "latinized" field in the mutation.
func (m *ArticleMutation) Latinized() (r string, exists bool) {
	v := m.latinized
	if v == nil {
		return
	}
	return *v, true
}

// OldLatinized returns the old "latinized" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldLatinized(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLatinized is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLatinized requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLatinized: %w", err)
	}
	return oldValue.Latinized, nil
}

// ResetLatinized resets all changes to the "latinized" field.
func (m *ArticleMutation) ResetLatinized() {
	m.latinized = nil
}

// SetHTML sets the "html" field.
func (m *ArticleMutation) SetHTML(s string) {
	m.html = &s
}

// HTML returns the value of the "html" field in the mutation.
func (m *ArticleMutation) HTML() (r string, exists bool) {
	v := m.html
	if v == nil {
		return
	}
	return *v, true
}

// OldHTML returns the old "html" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldHTML(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHTML is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHTML requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHTML: %w", err)
	}
	return oldValue.HTML, nil
}

// ResetHTML resets all changes to the "html" field.
func (m *ArticleMutation) ResetHTML() {
	m.html = nil
}

// SetShortText sets the "short_text" field.
func (m *ArticleMutation) SetShortText(s string) {
	m.short_text = &s
}

// ShortText returns the value of the "short_text" field in the mutation.
func (m *ArticleMutation) ShortText() (r string, exists bool) {
	v := m.short_text
	if v == nil {
		return
	}
	return *v, true
}

// OldShortText returns the old "short_text" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldShortText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldShortText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldShortText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldShortText: %w", err)
	}
	return oldValue.ShortText, nil
}

// ResetShortText resets all changes to the "short_text" field.
func (m *ArticleMutation) ResetShortText() {
	m.short_text = nil
}

// SetLang sets the "lang" field.
func (m *ArticleMutation) SetLang(s string) {
	m.lang = &s
}

// Lang returns the value of the "lang" field in the mutation.
func (m *ArticleMutation) Lang() (r string, exists bool) {
	v := m.lang
	if v == nil {
		return
	}
	return *v, true
}

// OldLang returns the old "lang" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldLang(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLang is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLang requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLang: %w", err)
	}
	return oldValue.Lang, nil
}

// ResetLang resets all changes to the "lang" field.
func (m *ArticleMutation) ResetLang() {
	m.lang = nil
}

// SetTags sets the "tags" field.
func (m *ArticleMutation) SetTags(s string) {
	m.tags = &s
}

// Tags returns the value of the "tags" field in the mutation.
func (m *ArticleMutation) Tags() (r string, exists bool) {
	v := m.tags
	if v == nil {
		return
	}
	return *v, true
}

// OldTags returns the old "tags" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldTags(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTags is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTags requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTags: %w", err)
	}
	return oldValue.Tags, nil
}

// ClearTags clears the value of the "tags" field.
func (m *ArticleMutation) ClearTags() {
	m.tags = nil
	m.clearedFields[article.FieldTags] = struct{}{}
}

// TagsCleared returns if the "tags" field was cleared in this mutation.
func (m *ArticleMutation) TagsCleared() bool {
	_, ok := m.clearedFields[article.FieldTags]
	return ok
}

// ResetTags resets all changes to the "tags" field.
func (m *ArticleMutation) ResetTags() {
	m.tags = nil
	delete(m.clearedFields, article.FieldTags)
}

// SetB17Account sets the "b17_account" field.
func (m *ArticleMutation) SetB17Account(s string) {
	m.b17_account = &s
}

// B17Account returns the value of the "b17_account" field in the mutation.
func (m *ArticleMutation) B17Account() (r string, exists bool) {
	v := m.b17_account
	if v == nil {
		return
	}
	return *v, true
}

// OldB17Account returns the old "b17_account" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldB17Account(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldB17Account is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldB17Account requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldB17Account: %w", err)
	}
	return oldValue.B17Account, nil
}

// ClearB17Account clears the value of the "b17_account" field.
func (m *ArticleMutation) ClearB17Account() {
	m.b17_account = nil
	m.clearedFields[article.FieldB17Account] = struct{}{}
}

// B17AccountCleared returns if the "b17_account" field was cleared in this mutation.
func (m *ArticleMutation) B17AccountCleared() bool {
	_, ok := m.clearedFields[article.FieldB17Account]
	return ok
}

// ResetB17Account resets all changes to the "b17_account" field.
func (m *ArticleMutation) ResetB17Account() {
	m.b17_account = nil
	delete(m.clearedFields, article.FieldB17Account)
}

// SetSource sets the "source" field.
func (m *ArticleMutation) SetSource(s string) {
	m.source = &s
}

// Source returns the value of the "source" field in the mutation.
func (m *ArticleMutation) Source() (r string, exists bool) {
	v := m.source
	if v == nil {
		return
	}
	return *v, true
}

// OldSource returns the old "source" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldSource(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSource is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSource requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSource: %w", err)
	}
	return oldValue.Source, nil
}

// ClearSource clears the value of the "source" field.
func (m *ArticleMutation) ClearSource() {
	m.source = nil
	m.clearedFields[article.FieldSource] = struct{}{}
}

// SourceCleared returns if the "source" field was cleared in this mutation.
func (m *ArticleMutation) SourceCleared() bool {
	_, ok := m.clearedFields[article.FieldSource]
	return ok
}

// ResetSource resets all changes to the "source" field.
func (m *ArticleMutation) ResetSource() {
	m.source = nil
	delete(m.clearedFields, article.FieldSource)
}

// SetIsPublished sets the "is_published" field.
func (m *ArticleMutation) SetIsPublished(b bool) {
	m.is_published = &b
}

// IsPublished returns the value of the "is_published" field in the mutation.
func (m *ArticleMutation) IsPublished() (r bool, exists bool) {
	v := m.is_published
	if v == nil {
		return
	}
	return *v, true
}

// OldIsPublished returns the old "is_published" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldIsPublished(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsPublished is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsPublished requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsPublished: %w", err)
	}
	return oldValue.IsPublished, nil
}

// ResetIsPublished resets all changes to the "is_published" field.
func (m *ArticleMutation) ResetIsPublished() {
	m.is_published = nil
}

// SetCdate sets the "cdate" field.
func (m *ArticleMutation) SetCdate(t time.Time) {
	m.cdate = &t
}

// Cdate returns the value of the "cdate" field in the mutation.
func (m *ArticleMutation) Cdate() (r time.Time, exists bool) {
	v := m.cdate
	if v == nil {
		return
	}
	return *v, true
}

// OldCdate returns the old "cdate" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldCdate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCdate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCdate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCdate: %w", err)
	}
	return oldValue.Cdate, nil
}

// ResetCdate resets all changes to the "cdate" field.
func (m *ArticleMutation) ResetCdate() {
	m.cdate = nil
}

// SetUdate sets the "udate" field.
func (m *ArticleMutation) SetUdate(t time.Time) {
	m.udate = &t
}

// Udate returns the value of the "udate" field in the mutation.
func (m *ArticleMutation) Udate() (r time.Time, exists bool) {
	v := m.udate
	if v == nil {
		return
	}
	return *v, true
}

// OldUdate returns the old "udate" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldUdate(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUdate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUdate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUdate: %w", err)
	}
	return oldValue.Udate, nil
}

// ClearUdate clears the value of the "udate" field.
func (m *ArticleMutation) ClearUdate() {
	m.udate = nil
	m.clearedFields[article.FieldUdate] = struct{}{}
}

// UdateCleared returns if the "udate" field was cleared in this mutation.
func (m *ArticleMutation) UdateCleared() bool {
	_, ok := m.clearedFields[article.FieldUdate]
	return ok
}

// ResetUdate resets all changes to the "udate" field.
func (m *ArticleMutation) ResetUdate() {
	m.udate = nil
	delete(m.clearedFields, article.FieldUdate)
}

// SetPubDate sets the "pub_date" field.
func (m *ArticleMutation) SetPubDate(t time.Time) {
	m.pub_date = &t
}

// PubDate returns the value of the "pub_date" field in the mutation.
func (m *ArticleMutation) PubDate() (r time.Time, exists bool) {
	v := m.pub_date
	if v == nil {
		return
	}
	return *v, true
}

// OldPubDate returns the old "pub_date" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldPubDate(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPubDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPubDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPubDate: %w", err)
	}
	return oldValue.PubDate, nil
}

// ClearPubDate clears the value of the "pub_date" field.
func (m *ArticleMutation) ClearPubDate() {
	m.pub_date = nil
	m.clearedFields[article.FieldPubDate] = struct{}{}
}

// PubDateCleared returns if the "pub_date" field was cleared in this mutation.
func (m *ArticleMutation) PubDateCleared() bool {
	_, ok := m.clearedFields[article.FieldPubDate]
	return ok
}

// ResetPubDate resets all changes to the "pub_date" field.
func (m *ArticleMutation) ResetPubDate() {
	m.pub_date = nil
	delete(m.clearedFields, article.FieldPubDate)
}

// SetIsProcessed sets the "is_processed" field.
func (m *ArticleMutation) SetIsProcessed(b bool) {
	m.is_processed = &b
}

// IsProcessed returns the value of the "is_processed" field in the mutation.
func (m *ArticleMutation) IsProcessed() (r bool, exists bool) {
	v := m.is_processed
	if v == nil {
		return
	}
	return *v, true
}

// OldIsProcessed returns the old "is_processed" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldIsProcessed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsProcessed is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsProcessed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsProcessed: %w", err)
	}
	return oldValue.IsProcessed, nil
}

// ResetIsProcessed resets all changes to the "is_processed" field.
func (m *ArticleMutation) ResetIsProcessed() {
	m.is_processed = nil
}

// Where appends a list predicates to the ArticleMutation builder.
func (m *ArticleMutation) Where(ps ...predicate.Article) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ArticleMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Article).
func (m *ArticleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArticleMutation) Fields() []string {
	fields := make([]string, 0, 15)
	if m.email != nil {
		fields = append(fields, article.FieldEmail)
	}
	if m.user_name != nil {
		fields = append(fields, article.FieldUserName)
	}
	if m.title != nil {
		fields = append(fields, article.FieldTitle)
	}
	if m.latinized != nil {
		fields = append(fields, article.FieldLatinized)
	}
	if m.html != nil {
		fields = append(fields, article.FieldHTML)
	}
	if m.short_text != nil {
		fields = append(fields, article.FieldShortText)
	}
	if m.lang != nil {
		fields = append(fields, article.FieldLang)
	}
	if m.tags != nil {
		fields = append(fields, article.FieldTags)
	}
	if m.b17_account != nil {
		fields = append(fields, article.FieldB17Account)
	}
	if m.source != nil {
		fields = append(fields, article.FieldSource)
	}
	if m.is_published != nil {
		fields = append(fields, article.FieldIsPublished)
	}
	if m.cdate != nil {
		fields = append(fields, article.FieldCdate)
	}
	if m.udate != nil {
		fields = append(fields, article.FieldUdate)
	}
	if m.pub_date != nil {
		fields = append(fields, article.FieldPubDate)
	}
	if m.is_processed != nil {
		fields = append(fields, article.FieldIsProcessed)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArticleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case article.FieldEmail:
		return m.Email()
	case article.FieldUserName:
		return m.UserName()
	case article.FieldTitle:
		return m.Title()
	case article.FieldLatinized:
		return m.Latinized()
	case article.FieldHTML:
		return m.HTML()
	case article.FieldShortText:
		return m.ShortText()
	case article.FieldLang:
		return m.Lang()
	case article.FieldTags:
		return m.Tags()
	case article.FieldB17Account:
		return m.B17Account()
	case article.FieldSource:
		return m.Source()
	case article.FieldIsPublished:
		return m.IsPublished()
	case article.FieldCdate:
		return m.Cdate()
	case article.FieldUdate:
		return m.Udate()
	case article.FieldPubDate:
		return m.PubDate()
	case article.FieldIsProcessed:
		return m.IsProcessed()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArticleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case article.FieldEmail:
		return m.OldEmail(ctx)
	case article.FieldUserName:
		return m.OldUserName(ctx)
	case article.FieldTitle:
		return m.OldTitle(ctx)
	case article.FieldLatinized:
		return m.OldLatinized(ctx)
	case article.FieldHTML:
		return m.OldHTML(ctx)
	case article.FieldShortText:
		return m.OldShortText(ctx)
	case article.FieldLang:
		return m.OldLang(ctx)
	case article.FieldTags:
		return m.OldTags(ctx)
	case article.FieldB17Account:
		return m.OldB17Account(ctx)
	case article.FieldSource:
		return m.OldSource(ctx)
	case article.FieldIsPublished:
		return m.OldIsPublished(ctx)
	case article.FieldCdate:
		return m.OldCdate(ctx)
	case article.FieldUdate:
		return m.OldUdate(ctx)
	case article.FieldPubDate:
		return m.OldPubDate(ctx)
	case article.FieldIsProcessed:
		return m.OldIsProcessed(ctx)
	}
	return nil, fmt.Errorf("unknown Article field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case article.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case article.FieldUserName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserName(v)
		return nil
	case article.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case article.FieldLatinized:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLatinized(v)
		return nil
	case article.FieldHTML:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHTML(v)
		return nil
	case article.FieldShortText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetShortText(v)
		return nil
	case article.FieldLang:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLang(v)
		return nil
	case article.FieldTags:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTags(v)
		return nil
	case article.FieldB17Account:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetB17Account(v)
		return nil
	case article.FieldSource:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSource(v)
		return nil
	case article.FieldIsPublished:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsPublished(v)
		return nil
	case article.FieldCdate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCdate(v)
		return nil
	case article.FieldUdate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUdate(v)
		return nil
	case article.FieldPubDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPubDate(v)
		return nil
	case article.FieldIsProcessed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsProcessed(v)
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArticleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArticleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Article numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArticleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(article.FieldTags) {
		fields = append(fields, article.FieldTags)
	}
	if m.FieldCleared(article.FieldB17Account) {
		fields = append(fields, article.FieldB17Account)
	}
	if m.FieldCleared(article.FieldSource) {
		fields = append(fields, article.FieldSource)
	}
	if m.FieldCleared(article.FieldUdate) {
		fields = append(fields, article.FieldUdate)
	}
	if m.FieldCleared(article.FieldPubDate) {
		fields = append(fields, article.FieldPubDate)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArticleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArticleMutation) ClearField(name string) error {
	switch name {
	case article.FieldTags:
		m.ClearTags()
		return nil
	case article.FieldB17Account:
		m.ClearB17Account()
		return nil
	case article.FieldSource:
		m.ClearSource()
		return nil
	case article.FieldUdate:
		m.ClearUdate()
		return nil
	case article.FieldPubDate:
		m.ClearPubDate()
		return nil
	}
	return fmt.Errorf("unknown Article nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArticleMutation) ResetField(name string) error {
	switch name {
	case article.FieldEmail:
		m.ResetEmail()
		return nil
	case article.FieldUserName:
		m.ResetUserName()
		return nil
	case article.FieldTitle:
		m.ResetTitle()
		return nil
	case article.FieldLatinized:
		m.ResetLatinized()
		return nil
	case article.FieldHTML:
		m.ResetHTML()
		return nil
	case article.FieldShortText:
		m.ResetShortText()
		return nil
	case article.FieldLang:
		m.ResetLang()
		return nil
	case article.FieldTags:
		m.ResetTags()
		return nil
	case article.FieldB17Account:
		m.ResetB17Account()
		return nil
	case article.FieldSource:
		m.ResetSource()
		return nil
	case article.FieldIsPublished:
		m.ResetIsPublished()
		return nil
	case article.FieldCdate:
		m.ResetCdate()
		return nil
	case article.FieldUdate:
		m.ResetUdate()
		return nil
	case article.FieldPubDate:
		m.ResetPubDate()
		return nil
	case article.FieldIsProcessed:
		m.ResetIsProcessed()
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArticleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArticleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArticleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArticleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArticleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArticleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArticleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Article unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArticleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Article edge %s", name)
}
