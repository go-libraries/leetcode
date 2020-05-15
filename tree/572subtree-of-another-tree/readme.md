# 子树

给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

         3
        / \
       4   5
      / \
     1   2
给定的树 t：

       4 
      / \
     1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

         3
        / \
       4   5
      / \
     1   2
        /
       0
给定的树 t：

       4
      / \
     1   2
返回 false。

# 暴力破解 DFS
DFS 枚举 ss 中的每一个节点，判断这个点的子树是否和 tt 相等。

如何判断一个节点的子树是否和 tt 相等呢，
我们又需要做一次 DFS 来检查，
即让两个指针一开始先指向该节点和 tt 的根，
然后「同步移动」两根指针来「同步遍历」这两棵树，判断对应位置是否相等。

O(|s|x|t|)
O(max{ds,dt})

# DFS 序列上做串匹配

这个方法需要我们先了解一个「小套路」：一棵子树上的点在 DFS 序列（即先序遍历）中是连续的。了解了这个「小套路」之后，我们可以确定解决这个问题的方向就是：把 ss 和 tt 先转换成 DFS 序，然后看 tt 的 DFS 序是否是 ss 的 DFS 序的「子串」。

这样做正确吗？ 假设 ss 由两个点组成，11 是根，22 是 11 的左孩子；tt 也由两个点组成，11 是根，22 是 11 的右孩子。这样一来 ss 和 tt 的 DFS 序相同，可是 tt 并不是 ss 的某一棵子树。由此可见「ss 的 DFS 序包含 tt 的 DFS 序」是「tt 是 ss 子树」的 必要不充分条件，所以单纯这样做是不正确的。

为了解决这个问题，我们可以引入两个空值 lNull 和 rNull，当一个节点的左孩子或者右孩子为空的时候，就插入这两个空值，这样 DFS 序列就唯一对应一棵树。处理完之后，就可以通过判断 「ss 的 DFS 序包含 tt 的 DFS 序」来判断答案。


# hash

考虑把每个子树都映射成一个唯一的数，如果 tt 对应的数字和 ss 中任意一个子树映射的数字相等，则 tt 是 ss 的某一棵子树。如何映射呢？我们可以定义这样的哈希函数：

f_o = v_o + 31 . f_l . p(s_l) + 179 . f_r . p(s_r)


这里 f_xf ​表示节点 xx 的哈希值，s_xs 表示节点 xx 对应的子树大小，v_xv 代表节点 xx 的 val，p(n)p(n) 表示第 nn 个素数，oo 表示当前节点，ll 和 rr 分别表示左右孩子。这个式子的意思是：当前节点 oo 的哈希值等于这个点的 val 加上 3131 倍左子树的哈希值乘以第 s_ls 个素数，再加上 179179 倍右子树的哈希值乘以第 s_rs 个素数。这里的 3131 和 179179 这两个数字只是为了区分左右子树，你可以自己选择你喜欢的权值。

这样做为什么可行呢？ 回到我们的初衷，我们希望把每个子树都映射成一个唯一的数，这样真的能够确保唯一吗？实际上未必。但是我们在这个哈希函数中考虑到每个点的 val、子树哈希值、子树大小以及左右子树的不同权值，所以这些因素共同影响一个点的哈希值，所以出现冲突的几率较小，一般我们可以忽略。当然你也可以设计你自己的哈希函数，只要考虑到这些因素，就可以把冲突的可能性设计得比较小。可是如果还是出现了冲突怎么办呢？ 我们可以设计两个哈希函数 f_1f 
​	
当然，为了减少冲突，你也可以设计「三哈希」、「四哈希」等，可是这样编程的复杂度就会增加。实际上，一般情况下，只要运气不是太差，一个哈希函数就足够了。

我们可以用「埃氏筛法」或者「欧拉筛法」求出前 argπ(max{∣s∣,∣t∣}) 个素数（其中π(x) 表示 xx 以内素数个数，argπ(x) 为它的反函数，表示有多少以内包含 xx 个素数，这个映射是不唯一的，我们取最小值），然后 DFS 计算哈希值，最后比较 ss 的所有子树是否有和 tt 相同的哈希值即可。
