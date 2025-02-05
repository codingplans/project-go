#include <iostream>
#include <memory>

// 这段代码展示了三种智能指针的基本使用
using namespace std;

int main() {
    // shared_ptr
    auto sp = make_shared<int>(42);
    cout << *sp << endl;

    // unique_ptr
    unique_ptr<int> up = make_unique<int>(100);
    cout << *up << endl;

    // weak_ptr
    weak_ptr<int> wp = sp;
    if (auto tmp = wp.lock()) {
        cout << *tmp << endl;
    }

    return 0;
}

/*
  关键点：

  shared_ptr允许多个指针共享同一对象
  unique_ptr保证对象只有一个所有者
  weak_ptr通过lock()安全访问shared_ptr管理的对象，避免悬空引用*/