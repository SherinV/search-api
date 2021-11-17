package schema

import (
	"context"
	"strconv"

	db "github.com/SherinV/search-api/database"
	"github.com/SherinV/search-api/graph/model"
	klog "k8s.io/klog/v2"
)

func SearchComplete(ctx context.Context, property string, srchInput *model.SearchInput, limit *int) ([]*string, error) {
	query := searchCompleteQuery(ctx, property, srchInput, limit)
	klog.Infof("SearchComplete Query: ", query)
	return searchCompleteResults(query)
}

func searchCompleteQuery(ctx context.Context, property string, input *model.SearchInput, limit *int) string {
	var selectClause, limitClause, limitStr, query string
	if property != "" {
		klog.Infof("property: %s and limit:%d", property, limit)
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

	return ""
}

func searchCompleteResults(query string) ([]*string, error) {

	pool := db.GetConnection()
	//TODO: Handle error
	rows, _ := pool.Query(context.Background(), query)
	defer rows.Close()
	var srchCompleteOut []*string
	prop := ""
	for rows.Next() {
		_ = rows.Scan(&prop)
		tmpProp := prop
		srchCompleteOut = append(srchCompleteOut, &tmpProp)
		klog.Info("Property: ", prop, tmpProp)
	}
	return srchCompleteOut, nil
}
