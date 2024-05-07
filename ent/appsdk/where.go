// Code generated by ent, DO NOT EDIT.

package appsdk

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldUpdatedAt, v))
}

// App applies equality check predicate on the "app" field. It's identical to AppEQ.
func App(v uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldApp, v))
}

// Sdk applies equality check predicate on the "sdk" field. It's identical to SdkEQ.
func Sdk(v uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldSdk, v))
}

// SdkKey applies equality check predicate on the "sdk_key" field. It's identical to SdkKeyEQ.
func SdkKey(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldSdkKey, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLTE(FieldUpdatedAt, v))
}

// AppEQ applies the EQ predicate on the "app" field.
func AppEQ(v uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldApp, v))
}

// AppNEQ applies the NEQ predicate on the "app" field.
func AppNEQ(v uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNEQ(FieldApp, v))
}

// AppIn applies the In predicate on the "app" field.
func AppIn(vs ...uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIn(FieldApp, vs...))
}

// AppNotIn applies the NotIn predicate on the "app" field.
func AppNotIn(vs ...uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotIn(FieldApp, vs...))
}

// AppIsNil applies the IsNil predicate on the "app" field.
func AppIsNil() predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIsNull(FieldApp))
}

// AppNotNil applies the NotNil predicate on the "app" field.
func AppNotNil() predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotNull(FieldApp))
}

// SdkEQ applies the EQ predicate on the "sdk" field.
func SdkEQ(v uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldSdk, v))
}

// SdkNEQ applies the NEQ predicate on the "sdk" field.
func SdkNEQ(v uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNEQ(FieldSdk, v))
}

// SdkIn applies the In predicate on the "sdk" field.
func SdkIn(vs ...uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIn(FieldSdk, vs...))
}

// SdkNotIn applies the NotIn predicate on the "sdk" field.
func SdkNotIn(vs ...uuid.UUID) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotIn(FieldSdk, vs...))
}

// SdkIsNil applies the IsNil predicate on the "sdk" field.
func SdkIsNil() predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIsNull(FieldSdk))
}

// SdkNotNil applies the NotNil predicate on the "sdk" field.
func SdkNotNil() predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotNull(FieldSdk))
}

// SdkKeyEQ applies the EQ predicate on the "sdk_key" field.
func SdkKeyEQ(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEQ(FieldSdkKey, v))
}

// SdkKeyNEQ applies the NEQ predicate on the "sdk_key" field.
func SdkKeyNEQ(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNEQ(FieldSdkKey, v))
}

// SdkKeyIn applies the In predicate on the "sdk_key" field.
func SdkKeyIn(vs ...string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldIn(FieldSdkKey, vs...))
}

// SdkKeyNotIn applies the NotIn predicate on the "sdk_key" field.
func SdkKeyNotIn(vs ...string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldNotIn(FieldSdkKey, vs...))
}

// SdkKeyGT applies the GT predicate on the "sdk_key" field.
func SdkKeyGT(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGT(FieldSdkKey, v))
}

// SdkKeyGTE applies the GTE predicate on the "sdk_key" field.
func SdkKeyGTE(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldGTE(FieldSdkKey, v))
}

// SdkKeyLT applies the LT predicate on the "sdk_key" field.
func SdkKeyLT(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLT(FieldSdkKey, v))
}

// SdkKeyLTE applies the LTE predicate on the "sdk_key" field.
func SdkKeyLTE(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldLTE(FieldSdkKey, v))
}

// SdkKeyContains applies the Contains predicate on the "sdk_key" field.
func SdkKeyContains(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldContains(FieldSdkKey, v))
}

// SdkKeyHasPrefix applies the HasPrefix predicate on the "sdk_key" field.
func SdkKeyHasPrefix(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldHasPrefix(FieldSdkKey, v))
}

// SdkKeyHasSuffix applies the HasSuffix predicate on the "sdk_key" field.
func SdkKeyHasSuffix(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldHasSuffix(FieldSdkKey, v))
}

// SdkKeyEqualFold applies the EqualFold predicate on the "sdk_key" field.
func SdkKeyEqualFold(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldEqualFold(FieldSdkKey, v))
}

// SdkKeyContainsFold applies the ContainsFold predicate on the "sdk_key" field.
func SdkKeyContainsFold(v string) predicate.AppSdk {
	return predicate.AppSdk(sql.FieldContainsFold(FieldSdkKey, v))
}

// HasActiveCode applies the HasEdge predicate on the "active_code" edge.
func HasActiveCode() predicate.AppSdk {
	return predicate.AppSdk(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActiveCodeTable, ActiveCodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActiveCodeWith applies the HasEdge predicate on the "active_code" edge with a given conditions (other predicates).
func HasActiveCodeWith(preds ...predicate.ActiveCodeInfo) predicate.AppSdk {
	return predicate.AppSdk(func(s *sql.Selector) {
		step := newActiveCodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAppInfo applies the HasEdge predicate on the "app_info" edge.
func HasAppInfo() predicate.AppSdk {
	return predicate.AppSdk(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppInfoTable, AppInfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppInfoWith applies the HasEdge predicate on the "app_info" edge with a given conditions (other predicates).
func HasAppInfoWith(preds ...predicate.AppInfo) predicate.AppSdk {
	return predicate.AppSdk(func(s *sql.Selector) {
		step := newAppInfoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSdkInfo applies the HasEdge predicate on the "sdk_info" edge.
func HasSdkInfo() predicate.AppSdk {
	return predicate.AppSdk(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SdkInfoTable, SdkInfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSdkInfoWith applies the HasEdge predicate on the "sdk_info" edge with a given conditions (other predicates).
func HasSdkInfoWith(preds ...predicate.SdkInfo) predicate.AppSdk {
	return predicate.AppSdk(func(s *sql.Selector) {
		step := newSdkInfoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppSdk) predicate.AppSdk {
	return predicate.AppSdk(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppSdk) predicate.AppSdk {
	return predicate.AppSdk(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppSdk) predicate.AppSdk {
	return predicate.AppSdk(sql.NotPredicates(p))
}
