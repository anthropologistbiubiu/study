# 写一个 bash脚本以输出数字 0 到 500 中 7 的倍数(0 7 14 21…)的命令


num=0


while (($num<=500))
do
if (($num%7==0))
then
    echo $num
fi
    let "num++"
done


# chmod u=wrx Shell/day03/day03.sh 