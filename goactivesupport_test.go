package goactivesupport

import (
  "fmt"
  "os"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
  os.Exit(m.Run())
}

func TestNumberToStrFast(t *testing.T) {
  testCases := []struct {
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
    {int64(2229212), "2229212"},
  }

  for _, c := range testCases {
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

func TestMethodName(t *testing.T) {
  f := func() {}
  assert.Equal(t, "github.com/ipoval/goactivesupport.TestMethodName.func1", MethodName(f))
}

func TestStringValidLuhnCreditCardNumber(t *testing.T) {
  testCases := []struct {
    description string
    input       string
    ok          bool
  }{
    {"single digit strings can not be valid", "1", false,},
    {"a single zero is invalid", "0", false,},
    {"a simple valid SIN that remains valid if reversed", "059", true,},
    {"a simple valid SIN that becomes invalid if reversed", "59", true,},
    {"a valid Canadian SIN", "055 444 285", true,},
    {"invalid Canadian SIN", "055 444 286", false,},
    {"invalid credit card", "8273 1232 7352 0569", false,},
    {"valid number with an even number of digits", "095 245 88", true,},
    {"valid number with an odd number of spaces", "234 567 891 234", true,},
    {"valid strings with a non-digit added at the end become invalid", "059a", false,},
    {"valid strings with punctuation included become invalid", "055-444-285", false,},
    {"valid strings with symbols included become invalid", "055# 444$ 285", false,},
    {"single zero with space is invalid", " 0", false,},
    {"more than a single zero is valid", "0000 0", true,},
    {"input digit 9 is correctly converted to output digit 9", "091", true,},
    {"using ascii value for non-doubled non-digit isn't allowed", "055b 444 285", false,},
    {"using ascii value for doubled non-digit isn't allowed", ":9", false,},
  }

  for _, test := range testCases {
    if ok := StringValidLuhnCreditCardNumber(test.input); ok != test.ok {
      t.Fatalf("Valid(%s): %s\n\t Expected: %t\n\t Got: %t", test.input, test.description, test.ok, ok)
    }
  }
}

func BenchmarkStringValidLuhnCreditCardNumber(b *testing.B) {
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    StringValidLuhnCreditCardNumber("2323 2005 7766 3554")
  }
}

func TestBcryptHashStr(t *testing.T) {
  word := "test"
  hash, err := BcyptHashStr(word)
  assert.NoError(t, err)
  assert.NotEqual(t, word, hash)
  assert.Len(t, hash, 60)
}
