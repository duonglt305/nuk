package db

import (
	"fmt"
	"strings"
)

type Specification interface {
	GetQuery(idx ...int) string
	GetValues() []any
}

type joinSpecification struct {
	specs     []Specification
	separator string
}

func (s joinSpecification) GetQuery(idx ...int) string {
	var queries []string
	for ind, spec := range s.specs {
		queries = append(queries, spec.GetQuery(ind+1))
	}
	return strings.Join(queries, fmt.Sprintf(" %s ", s.separator))
}

func (s joinSpecification) GetValues() []any {
	var values []any
	for _, spec := range s.specs {
		values = append(values, spec.GetValues()...)
	}
	return values
}

type notSpecification struct {
	Specification
}

func (s notSpecification) GetQuery(idx ...int) string {
	return fmt.Sprintf("NOT (%s)", s.Specification.GetQuery())
}

type groupSpecification struct {
	specs []Specification
}

func (s groupSpecification) GetQuery(idx ...int) string {
	var queries []string
	for _, spec := range s.specs {
		queries = append(queries, spec.GetQuery())
	}
	return fmt.Sprintf("(%s)", strings.Join(queries, " "))
}

func (s groupSpecification) GetValues() []any {
	var values []any
	for _, spec := range s.specs {
		values = append(values, spec.GetValues()...)
	}
	return values
}

type binarySpecification[T any] struct {
	field    string
	operator string
	value    T
}

func (s binarySpecification[T]) GetQuery(idx ...int) string {
	ind := 1
	if len(idx) > 0 {
		ind = idx[0]
	}
	return fmt.Sprintf("%s %s $%d", s.field, s.operator, ind)
}

func (s binarySpecification[T]) GetValues() []any {
	return []any{s.value}
}

func Eq[T any](field string, value T) Specification {
	return binarySpecification[T]{field, "=", value}
}

func Ne[T any](field string, value T) Specification {
	return binarySpecification[T]{field, "!=", value}
}

func Gt[T any](field string, value T) Specification {
	return binarySpecification[T]{field, ">", value}
}

func Gte[T any](field string, value T) Specification {
	return binarySpecification[T]{field, ">=", value}
}

func Lt[T any](field string, value T) Specification {

	return binarySpecification[T]{field, "<", value}
}

func Lte[T any](field string, value T) Specification {
	return binarySpecification[T]{field, "<=", value}
}

func And(specs ...Specification) Specification {
	return joinSpecification{specs, "AND"}
}

func Or(specs ...Specification) Specification {
	return joinSpecification{specs, "OR"}
}

func Not(spec Specification) Specification {
	return notSpecification{spec}
}

// Group groups the specifications
func Group(spec Specification) Specification {
	return groupSpecification{[]Specification{spec}}
}
