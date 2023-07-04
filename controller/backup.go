package controller

import (
	"encoding/json"
	"html/template"
	"kvm_backup/common"
	"kvm_backup/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TaskGetIndex(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.Query("page"))
	var count int64
	var tInfo map[string]any
	var mtask []map[string]any
	if id > 0 {
		model.GetTaskDetail(&tInfo, id)
	}
	size := 10
	model.GetAllTask(&mtask, &count, size, (page-1)*size)
	pagination := common.NewPagination(c.Request, count, size, "page", page)
	c.HTML(http.StatusOK, "tasklist.tpl", gin.H{
		"list":  mtask,
		"count": count,
		"info":  tInfo,
		"pages": template.HTML(pagination.Pages()),
	})
}

func TaskPostIndex(c *gin.Context) {
	s := model.TBackupTask{}
	s.Ip = c.PostForm("ip")
	s.Name = c.PostForm("name")
	s.ScheduleType = c.PostForm("scheduleType")
	if s.ScheduleType == "cron" {
		s.CronExpression = c.PostForm("expression")
	} else if s.ScheduleType == "at" {
		s.AtTime = c.PostForm("expression")
	}
	s.RetentionPeriod, _ = strconv.Atoi(c.PostForm("retentionPeriod"))
	s.Status, _ = strconv.Atoi(c.PostForm("status"))
	id, _ := strconv.Atoi(c.PostForm("id"))
	if id > 0 {
		s.ID = id
	}
	// fmt.Printf("sNew var =%#v\n", s)
	if model.SaveTask(&s) != nil {
		c.HTML(http.StatusOK, "tip.tpl", gin.H{
			"tip": "任务操作失败",
		})
		return
	}

	//记录操作日志
	log := make(map[string]any)
	log["conent"] = s

	tcontent, _ := json.Marshal(s)

	if insertLog(s.ID, c.ClientIP(), model.TTask, "test", tcontent) != nil {
		c.HTML(http.StatusOK, "tip.tpl", gin.H{
			"tip": "日志添加失败",
		})
		return
	}
	c.Redirect(302, "/kvm_backup/task/")
}
func LogGetIndex(c *gin.Context) {
	//GET请求参数获取
	bid, _ := strconv.Atoi(c.Param("bid"))

	//POST请求参数获取
	ctime := c.PostForm("ctime")

	//获取日志列表
	var mlog []map[string]any
	var count int64
	model.GetLogs(&mlog, &count, &bid, &ctime, model.PageSize)
	sbids := common.Map2Index(&mlog, "backup_id")

	//获取task列表
	var mtask []map[string]any
	model.GetTaskByIds(&sbids, &mtask)
	mitask := common.IndexMap(&mtask, "id")
	// fmt.Printf("sNew var =%#v\n%#v\n", len(mitask), mitask)
	//对task与log信息进行组合
	if len(mlog) > 0 {
		for _, value := range mlog {
			value["ip"] = mitask[value["backup_id"]]["ip"]
			value["name"] = mitask[value["backup_id"]]["name"]
		}
	}
	c.HTML(http.StatusOK, "loglist.tpl", gin.H{
		"loglist": mlog,
		"bid":     bid,
		"count":   count,
		"ctime":   ctime,
	})
}

func insertLog(bid int, cip string, stable string, operator string, tcontent []byte) (err error) {
	s := model.TBackupLog{}
	s.BackupId = bid
	s.ClientIp = cip
	s.LogTable = stable
	s.Operator = operator
	s.Content = string(tcontent)
	return model.InsertLog(&s)
}
