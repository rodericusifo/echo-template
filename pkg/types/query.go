package types

type SelectOperation struct {
	Field    string
	Operator string
	Alias    string
}

type SearchOperation struct {
	Field    string
	Operator string
	Value    any
}

type JoinOperation struct {
	Relation string
	Selects  []SelectOperation
	Searches [][]SearchOperation
}

type OrderOperation struct {
	Field      string
	Descending bool
}

type GroupOperation struct {
	Field    string
	Operator string
}

type Query struct {
	WithDeleted bool
	Distinct    bool
	Selects     []SelectOperation
	Searches    [][]SearchOperation
	Joins       []JoinOperation
	Orders      []OrderOperation
	Groups      []GroupOperation
	Offset      int
	Limit       int
}
