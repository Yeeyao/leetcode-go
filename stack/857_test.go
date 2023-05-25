package stack

/*
leetcode++ 介绍是使用 q/w 表示卷王指数，这个数值越大，表示干越多活，拿越少工资。因为要保证 k 个人中每个人至少需要有最低工资，因此这个数值必须是 k 人里面最小的
使用元组 eff[(q / w, q, w)] 保存，然后根据 -q/w 进行排序
因为需要 k 个工人，就从第一个开始到 k 进行计算，类似滑动窗口，每个窗口的最后一个的 q/w 就是 k 个里面最小的，计算所有窗口的总工资然后取最小值
有点暴力解法的味道

这里提到窗口的大小固定，因此就有类似固定堆的解题思路

```python
class Solution:
    def mincostToHireWorkers(self, quality: List[int], wage: List[int], K: int) -> float:
        eff = [(q / w, q, w) for a, b in zip(quality, wage)]
        eff.sort(key=lambda a: -a[0])
        ans = float('inf')
        for i in range(K-1, len(eff)):
            h = []
            k = K - 1
            rate, _, total = eff[i]
            # 找出工作效率比它高的 k 个人，这 k 个人的工资尽可能低。
            # 由于已经工作效率倒序排了，因此前面的都是比它高的，然后使用堆就可得到 k 个工资最低的。
            for j in range(i):
                heapq.heappush(h, eff[j][1] / rate)
            while k > 0:
                total += heapq.heappop(h)
                k -= 1
            ans = min(ans, total)
        return ans
*/

/*
雇佣 K 名工人的最低成本

There are n workers. You are given two integer arrays quality and wage where quality[i] is the quality of the ith worker and wage[i] is the minimum wage expectation for the ith worker.

We want to hire exactly k workers to form a paid group. To hire a group of k workers, we must pay them according to the following rules:

    Every worker in the paid group should be paid in the ratio of their quality compared to other workers in the paid group.
    Every worker in the paid group must be paid at least their minimum wage expectation.

Given the integer k, return the least amount of money needed to form a paid group satisfying the above conditions. Answers within 10-5 of the actual answer will be accepted.


N 名工人，第 i 个的工作质量是 quality[i]，最低期望工资是 wage[i]

现在我们想雇佣K名工人组成一个工资组。在雇佣一组 K 名工人时，我们必须按照下述规则向他们支付工资：

对工资组中的每名工人，应当按其工作质量与同组其他工人的工作质量的比例来支付工资。
工资组中的每名工人至少应当得到他们的最低期望工资。
返回组成一个满足上述条件的工资组至少需要多少钱。

这里要求按照工作质量与其他人的工作质量比例来支付工资，同时需要每个工人满足最低工资期望。
可以使用 w[i]/q[i] 表示单位质量所需要的工资，因为后一个条件需要满足，因此需要找到 w[i]/q[i] 最高的工人，这里只找出最大值就好了，同时，这里只需要找到
k 个工人里面的 w/q 的最大值，同时需要总工资最小，因此需要 q 尽量小

```


*/
