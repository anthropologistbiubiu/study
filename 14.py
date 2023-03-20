import sys

current_users=["Niuniu","Niumei","GURR","LOLO"]
new_users=["GurR","Niu Ke Le","LoLo","Tuo Chi"]

for name in new_users:
    if current_users.count(name) > 0:
        print("The user name %s has already been registered! Please change it and try again!" % name)
    else:
        print("Congratulations, the user %s Niu Ke Le is available!'" % name) 

names=list(input().split())
course=list(input().split())
tuples=zip(names,course)
print(list(tuples))
print(dict(zip(names,course)))