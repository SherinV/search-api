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
	val1["namespace"] = "test"
	val1["label"] = "app=search; component:search-ui"
	val1["_uid"] = "uid_pod1"
	val1["cluster"] = "local-cluster"
	val1["kind"] = "pod"
	val1["status"] = "Running"
	val1["restarts"] = "0"
	val1["hostIP"] = "10.0.143.101"
	val1["podIP"] = "10.128.0.201"
	val1["created"] = "2021-10-27T04:55:34Z"

	val2 := make(map[string]interface{})
	val2["name"] = "search-api"
	val2["label"] = "app=search; component:search-api"
	val2["_uid"] = "uid_pod2"
	val2["namespace"] = "test"
	val2["cluster"] = "local-cluster"
	val2["kind"] = "pod"
	val2["status"] = "ContainerCreating"
	val2["restarts"] = "1"
	val2["hostIP"] = "10.0.143.102"
	val2["podIP"] = "10.128.0.202"
	val2["created"] = "2021-10-26T04:55:34Z"

	items[0] = val1
	items[1] = val2

	srchrelatedresult := make([]*model.SearchRelatedResult, 0)
	count := 2
	clustercount := 1

	srchrelatedresult1 := model.SearchRelatedResult{Kind: "Node", Count: &count}
	srchrelatedresult2 := model.SearchRelatedResult{Kind: "Cluster", Count: &clustercount}
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult1)
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult2)

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
