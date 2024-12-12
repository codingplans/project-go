package main

import (
	"fmt"
	"os/exec"
)

type Parser struct {
	str      string
	idx      int
	depth    int
	maxDepth int
}

func execPy(input string) {
	cmd := exec.Command("/opt/homebrew/anaconda3/envs/py311/bin/python", "test/script.py", input)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		return
	}
	fmt.Println(string(output))
}

//
// func (p *Parser) parseDict() (map[string]interface{}, error) {
// 	if p.str == "{}" || !strings.Contains(p.str, ":") {
// 		return map[string]interface{}{}, nil
// 	}
//
// 	dataNow := make(map[string]interface{})
// 	needKey := true
// 	var key string
//
// 	for p.idx < len(p.str) {
// 		c := p.str[p.idx]
// 		switch c {
// 		case ':':
// 			if needKey {
// 				key = strings.TrimSpace(p.str[:p.idx])
// 				dataNow[key] = nil
// 				p.idx++
// 				needKey = false
// 			}
// 		case '{':
// 			p.depth++
// 			if p.depth <= p.maxDepth {
// 				p.idx++
// 				innerData, err := p.parseDict()
// 				if err != nil {
// 					return nil, err
// 				}
// 				dataNow[key] = innerData
// 				needKey = true
// 			} else {
// 				return dataNow, nil
// 			}
// 		case '[':
// 			if !needKey {
// 				p.depth++
// 				if p.depth <= p.maxDepth {
// 					p.idx++
// 					innerData, err := p.parseList()
// 					if err != nil {
// 						return nil, err
// 					}
// 					dataNow[key] = innerData
// 					needKey = true
// 				} else {
// 					return dataNow, nil
// 				}
// 			}
// 		case ',':
// 			if !needKey {
// 				value := strings.TrimSpace(p.str[:p.idx])
// 				dataNow[key] = value
// 				needKey = true
// 				p.idx++
// 			}
// 		case '}':
// 			value := strings.TrimSpace(p.str[:p.idx])
// 			if key != "" {
// 				dataNow[key] = value
// 			} else {
// 				if value == "" {
// 					dataNow = map[string]interface{}{}
// 				} else {
// 					dataNow = map[string]interface{}{
// 						"value": value,
// 					}
// 				}
// 			}
// 			p.depth--
// 			return dataNow, nil
// 		default:
// 			p.idx++
// 		}
// 	}
//
// 	return dataNow, nil
// }
//
// func (p *Parser) parseList() ([]interface{}, error) {
// 	if p.str == "[]" {
// 		return []interface{}{}, nil
// 	}
//
// 	dataNow := []interface{}{}
// 	for p.idx < len(p.str) {
// 		c := p.str[p.idx]
// 		switch c {
// 		case '{':
// 			p.depth++
// 			if p.depth <= p.maxDepth {
// 				p.idx++
// 				innerData, err := p.parseDict()
// 				if err != nil {
// 					return nil, err
// 				}
// 				dataNow = append(dataNow, innerData)
// 			} else {
// 				return dataNow, nil
// 			}
// 		case '[':
// 			p.depth++
// 			if p.depth <= p.maxDepth {
// 				p.idx++
// 				innerData, err := p.parseList()
// 				if err != nil {
// 					return nil, err
// 				}
// 				dataNow = append(dataNow, innerData)
// 			} else {
// 				return dataNow, nil
// 			}
// 		case ',':
// 			value := strings.TrimSpace(p.str[:p.idx])
// 			dataNow = append(dataNow, value)
// 			p.idx++
// 		case ']':
// 			value := strings.TrimSpace(p.str[:p.idx])
// 			if value != "" {
// 				dataNow = append(dataNow, value)
// 			}
// 			p.depth--
// 			return dataNow, nil
// 		default:
// 			p.idx++
// 		}
// 	}
//
// 	return dataNow, nil
// }

func _main() {
	str := `{env:{stock_id:000001,depth_time:101610000,ask1_price:101800,ask1_vol:199900,bid1_price:101700,bid1_vol:83500,volume_left:100000,time_left:4130000,model_input:[0.0],model_output:{prob_act:{1:1.0},exec_act:1},order:{order_id:7821670000100010001,price:101800,volume:100,side:B},supp_info:{is_model:1}}`
	execPy(str)

	str = `{env:{stock_id:002594,depth_time:93027000,ask1_price:2535700,ask1_vol:200,bid1_price:2532300,bid1_vol:200,volume_left:9400,time_left:7173000},model_input:[2535700,200,2532300,200,2534089.5349,2534210.9415,2534594.3359],model_output:{prob_act:{0:1.0000,1:0.0000},exec_act:0},order:{order_id:4293170001000010001,price:2532300,volume:100,side:S},supp_info:{is_model:1}}`
	execPy(str)

	str = `{env:{stock_id:300519,depth_time:101610000,ask1_price:103500,ask1_vol:200,bid1_price:103400,bid1_vol:100,volume_left:50000,time_left:4130000,model_input:[0.0],model_output:{prob_act:{1:1.0},exec_act:1},order:{order_id:7821670000100020001,price:103500,volume:100,side:B},supp_info:{is_model:1}}`
	execPy(str)

}
