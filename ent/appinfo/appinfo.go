// Code generated by ent, DO NOT EDIT.

package appinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the appinfo type in the database.
	Label = "app_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAppName holds the string denoting the app_name field in the database.
	FieldAppName = "app_name"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldAppCategory holds the string denoting the app_category field in the database.
	FieldAppCategory = "app_category"
	// FieldUseIndustry holds the string denoting the use_industry field in the database.
	FieldUseIndustry = "use_industry"
	// FieldAppCategoryName holds the string denoting the app_category_name field in the database.
	FieldAppCategoryName = "app_category_name"
	// FieldUseIndustryName holds the string denoting the use_industry_name field in the database.
	FieldUseIndustryName = "use_industry_name"
	// FieldAppKey holds the string denoting the app_key field in the database.
	FieldAppKey = "app_key"
	// FieldAppSecret holds the string denoting the app_secret field in the database.
	FieldAppSecret = "app_secret"
	// EdgeAppSdk holds the string denoting the app_sdk edge name in mutations.
	EdgeAppSdk = "app_sdk"
	// Table holds the table name of the appinfo in the database.
	Table = "app_info"
	// AppSdkTable is the table that holds the app_sdk relation/edge.
	AppSdkTable = "app_sdk"
	// AppSdkInverseTable is the table name for the AppSdk entity.
	// It exists in this package in order to avoid circular dependency with the "appsdk" package.
	AppSdkInverseTable = "app_sdk"
	// AppSdkColumn is the table column denoting the app_sdk relation/edge.
	AppSdkColumn = "app"
)

// Columns holds all SQL columns for appinfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldAppName,
	FieldSummary,
	FieldAppCategory,
	FieldUseIndustry,
	FieldAppCategoryName,
	FieldUseIndustryName,
	FieldAppKey,
	FieldAppSecret,
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
)

// OrderOption defines the ordering options for the AppInfo queries.
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

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByAppName orders the results by the app_name field.
func ByAppName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppName, opts...).ToFunc()
}

// BySummary orders the results by the summary field.
func BySummary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSummary, opts...).ToFunc()
}

// ByAppCategory orders the results by the app_category field.
func ByAppCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppCategory, opts...).ToFunc()
}

// ByUseIndustry orders the results by the use_industry field.
func ByUseIndustry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUseIndustry, opts...).ToFunc()
}

// ByAppCategoryName orders the results by the app_category_name field.
func ByAppCategoryName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppCategoryName, opts...).ToFunc()
}

// ByUseIndustryName orders the results by the use_industry_name field.
func ByUseIndustryName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUseIndustryName, opts...).ToFunc()
}

// ByAppKey orders the results by the app_key field.
func ByAppKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppKey, opts...).ToFunc()
}

// ByAppSecret orders the results by the app_secret field.
func ByAppSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppSecret, opts...).ToFunc()
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
