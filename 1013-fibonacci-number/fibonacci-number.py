class Solution:
    def fib(self, n: int) -> int:
        result =n
        a,b=0,1
        for k in range (2,n+1):
            result=a+b
            a,b=b,result
        return result
        