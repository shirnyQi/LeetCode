//
// Created by cloudy on 2022/6/21.
//

#include "gtest/gtest.h"
#include "codec.h"
#include <iostream>

TEST(test, codec)
{
    TreeNode* leftNode = new TreeNode(1);
    TreeNode* rightNode = new TreeNode(3);
    TreeNode* root = new TreeNode(2);
    root->left = leftNode;
    root->right = rightNode;

    Codec* ser = new Codec();
    Codec* deser = new Codec();
    string tree = ser->serialize(root);
    std::cout << tree << std::endl;
    TreeNode* ans = deser->deserialize(tree);
}