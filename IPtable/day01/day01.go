package main

// ip白名单就是防御为未知IP的SYN攻击

// mysql 和 redis 的缓存针对已知的数据查询  缓存查不到；数据库一定是可以查到；

// 因此二者之间是对立的设计。

// 因此可以使用reidis 存储白名单；再使用原生map 来做映射关系。

// 因此关于redis 是存储关于配置信息；mysql 是存储结构化的数据资料的；不是所有的redis 都要落库；
func main() {

	return 0

}
