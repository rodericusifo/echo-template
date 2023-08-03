package util

import (
	"fmt"
	"strings"

	"github.com/rodericusifo/echo-template/pkg/types"
)

func GenerateSQLSelectQuerySlice(tableAlias string, selects []types.SelectOperation) []string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	for _, s := range selects {
		fieldSelectStr := ""
		if s.Function != "" {
			function := s.Function
			fieldStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)
			fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
		} else {
			fieldSelectStr = fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)
		}

		if s.Alias != "" {
			querySlice = append(querySlice, fmt.Sprintf(`%s AS "%s"`, fieldSelectStr, s.Alias))
		} else {
			querySlice = append(querySlice, fieldSelectStr)
		}
	}

	return querySlice
}

func GenerateSQLWhereQueryStringAndBindValues(tableAlias string, searches [][]types.SearchOperation) (string, []any) {
	queryString := ""
	bindValues := make([]any, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	for indexOuter, searchOuter := range searches {
		if indexOuter > 0 {
			queryString += " OR "
		}
		for indexInner, searchInner := range searchOuter {
			if indexInner > 0 {
				queryString += " AND "
			}
			if searchInner.Value != nil {
				queryString += fmt.Sprintf(`"%s"."%s" %s ?`, tableAlias, searchInner.Field, searchInner.Operator)
				bindValues = append(bindValues, searchInner.Value)
			} else {
				queryString += fmt.Sprintf(`"%s"."%s" %s`, tableAlias, searchInner.Field, searchInner.Operator)
			}
		}
	}

	return queryString, bindValues
}

func GenerateSQLOrderQueryString(tableAlias string, orders []types.OrderOperation) string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	for _, order := range orders {
		if order.Descending {
			querySlice = append(querySlice, fmt.Sprintf(`"%s"."%s" DESC`, tableAlias, order.Field))
		} else {
			querySlice = append(querySlice, fmt.Sprintf(`"%s"."%s"`, tableAlias, order.Field))
		}
	}

	return strings.Join(querySlice, ",")
}

func GenerateSQLGroupQueryString(tableAlias string, groups []types.GroupOperation) string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	for _, group := range groups {
		fieldSelectStr := ""
		if group.Function != "" {
			function := group.Function
			fieldStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, group.Field)
			fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
		} else {
			fieldSelectStr = fmt.Sprintf(`"%s"."%s"`, tableAlias, group.Field)
		}
		querySlice = append(querySlice, fieldSelectStr)
	}

	return strings.Join(querySlice, ",")
}
