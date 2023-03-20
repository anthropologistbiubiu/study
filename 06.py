course=0.0
count=0
while True:
    if input()=="False":
        print(course)
        print(count)
        print(course/count)
        break
    else:
        course+=float(input())
        count+=1
