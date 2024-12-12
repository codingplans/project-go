
import pandas as pd

def process_csv(input_file, output_file):
    # 读取 CSV 文件
    df = pd.read_csv(input_file, header=None)


    df[2] = (df[2].astype(float) * 10000).round().astype(int)

    # 保存修改后的数据到新的 CSV 文件
    df.to_csv(output_file, index=False, header=False)

    
# 示例调用
input_file = 'pre_close_20240930.csv'  # 替换为您的 CSV 文件路径
output_file = 'output.csv'  # 替换为您希望保存的 CSV 文件路径

process_csv(input_file, output_file)
