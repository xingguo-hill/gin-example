# info
gin学习框架
以虚机kvm备份与日志记录为例，对gin框架进行使用了解

# 配置项目
<ol>
<li>conf/db.sql导入数据库表</li>
<li>conf/db.toml 配置数据库连接</li>
<li>conf/server.toml 配置服务运行配置</li>
</ol>

# 项目启动
<ul>
<li>debug模式：go run main.go web</li>
<li>二进制运行 go build;./kvm_backup web</li>
</ul>
