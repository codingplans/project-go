
#include <iostream>
#include <cstring>

class Person {
    char* name;
    int age;
public:
    Person(const char* n, int a) {
        name = new char[strlen(n) + 1];  // 堆内存分配
        strcpy(name, n);
        age = a;
    }
    ~Person() {
        delete[] name;  // 释放堆内存
    }
};

void example() {
    int x = 10;              // 栈内存
    Person* p = new Person("Tom", 20);  // p在栈上,对象在堆上

    int array[5];           // 栈上的数组
    int* dynArray = new int[5];  // 堆上的数组

    delete p;              // 释放堆内存
    delete[] dynArray;     // 释放数组
}  // 函数结束时栈内存自动释放

int main(){
    example();
    printf("main函数结束\n");
    return 0;
}

//int main() {
//    example();
//    return 0;
//}