# 题目描述：
# 盗贼面对n个宝箱，第i个宝箱需要time[i]时间打开，包含value[i]价值的宝物。
# 盗贼有两种技能：
# 1. 普通开锁：花费time[i]时间打开第i个宝箱，获得value[i]价值
# 2. 影分身：消耗50能量，可以同时打开连续的k个宝箱（k=2或3），时间为max(time[i:i+k])
#
# 盗贼总共有energy点能量，每单位时间消耗1点能量。
# 问：在能量耗尽前能获得的最大价值是多少？
#
#
# 示例：
# 输入：time=[3,5,2,4], value=[10,20,8,15], energy=100, k可选2或3
# 输出：最优策略下的最大价值




# 题目描述：
# 有n个敌人排成一排，第i个敌人的血量为hp[i]。魔法师有两种攻击方式：
# 1. 火球术：消耗3mp，对一个敌人造成2点伤害
# 2. 连锁闪电：消耗8mp，对连续的3个敌人分别造成3、4、3点伤害（从左到右）
#
# 敌人死亡后会立即消失，后面的敌人会向前补位。
# 问：击败所有敌人最少需要多少mp？
#
# 示例：
# 输入：hp = [4, 6, 5, 3]
# 输出：使用连锁闪电攻击位置1-3（造成3,4,3伤害）后，敌人血量变为[1,2,2,3]
# 再用火球术清理剩余血量，总共需要的mp。

def mincost(n):
    arr = [i for i in range(n + 1)]

    # arr = [0, 4, 6, 5, 2, 2, 9]

    def dp(x, arr) -> int:
        if x >= len(arr):
            return 0
        if arr[x] <= 0:
            return dp(x + 1, arr)

        mins = float('inf')
        tmparr = arr.copy()
        tmparr[x] -= 2
        mins = min(mins, 3 + dp(x, tmparr))
        tmparr = arr.copy()
        tmparr[x] -= 3
        if x + 1 < len(arr):
            tmparr[x + 1] -= 4
        if x + 2 < len(arr):
            tmparr[x + 2] -= 3
        mins = min(mins, 8 + dp(x, tmparr))
        return int(mins)

    return dp(1, arr)


for i in range(1, 10):
    print(f"目前有{10 - i}个怪物，最少花费{mincost(10 - i)}MP")
