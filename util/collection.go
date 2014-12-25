package util

func GetMapValues(m map[string]interface{}) []interface{} {
	v := make([]interface{}, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	return v
}
