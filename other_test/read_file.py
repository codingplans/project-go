import tailer

# 日志文件路径
log_file_path = 'random_data.txt'

# 定义一个函数来处理新行
def process_line(line):
    print(line.strip())

# 检查日志文件是否存在
try:
    with open(log_file_path, 'rb') as f:
        # 获取文件的当前大小，作为游标的初始值
        cursor = f.tell()
except FileNotFoundError:
    # 如果文件不存在，游标设为0，从头开始监听
    cursor = 0

print(cursor)
# 创建tailer对象，从指定的游标位置开始监听
t = tailer.Tailer(log_file_path, cursor=20, delay=1, use_polling=True)

# 监听文件，当有新内容时调用process_line函数
for line in t:
    process_line(line)