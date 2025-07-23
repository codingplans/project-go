import numpy as np
import pandas as pd


# 生成30天模拟收盘价数据
def generate_sample_data():
    """生成30天的模拟股价数据"""
    dates = pd.date_range(start='2025-04-11', periods=30, freq='D')

    # 模拟股价走势（添加一些随机波动和趋势）
    np.random.seed(42)  # 确保结果可重现
    base_price = 100.0
    prices = [base_price]

    for i in range(1, 30):
        # 添加趋势和随机波动
        trend = 0.1 * np.sin(i * 0.2)  # 正弦趋势
        noise = np.random.normal(0, 1.5)  # 随机噪声
        change = trend + noise
        new_price = prices[-1] * (1 + change / 100)
        prices.append(max(new_price, 50))  # 防止价格过低

    return pd.DataFrame({
        'Date': dates,
        'Close': prices
    })


# 使用pandas计算MACD的几种方法
def calculate_macd_pandas_basic(df, fast=12, slow=26, signal=9):
    """使用pandas基本方法计算MACD"""
    data = df.copy()

    # 方法1：使用ewm()函数计算EMA
    data['EMA_12'] = data['Close'].ewm(span=fast).mean()
    data['EMA_26'] = data['Close'].ewm(span=slow).mean()

    # 计算MACD线（DIF）
    data['MACD'] = data['EMA_12'] - data['EMA_26']

    # 计算信号线（DEA）
    data['Signal'] = data['MACD'].ewm(span=signal).mean()

    # 计算柱状图
    data['Histogram'] = data['MACD'] - data['Signal']

    return data


def calculate_macd_pandas_advanced(df, fast=12, slow=26, signal=9):
    """使用pandas高级方法计算MACD"""
    data = df.copy()

    # 方法2：使用adjust参数控制初始化方式
    data['EMA_12_adj'] = data['Close'].ewm(span=fast, adjust=False).mean()
    data['EMA_26_adj'] = data['Close'].ewm(span=slow, adjust=False).mean()

    # 方法3：使用alpha参数（与span等价）
    alpha_fast = 2 / (fast + 1)
    alpha_slow = 2 / (slow + 1)
    alpha_signal = 2 / (signal + 1)

    data['EMA_12_alpha'] = data['Close'].ewm(alpha=alpha_fast, adjust=False).mean()
    data['EMA_26_alpha'] = data['Close'].ewm(alpha=alpha_slow, adjust=False).mean()

    # 计算MACD
    data['MACD_adj'] = data['EMA_12_adj'] - data['EMA_26_adj']
    data['MACD_alpha'] = data['EMA_12_alpha'] - data['EMA_26_alpha']

    # 计算信号线
    data['Signal_adj'] = data['MACD_adj'].ewm(span=signal, adjust=False).mean()
    data['Signal_alpha'] = data['MACD_alpha'].ewm(alpha=alpha_signal, adjust=False).mean()

    # 计算柱状图
    data['Histogram_adj'] = data['MACD_adj'] - data['Signal_adj']
    data['Histogram_alpha'] = data['MACD_alpha'] - data['Signal_alpha']

    return data


def calculate_macd_talib_style(df, fast=12, slow=26, signal=9):
    """模拟TA-Lib风格的MACD计算"""
    data = df.copy()

    # 使用com参数（center of mass）
    # com = (span - 1) / 2
    com_fast = (fast - 1) / 2
    com_slow = (slow - 1) / 2
    com_signal = (signal - 1) / 2

    data['EMA_12_com'] = data['Close'].ewm(com=com_fast, adjust=False).mean()
    data['EMA_26_com'] = data['Close'].ewm(com=com_slow, adjust=False).mean()

    data['MACD_com'] = data['EMA_12_com'] - data['EMA_26_com']
    data['Signal_com'] = data['MACD_com'].ewm(com=com_signal, adjust=False).mean()
    data['Histogram_com'] = data['MACD_com'] - data['Signal_com']

    return data


def compare_methods(df):
    """比较不同计算方法的结果"""
    basic = calculate_macd_pandas_basic(df)
    advanced = calculate_macd_pandas_advanced(df)
    talib_style = calculate_macd_talib_style(df)

    # 合并结果进行比较
    comparison = pd.DataFrame({
        'Date': df['Date'],
        'Close': df['Close'],
        'MACD_basic': basic['MACD'],
        'MACD_adjust': advanced['MACD_adj'],
        'MACD_alpha': advanced['MACD_alpha'],
        'MACD_com': talib_style['MACD_com'],
        'Signal_basic': basic['Signal'],
        'Signal_adjust': advanced['Signal_adj'],
        'Histogram_basic': basic['Histogram'],
        'Histogram_adjust': advanced['Histogram_adj']
    })

    return comparison


def plot_macd(df, title="MACD分析"):
    """绘制MACD图表"""
    fig, (ax1, ax2) = plt.subplots(2, 1, figsize=(12, 8), sharex=True)

    # 上图：价格和EMA
    ax1.plot(df['Date'], df['Close'], label='收盘价', linewidth=2)
    ax1.plot(df['Date'], df['EMA_12'], label='EMA12', alpha=0.7)
    ax1.plot(df['Date'], df['EMA_26'], label='EMA26', alpha=0.7)
    ax1.set_title(f'{title} - 价格与EMA')
    ax1.legend()
    ax1.grid(True, alpha=0.3)

    # 下图：MACD
    ax2.plot(df['Date'], df['MACD'], label='MACD', linewidth=2)
    ax2.plot(df['Date'], df['Signal'], label='Signal', linewidth=2)
    ax2.bar(df['Date'], df['Histogram'], label='Histogram', alpha=0.6, width=0.8)
    ax2.axhline(y=0, color='black', linestyle='-', alpha=0.3)
    ax2.set_title(f'{title} - MACD指标')
    ax2.legend()
    ax2.grid(True, alpha=0.3)

    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.show()


def main():
    # 生成30天数据
    df = generate_sample_data()

    print("=== 30天收盘价数据 ===")
    print(df.head(10))
    print(f"...共{len(df)}天数据")
    print()

    print("=== Pandas EWM参数说明 ===")
    print("span=12 等价于 alpha=2/(12+1)=0.1538")
    print("span=26 等价于 alpha=2/(26+1)=0.0741")
    print("adjust=True(默认): 使用调整因子，适合历史分析")
    print("adjust=False: 不使用调整因子，适合实时计算")
    print()

    # 基本MACD计算
    basic_result = calculate_macd_pandas_basic(df)

    print("=== 基本MACD计算结果(最后10天) ===")
    columns_to_show = ['Date', 'Close', 'EMA_12', 'EMA_26', 'MACD', 'Signal', 'Histogram']
    print(basic_result[columns_to_show].tail(10).round(4))
    print()

    # 高级方法比较
    print("=== 不同计算方法对比(最后5天) ===")
    comparison = compare_methods(df)
    print(comparison.tail(5).round(4))
    print()

    # 显示当前MACD信号
    latest = basic_result.iloc[-1]
    print("=== 最新MACD分析 ===")
    print(f"日期: {latest['Date'].strftime('%Y-%m-%d')}")
    print(f"收盘价: {latest['Close']:.2f}")
    print(f"MACD: {latest['MACD']:.4f}")
    print(f"Signal: {latest['Signal']:.4f}")
    print(f"Histogram: {latest['Histogram']:.4f}")

    # 判断信号
    if latest['MACD'] > latest['Signal']:
        if latest['Histogram'] > 0:
            signal_status = "多头信号"
        else:
            signal_status = "多头减弱"
    else:
        if latest['Histogram'] < 0:
            signal_status = "空头信号"
        else:
            signal_status = "空头减弱"

    print(f"信号状态: {signal_status}")

    # 检查金叉死叉
    if len(basic_result) >= 2:
        prev_macd = basic_result.iloc[-2]['MACD']
        prev_signal = basic_result.iloc[-2]['Signal']
        curr_macd = latest['MACD']
        curr_signal = latest['Signal']

        if prev_macd <= prev_signal and curr_macd > curr_signal:
            print("🚀 金叉信号！MACD上穿Signal线")
        elif prev_macd >= prev_signal and curr_macd < curr_signal:
            print("⚠️ 死叉信号！MACD下穿Signal线")

    print()

    # pandas计算性能测试
    print("=== Pandas计算性能优势 ===")
    import time

    # 生成更大的数据集进行性能测试
    large_df = generate_sample_data()
    # 扩展到1000天数据
    dates_extended = pd.date_range(start='2021-01-01', periods=1000, freq='D')
    np.random.seed(42)
    prices_extended = np.random.lognormal(np.log(100), 0.02, 1000)
    large_df = pd.DataFrame({'Date': dates_extended, 'Close': prices_extended})

    start_time = time.time()
    for _ in range(100):  # 重复100次
        calculate_macd_pandas_basic(large_df)
    pandas_time = time.time() - start_time

    print(f"Pandas计算1000天数据×100次: {pandas_time:.4f}秒")
    print("Pandas优势: 向量化计算，内存高效，代码简洁")
    print()

    print("=== 常用Pandas MACD函数总结 ===")
    print("1. df['Close'].ewm(span=12).mean() - 基本EMA计算")
    print("2. df['Close'].ewm(span=12, adjust=False).mean() - 实时EMA")
    print("3. df['Close'].ewm(alpha=0.1538).mean() - 使用alpha参数")
    print("4. df['Close'].ewm(com=5.5).mean() - 使用center of mass")
    print("5. df['MACD'].ewm(span=9).mean() - 信号线计算")

    # 绘制图表（注释掉以避免在命令行环境中出错）
    # plot_macd(basic_result, "30天MACD分析")


if __name__ == "__main__":
    main()