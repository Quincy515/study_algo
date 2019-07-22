# 1. 玩转图论算法

1. [图论知识点概要](#0tO6f)
2. [图论的应用](#rAaLD)

# <a name="0tO6f"></a>图论知识点概要

图论<br />                 ↓<br />    关于**图**的算法<br />     ↓<br />**图**是一种数据结构<br />
<br />![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563764341370-b1d10ab7-6783-41fc-8847-8e688de04393.png#align=left&display=inline&height=142&name=image.png&originHeight=480&originWidth=1510&size=167713&status=done&width=448)               ![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563764361157-361bd086-bfa9-4a3e-9b58-eaf4e65726d6.png#align=left&display=inline&height=142&name=image.png&originHeight=776&originWidth=926&size=189383&status=done&width=169)<br />
<br />图这种数据结构本身很简单，基于图这种数据结构的算法有难度。<br />
<br />**面试重点**<br />

- 图的表示
- 图的遍历
  - 深度优先：连通性、路径、二分图检测、环的检测、floodfill
  - 广度优先：无权图的最短路径
- 如何使用图论对问题建模

<br />**其他**

- 欧拉路径
- 哈密尔顿路径，状态压缩
- 桥
- 割点

<br />**图论经典算法**

- 有向图算法 DAG，环检测，拓扑排序，强连通分量
- 最小生成树 Kruskal，Prim
- 最短路径 Dijkstra，Floyed，Bellman-Ford
- 网络流 最大流-最小割，Ford-Fulkerson

# <a name="rAaLD"></a>图论的应用

- 地理路线(地图)
- 路由路线
- 互联网连接-Google PageRank，爬虫
- 论文引用
- 社交网络
- 规划
- 匹配问题
- 编译原理，作用域、引用依赖关系
- 其他算法的基石：状态机 - KMP 算法
- 对大脑建模

图论：离散数学的一个分支，偏算法，而非数学证明，必要的定理需要学习，但整体以**实现算法**为主。
