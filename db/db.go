package db

type FilterPredicateType int

type NumericOperation int
type StringOperation int
type BooleanOperation int
type ComplexOperation int

type Direction int
type Aggregation int
type DataType int

const (
	FilterPredicateString  FilterPredicateType = 0
	FilterPredicateNumeric FilterPredicateType = 1
	FilterPredicateBoolean FilterPredicateType = 2
	FilterPredicateComplex FilterPredicateType = 3

	NumericEqual          NumericOperation = 0
	NumericNotEqual       NumericOperation = 1
	NumericGreater        NumericOperation = 2
	NumericLess           NumericOperation = 3
	NumericGreaterOrEqual NumericOperation = 4
	NumericLessOrEqual    NumericOperation = 5

	StringEqual       StringOperation = 0
	StringNotEqual    StringOperation = 1
	StringStartsWith  StringOperation = 2
	StringEndsWith    StringOperation = 3
	StringContains    StringOperation = 4
	StringNotContains StringOperation = 5
	StringIn          StringOperation = 6
	StringNotIn       StringOperation = 7

	BooleanEqual    BooleanOperation = 0
	BooleanNotEqual BooleanOperation = 1

	ComplexAnd ComplexOperation = 0
	ComplexOr  ComplexOperation = 1

	ASC  Direction = 0
	DESC Direction = 1

	MIN   Aggregation = 0
	MAX   Aggregation = 1
	AVG   Aggregation = 2
	SUM   Aggregation = 3
	COUNT Aggregation = 4
	NONE  Aggregation = 5

	STRING  DataType = 0
	LONG    DataType = 0
	BOOLEAN DataType = 0
	DOUBLE  DataType = 0
	JSON    DataType = 0
)

type Client interface {
}
