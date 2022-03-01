package ureg_test

import (
	"github.com/sun-fight/tools/utils/ureg"
	"testing"
)

func TestRegNickname(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "测试空格",
			input:    "a  ",
			expected: false,
		},
		{
			name:     "测试中文",
			input:    "  中文",
			expected: false,
		},
		{
			name:     "测试英文",
			input:    "  wer",
			expected: false,
		},
		{
			name:     "测试长度",
			input:    "a1234567891111111",
			expected: false,
		},
		{
			name:     "测试长度",
			input:    "1234567899876543",
			expected: true,
		},
		{
			name:     "测试正常",
			input:    "aws",
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := ureg.Username(c.input)
			if e, a := c.expected, actual; e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
