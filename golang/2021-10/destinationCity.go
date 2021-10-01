package _021_10

// Destination City
// leetcode: https://leetcode-cn.com/problems/destination-city/
func DestinationCity(paths [][]string) string {
	m := make(map[string]string)
	for _, path := range paths {
		m[path[0]] = path[1]
	}
	for _, path := range paths {
		if _, has := m[path[1]]; !has {
			return path[1]
		}
	}
	return ""
}