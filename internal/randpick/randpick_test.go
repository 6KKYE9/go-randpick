package randpick

import (
	"reflect"
	"testing"
)

func TestPickBasic(t *testing.T) {
	items := []string{"a", "b", "c"}
	picked, err := Pick(items, 2, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(picked) != 2 {
		t.Fatalf("应挑出 2 项，实际 %d", len(picked))
	}
}

func TestPickNoDup(t *testing.T) {
	items := []string{"a", "b", "c"}
	picked, err := Pick(items, 3, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(picked) != 3 {
		t.Fatalf("应挑出 3 项")
	}
	// 三项全取时应是原序列的一个排列
	m := map[string]bool{}
	for _, v := range picked {
		m[v] = true
	}
	if len(m) != 3 {
		t.Fatal("去重模式下不应有重复")
	}
}

func TestPickNoDupExceeds(t *testing.T) {
	items := []string{"a", "b"}
	if _, err := Pick(items, 3, true); err == nil {
		t.Fatal("去重且数量超限应报错")
	}
}

func TestEmptyList(t *testing.T) {
	var items []string
	if _, err := Pick(items, 1, false); err == nil {
		t.Fatal("空列表应报错")
	}
}

func TestZeroCount(t *testing.T) {
	if _, err := Pick([]string{"a"}, 0, false); err == nil {
		t.Fatal("0 应报错")
	}
	if _, err := Pick([]string{"a"}, -1, false); err == nil {
		t.Fatal("负数应报错")
	}
}

// 去重挑 3 次，20 轮后不应该全一样
func TestRandomness(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	sets := map[string]int{}
	for i := 0; i < 20; i++ {
		p, err := Pick(items, 3, true)
		if err != nil {
			t.Fatal(err)
		}
		key := ""
		for _, v := range p {
			key += string(rune('0' + v))
		}
		sets[key]++
	}
	if len(sets) < 2 {
		t.Fatalf("20 轮随机挑 3 应出现多种排列，实际只有 %d 种", len(sets))
	}
}

func TestLines(t *testing.T) {
	text := "alpha\nbeta\ngamma\n\n"
	lines, err := Lines(text, 2, true, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("应抽出 2 行，实际 %d", len(lines))
	}
	// 空行已跳过
	for _, l := range lines {
		if l == "" {
			t.Fatal("不应有空行")
		}
	}
}

func TestLinesWithCRLF(t *testing.T) {
	text := "a\r\nb\r\nc\r\n"
	lines, err := Lines(text, 3, true, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 3 {
		t.Fatalf("应抽出 3 行，实际 %d", len(lines))
	}
}

func TestLinesSorted(t *testing.T) {
	text := "c\na\nb\n"
	lines, err := Lines(text, 3, true, true)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(lines, want) {
		t.Fatalf("排序不对: %v", lines)
	}
}

func TestPickWithDup(t *testing.T) {
	items := []string{"x", "y"}
	picked, err := Pick(items, 5, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(picked) != 5 {
		t.Fatalf("允许重复时也应挑出指定数量，实际 %d", len(picked))
	}
}

func TestGenericTypes(t *testing.T) {
	// 确保泛型在 int 上也能用
	ints := []int{10, 20, 30}
	picked, err := Pick(ints, 2, true)
	if err != nil {
		t.Fatal(err)
	}
	if len(picked) != 2 {
		t.Fatal("泛型 int 挑错了")
	}
}
