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
    <a class="nav-link" aria-current="page" href="/kvm_backup/task/">备份任务配置</a>
  </li>
  <li class="nav-item">
    <a class="nav-link active" aria-current="page">操作日志</a>
  </li>
</ul>
  <form action="/kvm_backup/log/" method="post">
  <div class="mb-3 mt-3 ms-2">
      <input type="hidden" name="bid" value="{{.bid}}">
      <label>时间:</label>
      <input type="text" name="ctime" value="{{.ctime}}" placeholder="YYYY-MM-DD">
   <input type="submit" value="查询"> 
  </div>

  </form> 
  <div class="bg-primary ms-3 text-center">总记录数：{{.count}}</div>
  <table class="table ms-3">
      <tr class="table-primary"><th>虚机IP</th><th>名称</th><th>ClientIp</th><th>操作人</th><th>相关表</th><th>创建时间</th><th>内容</th></tr>
      {{if gt .count 0}}
        {{ range .loglist }}
        <tr><td>{{.ip}}</td><td>{{.name}}</td><td>{{.client_ip}}</td><td>{{.operator}}</td><td>{{.log_table}}</td><td>{{.ctime|formatDate}}</td><td>{{.content}}</td></tr>
        {{end}}
      {{end}}
  </table>
</body>
<style>

</style>
