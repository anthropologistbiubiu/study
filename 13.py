a = ["Niumei", "Niu Ke Le", "GURR", "LOLO"]
z = {"Niumei": "Nowcoder", "GURR": "Huawei"}

for name in a :
    if  name in z:
        print("Hi, %s! Thank you for participating in our graduation survey!" % name)
    else:
        print("Hi, %s! Could you take part in our graduation survey?" %name)