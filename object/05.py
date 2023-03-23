
# 继承方法
class Animal:
    a_type = "哺乳动物"
    def __init__(self,name,age,sex):
       self.name = name 
       self.age = age
       self.sex = sex
    def eat(self):
        print("%s is eating" % self.name) 

class Person(Animal):
    a_type = "哺乳高等动物"
    def __init__(self, name, age, sex,hobby):
        self.hobby = hobby
        Animal.__init__(self,name, age, sex) # 人在优雅的吃饭
    def talk(self):
        print("person %s is talking..." % self.name)
    def eat(self):
        Animal.eat(self)
        print("人在优雅的吃饭...")


class Dog(Animal):
    def chase_rabbit(self):
        print("狗在追兔子...")
# 先执行父类别执行方法 再去执行子类执行方法

# animal =Animal()
p = Person("sunweiming",18,"M","pingpong")
p.talk()
p.eat()