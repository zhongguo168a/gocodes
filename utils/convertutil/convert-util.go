package convertutil

import "strconv"

func ConvertInt2String(ival interface{}) string {
	val := ""
	switch x := ival.(type) {
	case int:
		val = strconv.Itoa(x)
	case int8:
		val = strconv.Itoa(int(x))
	case int16:
		val = strconv.Itoa(int(x))
	case int32:
		val = strconv.Itoa(int(x))
	case int64:
		val = strconv.Itoa(int(x))
	}
	return val
}

func Convert2String(ival interface{}) string {
	val := ""
	switch x := ival.(type) {
	case string:
		val = x
	case int:
		val = strconv.Itoa(x)
	case int8:
		val = strconv.Itoa(int(x))
	case int16:
		val = strconv.Itoa(int(x))
	case int32:
		val = strconv.Itoa(int(x))
	case int64:
		val = strconv.Itoa(int(x))
	case uint:
		val = strconv.Itoa(int(x))
	case uint8:
		val = strconv.Itoa(int(x))
	case uint16:
		val = strconv.Itoa(int(x))
	case uint32:
		val = strconv.Itoa(int(x))
	case uint64:
		val = strconv.Itoa(int(x))
	case float32:
		val = strconv.Itoa(int(x))
	case float64:
		val = strconv.Itoa(int(x))
	}
	return val
}

func Convert2Int64(ival interface{}) int64 {
	val := int64(0)
	switch data := ival.(type) {
	case int:
		val = int64(data)
	case int8:
		val = int64(data)
	case int16:
		val = int64(data)
	case int32:
		val = int64(data)
	case int64:
		val = data
	case uint:
		val = int64(data)
	case uint8:
		val = int64(data)
	case uint16:
		val = int64(data)
	case uint32:
		val = int64(data)
	case uint64:
		val = int64(data)
	case float32:
		val = int64(data)
	case float64:
		val = int64(data)
	case string:
		i, err := strconv.Atoi(data)
		if err != nil {
			panic(err)
		}
		val = int64(i)
	default:
	}

	return val
}

func Convert2Float64(ival interface{}) float64 {
	val := float64(0)
	switch data := ival.(type) {
	case int:
		val = float64(data)
	case int8:
		val = float64(data)
	case int16:
		val = float64(data)
	case int32:
		val = float64(data)
	case int64:
		val = float64(data)
	case uint:
		val = float64(data)
	case uint8:
		val = float64(data)
	case uint16:
		val = float64(data)
	case uint32:
		val = float64(data)
	case uint64:
		val = float64(data)
	case float32:
		val = float64(data)
	case float64:
		val = data
	case string:
		i, err := strconv.ParseFloat(data, 64)
		if err != nil {
			panic(err)
		}
		val = float64(i)
	}

	return val
}
