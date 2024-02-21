package graph

import "github.com/rezendelc/full-cycle-examples/graph-ql/internal/database/category.go"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CategoryDB *database.Category
}
