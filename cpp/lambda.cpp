#include <iostream>
#include <memory>

//// 这段代码展示了三种智能指针的基本使用
//using namespace std;
//
//auto sum = [](int a, int b) { return a + b; };
//int result = sum(3, 4); // result = 7
//
//int main() {
//    int multiplier = 10;
//    auto multiply = [multiplier](int x) { return x * multiplier; };
//    int result = multiply(5); // result = 50
//    return 0;
//}

//g++ -std=c++14 -g lambda.cpp -o test

#include <algorithm>
#include <vector>
#include <string>

struct Employee {
    int id;
    std::string name;
    double salary;
};

// 泛型工厂函数：生成按任意成员排序的 Lambda
auto createComparator = [](auto member) {
    return [member](const auto& a, const auto& b) {
        return a.*member < b.*member; // 成员指针语法
    };
};

int main() {
    std::vector<Employee> staff{
        {3, "Alice", 85000.0},
        {1, "Bob", 92000.5},
        {2, "Charlie", 78000.0}
    };

    // 按 ID 排序
    std::sort(staff.begin(), staff.end(), createComparator(&Employee::id));
    printf("按 ID 排序 ID\tName\tSalary\n");
    for (const auto& emp : staff) {
        printf("%d\t%s\t%.1f\n", emp.id, emp.name.c_str(), emp.salary);
    }
    // 按薪水逆序排序（Lambda 嵌套）
    std::sort(staff.begin(), staff.end(),
        [](const auto& a, const auto& b) {
            return a.salary < b.salary; // 逆序排列
        });
    printf("按薪水逆序排序 ID\tName\tSalary\n");
    for (const auto& emp : staff) {
        printf("%d\t%s\t%.1f\n", emp.id, emp.name.c_str(), emp.salary);
    }

    int multiplier = 10;
    auto multiply = [multiplier](int x) { return x * multiplier; };
    int result = multiply(5); // result = 50
    printf("result = %d\n", result);
}