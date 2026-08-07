package randpick

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Pick 从切片里随机挑出 n 项。
// noDup 为真时不重复挑——同一项不会被选两次。
func Pick[T any](items []T, n int, noDup bool) ([]T, error) {
	if len(items) == 0 {
		return nil, errors.New("列表为空")
	}
	if n <= 0 {
		return nil, errors.New("挑选数量必须大于 0")
	}
	if noDup && n > len(items) {
		return nil, errors.New("去重模式下挑选数量不能超过总数")
	}
	if !noDup && n > len(items) {
		// 允许挑更多，只是会重复
	}

	if noDup {
		return pickNoDup(items, n)
	}
	return pickWithDup(items, n)
}

func pickNoDup[T any](items []T, n int) ([]T, error) {
	// Fisher-Yates 洗牌取前 n 项
	cp := make([]T, len(items))
	copy(cp, items)
	for i := len(cp) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return nil, err
		}
		cp[i], cp[j.Int64()] = cp[j.Int64()], cp[i]
	}
	return cp[:n], nil
}

func pickWithDup[T any](items []T, n int) ([]T, error) {
	out := make([]T, n)
	for i := 0; i < n; i++ {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(len(items))))
		if err != nil {
			return nil, err
		}
		out[i] = items[j.Int64()]
	}
	return out, nil
}

// Lines 从多行字符串里随机挑 N 行，保持原顺序可选。
// 空行自动跳过。
func Lines(text string, n int, noDup bool, sorted bool) ([]string, error) {
	raw := stringsSplit(text)
	var lines []string
	for _, l := range raw {
		if l != "" {
			lines = append(lines, l)
		}
	}
	picked, err := Pick(lines, n, noDup)
	if err != nil {
		return nil, err
	}
	if sorted {
		SortLines(picked)
	}
	return picked, nil
}

// 不能叫 strings.Split，因为要处理 \r\n
func stringsSplit(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			line := s[start:i]
			// 去掉末尾的 \r
			if len(line) > 0 && line[len(line)-1] == '\r' {
				line = line[:len(line)-1]
			}
			lines = append(lines, line)
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

// SortLines 排序输出。
func SortLines(lines []string) {
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i] > lines[j] {
				lines[i], lines[j] = lines[j], lines[i]
			}
		}
	}
}
