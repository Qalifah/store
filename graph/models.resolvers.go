package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
)

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
