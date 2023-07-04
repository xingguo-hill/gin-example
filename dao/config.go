package dao

import (
	"kvm_backup/common"

	"github.com/spf13/viper"
)

var D = viper.GetString
var S = viper.GetString

func init() {
	dViper := common.OSetConfig("conf/", "db")
	D = dViper.GetString

	sViper := common.OSetConfig("conf/", "service")
	S = sViper.GetString
}
