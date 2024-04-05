# printf命令与echo命令的区别


#   区别一：printf不会自动换行，echo自动换行
#   区别二：printf一般用于格式打印，echo用于标准输出


# 2.printf语法结构：printf "%-8s %-8s %-8s\n" 姓名 性别 体重kg
printf "%-8s %-8s %-8s\n" a 1 1kg
printf "%-8s %-8s %-8s\n" b 1 1kg
printf "%-8s %-8s %-8s\n" c 1 1kg

# "%-8s"表示一个宽度为8个字符的内容  "-"表示左对齐
# 使用单引号也可以
