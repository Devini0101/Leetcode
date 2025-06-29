class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        n=len(nums)
        for i in range(n):
            for j in range(i+1,n):
                #sum the nums in the position & compare
                if nums[i]+ nums[j] == target:
                   #posi founded 
                 return [i,j]
        
        