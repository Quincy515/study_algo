# 2. 图的表示

1. [图的分类](#JP2EU)
1. [图的基本概念](#HTLPH)
1. [图的基本表示：邻接矩阵](#6k8Gc)

<br />![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563764341370-b1d10ab7-6783-41fc-8847-8e688de04393.png#align=left&display=inline&height=142&name=image.png&originHeight=480&originWidth=1510&size=167713&status=done&width=448)               ![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563764361157-361bd086-bfa9-4a3e-9b58-eaf4e65726d6.png#align=left&display=inline&height=142&name=image.png&originHeight=776&originWidth=926&size=189383&status=done&width=169)

# <a name="JP2EU"></a>图的分类

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563766535453-6b53b546-7ca4-4139-9f66-97b319d8f40a.png#align=left&display=inline&height=277&name=image.png&originHeight=626&originWidth=756&size=130452&status=done&width=335)   ![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563766605906-53f02b52-0245-4902-81d9-3b1b485afd5b.png#align=left&display=inline&height=275&name=image.png&originHeight=620&originWidth=732&size=145741&status=done&width=325)

- 顶点 Vertex：0、1、2、3、4、5、6
- 边 Edge

无向图 Undirected Graph 顶点之间的边是没有方向的<br />有向图 Directed Graph

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563766686924-8e6010a8-dff8-478b-add2-538408a9d360.png#align=left&display=inline&height=254&name=image.png&originHeight=592&originWidth=768&size=134183&status=done&width=329)![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563766766865-4c30a2ca-bc31-43bc-9d31-0ac0ec0c5e2f.png#align=left&display=inline&height=259&name=image.png&originHeight=736&originWidth=852&size=198220&status=done&width=300)

地铁应该用无向图

社交网络 微信是好友关系是无向图，微博关注所以是有向图

无权图和有权图

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563766818720-d313e260-0085-49b8-9323-157005b95ed5.png#align=left&display=inline&height=237&name=image.png&originHeight=636&originWidth=766&size=155320&status=done&width=285)

所以可以分为四类\*\*

|      | 无向           | 有向           |
| ---- | -------------- | -------------- |
| 无权 | **无向无权图** | **有向无权图** |
| 有权 | **无向有权图** | **有向有权图** |

# <a name="HTLPH"></a>图的基本概念

**无向无权图**

- 两点相邻： 如果两个点之间有边，就是两点相邻，比如 0 和 1、0 和 3
- 点的邻边：和一个点所有相邻的边，就是点的邻边，01 所对应的边，和 03 所对应的边
- 路径 Path：从一点到达另一点所走过的路，0 到 5 的路径有 0165、0325、0345
- 环 Loop：从一个点经过若干条边，又回到这个点，01230
- 自环边： 0 属于自环边
- 平行边：3、4 属于平行边

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563767244934-219a5791-28ce-444b-ac3e-a9e8688406f8.png#align=left&display=inline&height=225&name=image.png&originHeight=608&originWidth=748&size=130135&status=done&width=277)                  ![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563767184623-3a16daa1-37e5-4dbd-a90d-0ac43ca107cc.png#align=left&display=inline&height=234&name=image.png&originHeight=670&originWidth=774&size=142249&status=done&width=270)<br />没有自环边，没有平行边的图，称为**简单图**，这里学习主要处理简单图

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563767629774-e12cc56e-c257-4965-963c-e19030c8930b.png#align=left&display=inline&height=209&name=image.png&originHeight=620&originWidth=950&size=157234&status=done&width=321)<br />一张图中相互连接可以相互抵达的顶点集合称为**联通分量**。上图 9 个顶点有 2 个联通分量 0 到 6 和 7 到 8

- 一个图的所有节点不一定全部相连
- 一个图可能有多个联通分量

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563767791451-aaaa45bc-0eca-4804-87c9-69e6f80b55fc.png#align=left&display=inline&height=229&name=image.png&originHeight=630&originWidth=760&size=129192&status=done&width=276)           ![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563768008795-f4546a54-481c-4daa-947f-f303a58a293f.png#align=left&display=inline&height=232&name=image.png&originHeight=624&originWidth=736&size=125552&status=done&width=274)

**无环图**：不能找到一条路径从某个顶点出发，绕一圈回到它的本身，**树是一种无环图**

数据结构中谈论的树，根节点是唯一的，而图中谈论的是树的样子，根节点是不唯一的。

一个图里无环但不一定是树，所以**无环图不一定是树**，如右上图，有两个联通分量的无环图就不是树。

所以**联通的无环图是树**<br />![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563768316935-38e09082-11e6-4fd2-8070-667409301cd8.png#align=left&display=inline&height=261&name=image.png&originHeight=724&originWidth=1748&size=336805&status=done&width=629)

左右两边的图顶点相同，区别在于边，右边的图的边是左边图的边的子集。

把左边图一些边删除，就可以得到右边的图，同时右边的图还是树的形状，所以这个过程叫做**联通图的生成树**

联通图的生成树，包含所有的顶点的树，边数最少要保证 V-1(V 个顶点)，如果大于 V-1 则包含环了就不是树了。

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563768455506-4e4d05d3-6164-42f6-9b4a-26288e640366.png#align=left&display=inline&height=269&name=image.png&originHeight=750&originWidth=1774&size=341917&status=done&width=636)<br />包含所有顶点，边数 V-1，一定是联通图的生成树吗？

**一个联通图一定有生成树，一个图一定有生成森林**

一个图不联通，可以先求出各个联通分量，每个联通分量部分的图，单独可以生成树。

一个顶点的度 degree，对于无向图，这个顶点相邻的边数，顶点 0 的度就是 2，顶点 2 的度就是 3。

# <a name="6k8Gc"></a>图的基本表示：邻接矩阵

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563769031884-30a437c2-6626-4105-8869-8341f975aa87.png#align=left&display=inline&height=317&name=image.png&originHeight=814&originWidth=1672&size=363441&status=done&width=652)

简单图没有自环边，所以主对角线为 0。简单图没有平行边，所以不会出现 `A[i][j] = 2`

对于无向图，矩阵关于主对角线对称，意味着 `A[0][3] = 1` ，所以 `A[3][0] = 1`

![image.png](https://cdn.nlark.com/yuque/0/2019/png/288708/1563771859568-16ff1b3d-3fa8-4ce1-be7a-77dac195cfc3.png#align=left&display=inline&height=256&name=image.png&originHeight=716&originWidth=1950&size=318172&status=done&width=696)

V 表示多少顶点，E 表示多少边，下面每两个顶点代表有一条边。

```txt
7 9
0 1
0 3
1 2
1 6
2 3
2 5
3 4
4 5
5 6
```

```go
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type AdjMatrix struct {
	V   int
	E   int
	adj [][]int
}

func NewAdjMatrix(name string) {
	result := make([]string, len(name))
	if contents, err := ioutil.ReadFile(name); err == nil {
		result = strings.Fields(string(contents))
	}
	V, _ := strconv.Atoi(result[0])
	adj := make([][]int, 0)
	for i := 0; i < V; i++ {
		tmp := make([]int, V)
		adj = append(adj, tmp)
	}

	E, _ := strconv.Atoi(result[1])
	for i := 0; i < 2*E; i++ {
		if i%2 == 0 {
			a, _ := strconv.Atoi(result[2+i])
			b, _ := strconv.Atoi(result[3+i])
			fmt.Printf("a: %d, b: %d\n", a, b)
			adj[a][b] = 1
			adj[b][a] = 1
		}
	}
	fmt.Printf("V: %d, E: %d, len: %d", V, E, len(result))

    /* 输出数组元素 */
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			fmt.Printf("adj[%d][%d] = %d\n", i, j, adj[i][j])
		}
	}
}

func main() {
	NewAdjMatrix("g.txt")
}
```
