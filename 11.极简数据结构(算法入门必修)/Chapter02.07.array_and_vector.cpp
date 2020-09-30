#include <iostream>
#include <vector>

using namespace std;

int main()
{
    const int n = 42; // 定义常量
    const int c = 1;  // 避免出现魔数

    // 定义
    int A1[n];
    vector<int> v1(n);

    // 初始化
    for (int i = 0; i < n; i++)
    {
        A1[i] = c;
    }
    vector<int> v2(n, c);

    // 列表初始化
    vector<int> v3 = {1, 2, 3, 4, 5};
    for (size_t i = 0; i < v3.size(); i++)
    {
        cout << v3[i] << " ";
    }
    cout << endl; // 1 2 3 4 5

    // 动态变化
    v3.push_back(6);
    cout << v3.size() << endl;                      // 6
    cout << v3.front() << " " << v3.back() << endl; // 1 6
    v3.pop_back();
    cout << v3.size() << endl;                      // 5
    cout << v3.front() << " " << v3.back() << endl; // 1 5

    // 迭代器
    for (auto iter = v3.begin(); iter != v3.end(); iter++)
    {
        *iter = c;
        cout << *iter << " ";
    }
    cout << endl; // 1 1 1 1 1

    return 0;
}