package tree

/*
[ref](https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/you-ya-de-trie-by-zhi-zhao-cchub/)
[ref](https://leetcode-cn.com/problems/implement-trie-prefix-tree/solution/shi-xian-trie-qian-zhui-shu-by-leetcode/)
Implement a trie with insert, search, and startsWith methods.
实现一个前缀树，以及它们的插入，搜索和 startswith 方法
所有输入都是小写以及非空字符串，这里创建节点是通过建立新的 Trie 完成，每个 Trie 有自己的 slice
根节点对应的 next 数组中存在 Trie 就表示后面还有字母
*/

// 每个直接初始化 26 个节点的 slice 以及一个 isEnd 表示当前节点是否是字符的停止位置
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
/*
这里需要遍历插入的字符串，如果当前字母的 next 数组是 nil 就初始化
这里需要每个字符都取处理，下一个字符后面如果没有节点就需要创建，最后标记 isEnd = true
*/

func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		// 当前遍历到的节点 next slice 没有指向
		if node.next[v] == nil {
			node.next[v] = &Trie{}
		}
		node = node.next[v]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
