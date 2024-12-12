from concurrent.futures import ThreadPoolExecutor, as_completed

import datetime
from time import sleep

groups = list(range(1, 1501))
max_workers = 15
batch_size = 200


def dynamic_indicator(name):
    print(f"计算+{name}")
    sleep(0.01)


def main():
    start_time = datetime.datetime.now()

    # 定义处理单个组的函数
    def process_group(rows):
        print(f"执行rows{rows}")
        for row in rows:
            dynamic_indicator(row)
        return

    if len(groups) > batch_size:
        # logger.info(f"开始批次执行200个一组进行,共计：{len(groups)}, {len(groups) / 200}")
        # 按批次分组
        batches = [groups[i:i + batch_size] for i in range(0, len(groups), batch_size)]

        with ThreadPoolExecutor(max_workers=max_workers) as executor:
            futures = [executor.submit(process_group, batch) for batch in batches]
            for future in as_completed(futures):
                future.result()  # 捕获可能的异常
    else:
        # 如果 group 数量少于 batch_size，则直接逐个处理
        for rows in groups:
            dynamic_indicator(rows)
    print(f"完成所有任务{datetime.datetime.now() - start_time}")
    return


if __name__ == '__main__':
    main()
