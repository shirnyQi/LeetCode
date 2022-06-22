//
// Created by cloudy on 2022/6/23.
//

#ifndef CPP_ISVALIDBST_H
#define CPP_ISVALIDBST_H

#include <limits.h>
#include <string>

// 98. 验证二叉搜索树

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    bool helper(TreeNode* root, long long lower, long long upper) {
        if (root == nullptr) {
            return true;
        }
        if (root -> val <= lower || root -> val >= upper) {
            return false;
        }
        return helper(root -> left, lower, root -> val) && helper(root -> right, root -> val, upper);
    }
    bool isValidBST(TreeNode* root) {
        return helper(root, LONG_MIN, LONG_MAX);
    }
};

#endif //CPP_ISVALIDBST_H
