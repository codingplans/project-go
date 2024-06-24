import time
import random
import string

# 定义文件路径
file_path = 'random_data.txt'


# 定义写入文件的函数
def write_random_data(file_path):
    word = ["error", "warn","repo","date"]
    # 生成随机字符串
    random_data = ''.join(random.choices(string.ascii_letters + string.digits, k=10)) + word[random.randint(0,3)]
    # 打开文件，追加数据
    with open(file_path, 'a') as file:
        file.write(random_data + '\n')
    print(f"随机数据已写入：{random_data}")

# 定义定时任务
def timed_task(interval, file_path):
    while True:
        write_random_data(file_path)
        # 等待指定的时间间隔
        time.sleep(interval)


# 设置定时任务的时间间隔，单位为秒
interval_seconds = 3

# 启动定时任务
timed_task(interval_seconds, file_path)
