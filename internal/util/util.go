package util

import "backupergo/internal/config"

// ConvertConfig преобразует структуру Config для использования в разных контекстах
func ConvertConfig(src config.Config) config.Config {
	return config.Config{
		MysqlPath: src.MysqlPath,
		DBQuery:   src.DBQuery,
		DBName:    src.DBName,
	}
}
