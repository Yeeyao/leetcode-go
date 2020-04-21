# 数组题目总结

- continue语句也是挺耗时间的

- 遍历的 two pointers 从两边向中间或者从中间向两边

- 注意循环的终止条件以及初始化值的处理

## 还需要理解

- 1329

## 部分

### 排序

#### 一个待处理数组，一个参考数组，统计前者的元素数量，然后根据特征提取元素

- 1122

- 特征数组处理，统计元素数量后合并产生结果数组

- 这里是先统计数量，再根据另一个数组的顺序来将第一个数组的元素保存

- 即先统计所有出现的元素的数量，之后，分两次将元素拿出来

#### 元素数值特征排序

- 905 奇数，偶数排序

#### 原数组是升序，需要对每个元素进行计算乘方再排序

##### 977

- 已排序数组的平方排序数组

1. 可以使用上面的归并方法然后再把结果数组翻转一下，类似 88

2. 因为有正负数一起排序，所以从中间找两个指针来往两边遍历处理，其中第一个指针可以是第一个大于等于0的元素，
第二个指针是第一个指针的上一个元素

#### 乱序数组找最小差值

- 1200 先排序，之后按顺序，一遍统计最小差值，同时将等于该值的序对保存

#### 判断数组是否升序或者降序 可以使用两个值记录

### 矩阵变换特征处理

- 1252 元素数值变换

- 832 元素位置交换

- 766 元素是否关于斜角对称判断，观察下标关系来比较

- 566 矩阵进行形状变换，计算出元素总数后，只需要通过列数来计算前后矩阵的行列

- 1260 矩阵元素循环移动，观察元素下标特征进行变换 类似 566 的思路

### 字符串字母统计

- 1002 公共字母统计，分别统计每个单词的字母出现次数，与公共字母出现次数比较 这里只需要用字母来作为数组下标

### 元素变换相关

#### 针对一个元素进行数值处理，之后再求和，需要保留每次处理完的数组。这里思路直接先求和，然后再进行元素处理后再求和

- 985

#### 计算子数组的和 需要得到特定的和 这里统计出现的次数

- 1013

#### 求和，然后判断使得两边可以相等的位置，这里不要直接暴力来处理，主要是边界情况的处理很麻烦

- 所以，直接让一边求总和，然后遍历元素后，已经求和的一边减去当前元素值，未求和的加上当前元素值，相等时得到结果

#### 求总体的和，然后对特定元素进行加减处理

- 643 not done

#### 递归移动，通过翻转数组来处理

- 189

### 数组特定元素统计

#### 统计特定元素的出现次数，然后将后面的元素都向前移动，最后在结尾补齐

- 283 26  

#### 只需要统计最大的元素，出现元素则将计数 +1，变成 0 则需要重新统计

- 169

#### 找出不存在的元素，将出现的元素取相反数，同时需要注意下标的对应关系以及取相反数后的判断。最后，统计所有非负数

- 448 1h

- 442 类似

#### 找出一个缺失的元素，直接求和计算两个和的差值。因为只需要找一个，所以直接联想到求和

- 268

#### 序列的组成处理，元素组成数组，判断最后的元素

- 717 30 min

#### 找每个元素的最大重复序列的起始和结束位置，找特征元素的遍历处理

- 830

#### 最大连续升序数列长度

- 674 类似 830 的思路，每一趟遍历统计一下长度，然后更新当前最大长度

#### 计算整除的元素 这里利用取余的性质，避免了计算溢出

- 1018

#### 点是否同一直线问题

- 1232 斜率的计算，使用乘法比较好

#### 找到两个数的和可以整除 60 结合题目 1 two sum 以及 1018

- 不要想着排序后再处理

1. 同时因为最大值问题以及余数的性质，可以将所有可能的数值都可以直接保存到同一个数组位置

    - 对每个元素，找到其对应的差值的数量，然后加到总数量上，最后，自己的对应数量递增

2. 直接全部统计好每个数值出现的次数，然后计算序对的数量，0 和 30 可以自己和自己配对的，需要单独计算

- 1010 not done

#### DP 动态规划 以及遍历数组赋值，这里是后面的值等于前面的总和，同时加上比较

- 746

- 1277 统计给定二维数组的正方形数量

#### 移除特定元素，需要记录元素出现次数，比较当前元素和目标元素，如果只有一个，则只需要一个循环，相同则递增，否则赋值

- 这里，可以统计唯一的元素的数量，也可以统计重复元素的数量

- 1089 27 217 这个是判断是否存在重复元素

- 如果是多个元素，则需要判断是否排序，同时，在一趟里面将一个元素处理完

- 26

#### 计算距离的处理，特殊情况需要处理好

- 849

#### 统计序对的特征，相反的序对数量，这里通过将序对的数值合并成一个数值并放到 hash table 中进行统计

- 1128

#### 利用数学以及观察检查是否存在序列使得元素和满足特定的值

- 840 not done

#### 数组的升序以及降序条件判断

- 941

#### 变相求最大公约数

- 914

#### 遍历数组找空位，分别向尾部和头部补 0，然后从第二个元素开始，判断三个元素是否都是 0 来进行空位计数
    
- 605 not done

#### 非递减数组判断

- 665 i - 2, i - 1, i，其中 i - 1 的元素大于 i 的元素

- 这里，可以改变的元素可以是 i 赋值给 i - 1 或者是 i - 1 赋值给 i

- 因为将上一个元素变小对结果没有影响，所以，尽量这样处理，其中，

    - 需要判断 i - 2 元素是否小于等于 i 元素

        - 如果是则可以安全地直接将 i 赋值 给 i - 1 其中，开头的两个元素直接可以这样赋值而不需要判断了

            - 这样判断的意义是，如果满足，则将 i - 1 变成 i 的元素，也不会对结果产生影响

    - 如果不满足，则只能将 i 元素的值变成 i - 元素的值

### 辅助数组

#### 数组的重新构建使用到辅助数组，同时，注意元素的添加处理，先把元素保存，如果是 0 则需要再次保存一个 0

- 1089

- 217 类似保存了已经出现的元素

- 26 类似的处理，只是这里是统计重复元素

#### 计算数学公式，通过一个元素，计算所需要的另外一个元素

- 利用辅助数组对其中一个数组的元素出现进行统计

- 针对第一个数组，在第二个数组中找所需要的值，类似 two-sum 的做法

#### 使用辅助数组和原数组进行计算，这里类似 dp 信息分别保存在两个数组里面

- 119

### 计算元素距离

#### 顺时针以及逆时针的距离计算，直接求总距离，然后计算一个反向就得到另一个方向的距离，最后比较

- 1184

#### 求两个满足和的元素的索引，使用另一个数组存储需要的数值

- 1

### Old

#### 遍历相关

- two pointers

### 121 122

- 遍历统计每个元素特征

- 找最大差值等，遍历然后比较最大值

- 总的差值，求和相邻元素的差值

### 53 643 985

- 最大子数组和

- 985 这里的是求和，然后元素数值变化，一般看到求和，需要思考是否需要先求和再做变化

### 717 35

- 子数组元素特征匹配

### 674 830

- 找符合特征的子数组

### 88

- 两个已排序数组的合并，从尾部到头部处理

### 605

- 花盆问题，遍历找空位

### 724

- 找中间数，使两边之和相等

### 4

- 遍历两个已排序数组然  后获得中间值

- 同时需要处理其中一个数组未空的情况

1. 可以使用顺序遍历两个数组并判断，但是边界条件也不难，就是需要思考

2. 也可以使用类似归并排序的思路，统计中间值。类似 88

### 1170

- 两个数组元素特征之比较，当有特定要求时，考虑排序等方法减少比较次数

### 888

- 思考一下，两个数组和的大小关系是否需要判断

### 1

- 使用map，找特征元素，2-sum问题

### 238

- 这里可以使用两个辅助数组来计算

- 也可以不使用辅助数组，直接两次遍历都乘

## brute force

### 找三个元素之最大乘积

- 找最大三个正数以及最小两个负数再比较 628

## 遍历两个数组

## 未分类

### 数组元素实现计算中的加法

- 66