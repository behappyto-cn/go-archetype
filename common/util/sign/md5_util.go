package signUtil

import (
	"crypto/md5"
	"fmt"
	"io"
)

// GetMd5ByStr 方式一
func GetMd5ByStr(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return str
	}
	arr := m.Sum(nil)

	return fmt.Sprintf("%x", arr)
}
