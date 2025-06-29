class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x<0:
            return False
        #0 until 9 are palindromes
        elif x>=0 and x<10:
            return True
        else:
         digit = str(x)
         posi=len(digit)
         #check from the middle
         for i in range (posi//2):
        #if the index that startes is diff from the one in the middle its already not a palindrome
            if digit[i] != digit[-i-1]:
                return False
        return True