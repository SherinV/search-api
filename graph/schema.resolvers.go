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
	fmt.Println("Received Search query with input", input)
	items := make([]map[string]interface{}, 2)
	val1 := make(map[string]interface{})
	val1["name"] = "search-ui"
	val1["label"] = "app=search; component:search-ui"
	val2 := make(map[string]interface{})
	val2["name"] = "search-api"
	val2["label"] = "app=search; component:search-api"

	items[0] = val1
	items[1] = val2

	srchrelatedresult := make([]*model.SearchRelatedResult, 0)
	count := 2
	srchrelatedresult1 := model.SearchRelatedResult{Kind: "Pod", Count: &count}
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult1)

	srchresult1 := model.SearchResult{
		Count:   &count,
		Items:   items,
		Related: srchrelatedresult,
	}
	srchResult := make([]*model.SearchResult, 0)
	srchResult = append(srchResult, &srchresult1)
	return srchResult, nil

	// r.searchresults = append(r.searchresults, &srchresult1)
	// return r.searchresults, nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	fmt.Println("Received Messages query")

	messages := make([]*model.Message, 0)
	kind := "Informational"
	desc := "Trial search-api"
	message1 := model.Message{ID: "1", Kind: &kind, Description: &desc}
	messages = append(messages, &message1)
	return messages, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchSchema(ctx context.Context) (map[string]interface{}, error) {
	fmt.Println("Received SearchSchema query")

	srchSchema := make(map[string]interface{})
	schema := [5]string{"kind", "name", "namespace", "cpu", "created"}
	srchSchema["allProperties"] = schema
	return srchSchema, nil
}

func (r *queryResolver) SavedSearches(ctx context.Context) ([]*model.UserSearch, error) {
	fmt.Println("Received SavedSearches query")

	savedSrches := make([]*model.UserSearch, 0)
	id := "1"
	name := "savedSrch1"
	srchText := "Trial savedSrch1"
	desc := "Trial search-api savedSrch1"
	savedSrch1 := model.UserSearch{ID: &id, Name: &name, Description: &desc, SearchText: &srchText}
	savedSrches = append(savedSrches, &savedSrch1)
	return savedSrches, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchComplete(ctx context.Context, property string, query *model.SearchInput, limit *int) ([]*string, error) {
	podKind := "pod"
	return []*string{&podKind}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
