//
// Created by cloudy on 2022/6/23.
//

#include "isValidBST.h"
#include "gtest/gtest.h"
#include <iostream>

TEST(test, isValidBST)
{
    TreeNode* leftNode = new TreeNode(1);
    TreeNode* rightNode = new TreeNode(3);
    TreeNode* root = new TreeNode(2);
    root->left = leftNode;
    root->right = rightNode;

    Solution* slt = new Solution();
    ASSERT_TRUE(slt->isValidBST(root));
}