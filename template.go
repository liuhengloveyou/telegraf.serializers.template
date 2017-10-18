package template

import (
	"bytes"
	"text/template"

	"github.com/influxdata/telegraf"
)

type TemplateSerializer struct {
	Template string
}

func NewTemplateSerializer(template string) (*TemplateSerializer, error) {
	return &TemplateSerializer{
		Template: template,
	}, nil
}

func (s *TemplateSerializer) Serialize(metric telegraf.Metric) (serialized []byte, err error) {
	var b bytes.Buffer

	tmpl, err := template.New("t").Parse(s.Template)
	if err != nil {
		return b.Bytes(), err
	}

	err = tmpl.Execute(&b, metric.Fields())
	if err != nil {
		return b.Bytes(), err
	}

	return b.Bytes(), nil
}
