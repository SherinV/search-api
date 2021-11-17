package schema

import (
	"context"
	"strconv"
	"strings"

	klog "k8s.io/klog/v2"

	db "github.com/SherinV/search-api/database"
	"github.com/SherinV/search-api/graph/model"
)

var trimAND string = " AND "

func Search(ctx context.Context, input []*model.SearchInput) ([]*model.SearchResult, error) {
	limit := 0
	srchResult := make([]*model.SearchResult, 0)

	if len(input) > 0 {
		for _, in := range input {
			query := searchQuery(ctx, "", in, &limit)
			klog.Infof("Search Query:", query)
			//TODO: Check error
			srchRes, _ := searchResults(query)
			srchResult = append(srchResult, srchRes)
		}
	}
	return srchResult, nil
}

func searchQuery(ctx context.Context, property string, input *model.SearchInput, limit *int) string {
	var selectClause, whereClause, limitClause, limitStr, query string

	selectClause = "SELECT uid, cluster, data FROM resources "
	limitClause = " LIMIT "

	whereClause = " WHERE "

	for i, filter := range input.Filters {
		klog.Infof("Filters%d: %+v", i, *filter)
		// TODO: To be removed when indexer handles this as adding lower hurts index scans

		whereClause = whereClause + "lower(data->> '" + filter.Property + "')"
		var values string
		for _, val := range filter.Values {
			klog.Infof("Filter value: %s", *val)
			values = values + "lower('" + *val + "'), "
			//TODO: Change logic if array of values
			//TODO: Remove lower() conversion once data is correctly loaded from indexer
			//SELECT count(uid) FROM resources  WHERE lower(data->> 'kind') IN (lower('pod')) ;
		}
		whereClause = whereClause + " IN (" + strings.TrimRight(values, ", ") + ")" + " AND "

	}
	if input.Limit != nil {
		limitStr = strconv.Itoa(*input.Limit)
	}
	if limitStr != "" {
		limitClause = " LIMIT " + limitStr
		query = selectClause + strings.TrimRight(whereClause, trimAND) + limitClause

	} else {
		query = selectClause + strings.TrimRight(whereClause, trimAND)
	}
	return query
}

func searchResults(query string) (*model.SearchResult, error) {

	pool := db.GetConnection()
	rows, _ := pool.Query(context.Background(), query)
	//TODO: Handle error
	defer rows.Close()
	var uid, cluster string
	var data map[string]interface{}
	items := []map[string]interface{}{}

	for rows.Next() {
		// rowValues, _ := rows.Values()
		err := rows.Scan(&uid, &cluster, &data)
		if err != nil {
			klog.Errorf("Error %s retrieving rows for query:%s", err.Error(), query)
		}

		// TODO: To be removed when indexer handles this
		currItem := make(map[string]interface{})
		for k, myInterface := range data {
			switch v := myInterface.(type) {
			case string:
				currItem[k] = strings.ToLower(v)
			default:
				// klog.Info("Not string type.", k, v)
				continue
			}

		}
		currUid := uid
		currItem["_uid"] = currUid
		currCluster := cluster
		currItem["cluster"] = currCluster
		items = append(items, currItem)
	}
	klog.Info("len items: ", len(items))
	totalCount := len(items)
	srchrelatedresult := make([]*model.SearchRelatedResult, 0)
	nodecount := 2
	clustercount := 1

	srchrelatedresult1 := model.SearchRelatedResult{Kind: "Node", Count: &nodecount}
	srchrelatedresult2 := model.SearchRelatedResult{Kind: "Cluster", Count: &clustercount}
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult1)
	srchrelatedresult = append(srchrelatedresult, &srchrelatedresult2)

	srchresult1 := model.SearchResult{
		Count:   &totalCount,
		Items:   items,
		Related: srchrelatedresult,
	}
	// srchResult := make([]*model.SearchResult, 0)
	// srchResult = append(srchResult, &srchresult1)
	return &srchresult1, nil
}
