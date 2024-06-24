package main

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/yaml.v3"
)

func flatten(prefix string, data interface{}, result map[string]string) {
	switch value := data.(type) {
	case map[interface{}]interface{}:
		for k, v := range value {
			flatten(prefix+fmt.Sprintf("%v.", k), v, result)
		}
	case map[string]interface{}:
		for k, v := range value {
			flatten(prefix+fmt.Sprintf("%v.", k), v, result)
		}
	case []interface{}:
		for i, v := range value {
			flatten(fmt.Sprintf("%s%d.", prefix, i), v, result)
		}
	default:
		result[prefix[:len(prefix)-1]] = fmt.Sprintf("%v", value)
	}
}

func masin() {
	yamlData := `
data_source:
  orders:
    from:
      type: postgres
      db_name: alpha_running_logs
      table: orders
      fields: [ order_id, side, price, volume, log_update_time ]
      position:
        300450: [100000000, 23.75]
        601288:
          - - 100000000
            - 23.75
            - aa: www
`

	var data map[string]interface{}

	err := yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	flatMap := make(map[string]string)
	flatten("", data, flatMap)

	for k, v := range flatMap {
		fmt.Printf("%s: %s\n", k, v)
	}

	// 将展平后的键值对转换回YAML格式
	flatYaml, err := yaml.Marshal(flatMap)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nFlattened YAML:\n%s\n", string(flatYaml))
	// flatMap := map[string]string{
	// 	"data_source.orders.from.type":                   "postgres",
	// 	"data_source.orders.from.db_name":                "alpha_running_logs",
	// 	"data_source.orders.from.table":                  "orders",
	// 	"data_source.orders.from.fields.0":               "order_id",
	// 	"data_source.orders.from.fields.1":               "side",
	// 	"data_source.orders.from.fields.2":               "price",
	// 	"data_source.orders.from.fields.3":               "volume",
	// 	"data_source.orders.from.fields.4":               "log_update_time",
	// 	"data_source.orders.from.position.300450.0":      "100000000",
	// 	"data_source.orders.from.position.300450.1":      "23.75",
	// 	"data_source.orders.from.position.601288.0.0":    "100000000",
	// 	"data_source.orders.from.position.601288.0.1":    "23.75",
	// 	"data_source.orders.from.position.601288.0.2.aa": "www",
	// }

}

// flattenMap 递归函数，用于将嵌套的map展平
func flattenMap(data map[string]interface{}, prefix string, flatMap map[string]interface{}) {
	for k, v := range data {
		fullKey := k
		if prefix != "" {
			fullKey = prefix + "." + k
		}

		switch v.(type) {
		case map[interface{}]interface{}:
			// 将map[interface{}]interface{}转换为map[string]interface{}
			nestedMap := make(map[string]interface{})
			for nk, nv := range v.(map[interface{}]interface{}) {
				nestedMap[nk.(string)] = nv
			}
			flattenMap(nestedMap, fullKey, flatMap)
		case map[string]interface{}:
			flattenMap(v.(map[string]interface{}), fullKey, flatMap)
		case []interface{}:
			flattenArray(v.([]interface{}), fullKey, flatMap)
		default:
			flatMap[fullKey] = v
		}
	}
}

// flattenArray 递归函数，用于将嵌套的数组展平
func flattenArray(data []interface{}, prefix string, flatMap map[string]interface{}) {
	for i, v := range data {
		fullKey := fmt.Sprintf("%s.%d", prefix, i)

		switch v.(type) {
		case map[interface{}]interface{}:
			// 将map[interface{}]interface{}转换为map[string]interface{}
			nestedMap := make(map[string]interface{})
			for nk, nv := range v.(map[interface{}]interface{}) {
				nestedMap[nk.(string)] = nv
			}
			flattenMap(nestedMap, fullKey, flatMap)
		case map[string]interface{}:
			flattenMap(v.(map[string]interface{}), fullKey, flatMap)
		case []interface{}:
			flattenArray(v.([]interface{}), fullKey, flatMap)
		default:
			flatMap[fullKey] = v
		}
	}
}
func setValue(data map[string]interface{}, path []string, value interface{}) {
	if len(path) == 1 {
		data[path[0]] = value
		return
	}

	key := path[0]
	if idx, err := strconv.Atoi(key); err == nil {
		if _, ok := data["array"]; !ok {
			data["array"] = make([]interface{}, idx+1)
		}

		arr := data["array"].([]interface{})
		if len(arr) <= idx {
			newArr := make([]interface{}, idx+1)
			copy(newArr, arr)
			arr = newArr
			data["array"] = arr
		}

		if arr[idx] == nil {
			arr[idx] = make(map[string]interface{})
		}
		setValue(arr[idx].(map[string]interface{}), path[1:], value)
	} else {
		if _, ok := data[key]; !ok {
			data[key] = make(map[string]interface{})
		}
		setValue(data[key].(map[string]interface{}), path[1:], value)
	}
}
