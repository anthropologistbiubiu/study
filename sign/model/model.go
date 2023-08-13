package model

import (
	"sign/db/dmysql"
)

// 为了逻辑更清晰吧  实现对持久化层的中转

func CreateTable() error {
	return dmysql.Orms.CreateTable()
}

func FindSecretKey(kind string) (string, error) {
	return dmysql.Orms.FindSecretKey(kind)
}
func UpdateSecretKey(kind string) (string, error) {
	return dmysql.Orms.UpdateSecretKey(kind)
}
func DeleteSecretKey(kind string) (string, error) {
	return dmysql.Orms.DeleteSecretKey(kind)
}
func InsertSecretKey(kind string) (int64, error) {
	return dmysql.Orms.InsertSecretKey(kind)
}
