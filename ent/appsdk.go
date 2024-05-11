// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
)

// AppSdk is the model entity for the AppSdk schema.
type AppSdk struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// app的id
	App string `json:"app,omitempty"`
	// sdk的id
	Sdk uuid.UUID `json:"sdk,omitempty"`
	// 分配给app的sdk
	SdkKey string `json:"sdk_key,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppSdkQuery when eager-loading is set.
	Edges        AppSdkEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AppSdkEdges holds the relations/edges for other nodes in the graph.
type AppSdkEdges struct {
	// ActiveCode holds the value of the active_code edge.
	ActiveCode []*ActiveCodeInfo `json:"active_code,omitempty"`
	// AppInfo holds the value of the app_info edge.
	AppInfo *AppInfo `json:"app_info,omitempty"`
	// SdkInfo holds the value of the sdk_info edge.
	SdkInfo *SdkInfo `json:"sdk_info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ActiveCodeOrErr returns the ActiveCode value or an error if the edge
// was not loaded in eager-loading.
func (e AppSdkEdges) ActiveCodeOrErr() ([]*ActiveCodeInfo, error) {
	if e.loadedTypes[0] {
		return e.ActiveCode, nil
	}
	return nil, &NotLoadedError{edge: "active_code"}
}

// AppInfoOrErr returns the AppInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppSdkEdges) AppInfoOrErr() (*AppInfo, error) {
	if e.AppInfo != nil {
		return e.AppInfo, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: appinfo.Label}
	}
	return nil, &NotLoadedError{edge: "app_info"}
}

// SdkInfoOrErr returns the SdkInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppSdkEdges) SdkInfoOrErr() (*SdkInfo, error) {
	if e.SdkInfo != nil {
		return e.SdkInfo, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: sdkinfo.Label}
	}
	return nil, &NotLoadedError{edge: "sdk_info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppSdk) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appsdk.FieldApp, appsdk.FieldSdkKey:
			values[i] = new(sql.NullString)
		case appsdk.FieldCreatedAt, appsdk.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case appsdk.FieldID, appsdk.FieldSdk:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppSdk fields.
func (as *AppSdk) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appsdk.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				as.ID = *value
			}
		case appsdk.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = value.Time
			}
		case appsdk.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				as.UpdatedAt = value.Time
			}
		case appsdk.FieldApp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app", values[i])
			} else if value.Valid {
				as.App = value.String
			}
		case appsdk.FieldSdk:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field sdk", values[i])
			} else if value != nil {
				as.Sdk = *value
			}
		case appsdk.FieldSdkKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sdk_key", values[i])
			} else if value.Valid {
				as.SdkKey = value.String
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppSdk.
// This includes values selected through modifiers, order, etc.
func (as *AppSdk) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// QueryActiveCode queries the "active_code" edge of the AppSdk entity.
func (as *AppSdk) QueryActiveCode() *ActiveCodeInfoQuery {
	return NewAppSdkClient(as.config).QueryActiveCode(as)
}

// QueryAppInfo queries the "app_info" edge of the AppSdk entity.
func (as *AppSdk) QueryAppInfo() *AppInfoQuery {
	return NewAppSdkClient(as.config).QueryAppInfo(as)
}

// QuerySdkInfo queries the "sdk_info" edge of the AppSdk entity.
func (as *AppSdk) QuerySdkInfo() *SdkInfoQuery {
	return NewAppSdkClient(as.config).QuerySdkInfo(as)
}

// Update returns a builder for updating this AppSdk.
// Note that you need to call AppSdk.Unwrap() before calling this method if this AppSdk
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AppSdk) Update() *AppSdkUpdateOne {
	return NewAppSdkClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AppSdk entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AppSdk) Unwrap() *AppSdk {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppSdk is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AppSdk) String() string {
	var builder strings.Builder
	builder.WriteString("AppSdk(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("created_at=")
	builder.WriteString(as.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(as.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app=")
	builder.WriteString(as.App)
	builder.WriteString(", ")
	builder.WriteString("sdk=")
	builder.WriteString(fmt.Sprintf("%v", as.Sdk))
	builder.WriteString(", ")
	builder.WriteString("sdk_key=")
	builder.WriteString(as.SdkKey)
	builder.WriteByte(')')
	return builder.String()
}

// AppSdks is a parsable slice of AppSdk.
type AppSdks []*AppSdk
