import warnings

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

warnings.filterwarnings('ignore')

# 设置中文字体
# plt.rcParams['font.sans-serif'] = ['SimHei', 'DejaVu Sans']
# plt.rcParams['axes.unicode_minus'] = False

plt.rcParams['font.family'] = 'Heiti TC'

plt.rcParams['axes.unicode_minus'] = False  # 正确显示负号
# 创建示例数据
np.random.seed(42)  # 设置随机种子以便结果可重现
data = {
    'date': pd.date_range('2023-01-01', periods=100),
    'open': np.random.randn(100).cumsum() + 100,
    'high': np.random.randn(100).cumsum() + 102,
    'low': np.random.randn(100).cumsum() + 98,
    'close': np.random.randn(100).cumsum() + 100,
    'volume': np.random.randint(1000, 10000, 100)
}
df = pd.DataFrame(data)


# 数据预处理：确保OHLC数据的逻辑一致性
def clean_ohlc_data(df):
    """清洗OHLC数据，确保High >= max(Open, Close), Low <= min(Open, Close)"""
    df_clean = df.copy()

    # 确保High是当天的最高价
    df_clean['high'] = np.maximum.reduce([df_clean['open'], df_clean['high'],
                                          df_clean['low'], df_clean['close']])

    # 确保Low是当天的最低价
    df_clean['low'] = np.minimum.reduce([df_clean['open'], df_clean['high'],
                                         df_clean['low'], df_clean['close']])

    return df_clean


df = clean_ohlc_data(df)

print("原始数据概览：")
print(df.head())
print(f"\n数据形状: {df.shape}")


# a) 计算移动平均线（5日、20日）
def calculate_moving_averages(df, periods=[5, 20]):
    """计算移动平均线"""
    df_ma = df.copy()

    for period in periods:
        df_ma[f'MA_{period}'] = df_ma['close'].rolling(window=period).mean()
        df_ma[f'MA_{period}_volume'] = df_ma['volume'].rolling(window=period).mean()

    return df_ma


df = calculate_moving_averages(df)


# b) 计算RSI指标
def calculate_rsi(prices, period=14):
    """计算RSI相对强弱指数"""
    delta = prices.diff()
    gain = (delta.where(delta > 0, 0)).rolling(window=period).mean()
    loss = (-delta.where(delta < 0, 0)).rolling(window=period).mean()
    # RS = 平均涨幅 / 平均跌幅
    rs = gain / loss
    rsi = 100 - (100 / (1 + rs))

    return rsi


df['RSI'] = calculate_rsi(df['close'])


# c) 识别并处理异常值
def detect_outliers(df):
    """使用IQR方法检测异常值"""
    outliers_info = {}

    for col in ['open', 'high', 'low', 'close', 'volume']:
        Q1 = df[col].quantile(0.25)
        Q3 = df[col].quantile(0.75)
        IQR = Q3 - Q1

        lower_bound = Q1 - 1.5 * IQR
        upper_bound = Q3 + 1.5 * IQR

        outliers = df[(df[col] < lower_bound) | (df[col] > upper_bound)]
        outliers_info[col] = {
            'count': len(outliers),
            'indices': outliers.index.tolist(),
            'values': outliers[col].tolist()
        }

    return outliers_info


outliers = detect_outliers(df)
print("\n异常值检测结果：")
for col, info in outliers.items():
    print(f"{col}: {info['count']} 个异常值")


# 处理异常值：使用中位数替换
def handle_outliers(df, method='median'):
    """处理异常值"""
    df_cleaned = df.copy()

    for col in ['open', 'high', 'low', 'close', 'volume']:
        Q1 = df_cleaned[col].quantile(0.25)
        Q3 = df_cleaned[col].quantile(0.75)
        IQR = Q3 - Q1

        lower_bound = Q1 - 1.5 * IQR
        upper_bound = Q3 + 1.5 * IQR

        # 使用中位数替换异常值
        median_val = df_cleaned[col].median()
        df_cleaned.loc[(df_cleaned[col] < lower_bound) | (df_cleaned[col] > upper_bound), col] = median_val

    return df_cleaned


df_cleaned = handle_outliers(df)


# d) 计算技术指标：MACD、布林带
def calculate_macd(prices, fast=12, slow=26, signal=9):
    """计算MACD指标"""
    ema_fast = prices.ewm(span=fast).mean()
    ema_slow = prices.ewm(span=slow).mean()

    macd_line = ema_fast - ema_slow
    signal_line = macd_line.ewm(span=signal).mean()
    histogram = macd_line - signal_line

    return macd_line, signal_line, histogram


def calculate_bollinger_bands(prices, period=20, std_dev=2):
    """计算布林带"""
    ma = prices.rolling(window=period).mean()
    std = prices.rolling(window=period).std()

    upper_band = ma + (std * std_dev)
    lower_band = ma - (std * std_dev)

    return upper_band, ma, lower_band


# 计算MACD
df['MACD'], df['MACD_Signal'], df['MACD_Histogram'] = calculate_macd(df['close'])

# 计算布林带
df['BB_Upper'], df['BB_Middle'], df['BB_Lower'] = calculate_bollinger_bands(df['close'])


# 计算额外的技术指标
def calculate_additional_indicators(df):
    """计算额外的技术指标"""
    # 波动率 (历史波动率)
    df['Volatility'] = df['close'].pct_change().rolling(window=20).std() * np.sqrt(252)

    # 价格变化率
    df['Price_Change'] = df['close'].pct_change()
    df['Price_Change_5d'] = df['close'].pct_change(5)

    # 真实波动幅度均值 (ATR)
    high_low = df['high'] - df['low']
    high_close = np.abs(df['high'] - df['close'].shift())
    low_close = np.abs(df['low'] - df['close'].shift())

    true_range = np.maximum(high_low, np.maximum(high_close, low_close))
    df['ATR'] = true_range.rolling(window=14).mean()

    return df


df = calculate_additional_indicators(df)


# 数据质量报告
def generate_data_quality_report(df):
    """生成数据质量报告"""
    print("\n=== 数据质量报告 ===")
    print(f"数据总行数: {len(df)}")
    print(f"数据总列数: {len(df.columns)}")
    print(f"日期范围: {df['date'].min()} 到 {df['date'].max()}")

    print("\n缺失值统计:")
    missing_values = df.isnull().sum()
    print(missing_values[missing_values > 0])

    print("\n基本统计信息:")
    numeric_cols = ['open', 'high', 'low', 'close', 'volume']
    print(df[numeric_cols].describe())


generate_data_quality_report(df)


# 创建可视化
def create_visualizations(df):
    """创建数据可视化"""
    fig, axes = plt.subplots(4, 1, figsize=(15, 20))

    # 1. 价格走势和移动平均线
    axes[0].plot(df['date'], df['close'], label='收盘价', linewidth=1)
    axes[0].plot(df['date'], df['MA_5'], label='5日均线', alpha=0.8)
    axes[0].plot(df['date'], df['MA_20'], label='20日均线', alpha=0.8)
    axes[0].fill_between(df['date'], df['BB_Upper'], df['BB_Lower'], alpha=0.2, label='布林带')
    axes[0].set_title('股价走势与技术指标', fontsize=14)
    axes[0].legend()
    axes[0].grid(True, alpha=0.3)

    # 2. 成交量
    axes[1].bar(df['date'], df['volume'], alpha=0.6, color='orange')
    axes[1].plot(df['date'], df['MA_5_volume'], label='5日成交量均线', color='red')
    axes[1].set_title('成交量', fontsize=14)
    axes[1].legend()
    axes[1].grid(True, alpha=0.3)

    # 3. RSI指标
    axes[2].plot(df['date'], df['RSI'], label='RSI', color='purple')
    axes[2].axhline(y=70, color='r', linestyle='--', alpha=0.7, label='超买线(70)')
    axes[2].axhline(y=30, color='g', linestyle='--', alpha=0.7, label='超卖线(30)')
    axes[2].set_title('RSI相对强弱指数', fontsize=14)
    axes[2].set_ylim(0, 100)
    axes[2].legend()
    axes[2].grid(True, alpha=0.3)

    # 4. MACD指标
    axes[3].plot(df['date'], df['MACD'], label='MACD', color='blue')
    axes[3].plot(df['date'], df['MACD_Signal'], label='信号线', color='red')
    axes[3].bar(df['date'], df['MACD_Histogram'], label='MACD柱', alpha=0.6, color='gray')
    axes[3].set_title('MACD指标', fontsize=14)
    axes[3].legend()
    axes[3].grid(True, alpha=0.3)

    plt.tight_layout()
    plt.show()


create_visualizations(df)

print("\n=== 最终数据样本 ===")
important_cols = ['date', 'open', 'high', 'low', 'close', 'volume',
                  'MA_5', 'MA_20', 'RSI', 'MACD', 'BB_Upper', 'BB_Lower']
print(df[important_cols].tail(10))

# 保存处理后的数据
df.to_csv('processed_stock_data.csv', index=False, encoding='utf-8')
print("\n数据已保存到 'processed_stock_data.csv'")


# 创建技术指标汇总
def create_technical_summary(df):
    """创建技术指标汇总"""
    latest = df.iloc[-1]

    print("\n=== 最新技术指标汇总 ===")
    print(f"日期: {latest['date'].strftime('%Y-%m-%d')}")
    print(f"收盘价: {latest['close']:.2f}")
    print(f"5日均线: {latest['MA_5']:.2f}")
    print(f"20日均线: {latest['MA_20']:.2f}")
    print(f"RSI: {latest['RSI']:.2f}")
    print(f"MACD: {latest['MACD']:.4f}")
    print(f"布林带上轨: {latest['BB_Upper']:.2f}")
    print(f"布林带下轨: {latest['BB_Lower']:.2f}")
    print(f"ATR: {latest['ATR']:.2f}")
    print(f"波动率: {latest['Volatility']:.4f}")

    # 简单的技术分析建议
    print("\n=== 技术分析参考 ===")
    if latest['RSI'] > 70:
        print("RSI提示: 可能超买")
    elif latest['RSI'] < 30:
        print("RSI提示: 可能超卖")
    else:
        print("RSI提示: 正常区间")

    if latest['close'] > latest['MA_5'] > latest['MA_20']:
        print("均线提示: 多头排列")
    elif latest['close'] < latest['MA_5'] < latest['MA_20']:
        print("均线提示: 空头排列")
    else:
        print("均线提示: 震荡格局")


create_technical_summary(df)
