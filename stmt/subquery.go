package stmt

import (
	"github.com/DzananGanic/loukoum/token"
	"github.com/DzananGanic/loukoum/types"
)

// Exists is a subquery expression.
type Exists struct {
	Subquery Expression
}

// NewExists returns a new Exists instance.
func NewExists(value interface{}) Exists {
	return Exists{
		Subquery: NewExpression(value),
	}
}

func (Exists) expression() {}

// Write exposes statement as a SQL query.
func (exists Exists) Write(ctx types.Context) {
	ctx.Write(token.Exists.String())
	ctx.Write(" (")
	exists.Subquery.Write(ctx)
	ctx.Write(")")
}

// IsEmpty returns true if statement is undefined.
func (exists Exists) IsEmpty() bool {
	return false
}

func (Exists) selectExpression() {}

// Ensure that Exists is an Expression
var _ Expression = Exists{}

// NotExists is a subquery expression.
type NotExists struct {
	Subquery Expression
}

// NewNotExists returns a new NotExists instance.
func NewNotExists(value interface{}) NotExists {
	return NotExists{
		Subquery: NewExpression(value),
	}
}

func (NotExists) expression() {}

// Write exposes statement as a SQL query.
func (nexists NotExists) Write(ctx types.Context) {
	ctx.Write("NOT " + token.Exists.String())
	ctx.Write(" (")
	nexists.Subquery.Write(ctx)
	ctx.Write(")")
}

// IsEmpty returns true if statement is undefined.
func (nexists NotExists) IsEmpty() bool {
	return false
}

func (NotExists) selectExpression() {}

// Ensure that NotExists is an Expression
var _ Expression = NotExists{}
