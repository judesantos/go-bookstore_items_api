package queries

type EsQuery struct {
	Equals []FieldValue
}

type FieldValue struct {
	Field string
	Value interface{}
}
