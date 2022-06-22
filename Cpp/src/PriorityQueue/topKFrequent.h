//
// Created by cloudy on 2022/6/23.
//

#ifndef CPP_TOPKFREQUENT_H
#define CPP_TOPKFREQUENT_H

//347. 前 K 个高频元素
//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

//示例 1:
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]

//示例 2:
//输入: nums = [1], k = 1
//输出: [1]
//

#include <queue>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    static bool cmp(pair<int, int>& m, pair<int, int>& n) {
        return m.second > n.second;
    }

    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> occurrences;
        for (auto& v : nums) {
            occurrences[v]++;
        }

        // pair 的第一个元素代表数组的值，第二个元素代表了该值出现的次数
        // std::priority_queue:优先队列
        // priority_queue<Type, Container, Functional>, 其中Type为数据类型，Container为保存数据的容器，Functional为元素比较方式
        // decltype，在C++中，作为操作符，用于查询表达式的数据类型。
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(&cmp)> q(cmp);
        for (auto& [num, count] : occurrences) {
            if (q.size() == k) {
                // 如果优先队列满了，且队首的个数比待入队列的次数少，则替换
                if (q.top().second < count) {
                    q.pop();
                    q.emplace(num, count);
                }
            } else {
                // 优先队列没满则插入数据
                q.emplace(num, count);
            }
        }
        vector<int> ret;
        while (!q.empty()) {
            //emplace_back()是c++11的新特性,和push_back()的区别在于:
            //push_back()方法要调用构造函数和复制构造函数，这也就代表着要先构造一个临时对象，然后把临时的copy构造函数拷贝或者移动到容器最后面。
            //而emplace_back()在实现时，则是直接在容器的尾部创建这个元素，省去了拷贝或移动元素的过程
            ret.emplace_back(q.top().first);
            q.pop();
        }
        return ret;
    }
};

#endif //CPP_TOPKFREQUENT_H
