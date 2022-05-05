// Code generated by entc, DO NOT EDIT.

package stocks

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gushikem01/usa-kabu-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// NameJa applies equality check predicate on the "name_ja" field. It's identical to NameJaEQ.
func NameJa(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameJa), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Exchange applies equality check predicate on the "exchange" field. It's identical to ExchangeEQ.
func Exchange(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExchange), v))
	})
}

// ExchangeShortName applies equality check predicate on the "exchange_short_name" field. It's identical to ExchangeShortNameEQ.
func ExchangeShortName(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExchangeShortName), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymbol), v))
	})
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymbol), v...))
	})
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymbol), v...))
	})
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymbol), v))
	})
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymbol), v))
	})
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymbol), v))
	})
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymbol), v))
	})
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymbol), v))
	})
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymbol), v))
	})
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymbol), v))
	})
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymbol), v))
	})
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymbol), v))
	})
}

// NameJaEQ applies the EQ predicate on the "name_ja" field.
func NameJaEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameJa), v))
	})
}

// NameJaNEQ applies the NEQ predicate on the "name_ja" field.
func NameJaNEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameJa), v))
	})
}

// NameJaIn applies the In predicate on the "name_ja" field.
func NameJaIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNameJa), v...))
	})
}

// NameJaNotIn applies the NotIn predicate on the "name_ja" field.
func NameJaNotIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNameJa), v...))
	})
}

// NameJaGT applies the GT predicate on the "name_ja" field.
func NameJaGT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameJa), v))
	})
}

// NameJaGTE applies the GTE predicate on the "name_ja" field.
func NameJaGTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameJa), v))
	})
}

// NameJaLT applies the LT predicate on the "name_ja" field.
func NameJaLT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameJa), v))
	})
}

// NameJaLTE applies the LTE predicate on the "name_ja" field.
func NameJaLTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameJa), v))
	})
}

// NameJaContains applies the Contains predicate on the "name_ja" field.
func NameJaContains(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameJa), v))
	})
}

// NameJaHasPrefix applies the HasPrefix predicate on the "name_ja" field.
func NameJaHasPrefix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameJa), v))
	})
}

// NameJaHasSuffix applies the HasSuffix predicate on the "name_ja" field.
func NameJaHasSuffix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameJa), v))
	})
}

// NameJaIsNil applies the IsNil predicate on the "name_ja" field.
func NameJaIsNil() predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNameJa)))
	})
}

// NameJaNotNil applies the NotNil predicate on the "name_ja" field.
func NameJaNotNil() predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNameJa)))
	})
}

// NameJaEqualFold applies the EqualFold predicate on the "name_ja" field.
func NameJaEqualFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameJa), v))
	})
}

// NameJaContainsFold applies the ContainsFold predicate on the "name_ja" field.
func NameJaContainsFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameJa), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// ExchangeEQ applies the EQ predicate on the "exchange" field.
func ExchangeEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExchange), v))
	})
}

// ExchangeNEQ applies the NEQ predicate on the "exchange" field.
func ExchangeNEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExchange), v))
	})
}

// ExchangeIn applies the In predicate on the "exchange" field.
func ExchangeIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExchange), v...))
	})
}

// ExchangeNotIn applies the NotIn predicate on the "exchange" field.
func ExchangeNotIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExchange), v...))
	})
}

// ExchangeGT applies the GT predicate on the "exchange" field.
func ExchangeGT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExchange), v))
	})
}

// ExchangeGTE applies the GTE predicate on the "exchange" field.
func ExchangeGTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExchange), v))
	})
}

// ExchangeLT applies the LT predicate on the "exchange" field.
func ExchangeLT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExchange), v))
	})
}

// ExchangeLTE applies the LTE predicate on the "exchange" field.
func ExchangeLTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExchange), v))
	})
}

// ExchangeContains applies the Contains predicate on the "exchange" field.
func ExchangeContains(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExchange), v))
	})
}

// ExchangeHasPrefix applies the HasPrefix predicate on the "exchange" field.
func ExchangeHasPrefix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExchange), v))
	})
}

// ExchangeHasSuffix applies the HasSuffix predicate on the "exchange" field.
func ExchangeHasSuffix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExchange), v))
	})
}

// ExchangeEqualFold applies the EqualFold predicate on the "exchange" field.
func ExchangeEqualFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExchange), v))
	})
}

// ExchangeContainsFold applies the ContainsFold predicate on the "exchange" field.
func ExchangeContainsFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExchange), v))
	})
}

// ExchangeShortNameEQ applies the EQ predicate on the "exchange_short_name" field.
func ExchangeShortNameEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameNEQ applies the NEQ predicate on the "exchange_short_name" field.
func ExchangeShortNameNEQ(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameIn applies the In predicate on the "exchange_short_name" field.
func ExchangeShortNameIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExchangeShortName), v...))
	})
}

// ExchangeShortNameNotIn applies the NotIn predicate on the "exchange_short_name" field.
func ExchangeShortNameNotIn(vs ...string) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExchangeShortName), v...))
	})
}

// ExchangeShortNameGT applies the GT predicate on the "exchange_short_name" field.
func ExchangeShortNameGT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameGTE applies the GTE predicate on the "exchange_short_name" field.
func ExchangeShortNameGTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameLT applies the LT predicate on the "exchange_short_name" field.
func ExchangeShortNameLT(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameLTE applies the LTE predicate on the "exchange_short_name" field.
func ExchangeShortNameLTE(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameContains applies the Contains predicate on the "exchange_short_name" field.
func ExchangeShortNameContains(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameHasPrefix applies the HasPrefix predicate on the "exchange_short_name" field.
func ExchangeShortNameHasPrefix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameHasSuffix applies the HasSuffix predicate on the "exchange_short_name" field.
func ExchangeShortNameHasSuffix(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameEqualFold applies the EqualFold predicate on the "exchange_short_name" field.
func ExchangeShortNameEqualFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExchangeShortName), v))
	})
}

// ExchangeShortNameContainsFold applies the ContainsFold predicate on the "exchange_short_name" field.
func ExchangeShortNameContainsFold(v string) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExchangeShortName), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Stocks {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Stocks(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasStocks applies the HasEdge predicate on the "stocks" edge.
func HasStocks() predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StocksTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StocksTable, StocksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStocksWith applies the HasEdge predicate on the "stocks" edge with a given conditions (other predicates).
func HasStocksWith(preds ...predicate.Stocks) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StocksTable, StocksPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stocks) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stocks) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Stocks) predicate.Stocks {
	return predicate.Stocks(func(s *sql.Selector) {
		p(s.Not())
	})
}