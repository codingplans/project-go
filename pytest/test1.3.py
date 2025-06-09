import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

# 多只股票数据分析
# ```python
# # 题目：分析多只股票的相关性和表现
# # 假设有如下格式的数据：
# dates = pd.date_range('2023-01-01', periods=252)
# stocks = ['AAPL', 'GOOGL', 'MSFT', 'TSLA', 'AMZN']
# # 创建随机价格数据
# price_data = pd.DataFrame({
#     stock: 100 + np.random.randn(252).cumsum()
#     for stock in stocks
# }, index=dates)
# # 要求：
# # a) 计算各股票的日收益率
# # b) 计算相关系数矩阵
# # c) 找出表现最好和最差的股票
# # d) 计算滚动相关系数
# # e) 识别异常交易日

# 设置随机种子保证可复现
np.random.seed(42)

# 模拟数据
dates = pd.date_range('2023-01-01', periods=252)
stocks = ['AAPL', 'GOOGL', 'MSFT', 'TSLA', 'AMZN']
price_data = pd.DataFrame({
    stock: 100 + np.random.randn(252).cumsum()
    for stock in stocks
}, index=dates)

# a) 计算各股票的日收益率
returns = price_data.pct_change().dropna()

# b) 计算相关系数矩阵
correlation_matrix = returns.corr()


print("相关系数矩阵：\n", correlation_matrix)

# c) 找出表现最好和最差的股票（年内总收益率）
total_returns = (price_data.iloc[-1] / price_data.iloc[0]) - 1
best_stock = total_returns.idxmax()
worst_stock = total_returns.idxmin()
print(f"\n表现最好股票：{best_stock}，收益率：{total_returns[best_stock]:.2%}")
print(f"表现最差股票：{worst_stock}，收益率：{total_returns[worst_stock]:.2%}")

# d) 计算滚动相关系数（以 AAPL 为基准，窗口为20天）
rolling_corr = returns.rolling(window=20).corr(returns['AAPL'])
# 示例：提取 MSFT 与 AAPL 的滚动相关性
rolling_corr_msft = rolling_corr['MSFT']

# e) 识别异常交易日（收益率的z-score绝对值大于3）
z_scores = (returns - returns.mean()) / returns.std()
anomalies = (np.abs(z_scores) > 3)
anomalous_days = anomalies.any(axis=1)
print(f"\n发现异常交易日数量：{anomalous_days.sum()} 天")
print("异常交易日如下：\n", returns[anomalous_days])


# variance = returns.var()  # 默认是样本方差（n-1）
# returns.std() 样本标准差

plt.rcParams['font.family'] = 'Heiti TC'

# （可选）可视化
# 1. 收益率热力图
import seaborn as sns
plt.figure(figsize=(8,6))
sns.heatmap(correlation_matrix, annot=True, cmap='coolwarm')
plt.title("股票收益率相关系数矩阵")
plt.show()

# 2. 滚动相关图
plt.figure(figsize=(10,4))
rolling_corr_msft.plot()
plt.title("MSFT 与 AAPL 的20日滚动相关系数")
plt.ylabel("相关系数")
plt.grid(True)
plt.show()

# 3. 异常交易日标记图（以 AAPL 为例）
plt.figure(figsize=(12,5))
plt.plot(returns.index, returns['AAPL'], label='AAPL 日收益率')
plt.scatter(returns.index[anomalous_days], returns['AAPL'][anomalous_days], color='red', label='异常点')
plt.title("AAPL 收益率与异常交易日")
plt.legend()
plt.grid(True)
plt.show()