package gqlgen_repro

import (
	"context"
	"errors"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return []*Todo{{}}, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return nil, errors.New("oh no")
}
