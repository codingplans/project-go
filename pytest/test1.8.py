# 空问限制:262MB

# 问题描述
# 您正在帮助一家金融公司分析银行之间的转账费用。系统中有多个银行,每两个银行之间可能存在转账费用。您的任务是找到从指
# 的起始银行到目标银行的最小总转账费用。
# 输入
# start_bank:字符串,表示起始银行的名称。
# end_bank:字符串,表示目标银行的名称。
# fee_graph_from:字符串数组,表示每条转账路线的源银行。
# fee_graph_to:字符串数组,表示每条转账路线的目的银行。
# fee_graph_fee:浮点数数组,表示对应转账路线的费用,
# 每个索引i对应于从fee_graph_from[i] 到fee_graph_to[i]的一条转账路线,费用为fee_graph_fee[i]
# 输出
# 一个浮点数,表示从起始银行到目标银行的最小总转账费用。如果两银行之间没有路径,返回-1.0.
# 示例
# 输入:
# start_bank = "BankA";
# end_bank = "BankD";
# fee_graph_from = {"BankA", "BankA", "BankB", "BankB", "BankB", "BankC", "BankC", "BankC", "BankD", "BankD"};
# fee_graph_to = {"BankB", "BankC", "BankA", "BankC", "BankD","BankA", "BankB", "BankD", "BankB", "BankC"};
# fee_graph_fee = {5.0. 10.0, 5.0, 3.0, 10.0, 3.0,2.0.7.0.2.0}:
# 要找到从BankA到BankD的最小总转账费用,一条可能的路径是BankA->BankB ->BankC ->BankD,总费用为10.0.
# 输出:
# 10.0
# 约束条件
# fee_graph_from、fee_graph_to 和 fee_graph_fee 至少包含一条转账路线
# 转账费率图是连通的,从start_bank到end_bank一定有一条通路。
# 转账费用为正浮点数,
# 银行名称是唯一的字符串。
# 1示例1
# 收起
# 输入
# "BanKA","BankD",("BankA", "BankA", "BankB", "BankB", "BankC""BankC", "BankC", "BankD", "BankD"J.l"BankB", "BankC","BankA"
# "BankC", "BankD", "BankA", "BankB", "BankD", "BankB", "BankCCJ.15.0, 10.0, 5.0, 7.0, 10.0, 3.0, 2.0, 7.0,2.0
# 输出
# 10.00000

import heapq
from collections import defaultdict
from typing import List, Dict, Tuple


def find_min_transfer_cost(start_bank: str, end_bank: str,
                           fee_graph_from: List[str],
                           fee_graph_to: List[str],
                           fee_graph_fee: List[float]) -> float:
    """
    使用Dijkstra算法找到从起始银行到目标银行的最小转账费用

    Args:
        start_bank: 起始银行名称
        end_bank: 目标银行名称
        fee_graph_from: 转账路线源银行列表
        fee_graph_to: 转账路线目标银行列表
        fee_graph_fee: 对应转账费用列表

    Returns:
        最小转账费用，如果无路径则返回-1.0
    """
    # 构建邻接表表示的有向加权图
    graph = defaultdict(list)

    # 添加所有转账路线到图中
    for i in range(len(fee_graph_from)):
        from_bank = fee_graph_from[i]
        to_bank = fee_graph_to[i]
        cost = fee_graph_fee[i]
        graph[from_bank].append((to_bank, cost))

    # 如果起始银行不在图中，返回-1
    if start_bank not in graph:
        return -1.0

    # Dijkstra算法实现
    # 使用优先队列存储 (费用, 银行名称)
    priority_queue = [(0.0, start_bank)]
    # 记录到每个银行的最小费用
    distances = {start_bank: 0.0}
    # 记录已访问的银行
    visited = set()

    while priority_queue:
        current_cost, current_bank = heapq.heappop(priority_queue)

        # 如果已经访问过这个银行，跳过
        if current_bank in visited:
            continue

        # 标记为已访问
        visited.add(current_bank)

        # 如果到达目标银行，返回费用
        if current_bank == end_bank:
            return current_cost

        # 遍历当前银行的所有转账路线
        for next_bank, transfer_cost in graph[current_bank]:
            new_cost = current_cost + transfer_cost

            # 如果找到更短的路径，更新距离并加入队列
            if next_bank not in distances or new_cost < distances[next_bank]:
                distances[next_bank] = new_cost
                heapq.heappush(priority_queue, (new_cost, next_bank))

    # 如果无法到达目标银行
    return -1.0


def print_graph_analysis(fee_graph_from: List[str],
                         fee_graph_to: List[str],
                         fee_graph_fee: List[float]):
    """打印图的结构分析"""
    print("=== 转账网络图结构 ===")
    graph = defaultdict(list)

    for i in range(len(fee_graph_from)):
        from_bank = fee_graph_from[i]
        to_bank = fee_graph_to[i]
        cost = fee_graph_fee[i]
        graph[from_bank].append((to_bank, cost))

    for from_bank in sorted(graph.keys()):
        connections = []
        for to_bank, cost in graph[from_bank]:
            connections.append(f"{to_bank}({cost})")
        print(f"{from_bank} -> {', '.join(connections)}")


def find_all_paths(start_bank: str, end_bank: str,
                   fee_graph_from: List[str],
                   fee_graph_to: List[str],
                   fee_graph_fee: List[float]) -> List[Tuple[List[str], float]]:
    """找到所有可能的路径（用于调试和分析）"""
    graph = defaultdict(list)

    for i in range(len(fee_graph_from)):
        from_bank = fee_graph_from[i]
        to_bank = fee_graph_to[i]
        cost = fee_graph_fee[i]
        graph[from_bank].append((to_bank, cost))

    all_paths = []

    def dfs(current_bank, path, total_cost, visited):
        if current_bank == end_bank:
            all_paths.append((path.copy(), total_cost))
            return

        if len(path) > 10:  # 防止无限循环
            return

        for next_bank, cost in graph[current_bank]:
            if next_bank not in visited:
                path.append(next_bank)
                visited.add(next_bank)
                dfs(next_bank, path, total_cost + cost, visited)
                path.pop()
                visited.remove(next_bank)

    visited = {start_bank}
    dfs(start_bank, [start_bank], 0.0, visited)

    return sorted(all_paths, key=lambda x: x[1])


def main():
    print("=== 银行转账最小费用计算 ===\n")

    priority_queue = [(0.0, 222)]

    heapq.heappush(priority_queue, (2.2, 111))
    heapq.heappush(priority_queue, (2.3, 112))
    heapq.heappush(priority_queue, (2.4, 113))
    heapq.heappush(priority_queue, (2.5, 114))

    print(priority_queue)
    current_cost, current_bank = heapq.heappop(priority_queue)
    print(current_cost, current_bank)
    print(priority_queue)

    current_cost, current_bank = heapq.heappop(priority_queue)
    print(current_cost, current_bank)
    print(priority_queue)



    # 测试用例1：原始示例数据
    print("测试用例1:")
    start_bank = "BankA"
    end_bank = "BankD"
    fee_graph_from = ["BankA", "BankA", "BankB", "BankB", "BankB", "BankC", "BankC", "BankC", "BankD", "BankD"]
    fee_graph_to = ["BankB", "BankC", "BankA", "BankC", "BankD", "BankA", "BankB", "BankD", "BankB", "BankC"]
    fee_graph_fee = [5.0, 10.0, 5.0, 3.0, 10.0, 3.0, 2.0, 7.0, 2.0, 9.0]

    result = find_min_transfer_cost(start_bank, end_bank, fee_graph_from, fee_graph_to, fee_graph_fee)
    print(f"从 {start_bank} 到 {end_bank} 的最小转账费用: {result:.5f}")

    print_graph_analysis(fee_graph_from, fee_graph_to, fee_graph_fee)

    # 分析所有可能的路径
    print("\n=== 所有可能路径分析 ===")
    all_paths = find_all_paths(start_bank, end_bank, fee_graph_from, fee_graph_to, fee_graph_fee)
    for i, (path, cost) in enumerate(all_paths[:10]):  # 只显示前10条路径
        path_str = " -> ".join(path)
        print(f"路径{i + 1}: {path_str}, 费用: {cost}")

    print(f"\n最优路径费用: {all_paths[0][1] if all_paths else '无路径'}")

    # 测试用例2：简化版本，验证算法正确性
    print("\n" + "=" * 50)
    print("测试用例2 - 简化验证:")
    start_bank2 = "A"
    end_bank2 = "D"
    fee_graph_from2 = ["A", "A", "B", "C"]
    fee_graph_to2 = ["B", "C", "D", "D"]
    fee_graph_fee2 = [5.0, 10.0, 8.0, 3.0]

    result2 = find_min_transfer_cost(start_bank2, end_bank2, fee_graph_from2, fee_graph_to2, fee_graph_fee2)
    print(f"从 {start_bank2} 到 {end_bank2} 的最小转账费用: {result2:.5f}")

    print_graph_analysis(fee_graph_from2, fee_graph_to2, fee_graph_fee2)

    print("\n可能路径:")
    print("A -> B -> D: 5.0 + 8.0 = 13.0")
    print("A -> C -> D: 10.0 + 3.0 = 13.0")
    print("最优费用应该是: 13.0")


if __name__ == "__main__":
    main()