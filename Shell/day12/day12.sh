#! /bin/sh

# 示例8：使用模式读取字符串列表
# 在这里，“ /，/”模式用于根据逗号分割字符串值。
# Define a list of string variable
stringList=WordPress,Joomla,Magento

# Use comma as separator and apply as pattern
for val in ${stringList//,/ }
do
   echo $val
done
:<<!
输出hello world
!
echo "hello world"