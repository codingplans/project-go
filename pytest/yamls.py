import yaml

def flatten_dict(d, parent_key='', sep='.'):
    items = []
    for k, v in d.items():
        new_key = f"{parent_key}{sep}{k}" if parent_key else k
        if isinstance(v, dict):
            items.extend(flatten_dict(v, new_key, sep=sep).items())
        elif isinstance(v, list):
            for i, item in enumerate(v):
                if isinstance(item, dict):
                    items.extend(flatten_dict(item, f"{new_key}.{i}", sep=sep).items())
                else:
                    items.append((f"{new_key}.{i}", item))
        else:
            items.append((new_key, v))
    return dict(items)

def unflatten_dict(d, sep='.'):
    result_dict = {}
    for k, v in d.items():
        keys = k.split(sep)
        d = result_dict
        for key in keys[:-1]:
            d = d.setdefault(key, {})
        d[keys[-1]] = v
    return result_dict

# 示例 YAML 数据
yaml_data = '''
data_source:
  orders:
    from:
      type: true
      db_name: alpha_running_logs
      table: 
        '0': 123
        '1': 1233
      fields: [ order_id, side, price, volume, log_update_time ]
      position:
        "300450": [100000000, 23.75]
        000008: 
          - - 100000000
            - 23.75
            - aa: 
                www
'''

# data_source.orders.from.position.300450.#0


# 解析 YAML 数据
data = yaml.safe_load(yaml_data)

# 展平处理
flat_data = flatten_dict(data)

# 输出展平后的键值对
for k, v in flat_data.items():
    print(f"{k}: {v}")

# 将展平后的键值对转换回嵌套的 YAML 格式
nested_data = unflatten_dict(flat_data)
yaml_output = yaml.dump(nested_data, default_flow_style=False, sort_keys=False)
print("\nConverted back to YAML format:\n")
print(yaml_output)
