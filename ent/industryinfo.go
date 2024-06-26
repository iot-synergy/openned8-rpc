// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/iot-synergy/openned8-rpc/ent/industryinfo"
)

// IndustryInfo is the model entity for the IndustryInfo schema.
type IndustryInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 名字
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IndustryInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case industryinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case industryinfo.FieldName:
			values[i] = new(sql.NullString)
		case industryinfo.FieldCreatedAt, industryinfo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IndustryInfo fields.
func (ii *IndustryInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case industryinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ii.ID = uint64(value.Int64)
		case industryinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ii.CreatedAt = value.Time
			}
		case industryinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ii.UpdatedAt = value.Time
			}
		case industryinfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ii.Name = value.String
			}
		default:
			ii.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the IndustryInfo.
// This includes values selected through modifiers, order, etc.
func (ii *IndustryInfo) Value(name string) (ent.Value, error) {
	return ii.selectValues.Get(name)
}

// Update returns a builder for updating this IndustryInfo.
// Note that you need to call IndustryInfo.Unwrap() before calling this method if this IndustryInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ii *IndustryInfo) Update() *IndustryInfoUpdateOne {
	return NewIndustryInfoClient(ii.config).UpdateOne(ii)
}

// Unwrap unwraps the IndustryInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ii *IndustryInfo) Unwrap() *IndustryInfo {
	_tx, ok := ii.config.driver.(*txDriver)
	if !ok {
		panic("ent: IndustryInfo is not a transactional entity")
	}
	ii.config.driver = _tx.drv
	return ii
}

// String implements the fmt.Stringer.
func (ii *IndustryInfo) String() string {
	var builder strings.Builder
	builder.WriteString("IndustryInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ii.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ii.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ii.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ii.Name)
	builder.WriteByte(')')
	return builder.String()
}

// IndustryInfos is a parsable slice of IndustryInfo.
type IndustryInfos []*IndustryInfo
