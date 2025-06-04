import concurrent.futures

nums = list(range(1, 101))  # 假设要处理 100 个数
all_results = []

def f(x):
    return x * x


def main():
    global all_results
    
    with concurrent.futures.ProcessPoolExecutor(max_workers=20) as executor:
        for _ in range(10):  # 执行 10 次
            results = executor.map(f, nums)
            print(results)
        
    
    for result in all_results:
        print(f"{result}")


if __name__ == '__main__':
    main()

