from functools import lru_cache
from typing import List

#有几个怪物排成一排,第i个怪物的血量为i。小红有
# 两个技能可以打怪:
# 1.强力攻击：消耗1mp,对一只怪物造成1点伤害。
# 2.踏前斩：消耗5mp,对当前怪物造成1的伤害,同

# 时剑气将波及后两个怪物,对下一个怪物造成2点伤害,对下下个怪物造成3点伤害。如果一个怪物受伤后血量小于等于0,则怪物死亡。死亡后怪物的尸体依然占据一个位置,会被踏前斩的剑气打到。小红想知道,击杀全部怪物至少需要花费多少mp?
def min_mp_to_kill_monsters(n: int) -> int:
    # 初始化怪物血量
    hp = [i for i in range(n + 1)]  # hp[0] 无意义，hp[i] = i for i = 1 to n

    @lru_cache(None)
    def dp(i: int, hp_tuple: tuple) -> int:
        # 将hp数组转为元组以便记忆化
        hp = list(hp_tuple)
        # 如果已经处理到超出怪物数量，检查是否所有怪物都死亡
        if i > n:
            return 0 if all(h <= 0 for h in hp[1:]) else float('inf')

        # 如果当前怪物已死亡，跳到下一个
        if i <= n and hp[i] <= 0:
            return dp(i + 1, tuple(hp))

        # 初始化最小MP
        min_mp = float('inf')

        # 选择1：强力攻击
        hp_new = hp.copy()
        hp_new[i] -= 1
        min_mp = min(min_mp, 1 + dp(i, tuple(hp_new)))

        # 选择2：踏前斩
        hp_new = hp.copy()
        hp_new[i] -= 1
        if i + 1 <= n:
            hp_new[i + 1] -= 2
        if i + 2 <= n:
            hp_new[i + 2] -= 3
        min_mp = min(min_mp, 5 + dp(i, tuple(hp_new)))

        return min_mp

    # 初始血量元组
    initial_hp = tuple(hp)
    result = dp(1, initial_hp)
    return result


# 测试函数
def solve(n: int) -> int:
    return min_mp_to_kill_monsters(n)


# 示例测试
for n in range(1, 10):
    print(f"n={n}, min MP={solve(n)}")