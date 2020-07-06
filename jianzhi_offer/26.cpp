//
// Created by home on 2020/7/6.
//

/*
 * 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
 * 同 leetcode 426
 */
// Definition for a Node.
class Node {
public:
    int val;
    Node *left;
    Node *right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node *_left, Node *_right) {
        val = _val;
        left = _left;
        right = _right;
    }
};

/*
 * 遍历下去，然后找到左边的叶子节点，接着指向根，最后指向右节点，然后递归向上遍历直到遍历完
 * 即中序遍历
 */
class Solution {
public:
    Node *treeToDoublyList(Node *root) {
        if (root == nullptr) {
            return nullptr;
        }
        Node* pre = nullptr, *head = nullptr;
        helper(root, head, pre);
        head->left = pre;
        pre->right = head;
        return head;
    }
    void helper(Node* root, Node* &head, Node* &pre) {
        if (!root) {
            return;
        }
        helper(root->left, head, pre);
        if (!head) {
            head = root;
            pre = root;
        } else {
            // 左边建立双向
            pre->right = root;
            root->left = pre;
            pre = root;
        }
        helper(root->right, head, pre);
    }
};
