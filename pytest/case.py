import datetime

import numpy as np
import pandas as pd

def tt(d, a, b, c):
    print(d, a, b, c)


params = {"a": 1}
params['c'] = 2
params['b'] = 3
params['d'] = 4

# 执行给定的代码
v = dict.fromkeys(['k1','k2'],[])
v['k1'].append(666)
print(v)
v['k1'] = 777
print(v)
ss = '2024-10-30 djoiqwjdiq'
print(ss.replace('2024-10-30 ',''))
print(222)
print(f"{datetime.datetime.now().strftime('%Y-%m-%d ')}xxxxxxx")
# print(datetime.datetime.now(阿萨德asd).strftime(''))

# 分块读取CSV
# chunk_list = []
# for chunk in pd.read_csv('large_file.csv', chunksize=10000):
#     # 对每个chunk进行处理
#     processed_chunk = chunk.groupby('column').sum()
#     chunk_list.append(processed_chunk)
#
# result = pd.concat(chunk_list, ignore_index=True)


# 创建时间序列数据
dates = pd.date_range('2023-01-01', periods=100, freq='D')
ts = pd.Series(np.random.randn(100), index=dates)

print(np.random.randint(0, 10, (2, 3)))  # 生成 2x3 介于 [0,10) 之间的整数
print(np.random.randn(9))  # 生成 3 个服从标准正态分布的数

# 创建时间范围
time_range = pd.date_range(start='2023-01-01', end='2023-12-31', freq='D')
business_days = pd.date_range(start='2023-01-01', end='2023-12-31', freq='B')



def test_account_asset_v2():
    holdings = pd.DataFrame(
        {
            "instrument": ["300519", "601288", "cash"],
            "volume": [100, 200, 50000],
            "holding_cost": [1000, 2000, 50000],
            "log_update_time": [
                pd.to_datetime("2024-05-28T10:00:00.000000"),
                pd.to_datetime("2024-05-28T10:00:00.100000"),
                pd.to_datetime("2024-05-28T10:00:00.200000"),
            ],
        }
    )

    depth = pd.DataFrame(
        {
            "instrument": ["300519", "601288"],
            "lastprice": [10.5 * 1e4, 12.0 * 1e4],
            "depthmarkettime": [1000000, 1000000],
        }
    )
    print(depth.iloc[:, 0])
    print(depth.iloc[:, 2])
    df = pd.DataFrame({'A': [1, np.nan, 3], 'B': [4, 5, np.nan]})
    # 删除缺失值
    df.dropna()
    # 删除含 NaN 的行
    df.dropna(axis=1)  # 删除含 NaN 的列
    # 填充缺失值
    df.fillna(0)  # 用 0 填充
    df.fillna(df.mean())  # 用列的均值填充
    print(df)


    df = pd.DataFrame({'A': [1, 2, 2, 3], 'B': [4, 5, 5, 6]})
    df.drop_duplicates()  # 删除重复行
    df.drop_duplicates(subset=['A'], keep='first')  # 只考虑 'A' 列，保留第一条
    print(df)


    df = pd.DataFrame({'Category': ['A', 'B', 'A', 'B'], 'Value': [10, 20, 30, 40]})  # 按 'Category' 分组并求和
    print(df.groupby('Category').sum())  # 计算多个聚合指标
    print(df.groupby('Category').agg({'Value': ['sum', 'mean']}))
    # result = account_asset_v2(holdings, depth)
    # assert isinstance(result, np.ndarray)
    # assert len(result) == 5
    # assert np.allclose(result[0], 53450)
    # assert np.allclose(result[1], 50000)
    # assert np.allclose(result[2], 3450)
    # assert np.allclose(result[3], 2400)
    # assert np.allclose(result[4], 1050)
    #
    # empty_holdings = pd.DataFrame()
    # with pytest.raises(ValueError) as excinfo:
    #     account_asset_v2(empty_holdings, depth)
    # assert "some fields are not found in holdings columns" in str(excinfo.value)
    #
    # missing_depth = depth.drop(0)
    # result = account_asset_v2(holdings, missing_depth)
    # assert np.allclose(result[0], 52400)
    # assert np.allclose(result[1], 50000)
    # assert np.allclose(result[2], 2400)
    #
    # missing_holdings = holdings.drop(0)
    # result = account_asset_v2(missing_holdings, depth)
    # assert np.allclose(result[0], 52400)
    # assert np.allclose(result[1], 50000)
    # assert np.allclose(result[2], 2400)
