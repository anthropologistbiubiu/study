# linux 的输出打印


# printf命令与echo命令的区别

num=0
while (($num <=500))
do
if (($num % 7 == 0))
then 
    echo $num
    let "num+=7" 
fi
done