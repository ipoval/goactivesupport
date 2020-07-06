package concurrency

import (
  "testing"
)

var testCasesChanReduce = []struct {
  name       string
  reduceFunc func(...<-chan interface{}) <-chan interface{}
}{
  {
    name:       "chansReduceGoRoutines",
    reduceFunc: ChannelReduceMany,
  },
  {
    name:       "chansReduceGoReflection",
    reduceFunc: ChanReduceManyWithReflect,
  },
}

func TestChanReduceMany(t *testing.T) {
  for _, tc := range testCasesChanReduce {
    t.Run(tc.name, func(t *testing.T) {
      seen := make(map[int]bool)
      ch1, ch2, ch3 := ChannelWrap(1, 2, 3), ChannelWrap(4, 5, 6), ChannelWrap(7, 8, 9)

      mergedCh := tc.reduceFunc(ch1, ch2, ch3)

      for v := range mergedCh {
        if ok := seen[v.(int)]; ok {
          t.Errorf("saw %d at least twice", v)
        }
        seen[v.(int)] = true
      }

      for i := 1; i <= 9; i++ {
        if !seen[i] {
          t.Errorf("did not see %d", i)
        }
      }
    })
  }
}

func BenchmarkChanReduceMany(b *testing.B) {
  b.ResetTimer()

  for _, tc := range testCasesChanReduce {
    for n := 1; n < 1024; n *= 2 {
      chans := make([]<-chan interface{}, n)

      b.Run("test case " + tc.name, func(b *testing.B) {
        for i := 0; i < b.N; i++ {
          for i := range chans {
            chans[i] = ChannelWrap(1, 2, 3, 4, 5)
          }

          merged := tc.reduceFunc(chans...)
          for range merged {
          }
        }
      })
    }
  }
}
