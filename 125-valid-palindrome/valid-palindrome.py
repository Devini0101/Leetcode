class Solution:
    def isPalindrome(self, s: str) -> bool:
        compairArray=""
        for i in s:
            if i.isalnum():
                compairArray+=i.lower()
        return compairArray == compairArray[::-1]

        
        