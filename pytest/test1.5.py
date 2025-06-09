import pandas as pd


def count_salary_categories(accounts: pd.DataFrame) -> pd.DataFrame:
    # 根据工资列创建三个工资类别
    salary_categories = ['Low Salary', 'Average Salary', 'High Salary']

    # 创建一个字典，用于存储每个工资类别的账户数量
    counts = {category: 0 for category in salary_categories}

    # 遍历每一行数据，根据收入确定工资类别并进行计数
    for index, row in accounts.iterrows():
        income = row['income']
        if income < 20000:
            counts['Low Salary'] += 1
        elif 20000 <= income <= 50000:
            counts['Average Salary'] += 1
        else:
            counts['High Salary'] += 1

    # 创建结果DataFrame
    result = pd.DataFrame(counts.items(), columns=['category', 'accounts_count'])


    return result


df = pd.DataFrame(
    {'income': [15000, 30000, 45000, 60000, 20000, 70000, 25000, 80000],
     'account_name': ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H']}
)
print(count_salary_categories(df))
