package redis

import (
	"fmt"
	"strconv"
)

// parseDatabase 将配置中的 Database 字段解析为 redis DB 库号。
func parseDatabase(database string) (int, error) {
	// 空字符串表示使用默认 DB=0。
	if database == "" {
		return 0, nil
	}
	// 将字符串解析为整数 DB 库号。
	db, err := strconv.Atoi(database)
	// 解析失败时包装错误并返回。
	if err != nil {
		return 0, fmt.Errorf("invalid redis database: %w", err)
	}
	// 返回解析后的 DB 库号。
	return db, nil
}
