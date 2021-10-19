package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/SherinV/search-api/graph/generated"
	"github.com/SherinV/search-api/graph/model"
)

func (r *mutationResolver) DeleteSearch(ctx context.Context, resource *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SaveSearch(ctx context.Context, resource *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, input []*model.SearchInput) ([]*model.SearchResult, error) {
	// items := make([]string, 2)
	// items = append(items, "search-ui")
	// items = append(items, "search-api")

	srchrelatedresult := make([]*model.SearchRelatedResult, 2)
	count := 2
	srchrelatedresult1 := model.SearchRelatedResult{Kind: "Pod", Count: &count}
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult1)

	srchresult1 := model.SearchResult{
		Count:   &count,
		Related: srchrelatedresult,
	}
	srchResult := make([]*model.SearchResult, 1)
	srchResult = append(srchResult, &srchresult1)
	return srchResult, nil
	// panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
