package array

/*
146. LRU Cache
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present.
When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
The cache is initialized with a positive capacity.
Follow up:
Could you do both operations in O(1) time complexity?

[ref](https://leetcode-cn.com/problems/lru-cache/solution/lruhuan-cun-ji-zhi-by-leetcode-solution/)
实现一个数据结构来实现 LRU cache，支持 get 和 put 操作，get(key) 返回正数或者 -1 表示不存在，同时需要移动到开头
put(key,value)如果 key 不存在就保存，存在就更新数值，当 cache 满了，在插入一个新的数据前删除最近最少使用的
直接队列了 O(1) 是使用 hash 和双向链表
*/

// 构建双向链表节点
type CacheNode struct {
	Key  int
	Val  int
	Next *CacheNode
	Prev *CacheNode
}

type LRUCache struct {
	Cap   int
	Len   int
	Head  *CacheNode
	Tail  *CacheNode
	Cache map[int]*CacheNode
}

// 这里使用伪头部尾部
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Cache: map[int]*CacheNode{},
		Head:  &CacheNode{0, 0, nil, nil},
		Tail:  &CacheNode{0, 0, nil, nil},
		Cap:   capacity,
	}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}
	node := this.Cache[key]
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	// 之前不存在
	if _, ok := this.Cache[key]; !ok {
		node := &CacheNode{Key: key, Val: value}
		this.Cache[key] = node
		this.addToHead(node)
		this.Len++
		// 超过容量
		if this.Len > this.Cap {
			removed := this.removeTail()
			delete(this.Cache, removed.Key)
			this.Len--
		}
	} else {
		node := this.Cache[key]
		node.Val = value
		this.moveToHead(node)
	}
}

// 将节点添加到头部 put, get 时
func (this *LRUCache) addToHead(node *CacheNode) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

// 移除节点 moveToHead
func (this *LRUCache) removeNode(node *CacheNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// 将节点添加到头部 put get
func (this *LRUCache) moveToHead(node *CacheNode) {
	// 先从链表中移除该元素再移动到头部
	this.removeNode(node)
	this.addToHead(node)
}

// 移除尾部节点 超过容量
func (this *LRUCache) removeTail() *CacheNode {
	node := this.Tail.Prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
