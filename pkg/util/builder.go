package util

import (
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func BuildQuery(tableName string, db *gorm.DB, query *types.Query) *gorm.DB {
	q := db
	
	if len(query.Selects) > 0 {
		querySlice := GenerateSQLSelectQuerySlice(tableName, MergeSlices(true, query.Selects, constant.DEFAULT_SELECT_COLUMNS))
		if query.Distinct {
			q = q.Distinct(querySlice)
		} else {
			q = q.Select(querySlice)
		}
	}
	if len(query.Searches) > 0 {
		queryString, bindValues := GenerateSQLWhereQueryStringAndBindValues(tableName, query.Searches)
		q = q.Where(queryString, bindValues...)
	}
	if len(query.Joins) > 0 {
		for _, join := range query.Joins {
			if len(join.Selects) > 0 || len(join.Searches) > 0 {
				qj := db
				if len(join.Selects) > 0 {
					querySlice := GenerateSQLSelectQuerySlice(join.Relation, MergeSlices(true, join.Selects, constant.DEFAULT_JOIN_SELECT_COLUMNS))
					if query.Distinct {
						qj = qj.Distinct(querySlice)
					} else {
						qj = qj.Select(querySlice)
					}
				}
				if len(query.Searches) > 0 {
					queryString, bindValues := GenerateSQLWhereQueryStringAndBindValues(join.Relation, join.Searches)
					qj = qj.Where(queryString, bindValues...)
				}
				q = q.Joins(join.Relation, qj)
			} else {
				q = q.Joins(join.Relation)
			}
		}
	}
	if len(query.Orders) > 0 {
		queryString := GenerateSQLOrderQueryString(tableName, query.Orders)
		q = q.Order(queryString)
	}
	if len(query.Groups) > 0 {
		queryString := GenerateSQLGroupQueryString(tableName, query.Groups)
		q = q.Group(queryString)
	}
	if query.Limit != 0 {
		q = q.Limit(query.Limit)
	}
	if query.Offset != 0 {
		q = q.Offset(query.Offset)
	}
	if query.WithDeleted {
		q = q.Unscoped()
	}

	return q
}