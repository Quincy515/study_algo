#include <iostream>

using namespace std;

int main()
{
    const int n = 64;
    int x = 4;
    int y = 2;
    int M[n][n]; // 常量 n 避免魔数

    // 方法一：判断是否在对角线上
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (i == j)
                M[i][j] = x;
            else
                M[i][j] = y;
            // M[i][j] == (i == j) ? x : y;
        }
    }

    // 方法二：整体到局部
    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            M[i][j] = y;

    for (int i = 0; i < n; i++)
        M[i][i] = x;

    // 方法三：循环展开
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < i; j++)
            M[i][j] = y;
        M[i][i] = x;
        for (int j = i + 1; j < n; j++)
            M[i][j] = y;
    }

    // 打印输出
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
            cout << M[i][j];
        cout << endl;
    }
    return 0;
}