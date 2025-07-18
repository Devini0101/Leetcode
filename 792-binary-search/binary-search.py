class Solution:
    def search(self, nums: List[int], target: int) -> int:
        start=0
        end=len(nums)- 1 #o ultimo index da lista, já que ela comeca com 0

        while start<=end:
            mid=(start+end)//2
            if target>nums[mid]:
                start=mid + 1
            elif target<nums[mid]:
                end=mid - 1
            else:
                return mid
        return -1