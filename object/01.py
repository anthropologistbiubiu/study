# 定义一个类
class Dog:
    d_type = "京巴" # 类属性
    def sayhi(self):# 代表实际例子本身 定义方法
        print("hello i am a dog, my bite is ",self.d_type,self.name,self.age)
    def __init__(self,name,age):   # 初始化方法 ，构造函数 ，实际化会自动执行，进行初始化工作
        print("hahaha",name,age)
        self.name = name
        self.age = age 
    # 要想把name age 这两个值 真正的存到实例中 就要进行绑定

# 实际例化
d1 = Dog()
d2 =  Dog()
d1.sayhi()
Dog.d_type= "藏獒"  # 共有属性

# 构造函数
d3 = Dog("mjj",3)
d3.sayhi()


# 实例变量 成员变量 的引用

print(Dog.d_type)

Dog.d_type = "金毛" # 公共属性的访问

class People:
    nationnality = "CN"
    def __init__(self,name,age,sex,nationality) -> None:
        self.name = name
        self.age = age 
        self.sex =sex

p1 = People("mjj","22","M")
p2 = People("Alex","23","M")
p3 = People("Jack","25","F")
print(p1)
print(p1.nationnality)

p1.nationnality = "JP"

print(p1.nationnality) # 相当于给p1 实例创建了新的实例属性
    



