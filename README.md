## Telegraf 模板输出数据格式

最新文档请移步: [中文文档](https://www.sixianed.com/2017/10/18/Telegraf-Template-Output-Data-Formats/)

让 Telegraf 的输出格式支持GO模板（text/template）:

## 安装

1. 安装Telegraf
```
go get -d github.com/influxdata/telegraf
```
2. 下载插件源码
```
go get github.com/liuhengloveyou/telegraf.serializers.template
```
3. 注册插件到Telegraf

这里要吐槽一下Telegraf，插件注册做得不方便，需要改Telegraf源码：$GOPATH/src/github.com/influxdata/telegraf/plugins/serializers/registry.go 

```
// import插件源码
import (
	... ...
	"github.com/influxdata/telegraf/plugins/serializers/json"
	template "github.com/liuhengloveyou/telegraf.serializers.template"
)

// 注册插件
func NewSerializer(config *Config) (Serializer, error) {
	var err error
	var serializer Serializer
	switch config.DataFormat {
	case "influx":
		serializer, err = NewInfluxSerializer()
	case "graphite":
		serializer, err = NewGraphiteSerializer(config.Prefix, config.Template)
	case "json":
		serializer, err = NewJsonSerializer(config.TimestampUnits)
	case "template":
		serializer, err = template.NewTemplateSerializer(config.Template)
	default:
		err = fmt.Errorf("Invalid data format: %s", config.DataFormat)
	}
	return serializer, err
}
```

## 配置

```toml
[[outputs.file]]
  ## Files to write to, "stdout" is a specially handled file.
  files = ["stdout"]

  ## Data format to output.
  ## Each data format has its own unique set of configuration options, read
  ## more about them here:
  ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_OUTPUT.md
  data_format = "template"
  ## template (text/template)
  ## 例如nginx:
  ##log_format main '$time_iso8601 $server_name $remote_addr "$request" '
  ##  '$status $body_bytes_sent $request_time "$http_referer" '
  ##  '"$http_user_agent" "$http_x_forwarded_for"';
  template = "{{.time}} {{.server_name}} {{.client_ip}} \"{{.request}}\" {{.resp_code}} {{.resp_bytes}} {{printf \"%.3f\" .request_time}} \"{{.referer}}\" {{.user_agent}} \"{{.x_forwarded_for}}\"\n"
  
```
