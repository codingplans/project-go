
def simulate(N: int, suspension_days: list[list[int]], daily_limit: int = 500):
    """
    N: 每只股票的目标买入总量
    suspension_days: 三个子列表，分别是 A/B/C 股票的停牌日
    返回：三个子列表，分别是 A/B/C 每天的实际买入量（长度会各自截断到买满 N 为止）
    """
    # 参数配置
    bases     = [100, 200, 300]     # 原始基准
    cycles    = [1,   2,   3  ]     # 轮次：A 每 1 天，B 每 2 天，C 每 3 天
    miss_cnt  = [0,   0,   0  ]     # 累计的 miss 计数
    bought    = [0,   0,   0  ]     # 累计已买入
    plans     = [[],  [],  [] ]     # 最终要输出的“每天买入量”列表
    day = 1

    # 判断是否所有股票均已买满 N
    def done():
        return all(b >= N for b in bought)

    while not done():
        # —— 第一步：算出每只股票“原始累计计划”、“落后量”，并给出预期当天买入量 desired
        orig_cum = [0,0,0]
        lag      = [0,0,0]
        desired  = [0,0,0]
        for i in range(3):
            base = bases[i]
            cyc  = cycles[i]

            # 1) 计算到今天(含)原始应累计买入：scheduled_days * base
            #    scheduled_days = floor((day-1)/cyc)+1
            sched_days = ((day-1)//cyc) + 1
            orig_cum[i] = sched_days * base

            # 2) 当天“落后量” = 原始累计计划 - 已买量
            lag[i] = orig_cum[i] - bought[i]

            # 3) 如果是该股的“买入日” 并且当天 **不** 停牌，则给出预期购买量
            if (day-1) % cyc == 0 and day not in suspension_days[i]:
                tgt = int(base * (1 + 0.3 * miss_cnt[i]))           # 加速后的量
                # cap 在“原始累计 - 已买”以内
                tgt = min(tgt, orig_cum[i] - bought[i])
                desired[i] = max(tgt, 0)
            # 否则 desired[i] 保持 0

        # —— 第二步：做全体 500 股的流动性分配
        total_desired = sum(desired)
        actual = [0,0,0]
        if total_desired <= daily_limit:
            actual = desired.copy()
        else:
            # 按 lag 降序；lag 相同时按 i (A=0,B=1,C=2) 先后顺序
            order = sorted(range(3), key=lambda i: (-lag[i], i))
            rem   = daily_limit
            for i in order:
                take = min(desired[i], rem)
                actual[i] = take
                rem -= take
            # 其他股票 actual[i] = 0

        # —— 第三步：记录结果 & 更新累积 & miss 计数
        for i in range(3):
            plans[i].append(actual[i])
            bought[i] += actual[i]
            # 只有在“买入日”时，如果 actual < desired，就算一次 miss
            if (day-1) % cycles[i] == 0:
                if actual[i] < desired[i]:
                    miss_cnt[i] += 1

        day += 1

    return plans

if __name__ == "__main__":
    # —— 测试一下给的例子
    N = 1000
    case1 = simulate(N, [[1,2], [3], [1]])
    print("case1 A:", case1[0])
    print("case1 B:", case1[1])
    print("case1 C:", case1[2])

    case2 = simulate(N, [[],[],[]])
    print("case2 A:", case2[0])
    print("case2 B:", case2[1])
    print("case2 C:", case2[2])