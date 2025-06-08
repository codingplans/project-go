import numpy as np

# 给定股票价格数据
prices = np.array([100, 102, 98, 105, 107, 103, 99, 101, 104, 106])
print("原始价格数据:")
print(prices)
print()

# a) 日收益率（每日价格变化百分比）
# 公式：(今日价格 - 昨日价格) / 昨日价格
daily_returns = (prices[1:] - prices[:-1]) / prices[:-1]
daily_returns2 = (prices[1:]/prices[:-1]) -1
print(prices[:1])
print(prices[:-1])
print(prices[1:])
print("a) 日收益率:")
print(f"数值: {daily_returns}")
print(f"数值: {daily_returns2}")
print("百分比:", [f"{x*100:.2f}%" for x in daily_returns])
print(f"平均日收益率: {np.mean(daily_returns):.4f} ({np.mean(daily_returns)*100:.2f}%)")
print()

# b) 计算累计收益率
cumulative_returns = (prices / prices[0] - 1) * 100
print("b) 累计收益率 (%)：")
for i, cum_ret in enumerate(cumulative_returns):
    print(f"   第{i}天: {cum_ret:.2f}%")
print(f"   总累计收益率: {cumulative_returns[-1]:.2f}%")
print(type(enumerate(cumulative_returns)))
# c) 计算最大回撤
# 计算每个时点的累计最高价格
peak_prices = np.maximum.accumulate(prices)
# 计算回撤（从最高点下跌的百分比）
drawdowns = (prices - peak_prices) / peak_prices * 100
max_drawdown = np.min(drawdowns)

print("c) 回撤分析：")
print("   各时点回撤 (%)：")
for i, dd in enumerate(drawdowns):
    print(f"   第{i}天: {dd:.2f}%")
print(f"   最大回撤: {max_drawdown:.2f}%")
print()


print(f"收益率波动率（标准差）: {np.std(daily_returns, ddof=1):.2f}%")

print(f"最大单日涨幅: {np.max(daily_returns):.2f}%")
print(f"最大单日跌幅: {np.min(daily_returns):.2f}%")
print(f"正收益天数: {daily_returns > 0}")
print(f"正收益天数: {np.sum(daily_returns > 0)}")
print(f"负收益天数: {np.sum(daily_returns < 0)}")
print(f"胜率: {np.sum(daily_returns > 0) / len(daily_returns) * 100:.1f}%")




# 设置随机种子以便结果可重现
np.random.seed(42)

# 生成模拟数据
returns = np.random.randn(252, 5)  # 252个交易日，5只股票
weights = np.array([0.2, 0.3, 0.2, 0.15, 0.15])

print("投资组合风险计算")
print("=" * 50)

# 验证权重和为1
print(f"权重向量: {weights}")
print(f"权重总和: {weights.sum():.3f}")
print()

# a) 计算协方差矩阵
print("步骤 a) 计算协方差矩阵")
print("-" * 30)

# 方法1：使用numpy的cov函数（默认计算样本协方差）
cov_matrix = np.cov(returns.T)  # 转置因为cov函数默认每行是一个变量

print("协方差矩阵 (5x5):")
print(f"矩阵形状: {cov_matrix.shape}")
print(cov_matrix.round(4))
print()

# 显示协方差矩阵的含义
print("协方差矩阵解释:")
for i in range(5):
    for j in range(5):
        if i == j:
            print(f"资产{i+1}的方差: {cov_matrix[i,j]:.4f}")
        elif i < j:  # 只显示上三角矩阵的协方差
            print(f"资产{i+1}与资产{j+1}的协方差: {cov_matrix[i,j]:.4f}")

print()

# b) 计算投资组合方差
print("步骤 b) 计算投资组合方差")
print("-" * 30)

# 投资组合方差公式: σ²_p = w^T × Σ × w
# 其中 w 是权重向量，Σ 是协方差矩阵

portfolio_variance = np.dot(weights.T, np.dot(cov_matrix, weights))

print("投资组合方差计算过程:")
print(f"公式: σ²_p = w^T × Σ × w")
print(f"权重向量 w: {weights}")
print(f"投资组合方差: {portfolio_variance:.6f}")
print()

# 详细计算步骤
print("详细计算步骤:")
step1 = np.dot(cov_matrix, weights)
print(f"步骤1 - Σ × w: {step1.round(4)}")
step2 = np.dot(weights.T, step1)
print(f"步骤2 - w^T × (Σ × w): {step2:.6f}")
print()

# c) 计算投资组合标准差（波动率）
print("步骤 c) 计算投资组合标准差（波动率）")
print("-" * 30)

portfolio_std = np.sqrt(portfolio_variance)

print(f"投资组合标准差（波动率）: {portfolio_std:.6f}")
print(f"投资组合年化波动率: {portfolio_std * np.sqrt(252):.2%}")
print()

# 额外分析：各资产的个别风险
print("额外分析：个别资产风险")
print("-" * 30)

individual_stds = np.sqrt(np.diag(cov_matrix))
print("各资产标准差:")
for i, std in enumerate(individual_stds):
    print(f"资产{i+1}: {std:.4f}")

print(f"\n加权平均标准差: {np.sum(weights * individual_stds):.4f}")
print(f"投资组合标准差: {portfolio_std:.4f}")
print(f"分散化效应: {((np.sum(weights * individual_stds) - portfolio_std) / np.sum(weights * individual_stds)):.2%}")

# 验证计算的另一种方法
print("\n验证：使用矩阵运算的另一种表示")
print("-" * 30)

# 重新整理权重为列向量进行矩阵运算
w = weights.reshape(-1, 1)  # 列向量
portfolio_var_alt = (w.T @ cov_matrix @ w)[0, 0]
print(f"使用 @ 运算符计算的投资组合方差: {portfolio_var_alt:.6f}")
print(f"两种方法结果一致: {np.isclose(portfolio_variance, portfolio_var_alt)}")

# 相关系数矩阵
print("\n相关系数矩阵")
print("-" * 30)
correlation_matrix = np.corrcoef(returns.T)
print("相关系数矩阵:")
print(correlation_matrix.round(3))

# 风险贡献分析
print("\n风险贡献分析")
print("-" * 30)
risk_contributions = (weights * np.dot(cov_matrix, weights)) / portfolio_variance
print("各资产的风险贡献:")
for i, contrib in enumerate(risk_contributions):
    print(f"资产{i+1}: {contrib:.2%}")
print(f"风险贡献总和: {risk_contributions.sum():.2%}")









