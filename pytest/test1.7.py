def stock_trading_plan(N, suspension_days):
    """
    股票交易计划程序

    参数:
    N: 每只股票要买入的总量
    suspension_days: 停牌日期序列，格式为[[A股票停牌日], [B股票停牌日], [C股票停牌日]]

    返回:
    三个列表，分别表示A、B、C股票每日的买入量
    """

    # 股票配置
    stocks = {
        'A': {'daily_plan': 100, 'cycle': 1, 'priority': 1},
        'B': {'daily_plan': 200, 'cycle': 2, 'priority': 2},
        'C': {'daily_plan': 300, 'cycle': 3, 'priority': 3}
    }

    # 初始化状态
    stock_names = ['A', 'B', 'C']
    total_bought = {stock: 0 for stock in stock_names}
    acceleration_factor = {stock: 1.0 for stock in stock_names}
    daily_purchases = {stock: [] for stock in stock_names}

    # 转换停牌日期为集合，便于查找
    suspension_sets = [set(days) for days in suspension_days]

    daily_limit = 500
    day = 1

    # 继续交易直到所有股票都买完
    while any(total_bought[stock] < N for stock in stock_names):
        # 计算每只股票的计划买入量
        planned_purchases = {}

        for i, stock in enumerate(stock_names):
            if total_bought[stock] >= N:
                planned_purchases[stock] = 0
                continue

            # 检查是否停牌
            if day in suspension_sets[i]:
                planned_purchases[stock] = 0
                # 停牌导致加速
                acceleration_factor[stock] *= 1.3
            else:
                # 检查是否应该买入（根据周期）
                config = stocks[stock]
                if (day - 1) % config['cycle'] == 0:
                    # 计算计划买入量
                    base_amount = config['daily_plan']
                    accelerated_amount = int(base_amount * acceleration_factor[stock])

                    # 不能超过剩余需要买入的量
                    remaining = N - total_bought[stock]
                    planned_amount = min(accelerated_amount, remaining)
                    planned_purchases[stock] = planned_amount
                else:
                    planned_purchases[stock] = 0

        # 计算每只股票相对于理想进度的落后程度
        deficits = {}
        for i, stock in enumerate(stock_names):
            if total_bought[stock] >= N:
                deficits[stock] = 0
                continue

            # 计算理想进度
            config = stocks[stock]
            ideal_total = 0
            for d in range(1, day + 1):
                if (d - 1) % config['cycle'] == 0:
                    ideal_total += config['daily_plan']

            # 限制理想进度不超过总目标
            ideal_total = min(ideal_total, N)
            deficits[stock] = max(0, ideal_total - total_bought[stock])

        # 按优先级排序（落后程度越大优先级越高，相同时按股票优先级）
        stocks_to_buy = []
        for stock in stock_names:
            if planned_purchases[stock] > 0:
                stocks_to_buy.append((stock, deficits[stock], stocks[stock]['priority']))

        # 排序：首先按落后程度降序，然后按优先级升序
        stocks_to_buy.sort(key=lambda x: (-x[1], x[2]))

        # 分配每日限额
        remaining_limit = daily_limit
        actual_purchases = {stock: 0 for stock in stock_names}

        for stock, deficit, priority in stocks_to_buy:
            desired_amount = planned_purchases[stock]
            if desired_amount > 0 and remaining_limit > 0:
                actual_amount = min(desired_amount, remaining_limit)
                actual_purchases[stock] = actual_amount
                remaining_limit -= actual_amount
                total_bought[stock] += actual_amount

                # 检查是否需要调整加速因子
                if actual_amount < desired_amount:
                    # 未能足额买入，增加加速
                    acceleration_factor[stock] *= 1.3
                else:
                    # 检查是否追上计划，如果追上则恢复正常
                    if deficits[stock] <= actual_amount:
                        acceleration_factor[stock] = 1.0

        # 记录每日买入量
        for stock in stock_names:
            daily_purchases[stock].append(actual_purchases[stock])

        day += 1

        # 防止无限循环
        if day > 100:
            break

    # 返回结果
    return [daily_purchases[stock] for stock in stock_names]


def enhanced_stock_trading_plan(stock_configs, N_values, suspension_days, daily_limit=500):
    """
    增强版股票交易计划程序，支持更多股票和自定义配置

    参数:
    stock_configs: 股票配置字典，格式为 {'股票名': {'daily_plan': 日计划量, 'cycle': 周期, 'priority': 优先级}}
    N_values: 每只股票要买入的总量字典，格式为 {'股票名': 总量}
    suspension_days: 停牌日期序列字典，格式为 {'股票名': [停牌日期列表]}
    daily_limit: 每日买入限制

    返回:
    字典，每只股票对应一个每日买入量列表
    """

    stock_names = list(stock_configs.keys())
    total_bought = {stock: 0 for stock in stock_names}
    acceleration_factor = {stock: 1.0 for stock in stock_names}
    daily_purchases = {stock: [] for stock in stock_names}

    # 转换停牌日期为集合
    suspension_sets = {stock: set(suspension_days.get(stock, [])) for stock in stock_names}

    day = 1

    # 继续交易直到所有股票都买完
    while any(total_bought[stock] < N_values[stock] for stock in stock_names):
        # 计算每只股票的计划买入量
        planned_purchases = {}

        for stock in stock_names:
            N = N_values[stock]
            if total_bought[stock] >= N:
                planned_purchases[stock] = 0
                continue

            # 检查是否停牌
            if day in suspension_sets[stock]:
                planned_purchases[stock] = 0
                # 停牌导致加速
                acceleration_factor[stock] *= 1.3
            else:
                # 检查是否应该买入（根据周期）
                config = stock_configs[stock]
                if (day - 1) % config['cycle'] == 0:
                    # 计算计划买入量
                    base_amount = config['daily_plan']
                    accelerated_amount = int(base_amount * acceleration_factor[stock])

                    # 不能超过剩余需要买入的量
                    remaining = N - total_bought[stock]
                    planned_amount = min(accelerated_amount, remaining)
                    planned_purchases[stock] = planned_amount
                else:
                    planned_purchases[stock] = 0

        # 计算每只股票相对于理想进度的落后程度
        deficits = {}
        for stock in stock_names:
            N = N_values[stock]
            if total_bought[stock] >= N:
                deficits[stock] = 0
                continue

            # 计算理想进度
            config = stock_configs[stock]
            ideal_total = 0
            for d in range(1, day + 1):
                if (d - 1) % config['cycle'] == 0:
                    ideal_total += config['daily_plan']

            # 限制理想进度不超过总目标
            ideal_total = min(ideal_total, N)
            deficits[stock] = max(0, ideal_total - total_bought[stock])

        # 按优先级排序（落后程度越大优先级越高，相同时按股票优先级）
        stocks_to_buy = []
        for stock in stock_names:
            if planned_purchases[stock] > 0:
                stocks_to_buy.append((stock, deficits[stock], stock_configs[stock]['priority']))

        # 排序：首先按落后程度降序，然后按优先级升序
        stocks_to_buy.sort(key=lambda x: (-x[1], x[2]))

        # 分配每日限额
        remaining_limit = daily_limit
        actual_purchases = {stock: 0 for stock in stock_names}

        for stock, deficit, priority in stocks_to_buy:
            desired_amount = planned_purchases[stock]
            if desired_amount > 0 and remaining_limit > 0:
                actual_amount = min(desired_amount, remaining_limit)
                actual_purchases[stock] = actual_amount
                remaining_limit -= actual_amount
                total_bought[stock] += actual_amount

                # 检查是否需要调整加速因子
                if actual_amount < desired_amount:
                    # 未能足额买入，增加加速
                    acceleration_factor[stock] *= 1.3
                else:
                    # 检查是否追上计划，如果追上则恢复正常
                    if deficits[stock] <= actual_amount:
                        acceleration_factor[stock] = 1.0

        # 记录每日买入量
        for stock in stock_names:
            daily_purchases[stock].append(actual_purchases[stock])

        day += 1

        # 防止无限循环
        if day > 200:
            break

    return daily_purchases


def test_example():

    # case1:
    # input:
    # N: 1000,suspension_days: [[1, 2], [3], [1]]
    # output:
    # [[0,0,160,110,190,140,0,130,130,130,10],
    # [200,0,0,0,260,0,110,0,320,0,110],
    # [0,0,0,390,0,0,390,0,0,220,0]]
    #
    # case2:
    # input:
    # N:1000 suspension_days: [[], [], []]
    # [[100,100,100,100,100,100,100,100,100,100,0],
    # [200,0,200,0,200,0,90,0,260,0,50],
    # [200,0,0,390,0,0,310,0,0,100,0]]

    # trading_plan(1000, [[1, 2], [3], [1]])
    # # [[0,0,160,110,190,140,0,130,130,130,10],
    # # [200,0,0,0,260,0,110,0,320,0,110],
    # # [0,0,0,390,0,0,390,0,0,220,0]]
    # K.trading_plan(1000, [[], [], []])
    # # [
    # # [100,100,100,100,100,100,100,100,100,100,0],
    # # [200,0,200,0,200,0,90,0,260,0,50],
    # # [200,0,0,390,0,0,310,0,0,100,0]]
    """测试示例"""
    print("=== 测试原始示例 ===")
    N = 1000
    suspension_days = [[1, 2], [3], [1]]

    result = stock_trading_plan(N, suspension_days)

    print("最终结果:")
    print(f"A股票买入计划: {result[0]}")
    print(f"B股票买入计划: {result[1]}")
    print(f"C股票买入计划: {result[2]}")

    # 验证总量
    for i, stock in enumerate(['A', 'B', 'C']):
        total = sum(result[i])
        print(f"{stock}股票总买入量: {total}")

    print("\n=== 测试增强版本（支持更多股票）===")
    # 测试增强版本
    stock_configs = {
        'A': {'daily_plan': 100, 'cycle': 1, 'priority': 1},
        'B': {'daily_plan': 200, 'cycle': 2, 'priority': 2},
        'C': {'daily_plan': 300, 'cycle': 3, 'priority': 3},
        'D': {'daily_plan': 150, 'cycle': 4, 'priority': 4}
    }

    N_values = {'A': 800, 'B': 600, 'C': 900, 'D': 600}
    suspension_days_dict = {
        'A': [1, 2],
        'B': [3],
        'C': [1],
        'D': [5, 10]
    }

    enhanced_result = enhanced_stock_trading_plan(stock_configs, N_values, suspension_days_dict)

    print("增强版结果:")
    for stock, purchases in enhanced_result.items():
        total = sum(purchases)
        print(f"{stock}股票买入计划: {purchases[:15]}{'...' if len(purchases) > 15 else ''}")
        print(f"{stock}股票总买入量: {total}")


if __name__ == "__main__":
    test_example()