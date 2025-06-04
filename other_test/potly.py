import numpy as np
import matplotlib.pyplot as plt

# 设置参数
lambda_values = [0.5, 1, 2]  # 不同率参数lambda
x = np.linspace(0, 10, 1000)  # x轴范围

# 绘制图形
plt.figure(figsize=(8, 6))
for lam in lambda_values:
    # 指数分布的概率密度函数
    pdf = lam * np.exp(-lam * x)
    plt.plot(x, pdf, label=f'λ = {lam}')

# 添加标签和标题
plt.title('Exponential Distribution PDF')
plt.xlabel('x')
plt.ylabel('Probability Density')
plt.legend()
plt.grid(True)
plt.show()
