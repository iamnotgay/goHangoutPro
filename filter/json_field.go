package filter

import (
	"github.com/golang/glog"
)

type JsonFiledFilter struct {
	field  string
	target string
}

func NewJsonFiledFilter(config map[interface{}]interface{}) *JsonFiledFilter {
	plugin := &JsonFiledFilter{
		target: "",
	}

	if field, ok := config["field"]; ok {
		plugin.field = field.(string)
	} else {
		glog.Fatal("field must be set in Json filter")
	}
	if target, ok := config["target"]; ok {
		plugin.target = target.(string)
	}
	return plugin
}

func (plugin *JsonFiledFilter) Filter(event map[string]interface{}) (map[string]interface{}, bool) {
	var o = event[plugin.field]
	for k, v := range o.(map[string]interface{}) {
		event[k] = v
	}
	return event, true
}
