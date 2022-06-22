//
// Created by cloudy on 2022/6/23.
//

#ifndef CPP_CONNECT_H
#define CPP_CONNECT_H
// 116. 填充每个节点的下一个右侧节点指针
// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
// struct Node {
//    int val;
//    Node *left;
//    Node *right;
//    Node *next;
// }
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
// 初始状态下，所有 next 指针都被设置为 NULL。
//
// 示例 1：
// 输入：root = [1,2,3,4,5,6,7]
// 输出：[1,#,2,3,#,4,5,6,7,#]
// 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。

// 示例 2:
// 输入：root = []
// 输出：[]

#include <string>
#include <queue>

using namespace std;

class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
            : val(_val), left(_left), right(_right), next(_next) {}
};



class Solution {
public:
    Node* connect(Node* root) {
        if (root == nullptr) {
            return root;
        }

        // 初始化队列同时将第一层节点加入队列中，即根节点
        queue<Node*> Q;
        Q.push(root);

        // 外层的 while 循环迭代的是层数
        while (!Q.empty()) {

            // 记录当前队列大小
            int size = Q.size();

            // 遍历这一层的所有节点
            for(int i = 0; i < size; i++) {

                // 从队首取出元素
                Node* node = Q.front();
                Q.pop();

                // 连接
                if (i < size - 1) {
                    node->next = Q.front();
                }

                // 拓展下一层节点
                if (node->left != nullptr) {
                    Q.push(node->left);
                }
                if (node->right != nullptr) {
                    Q.push(node->right);
                }
            }
        }

        // 返回根节点
        return root;
    }
};


#endif //CPP_CONNECT_H
