# 动态域名解析

在本地局域网中运行此程序。此程序会查询本地局域网所属的公网IP（网络供应商分配的动态IP）， 然后通过阿里云API，将该IP绑定到阿里云指定的域名上，从而实现域名与IP的动态绑定。


[API参考文档](https://help.aliyun.com/document_detail/2355673.html?spm=a2c4g.2355665.0.0.60734802Zanhco)


## 1、获取阿里云接口请求Accesskey

+ AccessKey ID

+ AccessKey Secret

配置到脚本中 accessKeyId, accessKeySecret 变量

## 2、使用Accesskey查询子域名ID

```shell
# 设置环境变量
export ALIBABA_CLOUD_ACCESS_KEY_ID=$accessKeyId
export ALIBABA_CLOUD_ACCESS_KEY_SECRET=$accessKeySecret

# 查询域名所有的
./dynamic_ip queryRecords 7shu.co
```

```shell
+---+--------------------+------------+-------------+--------+------+----------------+---------------------------+
| # | RECORD ID          | RR         | DOMAIN NAME | STATUS | TYPE | VALUE          | UPDATE TIME               |
+---+--------------------+------------+-------------+--------+------+----------------+---------------------------+
| 0 | 8686xxxxxxxx803648 | lighthouse | 7shu.co     | ENABLE | A    | 81.70.166.8    | 2023-12-20T14:59:46+08:00 |
+---+--------------------+------------+-------------+--------+------+----------------+---------------------------+
| 1 | 8663xxxxxxxx228672 | live       | 7shu.co     | ENABLE | A    | 221.3.21.99    | 2024-01-19T11:00:52+08:00 |
+---+--------------------+------------+-------------+--------+------+----------------+---------------------------+
```

配置脚本中的 recordId, recordName, domainName 变量

## 3、配置动态IP绑定命令

配置脚本中的 cmd 变量

## 4、创建日志输出相关变量

logPath -- 日志输出目录
logFile -- 日志输出文件名

## 5、日志滚动配置

logrotateCfg --  logrotate 配置文件全路径

配置文件属性设置

```shell
chown root:root /path/to/logrotateCfg
chmod 644       /path/to/logrotateCfg
```

## 6、配置计划任务

```shell
sudo crontab -e

*/1 * * * * /home/tao/Desktop/tools/dynamic_ip/bind.sh > /tmp/dynamic_ip.log
```
