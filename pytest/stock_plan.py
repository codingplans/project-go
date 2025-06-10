from typing import List


# ＃描述
# 我们有一个交易计划，分别买入A,B,C三只股票各N股。考虑三只股票的流动性我们准备执行下列计划：
# - A股票日计划每天买入100股
# - B股票日计划隔天买入200股
# - C股票日计划每3天买入300股，即从第一天开始到第N天的买入总量的计划（总计划）为300（买入），300，300，600（买入），600，600，900（买入），…
# -为了避免对市场造成冲击，我们对流动性进行控制，每日买入股票的总数量需要控制在500股以内。当受到控制时，优先保障买入当前（当日买入前）落后总买入计划最多的股票，如果落后数量相同，优先保障A股票的交易，然后是B股票，最后是C股票。
# -每只股票都会发生某日停牌不能交易的情况。当发生这种情况的时候，当日无法买入。
# -当某只股票由于停牌或者流动性控制导致没有按照预期足额买入时，买入量开始落后于总买入计划。所以下一次的买入数量增加原始日计划买入量的30%。加速可以叠加，但买入量不能超过总计划的进度。当买入量追赶上原总计划后恢复到按原计划执行。
# 请编写程序输入为每只股票要买入的总量（N）和停牌日期序列（Suspension），输出每日的买入量。
#
# ＃示例
# ##＃ 输入：
# -买入量：1000，表示A，B，C三只股票各买入1000股
# - 停牌：［［1,2],［3],［1］]，分别标识A，B，C股票在第i日停牌，日期从1开始。示例中A股票第1日和第2日连续停牌，B股票第3日停牌，C股票第1日停牌。
# ### 输出:
# 输出为一个列表中包含三个整数列表,列表内每个列表分别对应A,B,C三只股票每日的买入股数的整型。
# [[ 0, 0, 160, 110, 190, 140, 0, 130, 130,130,10],#A股票买入计划
# [200, 0, 0, 0, 260, 110, 0, 320,0,110], #B股票买入计划
# [ 0, 0, 0, 390, 0, 0, 390, 0, 0,220,0]]  #C股票买入计划
#
# ###说明:
# -第一天,A股票停牌不买入,B股票按计划买入200,C股票停牌不买入
# -第二天,A股票停牌不买入,B股票按计划轮空不买入,C股票轮空不买入
# -第三天,A股票由于停牌两天,买入计划加速两次100+230%6100=160,故买入160,B股票停牌,C股票轮空均不买入
# -第四天,B股票轮空,由于每日500股的限制,A股票落后总计划140,C股票落后总计划300,优先买入C股票,C股票由于停牌一日买入量为300+30%*300=390。故A股票无法按原计划买入160,而是买入110。
# -第五天,A股票再次加速30%买入190,B股票停牌后加速为260,C股票轮空
# -第六日,A股票买入140追上计划,其他两只股票轮空
# -第七日,由于500限制A股票无法买入,C股票由于缺口大全额买入,B股票分得剩余额度
# ...
# -第十一日,买入最后的股票,交易计划结束
#
# 提示
# 请仔细思考并设计代码结构,让代码能够容易扩展(如支特更多股票等)
#
#
# 这里是两个例子，请使该代码能通过测试：
# ```
# case1:
# input:
# N: 1000,suspension_days: [[1, 2], [3], [1]]
# output:
# [[0,0,160,110,190,140,0,130,130,130,10],
# [200,0,0,0,260,0,110,0,320,0,110],
# [0,0,0,390,0,0,390,0,0,220,0]]
# case2:
# input:
# N:1000 suspension_days: [[], [], []]
# [[100,100,100,100,100,100,100,100,100,100,0],
# [200,0,200,0,200,0,90,0,260,0,50],
# [200,0,0,390,0,0,310,0,0,100,0]]
# ```


class Solution:
    def __init__(self):
        self.plan = 1000
        self.day_max = 500
        self.has_traded = {1: [],
                           2: [],
                           3: []}
        self.suspension = {}
        self.need_supp_days = {1: {},
                               2: {},
                               3: {}}
        self.stockConfig = {1: {'day_volume': 100,
                                'step': 0,
                                'remainder': 0},
                            2: {'day_volume': 200,
                                'step': 2,
                                'remainder': 1},
                            3: {'day_volume': 300,
                                'step': 3,
                                'remainder': 1}, }

    def trading_plan(self, N: int, Suspension: List[List[int]]):
        self.plan = N
        for stock, suspension in enumerate(Suspension):
            self.suspension[stock] = suspension

        for day in range(1, int(self.plan / self.stockConfig[1]['day_volume']) + 3):
            current_day = {}
            current_supp_day = {}

            for stock in [1, 2, 3]:
                # 计划买入量
                volume = self.calculate_stock(stock, day, Suspension[stock - 1])
                # 累计计划买入总量
                plan_volume = self.plan_volume(stock, day, Suspension[stock - 1]) + volume
                # 当日需要补量
                supp_volume = self.supplement_volume(stock, day, Suspension[stock - 1])
                # 存储补量到字典中
                current_supp_day[stock] = supp_volume
                # 验算当天是否超额，重新梳理，按照顺序
                current_supp_day_new = self.filter_full_day(current_supp_day.copy())
                # 当日计划买入的总量
                total_volume = volume + current_supp_day_new[stock]
                # 之前已买入的累计总量
                pre_total_volume = self.pre_total_volume(stock)

                # 之前累计量和当日总量超过计划累计总量 则消减
                if pre_total_volume + total_volume > plan_volume:
                    total_volume = plan_volume - pre_total_volume
                    # 表示之前的欠额都补齐了，全部消掉
                    self.need_supp_days[stock] = {}
                # 判断当日总量加之前总量是否超过 总计划，若超过就再消减，且进行标注
                if total_volume + pre_total_volume > self.plan:
                    total_volume = self.plan - pre_total_volume
                    self.stockConfig[stock]['done'] = True
                # 最后得到当日真正要买入量
                current_day[stock] = total_volume
                # 重新验算是否有超额情况，按照股票顺序和欠款顺序 验算
                current_day_new = self.full_day_order_by_stock(current_day.copy(), current_supp_day_new)
                # 如果当日停牌，进行标注
                if self.is_stop(day, Suspension[stock - 1]):
                    self.need_supp_days[stock][day] = True
            # 梳理当天所有票是否有补充逻辑，有的话进行标注，用于后续计算补充量
            for stock, volume in current_day_new.items():
                if current_day[stock] != volume:
                    self.need_supp_days[stock][day] = True
            # 填充所有票的当日买入量
            for stock, traded in current_day_new.items():
                self.has_traded[stock].append(current_day_new[stock])
            print(f"当前{day}计划买入{current_day_new}")
        self.print_output()
        return list(self.has_traded.values())

    def print_output(self):
        for stock, traded in self.has_traded.items():
            traded_str = "[" + ",".join(f"{x:>3}" for x in traded) + "]"
            print(f"总计{sum(traded)}股{stock}买入列表{traded_str:<50}")

    def plan_volume(self, stock, day, sup) -> int:
        plan_volume = 0
        for day_item in range(1, day):
            plan_volume += self.calculate_stock_v2(stock, day_item, sup)
        return plan_volume

    def pre_total_volume(self, stock: int) -> int:
        return sum(self.has_traded[stock])

    def filter_full_day(self, current_day: dict) -> dict:
        volume = sum(current_day.values())
        left_trade = self.day_max
        if volume > self.day_max:
            sorted_data = dict(sorted(current_day.items(), key=lambda item: item[1], reverse=True))
            for stock, traded in sorted_data.items():
                print(stock, traded)
                if left_trade <= 0:
                    current_day[stock] = 0
                    continue
                if left_trade - traded < 0:
                    current_day[stock] = left_trade
                left_trade -= traded
        return current_day

    def full_day_order_by_stock(self, current_day, current_supp_day_new: dict) -> dict:
        volume = sum(current_day.values())
        left_trade = self.day_max
        if volume > self.day_max:
            sorted_data2 = dict(sorted(current_supp_day_new.items(), key=lambda item: item[1], reverse=True))
            sorted_data = current_day
            if sum(sorted_data2.values()) > 0:
                sorted_data_temp = {}
                for stock, traded in sorted_data2.items():
                    sorted_data_temp[stock] = sorted_data[stock]
                sorted_data = sorted_data_temp
            for stock, traded in sorted_data.items():
                if left_trade <= 0:
                    current_day[stock] = 0
                    continue
                if left_trade - traded < 0:
                    current_day[stock] = left_trade
                left_trade -= traded

        return current_day

    # 是否停牌
    def is_stop(self, day: int, suspension: list) -> bool:
        return day in suspension

    # 指定日期计划买入的数量
    def calculate_stock(self, stock, day: int, suspension: list) -> int:
        v = self.stockConfig[stock]['day_volume']
        if self.is_stop(day, suspension):
            v = 0
        # print(f"第{day}天，股票{stock},计划买入{v}")
        # 计算是否需要隔天处理
        if self.stockConfig[stock]['step'] == 0:
            return v
        # 如果是隔天买入，计算应该买入数量
        if day % self.stockConfig[stock]['step'] == self.stockConfig[stock]['remainder']:
            return v

        return 0

        # 指定日期计划买入的数量

    def calculate_stock_v2(self, stock, day: int, suspension: list) -> int:
        v = self.stockConfig[stock]['day_volume']
        # print(f"第{day}天，股票{stock},计划买入{v}")
        # 计算是否需要隔天处理
        if self.stockConfig[stock]['step'] == 0:
            return v
        # 如果是隔天买入，计算应该买入数量
        if day % self.stockConfig[stock]['step'] == self.stockConfig[stock]['remainder']:
            return v

        return 0

    # 补充买入
    def supplement_volume(self, stock: int, day: int, suspension: list, ) -> int:
        # 如果当天没有买入，则补充买入
        if self.is_stop(day, suspension):
            return 0

        # 如果是隔天买入，计算应该买入数量
        if self.stockConfig[stock]['step'] == 0 or day % self.stockConfig[stock]['step'] == self.stockConfig[stock]['remainder']:
            # stop_days = len([stop_day for stop_day in suspension if stop_day < day])
            stop_days = len(self.need_supp_days[stock])
            total = stop_days * self.stockConfig[stock]['day_volume'] * 0.3
            return int(total)
        return 0

    # 对比该日计划买入总配置，有超出则削减
    def compare_volume(self, list: []) -> list:
        return []

    # 校验当前票是否已不能在交易
    def is_max_volume(self, stock: int) -> int:
        if sum(self.has_traded[stock]) >= self.plan:
            return True
        return False


K = Solution()
# aa = K.filter_full_day({1: 160,
#                     2: 0,
#                     3: 390})
# aa = K.filter_full_day({1: 60,
#                     2: 60,
#                     3: 490})
# K.trading_plan(1000, [[1, 2], [3], [1]])
# 总计1计划买入列表[  0,  0,160,110,190,140,  0,130,130,130, 10]
# 总计2计划买入列表[200,  0,  0,  0,260,  0,110,  0,320,  0,110]
# 总计3计划买入列表[  0,  0,  0,390,  0,  0,390,  0,  0,220,  0]
# K.trading_plan(10000, [[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 2], [], []])
K.trading_plan(1, [])
# K.trading_plan(1000, [[], [], []])
# 总计1计划买入列表[100,100,100,100,100,100,100,100,100,100,  0]
# 总计2计划买入列表[200,  0,200,  0,200,  0, 90,  0,260,  0, 50]
# 总计3计划买入列表[200,  0,  0,390,  0,  0,310,  0,  0,100,  0]
# [
# [100,100,100,100,100,100,100,100,100,100,0],
# [200,0,200,0,200,0,90,0,260,0,50],
# [200,0,0,390,0,0,310,0,0,100,0]]

# K.trading_plan(1000, [[1,2,3,4,5], [], []])
