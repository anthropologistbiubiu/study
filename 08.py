# try:
#     fh = open("test.txt", "w")
#     fh.write("这是一个测试文件，用于测试异常!!")
# except IOError:
#     print "Error: 没有找到文件或读取文件失败"
# else:
#     print "内容写入文件成功"
#     fh.close()
# 

entry_form=("Niuniu","Niumei")
print(entry_form)
try:
    entry_form[1]="Niukele"
except TypeError:
    print("The entry form cannot be modified!")
else:
    print("ok!")

operator_dict={"<":"less than","==":"equal"}
print("Here is the origianl dict:")
for i in sorted(operator_dict):
    print("Operator %s means %s." % (i, operator_dict[i]))




my_dict_1={"name":"Niuniu","Student ID":1}
my_dict_2={"name":"Niumei","Student ID":2}
my_dict_3={"name":"Niu Ke Le","Student ID":3}
dict_list=[]
dict_list.append(my_dict_1)
dict_list.append(my_dict_2)
dict_list.append(my_dict_3)
print(dict_list)

# for value in dict_list:
#    print(value)
#    print(value["name"])
#    print(value["Student ID"])
#
new_dict={'a': ['apple', 'abandon', 'ant'], 'b': ['banana', 'bee', 'become'], 'c': ['cat', 'come'], 'd': 'down'}
letter=input()
word=input()
new_dict.setdefault(letter)
print(new_dict)