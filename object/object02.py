

class Dog:
    def __init__(self,name,breed,attack_val,life_val) -> None:
        self.name = name 
        self.breed = breed
        self.attack_val = attack_val
        self.life_val  = life_val
    def bite(self,person):
        person.life_val -= self.attack_val
        print("person.life_val is",person.life_val)


class Person:
    role = "person"
    def __init__(self,name,sex,attack_val,life_val) -> None:
        self.name = name
        self.sex=sex
        self.attack_val = attack_val
        self.life_val = life_val
    def attack(self,dog):
        dog.life_val -= self.attack_val
        print("dog.life_val is %",dog.life_val)

dog1 = Dog("mjj","二哈",30,100)
dog1 = Dog("马金毛","金毛",30,100)
p1 = Person("Alex","M",50,100)
p1.attack(dog1)
dog1.bite(p1)