package main

import (
	"fmt"
	"strings"
)

func buildQuery(filters []string, value string) string {
	baseQuery := "SELECT * FROM your_table"
	if len(filters) == 0 {
		return baseQuery
	}

	var conditions []string
	for _, iter := range filters {
		conditions = append(conditions, fmt.Sprintf("%s LIKE '%%%s%%'", iter, value))
	}

	whereClause := strings.Join(conditions, " OR ")
	return fmt.Sprintf("%s WHERE %s", baseQuery, whereClause)
}

func main() {

	fmt.Println(buildQuery([]string{"license_keys", "topcher"}, "val"))

}
