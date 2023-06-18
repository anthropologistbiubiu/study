

# 2. for … in



array=("sun" "wei" "ming")
for element in ${array[@]}
#也可以写成for element in ${array[*]}
do
    echo $element
done;

for e in ${array[@]}

do
  echo $e
done;