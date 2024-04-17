// Code generated by ent, DO NOT EDIT.

package appinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUserID, v))
}

// AppName applies equality check predicate on the "app_name" field. It's identical to AppNameEQ.
func AppName(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppName, v))
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldSummary, v))
}

// AppCategory applies equality check predicate on the "app_category" field. It's identical to AppCategoryEQ.
func AppCategory(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppCategory, v))
}

// UseIndustry applies equality check predicate on the "use_industry" field. It's identical to UseIndustryEQ.
func UseIndustry(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUseIndustry, v))
}

// AppCategoryName applies equality check predicate on the "app_category_name" field. It's identical to AppCategoryNameEQ.
func AppCategoryName(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppCategoryName, v))
}

// UseIndustryName applies equality check predicate on the "use_industry_name" field. It's identical to UseIndustryNameEQ.
func UseIndustryName(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUseIndustryName, v))
}

// AppKey applies equality check predicate on the "app_key" field. It's identical to AppKeyEQ.
func AppKey(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppKey, v))
}

// AppSecret applies equality check predicate on the "app_secret" field. It's identical to AppSecretEQ.
func AppSecret(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppSecret, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldUserID, v))
}

// AppNameEQ applies the EQ predicate on the "app_name" field.
func AppNameEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppName, v))
}

// AppNameNEQ applies the NEQ predicate on the "app_name" field.
func AppNameNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldAppName, v))
}

// AppNameIn applies the In predicate on the "app_name" field.
func AppNameIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldAppName, vs...))
}

// AppNameNotIn applies the NotIn predicate on the "app_name" field.
func AppNameNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldAppName, vs...))
}

// AppNameGT applies the GT predicate on the "app_name" field.
func AppNameGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldAppName, v))
}

// AppNameGTE applies the GTE predicate on the "app_name" field.
func AppNameGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldAppName, v))
}

// AppNameLT applies the LT predicate on the "app_name" field.
func AppNameLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldAppName, v))
}

// AppNameLTE applies the LTE predicate on the "app_name" field.
func AppNameLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldAppName, v))
}

// AppNameContains applies the Contains predicate on the "app_name" field.
func AppNameContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldAppName, v))
}

// AppNameHasPrefix applies the HasPrefix predicate on the "app_name" field.
func AppNameHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldAppName, v))
}

// AppNameHasSuffix applies the HasSuffix predicate on the "app_name" field.
func AppNameHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldAppName, v))
}

// AppNameEqualFold applies the EqualFold predicate on the "app_name" field.
func AppNameEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldAppName, v))
}

// AppNameContainsFold applies the ContainsFold predicate on the "app_name" field.
func AppNameContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldAppName, v))
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldSummary, v))
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldSummary, v))
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldSummary, vs...))
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldSummary, vs...))
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldSummary, v))
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldSummary, v))
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldSummary, v))
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldSummary, v))
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldSummary, v))
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldSummary, v))
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldSummary, v))
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldSummary, v))
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldSummary, v))
}

// AppCategoryEQ applies the EQ predicate on the "app_category" field.
func AppCategoryEQ(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppCategory, v))
}

// AppCategoryNEQ applies the NEQ predicate on the "app_category" field.
func AppCategoryNEQ(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldAppCategory, v))
}

// AppCategoryIn applies the In predicate on the "app_category" field.
func AppCategoryIn(vs ...int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldAppCategory, vs...))
}

// AppCategoryNotIn applies the NotIn predicate on the "app_category" field.
func AppCategoryNotIn(vs ...int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldAppCategory, vs...))
}

// AppCategoryGT applies the GT predicate on the "app_category" field.
func AppCategoryGT(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldAppCategory, v))
}

// AppCategoryGTE applies the GTE predicate on the "app_category" field.
func AppCategoryGTE(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldAppCategory, v))
}

// AppCategoryLT applies the LT predicate on the "app_category" field.
func AppCategoryLT(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldAppCategory, v))
}

// AppCategoryLTE applies the LTE predicate on the "app_category" field.
func AppCategoryLTE(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldAppCategory, v))
}

// UseIndustryEQ applies the EQ predicate on the "use_industry" field.
func UseIndustryEQ(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUseIndustry, v))
}

// UseIndustryNEQ applies the NEQ predicate on the "use_industry" field.
func UseIndustryNEQ(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldUseIndustry, v))
}

// UseIndustryIn applies the In predicate on the "use_industry" field.
func UseIndustryIn(vs ...int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldUseIndustry, vs...))
}

// UseIndustryNotIn applies the NotIn predicate on the "use_industry" field.
func UseIndustryNotIn(vs ...int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldUseIndustry, vs...))
}

// UseIndustryGT applies the GT predicate on the "use_industry" field.
func UseIndustryGT(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldUseIndustry, v))
}

// UseIndustryGTE applies the GTE predicate on the "use_industry" field.
func UseIndustryGTE(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldUseIndustry, v))
}

// UseIndustryLT applies the LT predicate on the "use_industry" field.
func UseIndustryLT(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldUseIndustry, v))
}

// UseIndustryLTE applies the LTE predicate on the "use_industry" field.
func UseIndustryLTE(v int64) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldUseIndustry, v))
}

// AppCategoryNameEQ applies the EQ predicate on the "app_category_name" field.
func AppCategoryNameEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppCategoryName, v))
}

// AppCategoryNameNEQ applies the NEQ predicate on the "app_category_name" field.
func AppCategoryNameNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldAppCategoryName, v))
}

// AppCategoryNameIn applies the In predicate on the "app_category_name" field.
func AppCategoryNameIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldAppCategoryName, vs...))
}

// AppCategoryNameNotIn applies the NotIn predicate on the "app_category_name" field.
func AppCategoryNameNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldAppCategoryName, vs...))
}

// AppCategoryNameGT applies the GT predicate on the "app_category_name" field.
func AppCategoryNameGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldAppCategoryName, v))
}

// AppCategoryNameGTE applies the GTE predicate on the "app_category_name" field.
func AppCategoryNameGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldAppCategoryName, v))
}

// AppCategoryNameLT applies the LT predicate on the "app_category_name" field.
func AppCategoryNameLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldAppCategoryName, v))
}

// AppCategoryNameLTE applies the LTE predicate on the "app_category_name" field.
func AppCategoryNameLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldAppCategoryName, v))
}

// AppCategoryNameContains applies the Contains predicate on the "app_category_name" field.
func AppCategoryNameContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldAppCategoryName, v))
}

// AppCategoryNameHasPrefix applies the HasPrefix predicate on the "app_category_name" field.
func AppCategoryNameHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldAppCategoryName, v))
}

// AppCategoryNameHasSuffix applies the HasSuffix predicate on the "app_category_name" field.
func AppCategoryNameHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldAppCategoryName, v))
}

// AppCategoryNameEqualFold applies the EqualFold predicate on the "app_category_name" field.
func AppCategoryNameEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldAppCategoryName, v))
}

// AppCategoryNameContainsFold applies the ContainsFold predicate on the "app_category_name" field.
func AppCategoryNameContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldAppCategoryName, v))
}

// UseIndustryNameEQ applies the EQ predicate on the "use_industry_name" field.
func UseIndustryNameEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldUseIndustryName, v))
}

// UseIndustryNameNEQ applies the NEQ predicate on the "use_industry_name" field.
func UseIndustryNameNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldUseIndustryName, v))
}

// UseIndustryNameIn applies the In predicate on the "use_industry_name" field.
func UseIndustryNameIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldUseIndustryName, vs...))
}

// UseIndustryNameNotIn applies the NotIn predicate on the "use_industry_name" field.
func UseIndustryNameNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldUseIndustryName, vs...))
}

// UseIndustryNameGT applies the GT predicate on the "use_industry_name" field.
func UseIndustryNameGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldUseIndustryName, v))
}

// UseIndustryNameGTE applies the GTE predicate on the "use_industry_name" field.
func UseIndustryNameGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldUseIndustryName, v))
}

// UseIndustryNameLT applies the LT predicate on the "use_industry_name" field.
func UseIndustryNameLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldUseIndustryName, v))
}

// UseIndustryNameLTE applies the LTE predicate on the "use_industry_name" field.
func UseIndustryNameLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldUseIndustryName, v))
}

// UseIndustryNameContains applies the Contains predicate on the "use_industry_name" field.
func UseIndustryNameContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldUseIndustryName, v))
}

// UseIndustryNameHasPrefix applies the HasPrefix predicate on the "use_industry_name" field.
func UseIndustryNameHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldUseIndustryName, v))
}

// UseIndustryNameHasSuffix applies the HasSuffix predicate on the "use_industry_name" field.
func UseIndustryNameHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldUseIndustryName, v))
}

// UseIndustryNameEqualFold applies the EqualFold predicate on the "use_industry_name" field.
func UseIndustryNameEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldUseIndustryName, v))
}

// UseIndustryNameContainsFold applies the ContainsFold predicate on the "use_industry_name" field.
func UseIndustryNameContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldUseIndustryName, v))
}

// AppKeyEQ applies the EQ predicate on the "app_key" field.
func AppKeyEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppKey, v))
}

// AppKeyNEQ applies the NEQ predicate on the "app_key" field.
func AppKeyNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldAppKey, v))
}

// AppKeyIn applies the In predicate on the "app_key" field.
func AppKeyIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldAppKey, vs...))
}

// AppKeyNotIn applies the NotIn predicate on the "app_key" field.
func AppKeyNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldAppKey, vs...))
}

// AppKeyGT applies the GT predicate on the "app_key" field.
func AppKeyGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldAppKey, v))
}

// AppKeyGTE applies the GTE predicate on the "app_key" field.
func AppKeyGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldAppKey, v))
}

// AppKeyLT applies the LT predicate on the "app_key" field.
func AppKeyLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldAppKey, v))
}

// AppKeyLTE applies the LTE predicate on the "app_key" field.
func AppKeyLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldAppKey, v))
}

// AppKeyContains applies the Contains predicate on the "app_key" field.
func AppKeyContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldAppKey, v))
}

// AppKeyHasPrefix applies the HasPrefix predicate on the "app_key" field.
func AppKeyHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldAppKey, v))
}

// AppKeyHasSuffix applies the HasSuffix predicate on the "app_key" field.
func AppKeyHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldAppKey, v))
}

// AppKeyEqualFold applies the EqualFold predicate on the "app_key" field.
func AppKeyEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldAppKey, v))
}

// AppKeyContainsFold applies the ContainsFold predicate on the "app_key" field.
func AppKeyContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldAppKey, v))
}

// AppSecretEQ applies the EQ predicate on the "app_secret" field.
func AppSecretEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEQ(FieldAppSecret, v))
}

// AppSecretNEQ applies the NEQ predicate on the "app_secret" field.
func AppSecretNEQ(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNEQ(FieldAppSecret, v))
}

// AppSecretIn applies the In predicate on the "app_secret" field.
func AppSecretIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldIn(FieldAppSecret, vs...))
}

// AppSecretNotIn applies the NotIn predicate on the "app_secret" field.
func AppSecretNotIn(vs ...string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldNotIn(FieldAppSecret, vs...))
}

// AppSecretGT applies the GT predicate on the "app_secret" field.
func AppSecretGT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGT(FieldAppSecret, v))
}

// AppSecretGTE applies the GTE predicate on the "app_secret" field.
func AppSecretGTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldGTE(FieldAppSecret, v))
}

// AppSecretLT applies the LT predicate on the "app_secret" field.
func AppSecretLT(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLT(FieldAppSecret, v))
}

// AppSecretLTE applies the LTE predicate on the "app_secret" field.
func AppSecretLTE(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldLTE(FieldAppSecret, v))
}

// AppSecretContains applies the Contains predicate on the "app_secret" field.
func AppSecretContains(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContains(FieldAppSecret, v))
}

// AppSecretHasPrefix applies the HasPrefix predicate on the "app_secret" field.
func AppSecretHasPrefix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasPrefix(FieldAppSecret, v))
}

// AppSecretHasSuffix applies the HasSuffix predicate on the "app_secret" field.
func AppSecretHasSuffix(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldHasSuffix(FieldAppSecret, v))
}

// AppSecretEqualFold applies the EqualFold predicate on the "app_secret" field.
func AppSecretEqualFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldEqualFold(FieldAppSecret, v))
}

// AppSecretContainsFold applies the ContainsFold predicate on the "app_secret" field.
func AppSecretContainsFold(v string) predicate.AppInfo {
	return predicate.AppInfo(sql.FieldContainsFold(FieldAppSecret, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppInfo) predicate.AppInfo {
	return predicate.AppInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppInfo) predicate.AppInfo {
	return predicate.AppInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppInfo) predicate.AppInfo {
	return predicate.AppInfo(sql.NotPredicates(p))
}
