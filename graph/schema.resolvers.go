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
	items := make([]map[string]interface{}, 2)
	val1 := make(map[string]interface{})
	val1["name"] = "search-ui"
	val1["label"] = [2]string{"app=search", "component:search-ui"}
	val2 := make(map[string]interface{})
	val2["name"] = "search-api"
	val2["label"] = [2]string{"app=search", "component:search-api"}

	items = append(items, val1)
	items = append(items, val2)

	srchrelatedresult := make([]*model.SearchRelatedResult, 2)
	count := 2
	srchrelatedresult1 := model.SearchRelatedResult{Kind: "Pod", Count: &count}
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult1)

	srchresult1 := model.SearchResult{
		Count:   &count,
		Items:   items,
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
