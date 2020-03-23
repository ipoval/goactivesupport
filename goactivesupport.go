package goactivesupport

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "strconv"
  "strings"
  "unicode"
)

// SCOPE: OS
// Return default value if given environment variable is blank
func GetenvDv(k string, dv interface{}) (interface{}, error) {
  if v := os.Getenv(k); v != "" {
    return v, nil
  }
  return dv, nil
}

// SCOPE: LOGGING
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

// SCOPE: FILES
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

// SCOPE: STRINGS
func IsLetter(c rune) bool {
  return unicode.IsLetter(c)
}

// SCOPE: NUMBERS
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
