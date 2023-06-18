



array=("sun" "wei" "ming")
i=0
while [ $i -lt ${#array[@]} ]
#当变量（下标）小于数组长度时进入循环体
do
    echo ${ #array[i] }
    #按下标打印数组元素
    let i++
done