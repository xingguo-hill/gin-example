package common

import (
	"time"

	"github.com/spf13/viper"
)

func FormatDateYs(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func OSetConfig(sConfDir string, sName string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(sName)
	config.AddConfigPath(sConfDir)
	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	} else {
		return config
	}
}

/**
 * @description: 从map字段中提取字段值，并设置为map索引,面向数据库查询结果集处理
 * @param {*[]map[string]any} arr
 * @param {string} index
 * @return map[any]map[string]any
 */
func IndexMap(m *[]map[string]any, skey string) map[any]map[string]any {
	mRes := make(map[any]map[string]any)
	if len(*m) > 0 {
		for _, value := range *m {
			mRes[value[skey]] = value
		}
	}
	return mRes
}

/**
 * @description: 从map字段中提取字段值，并设置为一维分片,面向数据库查询结果集处理
 * @param {*[]map[string]any} arr
 * @param {string} index
 * @return []any
 */
func Map2Index(m *[]map[string]any, skey string) []any {
	sRes := make([]any, 0, len(*m))
	if len(*m) > 0 {
		for _, value := range *m {
			sRes = append(sRes, value[skey])
		}
	}
	return sRes
}
