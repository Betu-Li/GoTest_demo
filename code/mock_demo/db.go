package mock_demo

// DB 数据接口
type DB interface {
	Get(key string) (int, error)
	Add(key string, value int) error
}

// GetFromDB 根据key从DB查询数据的函数
func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key); err == nil {
		return v
	}
	return -1
}
