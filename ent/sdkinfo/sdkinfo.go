// Code generated by ent, DO NOT EDIT.

package sdkinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	uuid "github.com/gofrs/uuid/v5"
)

const (
	// Label holds the string label denoting the sdkinfo type in the database.
	Label = "sdk_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldDownloadURL holds the string denoting the download_url field in the database.
	FieldDownloadURL = "download_url"
	// EdgeAppSdk holds the string denoting the app_sdk edge name in mutations.
	EdgeAppSdk = "app_sdk"
	// Table holds the table name of the sdkinfo in the database.
	Table = "sdk_info"
	// AppSdkTable is the table that holds the app_sdk relation/edge.
	AppSdkTable = "app_sdk"
	// AppSdkInverseTable is the table name for the AppSdk entity.
	// It exists in this package in order to avoid circular dependency with the "appsdk" package.
	AppSdkInverseTable = "app_sdk"
	// AppSdkColumn is the table column denoting the app_sdk relation/edge.
	AppSdkColumn = "sdk"
)

// Columns holds all SQL columns for sdkinfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldAvatar,
	FieldDesc,
	FieldDownloadURL,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the SdkInfo queries.
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

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByDownloadURL orders the results by the download_url field.
func ByDownloadURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadURL, opts...).ToFunc()
}

// ByAppSdkCount orders the results by app_sdk count.
func ByAppSdkCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppSdkStep(), opts...)
	}
}

// ByAppSdk orders the results by app_sdk terms.
func ByAppSdk(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppSdkStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAppSdkStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppSdkInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AppSdkTable, AppSdkColumn),
	)
}
