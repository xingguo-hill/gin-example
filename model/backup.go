package model

import (
	"kvm_backup/dao"
	"time"
)

const TTask = "t_backup_task"
const TStorage = "t_backup_storage"
const PageSize = 100

type TBackupTask struct {
	ID              int `gorm:"primaryKey"`
	Name            string
	Ip              string
	ScheduleType    string
	CronExpression  string
	AtTime          string
	RetentionPeriod int
	Status          int
	Utime           time.Time `gorm:"autoUpdateTime"`
	Ctime           time.Time `gorm:"autoCreateTime"`
}

type TBackupLog struct {
	ID       int `gorm:"primaryKey"`
	ClientIp string
	BackupId int
	LogTable string
	Operator string
	Content  string
	Ctime    time.Time
}

func SaveTask(s *TBackupTask) (err error) {
	if err = dao.DB.Omit("ctime", "utime").Save(&s).Error; err != nil {
		return err
	}
	return
}

func GetTaskDetail(s *map[string]any, id int) (err error) {
	if err = dao.DB.Model(TBackupTask{}).First(s, id).Error; err != nil {
		return err
	}
	return
}

func GetTaskByIds(sids *[]any, infos *[]map[string]any) (err error) {
	if err = dao.DB.Model(TBackupTask{}).Where("id in (?)", *sids).Find(infos).Error; err != nil {
		return err
	}
	return
}

func GetAllTask(infos *[]map[string]any, c *int64, size int, offset int) (err error) {
	if err = dao.DB.Model(TBackupTask{}).Order("id desc").Limit(size).Offset(offset).Find(infos).Count(c).Error; err != nil {
		return err
	}
	if err = dao.DB.Model(TBackupTask{}).Count(c).Error; err != nil {
		return err
	}
	return
}

func InsertLog(s *TBackupLog) (err error) {
	if err = dao.DB.Omit("ctime").Create(s).Error; err != nil {
		return err
	}
	return
}

func GetLogs(infos *[]map[string]any, c *int64, bakupId *int, ctime *string, limit int) (err error) {
	tx := dao.DB.Model(TBackupLog{}).Order("id desc")
	if *bakupId > 0 {
		tx.Where("backup_id=?", *bakupId)
	}
	if len(*ctime) > 0 {
		tx.Where("ctime like ?", *ctime+"%")
	}
	if err = tx.Find(infos).Count(c).Limit(limit).Error; err != nil {
		return err
	}
	return
}
