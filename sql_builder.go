package bsql

import "strings"

// SQLBuilder constructs SQL queries with placeholders and corresponding values.
type SQLBuilder struct {
	strBuilder *strings.Builder
	values     []interface{}
}

// NewSQLBuilder creates and returns a new SQLBuilder instance.
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{
		strBuilder: &strings.Builder{},
		values:     []interface{}{},
	}
}

// Reset clears the current SQL string and associated values.
func (sb *SQLBuilder) Reset() *SQLBuilder {
	sb.strBuilder.Reset()
	sb.values = sb.values[:0]
	return sb
}

// WriteIf conditionally writes a SQL part and its values to the builder.
func (sb *SQLBuilder) WriteIf(condition bool, sqlPart string, values ...interface{}) *SQLBuilder {
	if condition {
		sb.strBuilder.WriteString(sqlPart)
		sb.AddValues(values...)
	}
	return sb
}

// Write writes a SQL part and its values to the builder.
func (sb *SQLBuilder) Write(sqlPart string, values ...interface{}) *SQLBuilder {
	sb.strBuilder.WriteString(sqlPart)
	sb.AddValues(values...)
	return sb
}

// WriteString writes only the SQL part to the builder.
func (sb *SQLBuilder) WriteString(sqlPart string) *SQLBuilder {
	sb.strBuilder.WriteString(sqlPart)
	return sb
}

// AddValues appends values to the builder's value list.
func (sb *SQLBuilder) AddValues(values ...interface{}) *SQLBuilder {
	sb.values = append(sb.values, values...)
	return sb
}

// SQL returns the constructed SQL string.
func (sb *SQLBuilder) SQL() string {
	return sb.strBuilder.String()
}

// Values returns the list of values for the SQL placeholders.
func (sb *SQLBuilder) Values() []interface{} {
	return sb.values
}
