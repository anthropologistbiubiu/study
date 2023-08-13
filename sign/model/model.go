package model

import (
	"sign/db/dmysql"
)

// 为了逻辑更清晰吧  实现对持久化层的中转

func CreateTable() error {
	return dmysql.Orm.CreateTable()
}

func FindSecretKey(kind string) (string, error) {
	return dmysql.Orm.FindSecretKey(kind)
}
func UpdateSecretKey(kind string) (string, error) {
	return dmysql.Orm.UpdateSecretKey(kind)
}
func DeleteSecretKey(kind string) (string, error) {
	return dmysql.Orm.DeleteSecretKey()
}
func InsertSecretKey(kind string) (int64, error) {
	return dmysql.Orm.InsertSecretKey(kind)
}
