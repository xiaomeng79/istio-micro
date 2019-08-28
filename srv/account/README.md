## account服务

#### 介绍
这个服务是用户的账户服务,主要用来演示使用postgres,根据版本号升级对应的数据库变更记录

#### sql增量更新
1. 系统版本查看 `make next-version `下一个版本号
2. 在当前项目`sqlupdate`目录下`record.json`文件中增加下一个版本号的sql变更记录,变更sql写到文件如:`./sqlupdate/20190828_002.sql`
```bash
{"version":"0.0.2","author":"xiaomeng79","file":"./sqlupdate/20190828_002.sql","date":"2019-08-28"}
```
3. git 打标记 `git tag  0.0.2`
4. 编译程序

**详细实现在run.go下的execUpdateSql方法**

#### TODO
- 逻辑订阅