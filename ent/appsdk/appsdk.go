// Code generated by ent, DO NOT EDIT.

package appsdk

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the appsdk type in the database.
	Label = "app_sdk"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldApp holds the string denoting the app field in the database.
	FieldApp = "app"
	// FieldSdk holds the string denoting the sdk field in the database.
	FieldSdk = "sdk"
	// FieldSdkKey holds the string denoting the sdk_key field in the database.
	FieldSdkKey = "sdk_key"
	// EdgeActiveCode holds the string denoting the active_code edge name in mutations.
	EdgeActiveCode = "active_code"
	// EdgeAppInfo holds the string denoting the app_info edge name in mutations.
	EdgeAppInfo = "app_info"
	// EdgeSdkInfo holds the string denoting the sdk_info edge name in mutations.
	EdgeSdkInfo = "sdk_info"
	// Table holds the table name of the appsdk in the database.
	Table = "app_sdk"
	// ActiveCodeTable is the table that holds the active_code relation/edge.
	ActiveCodeTable = "active_code_info"
	// ActiveCodeInverseTable is the table name for the ActiveCodeInfo entity.
	// It exists in this package in order to avoid circular dependency with the "activecodeinfo" package.
	ActiveCodeInverseTable = "active_code_info"
	// ActiveCodeColumn is the table column denoting the active_code relation/edge.
	ActiveCodeColumn = "app_skd_id"
	// AppInfoTable is the table that holds the app_info relation/edge.
	AppInfoTable = "app_sdk"
	// AppInfoInverseTable is the table name for the AppInfo entity.
	// It exists in this package in order to avoid circular dependency with the "appinfo" package.
	AppInfoInverseTable = "app_info"
	// AppInfoColumn is the table column denoting the app_info relation/edge.
	AppInfoColumn = "app"
	// SdkInfoTable is the table that holds the sdk_info relation/edge.
	SdkInfoTable = "app_sdk"
	// SdkInfoInverseTable is the table name for the SdkInfo entity.
	// It exists in this package in order to avoid circular dependency with the "sdkinfo" package.
	SdkInfoInverseTable = "sdk_info"
	// SdkInfoColumn is the table column denoting the sdk_info relation/edge.
	SdkInfoColumn = "sdk"
)

// Columns holds all SQL columns for appsdk fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldApp,
	FieldSdk,
	FieldSdkKey,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// SdkKeyValidator is a validator for the "sdk_key" field. It is called by the builders before save.
	SdkKeyValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the AppSdk queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByApp orders the results by the app field.
func ByApp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApp, opts...).ToFunc()
}

// BySdk orders the results by the sdk field.
func BySdk(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSdk, opts...).ToFunc()
}

// BySdkKey orders the results by the sdk_key field.
func BySdkKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSdkKey, opts...).ToFunc()
}

// ByActiveCodeCount orders the results by active_code count.
func ByActiveCodeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActiveCodeStep(), opts...)
	}
}

// ByActiveCode orders the results by active_code terms.
func ByActiveCode(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActiveCodeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAppInfoField orders the results by app_info field.
func ByAppInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppInfoStep(), sql.OrderByField(field, opts...))
	}
}

// BySdkInfoField orders the results by sdk_info field.
func BySdkInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSdkInfoStep(), sql.OrderByField(field, opts...))
	}
}
func newActiveCodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActiveCodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ActiveCodeTable, ActiveCodeColumn),
	)
}
func newAppInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AppInfoTable, AppInfoColumn),
	)
}
func newSdkInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SdkInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SdkInfoTable, SdkInfoColumn),
	)
}
