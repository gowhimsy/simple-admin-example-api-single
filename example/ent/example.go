// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-example-api/ent/example"
)

// Example is the model entity for the Example schema.
type Example struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Example) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case example.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Example fields.
func (e *Example) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case example.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Example.
// This includes values selected through modifiers, order, etc.
func (e *Example) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// Update returns a builder for updating this Example.
// Note that you need to call Example.Unwrap() before calling this method if this Example
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Example) Update() *ExampleUpdateOne {
	return NewExampleClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Example entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Example) Unwrap() *Example {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Example is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Example) String() string {
	var builder strings.Builder
	builder.WriteString("Example(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Examples is a parsable slice of Example.
type Examples []*Example
