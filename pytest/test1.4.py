import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from datetime import datetime, time
import warnings

warnings.filterwarnings('ignore')

# 设置中文字体支持
plt.rcParams['font.sans-serif'] = ['SimHei', 'Arial Unicode MS', 'DejaVu Sans']
plt.rcParams['axes.unicode_minus'] = False

# 创建分钟级股票数据
np.random.seed(42)  # 设置随机种子以便复现
minute_data = pd.DataFrame({
    'timestamp': pd.date_range('2023-01-01 09:30', '2023-01-01 16:00', freq='1min'),
    'price': 100 + np.random.randn(391).cumsum() * 0.1,
    'volume': np.random.randint(100, 1000, 391)
})

print("原始分钟级数据样本：")
print(minute_data.head(10))
print(f"\n数据总量：{len(minute_data)} 条记录")
print(f"时间范围：{minute_data['timestamp'].min()} 到 {minute_data['timestamp'].max()}")


# =============================================================================
# a) 聚合为5分钟K线数据（OHLCV）
# =============================================================================

def create_5min_klines(df):
    """
    将分钟级数据聚合为5分钟K线数据
    """
    # 设置timestamp为索引
    df_temp = df.set_index('timestamp')

    # 使用resample进行5分钟聚合
    klines_5min = df_temp.resample('5min').agg({
        'price': ['first', 'max', 'min', 'last'],  # OHLC
        'volume': 'sum'  # 成交量求和
    }).round(4)

    # 重命名列
    klines_5min.columns = ['open', 'high', 'low', 'close', 'volume']

    # 重置索引
    klines_5min = klines_5min.reset_index()

    return klines_5min


# 生成5分钟K线数据
klines_5min = create_5min_klines(minute_data)

print("\n" + "=" * 60)
print("a) 5分钟K线数据（OHLCV）：")
print("=" * 60)
print(klines_5min.head(10))
print(f"\n5分钟K线数量：{len(klines_5min)} 条")


# =============================================================================
# b) 计算VWAP（成交量加权平均价格）
# =============================================================================

def calculate_vwap(df):
    """
    计算VWAP（Volume Weighted Average Price）
    VWAP = Σ(价格 × 成交量) / Σ(成交量)
    """
    df_temp = df.copy()

    # 计算价格×成交量
    df_temp['price_volume'] = df_temp['price'] * df_temp['volume']

    # 设置timestamp为索引进行重采样
    df_temp = df_temp.set_index('timestamp')

    # 5分钟聚合计算VWAP
    vwap_5min = df_temp.resample('5min').agg({
        'price_volume': 'sum',
        'volume': 'sum',
        'price': ['first', 'max', 'min', 'last']
    })

    # 计算VWAP
    vwap_5min['vwap'] = (vwap_5min[('price_volume', 'sum')] /
                         vwap_5min[('volume', 'sum')]).round(4)

    # 重新整理列结构
    result = pd.DataFrame({
        'timestamp': vwap_5min.index,
        'open': vwap_5min[('price', 'first')],
        'high': vwap_5min[('price', 'max')],
        'low': vwap_5min[('price', 'min')],
        'close': vwap_5min[('price', 'last')],
        'volume': vwap_5min[('volume', 'sum')],
        'vwap': vwap_5min['vwap']
    }).reset_index(drop=True)

    return result


# 计算包含VWAP的5分钟数据
klines_with_vwap = calculate_vwap(minute_data)

print("\n" + "=" * 60)
print("b) 包含VWAP的5分钟K线数据：")
print("=" * 60)
print(klines_with_vwap.head(10))

# 计算全天VWAP
total_price_volume = (minute_data['price'] * minute_data['volume']).sum()
total_volume = minute_data['volume'].sum()
daily_vwap = total_price_volume / total_volume

print(f"\n全天VWAP：{daily_vwap:.4f}")


# =============================================================================
# c) 处理缺失数据和交易时间过滤
# =============================================================================

def filter_trading_hours(df, start_time='09:30', end_time='15:00'):
    """
    过滤交易时间，处理缺失数据
    """
    df_filtered = df.copy()

    # 确保timestamp列为datetime类型
    if not pd.api.types.is_datetime64_any_dtype(df_filtered['timestamp']):
        df_filtered['timestamp'] = pd.to_datetime(df_filtered['timestamp'])

    # 提取时间部分进行过滤
    df_filtered['time'] = df_filtered['timestamp'].dt.time

    # 转换时间字符串为time对象
    start_time_obj = datetime.strptime(start_time, '%H:%M').time()
    end_time_obj = datetime.strptime(end_time, '%H:%M').time()

    # 过滤交易时间
    trading_hours_mask = (df_filtered['time'] >= start_time_obj) & (df_filtered['time'] <= end_time_obj)
    df_filtered = df_filtered[trading_hours_mask].copy()

    # 删除辅助列
    df_filtered = df_filtered.drop('time', axis=1)

    # 处理缺失数据
    # 检查是否有缺失值
    missing_data = df_filtered.isnull().sum()

    if missing_data.any():
        print("发现缺失数据：")
        print(missing_data[missing_data > 0])

        # 价格缺失值用前向填充
        df_filtered['price'] = df_filtered['price'].fillna(method='ffill')
        # 成交量缺失值用0填充
        df_filtered['volume'] = df_filtered['volume'].fillna(0)

    return df_filtered, missing_data


# 创建一些包含缺失数据的测试数据
test_data = minute_data.copy()
# 随机设置一些缺失值
np.random.seed(123)
missing_indices = np.random.choice(test_data.index, size=10, replace=False)
test_data.loc[missing_indices[:5], 'price'] = np.nan
test_data.loc[missing_indices[5:], 'volume'] = np.nan

print("\n" + "=" * 60)
print("c) 缺失数据处理和交易时间过滤：")
print("=" * 60)

# 处理缺失数据和过滤交易时间
filtered_data, missing_info = filter_trading_hours(test_data, '09:30', '15:00')

print("缺失数据统计：")
print(missing_info)

print(f"\n过滤后数据量：{len(filtered_data)} 条")
print(f"过滤前数据量：{len(test_data)} 条")

# 重新聚合处理后的数据
final_klines = calculate_vwap(filtered_data)

print("\n处理后的5分钟K线数据（前10条）：")
print(final_klines.head(10))

# =============================================================================
# 数据可视化
# =============================================================================

# 创建图表
fig, axes = plt.subplots(2, 2, figsize=(15, 10))
fig.suptitle('股票数据聚合分析', fontsize=16, fontweight='bold')

# 1. 原始分钟数据价格走势
axes[0, 0].plot(minute_data['timestamp'], minute_data['price'],
                linewidth=0.8, color='blue', alpha=0.7)
axes[0, 0].set_title('原始分钟级价格走势')
axes[0, 0].set_xlabel('时间')
axes[0, 0].set_ylabel('价格')
axes[0, 0].grid(True, alpha=0.3)
axes[0, 0].tick_params(axis='x', rotation=45)

# 2. 5分钟K线图
for i, row in final_klines.iterrows():
    color = 'red' if row['close'] >= row['open'] else 'green'
    # 绘制K线实体
    axes[0, 1].plot([row['timestamp'], row['timestamp']],
                    [row['low'], row['high']], color='black', linewidth=1)
    # 绘制开盘收盘价格矩形
    height = abs(row['close'] - row['open'])
    bottom = min(row['open'], row['close'])
    axes[0, 1].bar(row['timestamp'], height, bottom=bottom,
                   width=pd.Timedelta(minutes=3), color=color, alpha=0.7)

axes[0, 1].set_title('5分钟K线图')
axes[0, 1].set_xlabel('时间')
axes[0, 1].set_ylabel('价格')
axes[0, 1].grid(True, alpha=0.3)
axes[0, 1].tick_params(axis='x', rotation=45)

# 3. VWAP vs 收盘价对比
axes[1, 0].plot(final_klines['timestamp'], final_klines['close'],
                label='收盘价', linewidth=2, color='blue')
axes[1, 0].plot(final_klines['timestamp'], final_klines['vwap'],
                label='VWAP', linewidth=2, color='orange', linestyle='--')
axes[1, 0].axhline(y=daily_vwap, color='red', linestyle=':',
                   linewidth=2, label=f'全天VWAP: {daily_vwap:.2f}')
axes[1, 0].set_title('VWAP vs 收盘价对比')
axes[1, 0].set_xlabel('时间')
axes[1, 0].set_ylabel('价格')
axes[1, 0].legend()
axes[1, 0].grid(True, alpha=0.3)
axes[1, 0].tick_params(axis='x', rotation=45)

# 4. 成交量分布
axes[1, 1].bar(final_klines['timestamp'], final_klines['volume'],
               width=pd.Timedelta(minutes=3), alpha=0.7, color='purple')
axes[1, 1].set_title('5分钟成交量分布')
axes[1, 1].set_xlabel('时间')
axes[1, 1].set_ylabel('成交量')
axes[1, 1].grid(True, alpha=0.3)
axes[1, 1].tick_params(axis='x', rotation=45)

plt.tight_layout()
plt.show()

# =============================================================================
# 统计摘要
# =============================================================================

print("\n" + "=" * 60)
print("数据聚合统计摘要：")
print("=" * 60)

summary_stats = {
    '原始数据点数': len(minute_data),
    '5分钟K线数': len(final_klines),
    '聚合比例': f"{len(minute_data) / len(final_klines):.1f}:1",
    '价格范围': f"{minute_data['price'].min():.2f} - {minute_data['price'].max():.2f}",
    '全天VWAP': f"{daily_vwap:.4f}",
    '平均5分钟VWAP': f"{final_klines['vwap'].mean():.4f}",
    '总成交量': f"{minute_data['volume'].sum():,}",
    '平均5分钟成交量': f"{final_klines['volume'].mean():.0f}"
}

for key, value in summary_stats.items():
    print(f"{key:15}: {value}")

print("\n" + "=" * 60)
print("聚合完成！以上展示了完整的数据处理流程：")
print("1. 分钟数据聚合为5分钟K线（OHLCV）")
print("2. VWAP计算和分析")
print("3. 缺失数据处理和交易时间过滤")
print("4. 数据可视化和统计分析")
print("=" * 60)