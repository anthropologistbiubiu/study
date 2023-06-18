# shell遍历数组3种方法



array=(192.168.0.1 192.168.0.2 192.168.0.3)


for(( i=0;i<${#array[@]};i++))
#${#array[@]}获取数组长度用于循环
do
    echo ${array[i]};
done;


array1=("sun" "wei" "ming")
for (( i=0;i<${#array1[@]};i++))

do
    echo ${array1[i]};
done;