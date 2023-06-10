

# 输出空行和删除空行
# sed -n '/^$/=' test.txt

sed -i '.bak' '/^$/d'  test.txt  # 苹果电脑强制备份