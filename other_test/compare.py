import pandas as pd
import argparse

def compare_csv(file1, file2, key_columns, compare_columns):
    df1 = pd.read_csv(file1)
    df2 = pd.read_csv(file2)
    
    # 仅保留关键列和对比列
    df1 = df1[key_columns + compare_columns]
    df2 = df2[key_columns + compare_columns]
    
    # 设置关键列作为索引
    # df1.set_index(key_columns, inplace=True)
    # df2.set_index(key_columns, inplace=True)
    # # 确保索引和列完全匹配
    # df1, df2 = df1.align(df2, join='outer', axis=0, fill_value="N/A")

    diff_df = df1.merge(df2, on=[key_columns], how='outer', indicator=True)
    # 进行比较
    diff = df1.compare(df2, align_axis=1)
    pd.concat([df1, df2]).drop_duplicates(subset=key_columns, keep=False)
    if diff.empty:
        print("No differences found.")
    else:
        print("Differences found:")
        print(diff)
        diff.to_csv("differences.csv")
        print("Differences saved to 'differences.csv'")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Compare two CSV files ba"
                                                 "sed on specified columns.")
    parser.add_argument("file1", help="Path to the first CSV file")
    parser.add_argument("file2", help="Path to the second CSV file")
    parser.add_argument("--keys", nargs='+', required=True, help="Key columns for identifying rows")
    parser.add_argument("--columns", nargs='+', required=True, help="Columns to compare")
    
    args = parser.parse_args()
    compare_csv(args.file1, args.file2, args.keys, args.columns)

