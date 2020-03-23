package goactivesupport

import (
  "fmt"
  "testing"
)

func TestNumberToStrFast(t *testing.T) {
  table := []struct {
    given interface{}
    want  string
  }{
    {0, "0"},
    {101, "101"},
    {-202, "-202"},
    {0.0, "0"},
    {0.1, "0.1"},
    {101.101, "101.101"},
    {-202.202, "-202.202"},
  }

  for _, c := range table {
    result := NumberToStrFast(c.given)
    if result != c.want {
      t.Errorf("mismatch given: %v, want: %v, result: %v", c.given, c.want, result)
    }
  }
}

func BenchmarkNumberToStrFast(b *testing.B) {
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    NumberToStrFast(11111)
  }
}

func BenchmarkSprintf(b *testing.B) {
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    fmt.Sprintf("%d", 11111)
  }
}
