class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        #set will remove the duplicate values and list will show both sets on a array
        return list(set(nums1) & set(nums2))
        
