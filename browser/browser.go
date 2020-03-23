package browser

import (
  "math/rand"
)

type userAgent string

var UserAgent userAgent

func (ua userAgent) Chrome() string {
  return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
}

func (ua userAgent) Random() string {
  const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
  result := make([]byte, rand.Intn(10)+10)
  for i := range result {
    result[i] = letters[rand.Intn(len(letters))]
  }
  return string(result)
}
