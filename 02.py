import sys
# 字符串转数字
#  十六进制转十进制
num=input()
print(int(num,16))

# 二进制转十进制
val=input()
print(int(val,2))

# 数字转字符串
print("{:b}".format(10)) # 1010
print("{:#b}".format(10)) # 0b1010

print("{:x}".format(10)) # a
print("{:#x}".format(10)) # 0xa







