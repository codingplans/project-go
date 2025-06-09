import datetime

import pandas as pd

# 读取车辆数据
df = pd.read_csv('usedcar_processed.csv')

# 清洗发布日和价格为空的数据
filter_data = df.dropna(subset=["posting_date", 'price'])
# 可以根据需要再次清理其他字段为空，或者特殊值的数据

# 去掉价格为零的数据，显然是不合理的数据
filter_data = filter_data.loc[filter_data['price'] > 0]

print("价格字段基本信息：")
print(filter_data['price'].describe())
print(f"缺失值数量：{filter_data['price'].isnull().sum()}")
print(f"零值数量：{(filter_data['price'] == 0).sum()}")
print(f"负值数量：{(filter_data['price'] < 0).sum()}")

# 使用IQR方法去除异常值
Q1 = filter_data['price'].quantile(0.25)
Q3 = filter_data['price'].quantile(0.75)
IQR = Q3 - Q1
lower_bound = Q1 - 1.5 * IQR
upper_bound = Q3 + 1.5 * IQR

price_col = 'price'
# 过滤异常值
cleaned_df = filter_data[
    (filter_data[price_col] >= lower_bound) &
    (filter_data[price_col] <= upper_bound)
    ]

from sklearn.preprocessing import MinMaxScaler, StandardScaler, RobustScaler

# 3. 数据标准化
# 方法1：MinMax标准化
scaler_minmax = MinMaxScaler()
df_minmax = cleaned_df.copy()
df_minmax['price_normalized'] = scaler_minmax.fit_transform(cleaned_df[['price']])

# 方法2：标准化
scaler_standard = StandardScaler()
df_standard = cleaned_df.copy()
df_standard['price_normalized'] = scaler_standard.fit_transform(cleaned_df[['price']])

# 方法3：鲁棒标准化
scaler_robust = RobustScaler()
df_robust = cleaned_df.copy()
df_robust['price_normalized'] = scaler_robust.fit_transform(cleaned_df[['price']])


def calculate(filter_data):
    # 将数据中发布日期转为datetime格式，用于时间过滤
    filter_data.loc[:, 'posting_date'] = pd.to_datetime(
        filter_data['posting_date'],
        format='%Y-%m-%dT%H:%M:%S%z',
        utc=True  # 显式指定转换为 UTC
    )
    now = datetime.datetime.now(datetime.timezone.utc)  # 统一用 UTC 时间
    # 设置10年前的时间点
    cutoff_date = datetime.datetime(now.year - 10, 1, 1, tzinfo=datetime.timezone.utc)
    # 执行过滤
    filter_year = filter_data[filter_data['posting_date'] > cutoff_date]
    # 根据生产厂商分组后，计算价格的中位数和均值，保留两位小数。
    price_df = filter_year.groupby('manufacturer').agg({'price': ['median', 'mean']}).round(2)
    # 重新命名列
    price_df.columns = ['median', 'mean']
    for producer, row in price_df.iterrows():
        print(f"生产商: {producer}，最近10年价格的中位数：{row['median']},平均数：{row['mean']}")
    print()
    print()


calculate(filter_data)
calculate(cleaned_df)
