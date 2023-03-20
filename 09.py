import sys
# hoby={'Allen':['red', 'blue', 'yellow'],'Tom':['green', 'white', 'blue'],'Andy':['black', 'pink']}
# 
# print(sorted(hoby))
# print(hoby)
# 
# for name in hoby:
#     print("%s's favorite colors are:" % name)
#     for h in hoby[name]:
#         print(h)
# 

# 用字典记述
ra=input()
st_dict={}
for i in list(ra):
    if  i in st_dict:
        st_dict[i]=st_dict[i]+1
    else:
        st_dict[i]=1
print(st_dict)