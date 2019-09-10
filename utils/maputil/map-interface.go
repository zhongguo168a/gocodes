package maputil

import "sort"

func Keys(t map[string]interface{}) (r []string) {

	for key := range t {
		r = append(r, key)
	}

	sort.Strings(r)
	return
}

func Set(t map[string]interface{}, key string, val interface{}) {
	old := t[key]
	if old == nil {
		switch data := val.(type) {
		case map[string]interface{}:
			m := map[string]interface{}{}
			t[key] = m
			for dk, dv := range data {
				m[dk] = dv
			}
		default:
			t[key] = val
		}
	} else {
		switch data := old.(type) {
		case map[string]interface{}:
			if val == nil {
				t[key] = val
			} else {
				m := val.(map[string]interface{})
				for dk, dv := range m {
					data[dk] = dv
				}
			}
		default:
			t[key] = val
		}
	}
}

func Has(t map[string]interface{}, key string) bool {
	_, ok := t[key]
	return ok
}

func Interface(t map[string]interface{}, key string) (r interface{}) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		return
	}

	r = ir
	return
}

func String(t map[string]interface{}, key string) (r string) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		return
	}

	r = ir.(string)
	return
}

func Int(t map[string]interface{}, key string) (r int) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		return
	}

	r = ir.(int)
	return
}

func Float64(t map[string]interface{}, key string) (r float64) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		return
	}

	r = ir.(float64)
	return
}

func Bool(t map[string]interface{}, key string) (r bool) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		return
	}

	r = ir.(bool)
	return
}

func Map(t map[string]interface{}, key string) (r map[string]interface{}) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		r = map[string]interface{}{}
		t[key] = r
		return
	}

	if ir == nil {
		return
	}

	r = ir.(map[string]interface{})

	return
}

func Array(t map[string]interface{}, key string) (r []interface{}) {
	if t == nil {
		return
	}

	ir, ok := t[key]
	if ok == false {
		r = []interface{}{}
		t[key] = r
		return
	}

	if ir == nil {
		return
	}

	r = ir.([]interface{})

	return
}
