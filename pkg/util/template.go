// Copyright 2019 Squeeze Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"encoding/json"
	"fmt"
	"text/template"
)

func NewTemplate(outputTmpl string) *template.Template {
	return template.Must(template.New("tmpl").Funcs(tmplFuncMap).Parse(outputTmpl))
}

var tmplFuncMap = template.FuncMap{
	"formatNumber":       formatNumber,
	"formatNumberInt":    formatNumberInt,
	"formatNumberUint64": formatNumberUint64,
	"formatNumberInt64":  formatNumberInt64,
	"jsonify":            jsonify,
}

func jsonify(v interface{}) string {
	d, _ := json.Marshal(v)
	return string(d)
}

func formatNumber(duration float64) string {
	return fmt.Sprintf("%4.4f", duration)
}

func formatNumberInt(duration int) string {
	return fmt.Sprintf("%d", duration)
}

func formatNumberUint64(i uint64) string {
	return fmt.Sprintf("%d", i)
}

func formatNumberInt64(i int64) string {
	return fmt.Sprintf("%d", i)
}
