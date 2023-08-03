package types

type SelectOperation struct {
	Field    string
	Alias    string
	Function string
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
	Function string
}

type Query struct {
	Selects     []SelectOperation
	Searches    [][]SearchOperation
	Joins       []JoinOperation
	Orders      []OrderOperation
	Groups      []GroupOperation
	Distinct    bool
	WithDeleted bool
	Limit       int
	Offset      int
}
