// Code generated by ent, DO NOT EDIT.

package autofieldmessage

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the autofieldmessage type in the database.
	Label = "auto_field_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// Table holds the table name of the autofieldmessage in the database.
	Table = "auto_field_messages"
)

// Columns holds all SQL columns for autofieldmessage fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
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

// OrderOption defines the ordering options for the AutoFieldMessage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}
