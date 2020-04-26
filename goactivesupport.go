package goactivesupport

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "reflect"
  "strconv"
  "strings"
  "unicode"
)

// Scope: OS
// Return default value if given environment variable is blank
func GetenvDv(k string, dv interface{}) (interface{}, error) {
  if v := os.Getenv(k); v != "" {
    return v, nil
  }
  return dv, nil
}

// Scope: Logging
func GetFileLogger(path string) *log.Logger {
  file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
  defer file.Close()

  if err != nil {
    log.Panicln("failed to open "+path, err.Error())
    return nil
  }

  var result *log.Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  return result
}

// Scope: Files
const (
  exitStatusCodeSuccess    = iota // 0
  exitStatusCodeError             // 1
  exitStatusCodeUnexpected        // 2
)

func WriteFile(path, msg string) {
  ioutil.WriteFile(path, []byte(msg), 0644)
}

func ReadFile(path string) string {
  bytes, _ := ioutil.ReadFile(path)
  return string(bytes)
}

func FileLineCount(filepath string) (int, error) {
  out, err := exec.Command("wc", "-l", filepath).Output()
  if err != nil {
    return 0, fmt.Errorf("could not run wc -l: %s", err)
  }

  // NOTE: \d\d\d filename.go
  count, err := strconv.Atoi(strings.Fields(string(out))[0])
  if err != nil {
    return 0, fmt.Errorf("could not convert wc -l output to integer: %s", err)
  }
  return count, nil
}

// Scope: Strings
func IsLetter(c rune) bool {
  return unicode.IsLetter(c)
}

func StringValidLuhnCreditCardNumber(str string) bool {
  str = strings.ReplaceAll(str, " ", "")
  l := len(str)
  if l < 2 {
    return false
  }

  sum := 0
  chars := strings.Split(str, "")

  for i, i2 := l-1, l-2; i > 0 || i2 > 0; i, i2 = i-2, i2-2 {
    sym, sym2 := chars[i], "0"
    if i2 >= 0 {
      sym2 = chars[i2]
    }

    digit, err := strconv.Atoi(sym)
    digit2, err2 := strconv.Atoi(sym2)
    if err != nil || err2 != nil {
      return false
    }

    if digit2 *= 2; digit2 > 9 {
      digit2 -= 9
    }

    sum += (digit + digit2)
  }

  return sum%10 == 0
}

// Scope: Numbers
const (
  _  = iota // ignore
  KB = 1 << (10 * iota)
  MB
  GB
  TB
  PG
  EB
  ZB
  YB
)

func NumberToHumanSize(n float64) string {
  return fmt.Sprintf("%.2fGB", n/GB)
}

func NumberToStrFast(n interface{}) string {
  v := reflect.ValueOf(n)
  switch v.Kind() {
  case reflect.Float32, reflect.Float64:
    return strconv.FormatFloat(v.Float(), 'f', -1, 64)
  case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64:
    return strconv.FormatInt(v.Int(), 10)
  default:
    return "nil"
  }
}
