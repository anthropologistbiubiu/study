grade=0.0
count=0
course=0
gpa={"A":4.0,"B":3.0,"C":2.0,"D":1.0,"F":0}
while True:
    if input()=="False":
        print(round(float(grade/count)),2)
        break
    else:
        m=gpa[input()]
        n=float(input())
        count+=1
