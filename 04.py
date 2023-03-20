import sys
group_list=['Tom', 'Allen', 'Jane', 'William', 'Tony' ]
print(group_list[:2])
print(group_list[1:4])
print(group_list[-2:])


names=input()
name=input()
names_list=[]
name_list=[]
for ns in names.split():
    names_list.append(ns)
for ns in name.split():
    name_list.append(ns)
for s in name_list:
    names_list.remove(s)
print(names_list)