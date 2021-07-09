package _021_07

import (
	"sort"
	"strconv"
)

// Display Table of Food Orders in a Restaurant
// leetcode: https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant/

func DisplayTable(orders [][]string) [][]string {
	ans := make([][]string, 1)
	foodMap := make(map[string][]string)
	tableMap := make(map[int]map[string]int)
	for _, order := range orders {
		foodMap[order[2]] = append(foodMap[order[2]], order[1])
		table, _ := strconv.Atoi(order[1])
		if tableMap[table] == nil {
			tableMap[table] = make(map[string]int)
		}
		tableMap[table][order[2]] ++
	}
	foodArr := make([]string, 0)
	tableArr := make([]int, 0)
	for food, _ := range foodMap {
		foodArr = append(foodArr, food)
	}
	sort.Strings(foodArr)
	foodArr = append([]string{"Table"}, foodArr...)
	ans[0] = foodArr
	for table, _ := range tableMap {
		tableArr = append(tableArr, table)
	}
	sort.Ints(tableArr)
	for _, table := range tableArr {
		tableFoods := []string{}
		tableFoods = append(tableFoods, strconv.Itoa(table))
		for i := 1; i < len(foodArr); i++ {
			num := tableMap[table][foodArr[i]]
			numStr := strconv.Itoa(num)
			tableFoods = append(tableFoods, numStr)
		}
		ans = append(ans, tableFoods)
	}
	return ans
}
