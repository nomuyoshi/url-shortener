package snowflake

import (
	"errors"
	"time"

	"github.com/nomuyoshi/snowflake"
)

var sf *snowflake.Snowflake

func SnowFlake() *snowflake.Snowflake {
	return sf
}

func SetSnowFlake() error {
	if sf != nil {
		return errors.New("既に設定済みです。後から変更できません")
	}
	echoTime, _ := time.Parse(time.DateTime, "2023-10-01 00:00:00")
	if s, err := snowflake.NewSnowflake(echoTime, 0, 0); err != nil {
		panic(err)
	} else {
		sf = s
	}
	return nil
}
