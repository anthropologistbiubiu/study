
### 关联关系

class Person:
    def __init__(self,name,age,sex):
       self.name = name 
       self.age = age
       self.sex = sex
       # self.relation = None
    def do_private_stuff(self):
        pass
    def get_my_parter(self):
        for i in self.relation.couple:
            if i.name != self.name: 
                print("my parter is %s" %i.name)


class Relation:
    """ 保存对象关系 """
    def __init__(self):
      self.couple = [] 
    def marry(self,p1,p2):
       self.couple = [p1,p2] 
       print(self.couple[0].name)
       print(self.couple[1].name)
       print("get married")
       p1.relation = self
       p2.relation = self
    def divorce(self):
       print(self.couple[0].name)
       print(self.couple[1].name)
       print("divorced")
       self.couple = [] 

sunweiming = Person("sunweiming",18,"M")
yujinling = Person("yujinling",21,"F")
# yujinling.parter = sunweiming 
# sunweiming.parter = None
couple1 = Relation()
couple1.marry(sunweiming,yujinling)

# print(sunweiming.get_my_parter())
print(yujinling.get_my_parter())
couple1.divorce()


