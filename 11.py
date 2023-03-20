import sys

nums=input()
num=""
for i in nums:
    num=num+str((int(i)+3)%9)
    print(i)
print(nums[0])
print(type(nums[0]))
print("0"+"1")
print(num[2]+num[3]+num[0]+num[1])