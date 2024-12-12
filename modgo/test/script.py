# script.py
import json
import sys
from typing import Any, Dict, List, Tuple


def dict_parser(str_: str, start_idx: int = 0, depth: int = 0, max_depth: int = 10) -> Tuple[Dict[str, Any], int, int]:
    """
    Parse the dictionary in the string.

    Parameters:
    str_ (str): The input string.
    start_idx (int): The starting index for processing.
    depth (int): The current depth.
    max_depth (int): The maximum depth.

    Returns:
    Tuple[Dict[str, Any], int, int]: The processed dictionary and the ending index.
    """
    if str_ == "{}" or ":" not in str_:
        return {}, start_idx + 1, depth

    need_key = True
    data_now = {}
    key = None
    i = start_idx
    while i < len(str_):
        c = str_[i]
        if c == ":" and need_key:
            end_idx = i
            need_key = False
            key = str_[start_idx:end_idx].strip()
            data_now[key] = None
            start_idx = i + 1

        elif c == "{":
            depth += 1
            if depth <= max_depth:
                inner_data, end_idx, depth = dict_parser(str_, i + 1, depth, max_depth)
                if key is not None:
                    data_now[key] = inner_data
                else:
                    data_now = inner_data
                i = end_idx + 1
                start_idx = i
                need_key = True
            else:
                return data_now, len(str_), depth

        elif c == "[" and not need_key:
            depth += 1
            if depth <= max_depth:
                inner_data, end_idx, depth = list_parser(str_, i + 1, depth, max_depth)
                if key is not None:
                    data_now[key] = inner_data
                else:
                    data_now = inner_data
                i = end_idx + 1
                start_idx = i
                need_key = True
            else:
                return data_now, len(str_), depth

        elif c == ",":
            end_idx = i
            if not need_key:
                value = str_[start_idx:end_idx].strip()
                data_now[key] = value
                need_key = True
                start_idx = i + 1

        elif c == "}":
            end_idx = i
            value = str_[start_idx:end_idx].strip()
            if key is not None:
                data_now[key] = value
            else:
                if not value:
                    data_now = {}
                else:
                    data_now = value
            depth -= 1

            return data_now, i + 1, depth

        i += 1

    return data_now, i, depth


def list_parser(str_: str, start_idx: int = 0, depth: int = 0, max_depth: int = 10) -> Tuple[List[Any], int, int]:
    """
    Parse the list in the string.

    Parameters:
    str_ (str): The input string.
    start_idx (int): The starting index for processing.
    depth (int): The current depth.
    max_depth (int): The maximum depth.

    Returns:
    Tuple[List[Any], int, int]: The processed list and the ending index.
    """
    if str_ == "[]":
        return [], start_idx + 1, depth
    data_now = []
    i = start_idx
    while i < len(str_):
        c = str_[i]
        if c == "{":
            depth += 1
            if depth <= max_depth:
                inner_data, end_idx, depth = dict_parser(str_, i + 1, depth, max_depth)
                data_now.append(inner_data)
                i = end_idx + 1
                start_idx = i
            else:
                return data_now, len(str_), depth
        elif c == "[":
            depth += 1
            if depth <= max_depth:
                inner_data, end_idx, depth = list_parser(str_, i + 1, depth, max_depth)
                data_now.extend(inner_data)
                i = end_idx + 1
                start_idx = i
            else:
                return data_now, len(str_), depth

        elif c == ",":
            end_idx = i
            value = str_[start_idx:end_idx].strip()
            data_now.append(value)
            start_idx = i + 1

        elif c == "]":
            end_idx = i
            value = str_[start_idx:end_idx].strip()
            if value:
                data_now.append(value)
            depth -= 1
            return data_now, i + 1, depth
        i += 1

    return data_now, i, depth

# _msg = "{env:{stock_id:000001,depth_time:101610000,ask1_price:101800,ask1_vol:199900,bid1_price:101700,bid1_vol:83500,volume_left:100000,time_left:4130000,model_input:[0.0],model_output:{prob_act:{1:1.0},exec_act:1},order:{order_id:7821670000100010001,price:101800,volume:100,side:B},supp_info:{{is_model:1}}}}}"
# other_dict=dict_parser(_msg, max_depth=10)[0]
# logger.info(other_dict)


def process_string(input_string):
    # 处理输入字符串的函数
    return input_string.upper()  # 示例：将字符串转换为大写

if __name__ == "__main__":
    input_string = sys.argv[1]  # 获取命令行参数
    result = dict_parser(input_string)
    print(json.dumps(result[0]))
