# 数组题目总结

这里按照目前做过的题目，对每题，选择相关类型的进行总结

进度：20211021 18-1995-300 334/1534 330  20211101 354-1996-1691 646-491

TODO: 712-583-72

## 分级别

1-4

1: 完全不熟悉，看答案还是很难懂

2: 看答案没有完全懂，需要花时间弄懂，大概半天

3: 看答案，然后比较快速就能看懂和理解

4: 自己思考，但是花了比较多时间才能做出来，大概半天。看答案能看懂，需要一点思考

5: 自己完成或者完全懂了

## 小结

continue 语句也是挺耗时间的

遍历的 two pointers 从两边向中间或者从中间向两边

注意循环的终止条件以及初始化值的处理

可以使用 struct slice 来替代 map 无法顺序遍历的问题

if value, ok := m[a]; ok {...} else {...} map 元素的查找

即使是 brute force 也可以考虑根据情况的剪枝

整体有序的相关情况算法，可以考虑是否可以利用局部有序

int 转换为 string 利用 strconv.Itoa(i)

任何递归的本质，实际上就是入栈出栈的过程。也就是说只要是递归的，都可以改成非递归，因此快排也可以通过栈来实现。

已排序数组的搜索，一般可以考虑使用二分搜索

很多的数据和很少数据比较时，可以先将很少的数据进行 hash，然后遍历很大的数据来进行比较 将原来的数据转换为 hash 后的数据

字符串哈希 hash 函数的本质是扫描字符串过程中, 根据之前的结果, 当前位置,当前字符的值使用一个公式计算出当前结果.

可以使用枚举回溯法的，思考能不能转换为利用 DP

## 题目

### xSum

#### 1. 2Sum

问题：给定数组和一个 target 数，求数组中两个数下标，要求它们的和等于 target

Tags：数组，map，求和

解法：使用一个 map 保存每个元素数值(key)和下标(index)，每个元素先找 map 中是否存在该元素，存在则表示找到满足的和，返回两个 index

#### 167 2Sum with sorted array

问题：给定已排序数组和一个 target 数，求数组中两个数下标，要求它们的和等于 target

Tags：已排序数组，求和，二分查找，双指针

解法：直接使用二分查找，初始化左右两个元素是头尾两个元素，然后 left < right 如果和大于 target 右边指针向左，小于则左边指针向右

#### 653 Two Sum IV - Input is a BST

问题：给定 BST 和一个 target 数，判断是否存在两个节点的数值和等于 target

Tags：tree，求和, bst

解法：

- 直接将二叉搜索树利用 slice 保存然后二分查找
- 直接遍历节点然后利用 map，像普通数组那样处理

#### 15 3Sum

问题：给定数组和一个 target 数，求数组中三个数，要求它们的和等于 0

Tags： 数组，哈希，求和，预排序，二分查找

解法：

- 先排序，然后判断第一个元素如果大于 0 就直接返回，因为全都大于 0 了，数组长度小于 3 也直接返回
- 第一个循环遍历每个元素数值 e 直到倒数第三个，如果当前元素和上一个元素相等就跳过
    - temp = 0 - e 就是二分查找需要找到的两个元素的和，从当前元素的下一个作为 left，最后一个元素作为 right 二分查找
    - 二分查找终止条件是 left < right 找到就保存到结果，然后需要继续找
        - 这里是找到的判断中，需要过滤左右两边相同的元素

#### 16. 3Sum Closest

问题：给定数组和一个 target 数，求数组中三个数，要求它们的和最接近 target，返回它们的和

Tags： 数组，哈希，求和，预排序，二分查找

解法：

- 先排序，然后遍历每个元素，注意这里还是需要跳过相同的元素，然后二分查找
- 如果结果等于 target 就返回 target，否则更新最接近的和，这里记录最接近的和以及最小差值辅助判断

#### 18 4Sum

Level: 2

问题：给定数组和一个 target 数，求数组中四个数，要求它们的和等于 target

Tags：数组，哈希，求和，预排序，二分查找

解法：

- 先排序，第一个循环遍历直到倒数第四个元素，第二个循环从第一个的下一个开始遍历，直到倒数第三个元素
- 上述循环每次遍历都要将重复元素进行过滤
- temp = target - nums[i] - nums[j]，然后二分查找左右两个元素之和等于 temp 就保存到结果，同时需要过滤相同元素

#### 454 4Sum II

问题：给定 4 个长度一样的数组 nums1, nums2, nums3, nums4，求出下标 (i, j, k, l) 的组合数量，使得 nums1[i]+nums2[j]+nums3[k]+nums4[l] = 0 其中 0 <=
i, j, k, l < n

Tags：求和，多个数组元素

解法: 直接将前两个数组的所有元素之和放到 map 同时计数，然后遍历后面两个数组的所有元素，找到 0-n3-n4 的下标的 map 对应的和的数量加上总数

#### 1534. Count Good Triplets

Level: 2

问题：给定整型数组 arr 以及三个整型 a, b, c 需要找到 good triplets 的数量，定义 arr[i], arr[j], arr[k]...

Tags: 元素组合，元素大小关系

解法：维护一个 arr[i] 频次数组的前缀和 sum，对一个二元组 (j,k) 我们可以 O(1) 得到答案为 sum[r]-sum[l-1]。如何维护频次数组存的数下标符合 i < j 的限制， 只需要从小到大枚举 j 每次 j
移动指针 + 1 时将 arr[j] 的值更新到 sum 这样保证枚举到 j 的时候，sum 数组存的数值下标满足限制

#### 1995. Count Special Quadruplets

Level: 2

问题： 给定一个索引从 0 开始的数组 nums，返回数组中不同的四元组数量（a,b,c,d）满足 nums[a] + nums[b] + nums[c] == nums[d] a < b < c < d

Tags：元素组合，元素大小关系

解法：直接使用 hashMap，然后用 i 下标的元素作为第三个元素，先计算后面的 d 的元素和 i 组成的数值的 hashMap，然后计算前面的 a+b 的数值

#### 334 Increasing Triplet Subsequence

Level: 3

问题：给定整型数组 nums，返回是否存在索引的三元组 (i, j, k) 存在 i < j < k 同时 nums[i] < nums[j] < nums[k] 如果不存在则返回 false

Tags: 数组元素关系

解法：这里从左到右遍历一次，记录下每个元素左边的最小值，从右到左遍历一次，记录每个元素右边的最大值。计算完其中一个后，另一个可以在遍历的时候提前判断 如果当前元素的左边最小值 < 当前元素值 < 当前元素右边最大值，则返回 true

#### 300. Longest Increasing Subsequence

Level: 3

问题：给定整型数组 nums，返回最长的严格递增子序列的长度

Tags: 数组元素关系，子序列

解法：

1. dp[i] 表示以第 i 个元素结尾的最长子序列长度，dp[i] = max(dp[j]) + 1 dp[0] = 1 res := 1 对每个 i 找到它之前的可以插入的序列的 最大的 dp 数值 maxDp，然后 dp[i] =
   maxDp + 1 然后就更新结果
2. d[i] 表示长度为 i 的最长子序列的最后一个元素的数值，如果 nums[i] > d[len] 就直接加入，否则每次都更新这个最小值 dp[1] = nums[0], res := 1

#### 673. Number of Longest Increasing Subsequence

题目： 给定一个整型数组 nums，返回最长的递增子序列的数量，序列需要严格递增 300 的变种。这里是需要统计数量。

tags：数组元素关系，子序列

题解：TODO

1. 同样 dp[i] 表示以第 i 个元素结尾的最长子序列长度 cnt[i] 表示长度为 i 的子序列数量。这里如果长度相同，则 cnt[i] += cnt[j] 否则 cnt[i] = cnt[j]...

#### 674. Longest Continuous Increasing Subsequence

题目： 给定一个整型数组 nums，返回最长的连续递增子序列的长度

tags：数组元素关系，子序列

题解： 直接遍历数组，每次遍历初始化临时长度为 1 直接向后遍历，满足当前元素大于上一个元素的序列，停止则更新最大长度

#### 712. Minimum ASCII Delete Sum for Two Strings

Level: 1

题目：给定两个字符串 s1, s2，需要将两个字符串构造成相同的字符串，只能删除字母，求最小删除全部字母的 ASCII 总和

tags: 字符串，dp，元素关系

题解：TODO
递归的分析情况见题目。dp[i][j] 表示 s1[i:] s2[j:] 的字符串的删除字母的总和 ASCII，dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))

#### 354. Russian Doll Envelopes TODO:

Level: 2

问题：给定有长宽的信封列表，返回可以套娃的信封最大数量

Tags: 数组元素关系

解法： 
    根据当前高度寻找插入的位置，如果找到了且在原来的数组之内，说明高度不是最高的，因此只需要更新该位置的高度
    如果找不到或者超过了数组，表示当前是新的高度，只需要追加到数组。因为当前高度是最高的，同时，宽度因为是递增的，也是最高的
    注意，这里宽度相同的情况下，因为高度最大的先加入了数组，因此后面的会直接更新高度，而不会追加

#### 1996. The Number of Weak Characters in the Game

Level: 2

问题：给定特性数组，每个特性有攻击力和防御力

Tags: 数组元素关系

解法：
根据当前高度寻找插入的位置，如果找到了且在原来的数组之内，说明高度不是最高的，因此只需要更新该位置的高度
如果找不到或者超过了数组，表示当前是新的高度，只需要追加到数组。因为当前高度是最高的，同时，宽度因为是递增的，也是最高的
注意，这里宽度相同的情况下，因为高度最大的先加入了数组，因此后面的会直接更新高度，而不会追加

#### 1691. Maximum Height by Stacking Cuboids

Level: 2

问题：给定多个立方体的三维数据，可以将立方体旋转。获取可以堆叠的最高的立方体的高度总和

解法：类似 646，因为可以旋转，因此可以在立方体本身进行旋转，将最高的作为高度。类似 646 的遍历，遍历每个元素，然后从第一个元素开始，判断能否放到该元素下堆叠

#### 646. Maximum Length of Pair Chain TODO:

Level: 3

问题：给定序对列1691. Maximum Height by Stacking Cuboids 表，定义序对没有交集就可以组成链表，求组成链表的最大长度

Tags: 序对，元素关系

解法：直接按照序对的右边界进行排序，然后 cur 初始化为最小整型，遍历列表，如果 cur < p[0] ，cur 更新为 p[1]，长度递增 1

#### 491. Increasing Subsequences

Level: 2

问题：给定整型数组，求数组的所有递增子序列

Tags: 元素关系，排列

解法：

#### 560 Subarray Sum Equals K

问题：给定整型数组以及整型 k 找到所有连续子数组的和等于 k 的数量

Tags：数组，求和，哈希，前缀和

解法：

- 使用 pre 保存当前的累加和，count 保存结果，建立 map m 保存每个累加和的出现次数
- 遍历每个元素，先将当前元素累加到 pre，然后 count 加上 m[pre-k] 的数量，m[pre]++，最后返回 count

### Combination Sum

#### 39. Combination Sum

问题：给定一个元素唯一的数组，以及一个 target，返回一个包含所有的唯一的组合的 candidate 数组，数组元素的和等于 target，数组的元素可以出现多次， 相同的元素出现的次数不同则组合不同。

Tags: 排列，DFS，暴力

解法：先对数组进行排序，这样可以根据数组和提前返回。DFS 下去，每个元素都可以作为开始，然后向后添加元素，当和满足了就可以保存结果并提前终止递归(因为数组是已排序的)
这里因为同一个元素可能被使用多次，因此解法中的递归处理，元素的开始还是当前的索引

#### 40. Combination Sum II

问题：给定一个数组，以及一个 target，返回一个 candidate 数组(每个元素只能使用一次，同时不能存在相同的组合)，数组元素的和等于 target

Tags: 排列，DFS，暴力

解法：相比 39 这里给定的数组元素可能是重复的，组合不能重复，需要先排序，然后递归中只有上一个元素和当前元素数值不同才会继续递归，同时元素开始是下一个索引

#### 216. Combination Sum III

问题：给定一个总和是 n，使用数字数量是 k，可以选择的数字是 1 到 9，求所有的 k 个数字然后得到和 n 的排列

Tags: 排列，DFS

解法：可选的数字是 1-9，每个数字只能使用一次，一定要使用 k 个数字。这里和前面的也是类似的，需要判断数量，大小，同样从 1 开始求和

#### 377. Combination Sum IV

问题：给定一个数组 nums 以及数字 target，使用 nums 中的元素组成排列的一个数组，其中数组的元素总和是 target

Tags: 排列，DP

解法：这里 dp[i] 表示和为 i 的排列数量，dp[0] = 1，当 1 <= i <= target 时，如果存在一种排列，元素之和是 i，则该排列最后一个元素一定是数组 nums 中的元素，假设该元素是 num，则有 num <=
i，对元素之和等于 i - num 的每种排列，在最后加上 num 就可以组成元素之和是 i 的排列，因此计算 dp[i] 的时候， 应该计算所有 dp[i-num] 之和

dp[0] = 1 遍历 i 从 1 到 target，对每个 i

- 遍历数组 nums 中每个元素 num，当 num <= i 时，将 dp[i-num] 的值加到 dp[i]
  最终 dp[target] 就是答案

#### 77. Combinations

问题： 给定两个整数 n 和 k，返回在 1 到 n 返回的 k 个数字的所有组合

Tags: 排列，DFS

解法：直接和 39 的类似，就是需要注意一下边界情况，判断当前数字和后面的数字是否足够

### sorted array

#### 4 Median of Two Sorted Arrays

问题：给定两个已排序数组，找到两个数组的中间大小元素，偶数则返回两个数的平均值

Tags: 已排序数组

TODO:[ref](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/)

#### 33. Search in Rotated Sorted Array

问题：有一个已排序的数组，在某个位置数组进行了旋转，比如 1，2，3，4 变成了 3，4，1，2。现在给定这样的数组和一个待查找的数字，找到该数字的位置，找不到则返回 -1

Tags：数组，已排序，找元素

解法：这里基本思路应该是二分查找，因为是已排序的数组，但是难点在于，确定当前二分查找的中间节点的位置，是处于反转位置的左边还是右边来判断左右边界的更新 同时，这里不是严格的在左边或者在右边，这个算法里面有体现

#### 81 Search in Rotated Sorted Array II

问题：类似 33，只是元素是允许重复出现的

Tags：数组，已排序，找元素

解法：这里在 33 的基础上，加上了判断左右之前需要先过滤一下相同的元素

#### 153 寻找旋转排序数组中的最小值

154 类似 153 只是有重复的元素

问题：排序后的数组，将后面的放到前面来实现旋转，找到最小值

Tags：数组，已排序，找元素

解法：直接二分找

#### 34 Find First and Last Position of Element in Sorted Array

问题：给定升序排列的数组，找到给定 target 数字的开始和结束位置，如果找不到返回 [-1,-1]

Tags：已排序数组，查找

解法：给定已排序的数组，然后查找，二分查找先找到左边界，这里判断一下，nums[mid] == target 也让 right = mid。 然后先判断 nums[left] 是否等于 target，等于就向右边找右边界

##### 35 Search Insert Position

问题：

Tags：

解法：

#### 278 First Bad Version 374 Guess Number Higher or Lower

两个都是二分查找的一个变种，没什么好说的

### water

#### 11 Container With Most Water

问题：求给定非负整型数组表示的高度组成的容器的最大值

Tags：数组，最大值，面积 TODO:还要看题解

解法：

- 初始化左右两个指针，然后循环 left < right 先用两个指针对应元素的较小值作为高度，然后更新较大面积
- 然后循环判断左指针指向的高度小于当前高度就不断右移，右指针同理

#### 42 Trapping Rain Wate

问题：给定非负整数的数组表示每个坝的高度，求所有坝最终存储的水的总量

Tags：数组，栈，最大值，DP

解法：TODO 本题以及 407

#### 238 Product of Array Except Self

问题：给定 n 个整数组成的数组，返回数组对应的是除了当前元素的所有其他元素的乘积

Tag：数组，累乘，数组元素变化，辅助数组

解法：

- 新建两个辅助数组，分别计算从左边和从右边到当前元素的累计乘积，这里，第一个的最左边和第二个的最右边默认都是 1
- 最终，结果数组的每个元素结果是辅助数组对应位置的乘积

### remove elements

#### 26 Remove Duplicates from Sorted Array

问题：给定已排序数组，移除数组中重复的元素，不使用额外的空间

Tag：数组，已排序，移除重复元素，不使用额外空间

解法：

- 一个计数器统计所有重复出现的元素个数，遍历每个元素
- 如果下标大于 0 同时上一个元素等于当前元素计数器递增，否则保存到 nums[i - count]

#### 27 Remove Element

问题：给定数组，将特定的元素从数组中移除，不使用额外的空间

Tags：数组，移除重复元素，不使用额外空间

解法：使用一个计数器 count 统计该元素出现次数，然后如果当前元素等于目标元素就将计数器+1，否则 nums[i - count] = nums[i]

#### 203 Remove Linked List Elements

问题：给定一个链表，移除给定数值的节点，然后返回新的头节点

Tag： 链表，移除特定元素

解法：普通判断就行了，然后需要注意处理下一个节点的连接问题，表头和表尾可能需要注意

#### 237 Delete Node in a Linked List

问题：给定一个链表，删除指定的节点，这里只给了需要删除的节点

Tag： 链表，移除特定元素

解法：题目不清晰，这里做法是将所有元素的数值向前移动，然后最后一个设置为 nil

#### 80 Remove Duplicates from Sorted Array2

问题：给定已排序数组，移除数组中重复的元素，需要注意这里允许每个元素最多出现两次，不使用额外的空间

Tag：数组，已排序，移除重复元素，不使用额外空间

解法：

- 类似 26 只是需要针对每个元素设置一个 isTwice 布尔值记录当前元素是否出现了两次
- 使用 i 记录最终数组的元素个数，循环遍历每个元素，
    - 如果 i 小于 2 或者当前元素大于上上个元素，将当前元素保存到 i 位置，i 递增
    - 否则直接什么都不做，因为元素遍历顺序是递增的，所以 i 大于上上个元素就表示已经满足不会重复出现超过两次了

### 元素组合

#### 78 Subsets

问题：给定一个元素都是唯一的整型数组，返回其所有子集

Tags：数组，数组子集

解法：这里元素都是唯一的，直接回溯就完事了。需要注意，这里是先保存临时结果，然后循环回溯，这里还需要注意 slice 的使用

#### 90 Subsets II

问题：给定可能包含重复元素的整型集合 nums，返回其所有子集，不能包含重复子集

Tag：数组，元素集合，元素子集

解法：先排序，然后按照排列处理，helper 函数内，需要注意这里保存的问题

- 遍历每个元素作为开头，然后每次先判断元素是否和上一个相同，相同则跳过，不重复开头元素
- 然后，重复元素自身的开头是需要保存的，所以添加判断是 开头或者不同(i != i-1)就需要递归

```go
*res = append(*res, temp)
*res = append(*res, append([]int{}, temp...))

// 注意这里的问题，第一个语句 temp 将会一直使用同一个 temp，同时没有触发扩容
```

### 元素排列

#### 31. Next Permutation

问题：给定整型数组，返回它的下一个排列，即下一个排列是第一个大于当前整型数组的排列

Tag：数组，排列，元素关系

解法：

- 从后面向前面找到 (i,j) 使得 nums[i] < nums[j] 就是需要交换的位置，同时 j 到最后的数组元素是降序的
- 从后面向前找到第一个元素 nums[k] 使得 nums[k] > nums[i]
- 将 nums[k] 和 nums[i] 进行交换
- 将 j 之后的所有元素升序排列，因为之前是降序的，所以这里直接反转就行
- 如果 2 中找不到，直接全部元素进行反转

#### 46. Permutations

问题：给定唯一元素集合，返回它们的所有排列

Tags：数组，排列，暴力

解法：

- 直接就是将元素随机打乱的思路，对每个位置，该位置之后的所有元素都和该位置交换
- 然后递归下去，下一次递归是当前的下一个位置作为和后面所有元素交换的位置
- 最后当位置等于数组长度就返回，需要注意每次递归需要还原
- *res = append(*res, append([]int{}, nums...)) 这样才保存到中间结果 TODO: 这里函数参数相关知识

#### 47. Permutations II

问题：给定元素集合，返回它们的所有排列，这里元素非唯一

Tags：数组，排列，暴力

解法：

- 先将元素进行排序，然后建立 visited 数组保存每个输入数组索引是否被使用，另一个临时结果数组
- 辅助函数参数是当前处理的索引，从 0 开始
    - 如果当前的位置等于输入数组长度，说明每个元素都找到位置，保存当前结果
    - 遍历每个元素
        - 如果当前元素索引已经被访问过，或者当前元素等于上个元素以及上个元素没有被访问过，直接跳过这次遍历
        - 否则将当前元素保存到临时结果数组，然后记录当前元素的索引已经访问，递归调用辅助函数，这里参数 + 1 然后两个数组还原

#### 60. Permutation Sequence

问题：给定 n 和 k，找到从 1 到 n 组成的数组中第 k 个排列，排列为字典序的

Tags：数组，排列，数学

解法：太复杂，直接看题目

### 合并已排序数据结构

#### 23 合并 k 个升序链表

问题：给定 k 个升序链表，将它们合并为一个

Tags：升序链表，合并链表

解法：暴力是将全部直接合并到一个链表。这里采用归并的方法来合并，先两个小的合并，然后继续两两合并

#### 88 合并已排序数组 见 link_list

问题：给定两个已排序数组 num1, num2 将两个数组按照大小升序合并到 num1，空间足够

Tags：数组，已排序，合并

解法：

- 因为是合并到其中一个，所以直接两个数组都从尾部向前遍历判断，将较大的元素保存到 num1 尾部
- 其中一个数组为空就直接将另一个复制到剩余位置上

#### 21 合并两个已排序列表 这里是链表形式

问题：合并两个已排序列表，返回一个新的已排序列表（链表）

Tags：链表，已排序，合并

解法：

- 直接新建一个 Head，然后使用两个指针遍历两个链表，将较小的放到 Head 后面并向前移动
- 如果其中一个已经遍历完就直接放到当前的 Head 后面

#### 977 已排序数组的平方

问题：给定一个升序的已排序数组，返回每个元素的平方组成的升序数组

Tags：数组，已排序，数组元素变化

解法：

- 直接思路是找到两个指针，指针指向非负数最小值和负数最大值，然后不断比较存放到结果数组并移动，同时需要判断是否遍历完
- 所以这里先遍历一次找到第一个非负数的位置，这里作为非负数最小值起点，该位置前一个位置作为负数最大值起点

#### 986 两个区间列表的交集区间列表

问题：给定两个区间列表，列表中区间不存在交集。找到这两个区间列表的交集列表

Tags：数组，区间，已排序

解法：

- 求完一个交集区间，较早结束的子区间，不可能再和其他子区间有重叠，它的指针要移动
- 较长的子区间还可能和别人重叠，它的指针暂时不动
- 时间和空间复杂度都是 O(M + N)

### DP

### 375 Guess Number Higher or Lower II

问题：猜数字，猜错需要给钱，求最少的钱

Tags：minMax，DP

算法：这里主要是推导出公式，然后根据公式编写代码，这里需要根据公式，确定有三个循环

### 464 Can I Win

问题：两个人选择数字加到一个公共和，某人选择数字加上去，超过或者等于某个数值就算赢，求先手者是否可以赢。两者都使用最优策略

Tags: DP

### 486 Predict the Winner

问题：P1，P2 从一个数组两端每次选择一个数字添加到自己的总和（初始为 0），P1 先选择，数组为空就停止。判断 P1 的总和是否大于等于 P2

Tags: minimax game, DP

### 查找第 k 个

#### 378 已排序 matrix 中查找第 k 个最小的元素

问题: 有 n * n 矩阵，每行和每列的元素都升序排列，找到第 k 个最小的元素

Tags: k, matrix, 已排序

解法: 这里从左下角到右上角，左边部分的总和都小于一个数字，因此可以使用二分方法来计算小于给定数值的元素的数量，然后和 k 比较。也可以使用优先队列方法

#### 373

问题: 给定两个已排序数组，分别从两个数组选择一个元素组成序对得到序对的和，找到第 k 个最小的和的序对

Tags: k, 已排序，序对

解法: 直接优先队列，需要注意，这里最开始将 nums1[i], nums2[0] i 从 [0,len(nums1)]，放入优先队列，然后选择元素出队列，下一个应该放入的元素是 top 的 nums1[i], nums2[j + 1]
，其中 i, j 都是 top 的元素

TODO: 为什么不能用二分？这里二分也需要将满足大小的元素先保存，然后再取 k 个，因为相同大小的元素数量有多个

#### 668

问题：对一个乘法表格，给定 m，n 表示 行列数量，返回第 k 个最小的表格元素

Tags: k, 已排序，乘积

解法：这里和 378 一样

#### 719.Find K-th Smallest Pair Distance

问题：给定一个数组以及 k，找到数组元素的距离中第 k 个最小的距离

Tag: 数组，排序，序对，第 k 个，距离

解法：先排序，直接二分，计算距离的上下界，然后计算小于等于给定距离的序对数量来二分

#### 658

问题：给定一个已排序数组以及 k 和 x，找到数组元素的距离 x 最小的前 k 个元素数组

Tag: 数组，排序，距离，第 k 个

解法：根据 x 和元素边界元素的关系，分为三种，第一种是 x 在左边界之外，则直接返回左边的前 k 个，第二种 x 在右边界之外，返回最后 k 个，第三种是 在中间，则需要找到 x
应该插入的位置，然后从这个位置向两边将距离较小的加入到结果数组，需要考虑两边到达边界的情况

#### 786

问题：对一个乘法表格，给定 m，n 表示 行列数量，返回第 k 个最小的表格元素

Tags: k, 已排序，分数

解法：类似 373 的两个数组，这里是同一个数组里面选择元素组成所求的数字，好像只有优先队列的做法，因为没法直接比较两个分数的大小，只能确定了第一个分数的分子后，选择下一个更大/小的