# 题目：给定一个股票价格数组，计算以下指标
import numpy as np

prices = np.array([100, 102, 98, 105, 107, 103, 99, 101, 104, 106])

# 要求计算：
# a) 日收益率（每日价格变化百分比）
# b) 累计收益率
# c) 最大回撤
# d) 夏普比率（假设无风险利率为2%）

# for i=0; i < len(prices);i++:
#     print(k)
#     print(prices[1])
#     # print(k)

daily_returns = (prices[1:]-prices[:-1])/prices[:-1] *100
print(daily_returns)

for i, ret in enumerate(daily_returns):
    print(f"   第{i+1}天: {ret:.2f}%")
print(f"   平均日收益率: {np.mean(daily_returns):.2f}%")
print()


cumulative_returns = (prices / prices[0] - 1) * 100
print("b) 累计收益率 (%)：")
for i, cum_ret in enumerate(cumulative_returns):
    print(f"   第{i}天: {cum_ret:.2f}%")
print(f"   总累计收益率: {cumulative_returns[-1]:.2f}%")
print()
