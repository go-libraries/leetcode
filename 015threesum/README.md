https://leetcode-cn.com/problems/3sum
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

def threeSum(self,nums):
    if len(nums) < 3
        return []
    nums.sort()
    res = set()
    for i,v in enumerate(nums[:-2]):
        if i>=1 and v ==nums[i-1]
            continue
        d = {}
        for x in nums[i+1:]:
            if x not in d:
                d[-v-x] =1
            else:
                res.add((v,-v-x,x))
    return map(list,res)