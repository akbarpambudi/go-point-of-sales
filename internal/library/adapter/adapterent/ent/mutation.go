// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/category"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/predicate"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/product"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/variant"
	"github.com/google/uuid"

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
	TypeCategory = "Category"
	TypeProduct  = "Product"
	TypeVariant  = "Variant"
)

// CategoryMutation represents an operation that mutates the Category nodes in the graph.
type CategoryMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Category, error)
	predicates    []predicate.Category
}

var _ ent.Mutation = (*CategoryMutation)(nil)

// categoryOption allows management of the mutation configuration using functional options.
type categoryOption func(*CategoryMutation)

// newCategoryMutation creates new mutation for the Category entity.
func newCategoryMutation(c config, op Op, opts ...categoryOption) *CategoryMutation {
	m := &CategoryMutation{
		config:        c,
		op:            op,
		typ:           TypeCategory,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCategoryID sets the ID field of the mutation.
func withCategoryID(id uuid.UUID) categoryOption {
	return func(m *CategoryMutation) {
		var (
			err   error
			once  sync.Once
			value *Category
		)
		m.oldValue = func(ctx context.Context) (*Category, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Category.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCategory sets the old Category of the mutation.
func withCategory(node *Category) categoryOption {
	return func(m *CategoryMutation) {
		m.oldValue = func(context.Context) (*Category, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CategoryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CategoryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Category entities.
func (m *CategoryMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *CategoryMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *CategoryMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CategoryMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CategoryMutation) ResetName() {
	m.name = nil
}

// Op returns the operation name.
func (m *CategoryMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Category).
func (m *CategoryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CategoryMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, category.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CategoryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case category.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CategoryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case category.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Category field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CategoryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case category.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Category field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CategoryMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CategoryMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CategoryMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Category numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CategoryMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CategoryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CategoryMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Category nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CategoryMutation) ResetField(name string) error {
	switch name {
	case category.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Category field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CategoryMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CategoryMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CategoryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CategoryMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CategoryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CategoryMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CategoryMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Category unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CategoryMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Category edge %s", name)
}

// ProductMutation represents an operation that mutates the Product nodes in the graph.
type ProductMutation struct {
	config
	op              Op
	typ             string
	id              *uuid.UUID
	name            *string
	categoryRef     *string
	clearedFields   map[string]struct{}
	variants        map[uuid.UUID]struct{}
	removedvariants map[uuid.UUID]struct{}
	clearedvariants bool
	done            bool
	oldValue        func(context.Context) (*Product, error)
	predicates      []predicate.Product
}

var _ ent.Mutation = (*ProductMutation)(nil)

// productOption allows management of the mutation configuration using functional options.
type productOption func(*ProductMutation)

// newProductMutation creates new mutation for the Product entity.
func newProductMutation(c config, op Op, opts ...productOption) *ProductMutation {
	m := &ProductMutation{
		config:        c,
		op:            op,
		typ:           TypeProduct,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProductID sets the ID field of the mutation.
func withProductID(id uuid.UUID) productOption {
	return func(m *ProductMutation) {
		var (
			err   error
			once  sync.Once
			value *Product
		)
		m.oldValue = func(ctx context.Context) (*Product, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Product.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProduct sets the old Product of the mutation.
func withProduct(node *Product) productOption {
	return func(m *ProductMutation) {
		m.oldValue = func(context.Context) (*Product, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProductMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProductMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Product entities.
func (m *ProductMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *ProductMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *ProductMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ProductMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Product entity.
// If the Product object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProductMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ProductMutation) ResetName() {
	m.name = nil
}

// SetCategoryRef sets the "categoryRef" field.
func (m *ProductMutation) SetCategoryRef(s string) {
	m.categoryRef = &s
}

// CategoryRef returns the value of the "categoryRef" field in the mutation.
func (m *ProductMutation) CategoryRef() (r string, exists bool) {
	v := m.categoryRef
	if v == nil {
		return
	}
	return *v, true
}

// OldCategoryRef returns the old "categoryRef" field's value of the Product entity.
// If the Product object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProductMutation) OldCategoryRef(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCategoryRef is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCategoryRef requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategoryRef: %w", err)
	}
	return oldValue.CategoryRef, nil
}

// ResetCategoryRef resets all changes to the "categoryRef" field.
func (m *ProductMutation) ResetCategoryRef() {
	m.categoryRef = nil
}

// AddVariantIDs adds the "variants" edge to the Variant entity by ids.
func (m *ProductMutation) AddVariantIDs(ids ...uuid.UUID) {
	if m.variants == nil {
		m.variants = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.variants[ids[i]] = struct{}{}
	}
}

// ClearVariants clears the "variants" edge to the Variant entity.
func (m *ProductMutation) ClearVariants() {
	m.clearedvariants = true
}

// VariantsCleared returns if the "variants" edge to the Variant entity was cleared.
func (m *ProductMutation) VariantsCleared() bool {
	return m.clearedvariants
}

// RemoveVariantIDs removes the "variants" edge to the Variant entity by IDs.
func (m *ProductMutation) RemoveVariantIDs(ids ...uuid.UUID) {
	if m.removedvariants == nil {
		m.removedvariants = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedvariants[ids[i]] = struct{}{}
	}
}

// RemovedVariants returns the removed IDs of the "variants" edge to the Variant entity.
func (m *ProductMutation) RemovedVariantsIDs() (ids []uuid.UUID) {
	for id := range m.removedvariants {
		ids = append(ids, id)
	}
	return
}

// VariantsIDs returns the "variants" edge IDs in the mutation.
func (m *ProductMutation) VariantsIDs() (ids []uuid.UUID) {
	for id := range m.variants {
		ids = append(ids, id)
	}
	return
}

// ResetVariants resets all changes to the "variants" edge.
func (m *ProductMutation) ResetVariants() {
	m.variants = nil
	m.clearedvariants = false
	m.removedvariants = nil
}

// Op returns the operation name.
func (m *ProductMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Product).
func (m *ProductMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProductMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, product.FieldName)
	}
	if m.categoryRef != nil {
		fields = append(fields, product.FieldCategoryRef)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProductMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case product.FieldName:
		return m.Name()
	case product.FieldCategoryRef:
		return m.CategoryRef()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProductMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case product.FieldName:
		return m.OldName(ctx)
	case product.FieldCategoryRef:
		return m.OldCategoryRef(ctx)
	}
	return nil, fmt.Errorf("unknown Product field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) SetField(name string, value ent.Value) error {
	switch name {
	case product.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case product.FieldCategoryRef:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategoryRef(v)
		return nil
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProductMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProductMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Product numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProductMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProductMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProductMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Product nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProductMutation) ResetField(name string) error {
	switch name {
	case product.FieldName:
		m.ResetName()
		return nil
	case product.FieldCategoryRef:
		m.ResetCategoryRef()
		return nil
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProductMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.variants != nil {
		edges = append(edges, product.EdgeVariants)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProductMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case product.EdgeVariants:
		ids := make([]ent.Value, 0, len(m.variants))
		for id := range m.variants {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProductMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedvariants != nil {
		edges = append(edges, product.EdgeVariants)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProductMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case product.EdgeVariants:
		ids := make([]ent.Value, 0, len(m.removedvariants))
		for id := range m.removedvariants {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProductMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedvariants {
		edges = append(edges, product.EdgeVariants)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProductMutation) EdgeCleared(name string) bool {
	switch name {
	case product.EdgeVariants:
		return m.clearedvariants
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProductMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Product unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProductMutation) ResetEdge(name string) error {
	switch name {
	case product.EdgeVariants:
		m.ResetVariants()
		return nil
	}
	return fmt.Errorf("unknown Product edge %s", name)
}

// VariantMutation represents an operation that mutates the Variant nodes in the graph.
type VariantMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	code           *string
	name           *string
	price          *float64
	addprice       *float64
	clearedFields  map[string]struct{}
	product        *uuid.UUID
	clearedproduct bool
	done           bool
	oldValue       func(context.Context) (*Variant, error)
	predicates     []predicate.Variant
}

var _ ent.Mutation = (*VariantMutation)(nil)

// variantOption allows management of the mutation configuration using functional options.
type variantOption func(*VariantMutation)

// newVariantMutation creates new mutation for the Variant entity.
func newVariantMutation(c config, op Op, opts ...variantOption) *VariantMutation {
	m := &VariantMutation{
		config:        c,
		op:            op,
		typ:           TypeVariant,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withVariantID sets the ID field of the mutation.
func withVariantID(id uuid.UUID) variantOption {
	return func(m *VariantMutation) {
		var (
			err   error
			once  sync.Once
			value *Variant
		)
		m.oldValue = func(ctx context.Context) (*Variant, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Variant.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withVariant sets the old Variant of the mutation.
func withVariant(node *Variant) variantOption {
	return func(m *VariantMutation) {
		m.oldValue = func(context.Context) (*Variant, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m VariantMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m VariantMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Variant entities.
func (m *VariantMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *VariantMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCode sets the "code" field.
func (m *VariantMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *VariantMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Variant entity.
// If the Variant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VariantMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *VariantMutation) ResetCode() {
	m.code = nil
}

// SetName sets the "name" field.
func (m *VariantMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *VariantMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Variant entity.
// If the Variant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VariantMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *VariantMutation) ResetName() {
	m.name = nil
}

// SetPrice sets the "price" field.
func (m *VariantMutation) SetPrice(f float64) {
	m.price = &f
	m.addprice = nil
}

// Price returns the value of the "price" field in the mutation.
func (m *VariantMutation) Price() (r float64, exists bool) {
	v := m.price
	if v == nil {
		return
	}
	return *v, true
}

// OldPrice returns the old "price" field's value of the Variant entity.
// If the Variant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *VariantMutation) OldPrice(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPrice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPrice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrice: %w", err)
	}
	return oldValue.Price, nil
}

// AddPrice adds f to the "price" field.
func (m *VariantMutation) AddPrice(f float64) {
	if m.addprice != nil {
		*m.addprice += f
	} else {
		m.addprice = &f
	}
}

// AddedPrice returns the value that was added to the "price" field in this mutation.
func (m *VariantMutation) AddedPrice() (r float64, exists bool) {
	v := m.addprice
	if v == nil {
		return
	}
	return *v, true
}

// ResetPrice resets all changes to the "price" field.
func (m *VariantMutation) ResetPrice() {
	m.price = nil
	m.addprice = nil
}

// SetProductID sets the "product" edge to the Product entity by id.
func (m *VariantMutation) SetProductID(id uuid.UUID) {
	m.product = &id
}

// ClearProduct clears the "product" edge to the Product entity.
func (m *VariantMutation) ClearProduct() {
	m.clearedproduct = true
}

// ProductCleared returns if the "product" edge to the Product entity was cleared.
func (m *VariantMutation) ProductCleared() bool {
	return m.clearedproduct
}

// ProductID returns the "product" edge ID in the mutation.
func (m *VariantMutation) ProductID() (id uuid.UUID, exists bool) {
	if m.product != nil {
		return *m.product, true
	}
	return
}

// ProductIDs returns the "product" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ProductID instead. It exists only for internal usage by the builders.
func (m *VariantMutation) ProductIDs() (ids []uuid.UUID) {
	if id := m.product; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetProduct resets all changes to the "product" edge.
func (m *VariantMutation) ResetProduct() {
	m.product = nil
	m.clearedproduct = false
}

// Op returns the operation name.
func (m *VariantMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Variant).
func (m *VariantMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *VariantMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.code != nil {
		fields = append(fields, variant.FieldCode)
	}
	if m.name != nil {
		fields = append(fields, variant.FieldName)
	}
	if m.price != nil {
		fields = append(fields, variant.FieldPrice)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *VariantMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case variant.FieldCode:
		return m.Code()
	case variant.FieldName:
		return m.Name()
	case variant.FieldPrice:
		return m.Price()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *VariantMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case variant.FieldCode:
		return m.OldCode(ctx)
	case variant.FieldName:
		return m.OldName(ctx)
	case variant.FieldPrice:
		return m.OldPrice(ctx)
	}
	return nil, fmt.Errorf("unknown Variant field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VariantMutation) SetField(name string, value ent.Value) error {
	switch name {
	case variant.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case variant.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case variant.FieldPrice:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Variant field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *VariantMutation) AddedFields() []string {
	var fields []string
	if m.addprice != nil {
		fields = append(fields, variant.FieldPrice)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *VariantMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case variant.FieldPrice:
		return m.AddedPrice()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *VariantMutation) AddField(name string, value ent.Value) error {
	switch name {
	case variant.FieldPrice:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Variant numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *VariantMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *VariantMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *VariantMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Variant nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *VariantMutation) ResetField(name string) error {
	switch name {
	case variant.FieldCode:
		m.ResetCode()
		return nil
	case variant.FieldName:
		m.ResetName()
		return nil
	case variant.FieldPrice:
		m.ResetPrice()
		return nil
	}
	return fmt.Errorf("unknown Variant field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *VariantMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.product != nil {
		edges = append(edges, variant.EdgeProduct)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *VariantMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case variant.EdgeProduct:
		if id := m.product; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *VariantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *VariantMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *VariantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedproduct {
		edges = append(edges, variant.EdgeProduct)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *VariantMutation) EdgeCleared(name string) bool {
	switch name {
	case variant.EdgeProduct:
		return m.clearedproduct
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *VariantMutation) ClearEdge(name string) error {
	switch name {
	case variant.EdgeProduct:
		m.ClearProduct()
		return nil
	}
	return fmt.Errorf("unknown Variant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *VariantMutation) ResetEdge(name string) error {
	switch name {
	case variant.EdgeProduct:
		m.ResetProduct()
		return nil
	}
	return fmt.Errorf("unknown Variant edge %s", name)
}
