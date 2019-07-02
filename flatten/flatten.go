package flatten

//Flatten nested array
func Flatten(list []interface{}) ([]int64, error) {
	var result []int64
	for _, data := range list {
		if num, ok := isANum(data); ok {
			result = append(result, num)
			continue
		}
		if err := extractElem(data, &result); err != nil {
			return result, err
		}
	}
	return result, nil
}

func extractElem(v interface{}, rs *[]int64) error {
	if data, ok := v.([]interface{}); ok {
		for _, el := range data {
			if num, ok := isANum(el); ok {
				*rs = append(*rs, num)
				continue
			}
			extractElem(el, rs)
		}
		return nil
	}
	if data, ok := v.([]int64); ok {
		for _, el := range data {
			*rs = append(*rs, el)
			continue
		}
	}
	return nil
}

func isANum(v interface{}) (int64, bool) {
	switch num := v.(type) {
	case int:
		return int64(num), true
	case int32:
		return int64(num), true
	case int16:
		return int64(num), true
	case int8:
		return int64(num), true
	case int64:
		return num, true
	default:
		return 0, false
	}
}
