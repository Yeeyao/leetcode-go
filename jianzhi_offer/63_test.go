package jianzhi_offer

import "testing"

func TestPro(t *testing.T) {
	t.Run("63  数据流中的中位数", func(t *testing.T) {
		root := Node{1, nil}
		get := solution(&root)
		want := nil
		if get != want {
			t.Errorf("got: %v, want: %v", get, want)
		}
	})
}

/*
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5

对数组进行排序然后就可以找到中位数
需要在添加元素的时候保持元素的有序性，二分查找然后保存以及后面的数据向后移动
使用堆来处理，用两个堆，各自保存一半的元素，
	A 小顶堆，长度为 N/2 或者 (N+1)/2 保存较大的部分元素
	B 大顶堆，长度为 N/2 或者 (N-1)/2 保存较小部分的元素
	当 n 为偶数，则中位数是 (a + b) / 2
	当 n 为奇数，则中位数是 a

大小顶堆的实现 可以使用 int slice 来实现二叉树，然后维护元素的大小
*/
func solution() {

}

type intMinHeap []int
