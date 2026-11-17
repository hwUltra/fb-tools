package utils

import (
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

// KebabCase 驼峰或下划线命名转短横线小写，如 MenuDemo → menu-demo
func KebabCase(s string) string {
	var b strings.Builder
	pendingSep := false
	for _, r := range strings.TrimSpace(s) {
		if r == '_' || r == '-' || r == ' ' {
			pendingSep = true
			continue
		}
		if r >= 'A' && r <= 'Z' {
			if b.Len() > 0 && !pendingSep {
				b.WriteByte('-')
			}
			r = unicode.ToLower(r)
		}
		if pendingSep && b.Len() > 0 {
			b.WriteByte('-')
		}
		b.WriteRune(r)
		pendingSep = false
	}
	return b.String()
}

// PluralKebab 短横线命名转英文复数，如 dict-item → dict-items、category → categories
// 接口路径与前端目录共用同一套复数规则，避免生成结果与现有路由风格不一致
func PluralKebab(s string) string {
	word := KebabCase(s)
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") ||
		strings.HasSuffix(word, "ch") || strings.HasSuffix(word, "sh") {
		return word + "es"
	}
	if len(word) > 1 && strings.HasSuffix(word, "y") && !strings.ContainsRune("aeiou", rune(word[len(word)-2])) {
		return word[:len(word)-1] + "ies"
	}
	return word + "s"
}

// SnakeString  XxYy to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// CamelString  xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// ObjectName  获取对象的类型名称
func ObjectName(i interface{}) string {
	v := reflect.ValueOf(i)
	t := v.Type()
	if t.Kind() == reflect.Func {
		return runtime.FuncForPC(v.Pointer()).Name()
	}
	return t.String()
}
