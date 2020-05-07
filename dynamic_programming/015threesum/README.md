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
    
    
本题的思路，就是“双指针法” + “将三数相加简化为两数相加”：

首先把原来的数组从小到大排序，排序后
先确定一个数，将这个数作为一个目标，比如数是nums[i]，那么目标就是让剩下的两个数的和等于-nums[i]
也就是nums[i] + 剩下两个数的和 = 0
在该数后面数字的一头一尾设定两个桩left/right，依次执行如下判断：
如果nums[i] + nums[left] + nums[right] < 0，说明nums[left]数字不够大，还可以把left向右移动，即left += 1
如果nums[i] + nums[left] + nums[right] > 0，说明nums[left]数字不够小，还可以把right向左移动，即right -= 1
如果nums[i] + nums[left] + nums[right] == 0，这就是我们要的解答，将对应的三个数字加到最终结果里面去。
等等，会不会有重复的解呢？当然有啦，比如官方的测试用例-2, 0, 0, 2, 2里面，i=0的时候，left=1/right=4还有left=2/right=3是一样的解。

所以咱们可以再用一个哈希表存储当前循环里已经存在的右边索引，nums[i]和nums[right]是绝对可以确定唯一nums[left]的。这感觉有点像两点确定一条直线...
这样每次得到解时，先判断nums[right]在不在哈希表，不在的话，就添加到最终结果。

最后把left和right都向中间移一格，即left += 1 right -= 1

等到left >= right的时候，说明已经查找完毕，可以选择下一个“目标”，即nums[i]了
依次遍历nums，将i逐渐+1,但是当nums[i]>0的时候，就可以结束了，因为三个正数相加不可能等于0;
此外，如果nums[i] == nums[i-1]，可以直接跳过这个nums[i]了，因为相同的事情已经做过了

