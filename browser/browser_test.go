package browser

import (
  "strings"
  "testing"
)

const (
  checkMark = "\u2713"
  ballotX   = "\u2717"
)

func TestUserAgent(t *testing.T) {
  t.Log("Given pre-defined UserAgent")
  {
    ua := UserAgent.Chrome()
    if !strings.HasPrefix(ua, "Mozilla/5.0") {
      t.Fatal("\tShould be able to access Chrome browser-like looking UA", ballotX)
    }
    t.Log("\tShould be able to access Chrome browser-like looking UA", checkMark)

    uaRand1 := UserAgent.Random()
    uaRand2 := UserAgent.Random()

    if len(uaRand1) < 10 || len(uaRand1) > 19 || uaRand1 == uaRand2 {
      t.Fatal("\tShould be able to access Random string looking UA", ballotX)
    }
    t.Log("\tShould be able to access Random string looking UA", checkMark)
  }
}
