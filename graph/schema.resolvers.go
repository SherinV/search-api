package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	klog "k8s.io/klog/v2"

	db "github.com/SherinV/search-api/database"
	"github.com/SherinV/search-api/graph/generated"
	"github.com/SherinV/search-api/graph/model"
)

func (r *mutationResolver) DeleteSearch(ctx context.Context, resource *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SaveSearch(ctx context.Context, resource *string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

//TODO: Have separate functions for each sub-api
func (r *queryResolver) QueryBuilder(ctx context.Context, property string, input []*model.SearchInput, limit *int) string {
	klog.Infoln("Received input: ", input, " and property: ", property, " with limit ", limit)
	var selectClause, whereClause, limitClause, limitStr, query string
	if property != "" {
		klog.Infof("property: %s and limit:%d", property, limit)
		// fmt.Sprintf(%d, ) Printf()
		if property == "cluster" {
			selectClause = "SELECT DISTINCT " + property + " FROM resources "
		} else {
			selectClause = "SELECT DISTINCT data->>'" + property + "' FROM resources "
		}
		if limit != nil {
			limitStr = strconv.Itoa(*limit)
		}

		if limitStr != "0" && limitStr != "" {
			limitClause = " LIMIT " + limitStr
			query = selectClause + limitClause

		} else {
			query = selectClause
		}
		klog.Info("SearchComplete Query: ", query)
		return query
	}
	if len(input) > 0 {
		selectClause = "SELECT uid, cluster, data FROM resources "
		limitClause = " LIMIT "

		for _, in := range input {
			whereClause = " WHERE "

			for i, filter := range in.Filters {
				klog.Infof("Filters%d: %+v", i, *filter)
				whereClause = whereClause + "lower(data->> '" + filter.Property + "')"
				var values string
				for _, val := range filter.Values {
					klog.Infof("Filter value: %s", *val)
					values = values + "lower('" + *val + "'), "
					//TODO: Change logic if array of values
					//SELECT count(uid) FROM resources  WHERE lower(data->> 'kind') IN (lower('pod')) ;
				}
				whereClause = whereClause + " IN (" + strings.TrimRight(values, ", ") + ")" + " AND "

			}
			if in.Limit != nil {
				limitStr = strconv.Itoa(*in.Limit)
			}
			if limitStr != "" {
				limitClause = " LIMIT " + limitStr
				query = selectClause + strings.TrimRight(whereClause, " AND ") + limitClause

			} else {
				query = selectClause + strings.TrimRight(whereClause, " AND ")
			}

		}
		klog.Info("Query: ", query)
	}

	return query
}

func (r *queryResolver) Search(ctx context.Context, input []*model.SearchInput) ([]*model.SearchResult, error) {
	var count int
	klog.Infof("Received Search query with input %+v\n", input)
	if len(input) > 0 {
		for idx, in := range input {
			klog.Infof("Received Search query with input %+v\n", *in)
			if in.Limit != nil {
				klog.Infof("Input %d: with limit: %d", idx, *in.Limit)
			}
			for i, comps := range in.Keywords {
				klog.Infof("Keywords%d: %s", i, *comps)
			}
			for i, filter := range in.Filters {
				klog.Infof("Filters%d: %+v", i, *filter)
				for _, val := range filter.Values {
					klog.Infof("Filter value: %s", *val)
				}
			}
			for i, kind := range in.RelatedKinds {
				klog.Infof("RelatedKinds%d: %s", i, *kind)
			}
		}
	}
	limit := 0
	query := r.QueryBuilder(ctx, "", input, &limit)
	klog.Infof("Search Query")
	pool := db.GetConnection()
	rows, _ := pool.Query(context.Background(), query)
	//TODO: Handle error
	defer rows.Close()
	var uid, cluster string
	var data map[string]interface{}
	items := []map[string]interface{}{}

	for rows.Next() {
		_ = rows.Scan(&uid, &cluster, &data)

		currItem := make(map[string]interface{})
		for k, myInterface := range data {
			// This will be removed when all this is handled in indexer
			switch v := myInterface.(type) {
			case string:
				// v is a string here
				fmt.Printf("String: %v", v)
				currItem[k] = strings.ToLower(v)
			default:
				klog.Info("Not string type.", k, v)
			}

		}
		currUid := uid
		currItem["_uid"] = currUid
		currCluster := cluster
		currItem["cluster"] = currCluster
		items = append(items, currItem)
		klog.Info("len items: ", len(items))
		klog.Info("currItem: ", currItem)

	}

	// val1 := make(map[string]interface{})
	// val1["name"] = "search-ui"
	// val1["namespace"] = "test"
	// val1["label"] = "app=search; component:search-ui"
	// val1["_uid"] = "uid_pod1"
	// val1["cluster"] = "local-cluster"
	// val1["kind"] = "pod"
	// val1["status"] = "Running"
	// val1["restarts"] = "0"
	// val1["hostIP"] = "10.0.143.101"
	// val1["podIP"] = "10.128.0.201"
	// val1["created"] = "2021-10-27T04:55:34Z"

	// val2 := make(map[string]interface{})
	// val2["name"] = "search-api"
	// val2["label"] = "app=search; component:search-api"
	// val2["_uid"] = "uid_pod2"
	// val2["namespace"] = "test"
	// val2["cluster"] = "local-cluster"
	// val2["kind"] = "pod"
	// val2["status"] = "ContainerCreating"
	// val2["restarts"] = "1"
	// val2["hostIP"] = "10.0.143.102"
	// val2["podIP"] = "10.128.0.202"
	// val2["created"] = "2021-10-26T04:55:34Z"

	// items[0] = val1
	// items[1] = val2

	srchrelatedresult := make([]*model.SearchRelatedResult, 0)
	count = 2
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
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	klog.Infoln("Received Messages query")

	messages := make([]*model.Message, 0)
	kind := "Informational"
	desc := "Trial search-api"
	message1 := model.Message{ID: "1", Kind: &kind, Description: &desc}
	messages = append(messages, &message1)
	return messages, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchSchema(ctx context.Context) (map[string]interface{}, error) {
	klog.Infoln("Received SearchSchema query")

	srchSchema := make(map[string]interface{})
	schema := [5]string{"kind", "name", "namespace", "cpu", "created"}
	srchSchema["allProperties"] = schema
	return srchSchema, nil
}

func (r *queryResolver) SavedSearches(ctx context.Context) ([]*model.UserSearch, error) {
	klog.Infoln("Received SavedSearches query")

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

func (r *queryResolver) SearchComplete(ctx context.Context, property string, srchInput *model.SearchInput, limit *int) ([]*string, error) {
	klog.Infof("Received SearchComplete query with input property **%s** and limit %d", property, limit)
	input := make([]*model.SearchInput, 0)
	input = append(input, srchInput)
	query := r.QueryBuilder(ctx, property, input, limit)
	klog.Infof("SearchComplete Query", query)
	in := srchInput
	klog.Infof("SearchComplete Input with limit: %d", *in.Limit)
	for i, comps := range in.Keywords {
		klog.Infof("SearchComplete Keywords%d: %s", i, *comps)
	}
	for i, filter := range in.Filters {
		klog.Infof("SearchComplete Filters%d: %s", i, *filter)
		for _, val := range filter.Values {
			klog.Infof("Filter value: %s", *val)

		}
	}
	for i, kinds := range in.RelatedKinds {
		klog.Infof("SearchComplete RelatedKinds%d: %s", i, *kinds)
	}

	klog.Infof("SearchComplete Query")
	pool := db.GetConnection()
	rows, _ := pool.Query(context.Background(), query)
	//TODO: Handle error
	defer rows.Close()
	var srchCompleteOut []*string
	prop := ""
	for rows.Next() {
		// rowValues, _ := rows.Values()
		_ = rows.Scan(&prop)
		tmpProp := prop
		// fmt.Println(rowValues)
		srchCompleteOut = append(srchCompleteOut, &tmpProp)
		klog.Info("Property: ", prop, tmpProp)
	}
	// podKind := "pod"
	// return []*string{&podKind}, nil
	return srchCompleteOut, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {

	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
