# 查字典
fruit_dic={'a': ['apple', 'abandon', 'ant'], 
     'b': ['banana', 'bee', 'become'], 
     'c': ['cat', 'come'], 'd': 'down'}
a=input()
for fruit in fruit_dic[a]:
    print(fruit,end=" ")

print("")
# 字符串类型的比较
string=input()
print(string.isalpha())
print(string.isdigit())
print(string.isspace())
# 字母转数字

alph=input()
print(ord(alph))
# 数字的二进制表示
num=input()
print(bin(int(num)))

#数字的八进制表示


#数字的十六进制的表示
num=input()
print("0x%x" % int(num))