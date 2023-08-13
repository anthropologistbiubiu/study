package dmysql

// 完成orm 的数据的持久化层操

// 在这里完成所有的增删改查的过程
func (g *Orm) CreateTable() error {
	return nil
}

func (g *Orm) InsertSecretKey(kind string) (int64, error) {
	return 0, nil
}
func (g *Orm) FindSecretKey(kind string) (string, error) {
	return "", nil
}

func (g *Orm) UpdateSecretKey(kind string) (string, error) {
	return "", nil
}

func (g *Orm) DeleteSecretKey(kind string) (string, error) {
	return "", nil
}
