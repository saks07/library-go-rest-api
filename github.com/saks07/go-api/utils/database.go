package utils

import "strings"

func QueryStringTable(queryString string, dbTable string) string {
	return strings.Replace(queryString, "{table}", dbTable, 1);
}