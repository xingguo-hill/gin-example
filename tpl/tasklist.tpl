<!doctype html>
<html>
<head>
  <link rel="stylesheet" href="/css/bootstrap.css">
  <link rel="stylesheet" href="/css/base.css">
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>虚机备份配置</title>
</head>
<body>
<ul class="nav nav-tabs">
  <li class="nav-item">
    <a class="nav-link active" aria-current="page">备份任务配置</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" aria-current="page" href="/kvm_backup/log/">操作日志</a>
  </li>
</ul>
  <form action="/kvm_backup/task/" method="post">
  <div class="mb-3 mt-3 ms-2">
      <input type="hidden" name="id" value="{{.info.id}}">
      <label>虚拟机IP:</label>
      <input type="text" name="ip" value="{{.info.ip}}" placeholder="内网服务ip">
      <label>虚拟机名称:</label>
      <input type="text" name="name" value="{{.info.name}}">
      <label>备份策略:</label>
      <select name="scheduleType" id="scheduleType">
        <option value="cron" {{if eq .info.schedule_type "cron"}} selected {{end}}>cron</option>
        <option value="at" {{if eq .info.schedule_type "at" }} selected {{end}}>at</option>
      </select>
      <input type="text" name="expression"  value='{{if eq .info.schedule_type "cron"}}{{.info.cron_expression}}{{else}}{{.info.at_time}}{{end}}' placeholder="备份策略:cron(周期)与linux计划任务相同，at(一次性)执行格式YYYY-MM-DD HH:MM:SS">
      <label>备份镜像保留天数（永久:0）</label>
      <input type="text" name="retentionPeriod" value="{{.info.retention_period}}" placeholder="天数">
      <label>状态:</label>
      <select name="status" id="status">
          <option value="0" {{if eq .info.status 0}} selected {{end}} >启用</option>
          <option value="1" {{if eq .info.status 1}} selected {{end}}>停用</option>
      </select>
   <input type="submit" value="{{if eq .info.id nil}}新建{{else}}修改{{end}}">
  </div>
  </form> 
  <div class="bg-primary ms-3 text-center">总记录数：{{.count}}</div>
  <table class="table ms-3">
      <tr class="table-primary"><th>IP</th><th>名称</th><th>备份类型</th><th>备份策略</th><th>更新时间</th><th>保留时间</th><th>状态</th><th>操作</th></tr>
      {{if gt .count 0}}
        {{ range .list }}
        <tr><td>{{.ip}}</td><td>{{.name}}</td><td>{{.schedule_type}}</td><td>{{if eq .schedule_type "cron"}}{{.cron_expression}}{{else}}{{.at_time}}{{end}}</td><td>{{.utime|formatDate}}</td><td>{{if .retention_period}}{{.retention_period}}{{else}}永久{{end}}</td><td>{{if eq .Status 0}}启用{{else}}停用{{end}}</td><td><a href="/kvm_backup/task/{{.id}}">修改 | <a href="/kvm_backup/log/{{.id}}">日志</a></td></tr>
        {{end}}
      {{end}}
  </table>
{{.pages}}
</body>
