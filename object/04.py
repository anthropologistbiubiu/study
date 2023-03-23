# 组合关系


# 由一堆组件构成一个完整的实体 组件本身就不自己允许 



class Dog:
    def __init__(self,name,breed,attack_val,life_val) -> None:
        self.name = name 
        self.breed = breed
        self.attack_val = attack_val
        self.life_val  = life_val
    def bite(self,person):
        person.life_val -= self.attack_val
        print("person.life_val is",person.life_val)


class Arms:
    def __init__(self):
        self.knife_attack_val=10
        self.gun_attack_val=20
    def knife(self,dog):
        dog.life_val -= self.knife_attack_val
    def gun(self,dog):
        dog.life_val -= self.gun_attack_val
    def log(self,dog):
        dog.life_val -= self.knife_attack_val
    def log(self,dog):
        print("dog's life_val is %d" %dog.life_val)
    
class Person:
    role = "person"
    def __init__(self,name,sex,attack_val,life_val) -> None:
        self.name = name
        self.sex=sex
        self.attack_val = attack_val
        self.life_val = life_val
        self.arm = Arms()
    def attack(self,dog):
        dog.life_val -= self.attack_val
        print("dog.life_val is ",dog.life_val)

p1 = Person("sunweiming","M",0,100)
d1 = Dog("二哈","二哈",5,40)
# arm1 = Arms()
# p1.arm = arm1
p1.arm.knife(d1)
p1.arm.log(d1)