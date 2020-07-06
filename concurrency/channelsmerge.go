// Merge 2 channels - https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308
// Merge n channesl - https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de
// https://youtu.be/2AulMm-hsdI

package concurrency

import (
  "fmt"
  "reflect"
  "sync"
)

// Wrap values into a channel
func ChannelWrap(payload ...interface{}) <-chan interface{} {
  c := make(chan interface{})

  go func() {
    for _, v := range payload {
      c <- v
    }
    close(c)
  }()

  return c
}

func ChannelReducePair(a, b <-chan interface{}) <-chan interface{} {
  merged := make(chan interface{})

  go func() {
    defer close(merged)

    for a != nil || b != nil {
      select {
      case v, ok := <-a:
        if !ok {
          a = nil // ! important to set to nil to make it a nil channel and avoid trying to select from closed channel
          fmt.Println("1st channel is empty")
          continue
        }
        merged <- v
      case v, ok := <-b:
        if !ok {
          b = nil // ! important to set to nil
          fmt.Println("2nd channel is empty")
          continue
        }
        merged <- v
      }
    }
  }()

  return merged
}

// How to use
//  a := ChannelWrap(1, 2, 3)
//  b := ChannelWrap(4, 5, 6)
//  c := ChannelWrap(7, 8, 9)
//  // d := ChannelReducePair(a, b)
//  d := ChannelReduceMany(a, b, c)
//  for v := range d {
//    fmt.Println(v)
//  }
func ChannelReduceMany(chans ...<-chan interface{}) <-chan interface{} {
  merged := make(chan interface{})

  go func() {
    var wg sync.WaitGroup
    wg.Add(len(chans))

    for _, ch := range chans {
      go func(ch <-chan interface{}) {
        for v := range ch {
          merged <- v
        }
        wg.Done()
      }(ch)
    }

    wg.Wait()
    close(merged)
  }()

  return merged
}

func ChanReduceManyWithReflect(chans ...<-chan interface{}) <-chan interface{} {
  mergeChan := make(chan interface{})

  go func() {
    defer close(mergeChan)

    reflectSelectCases := []reflect.SelectCase{}
    for _, ch := range chans {
      reflectSelectCases = append(reflectSelectCases, reflect.SelectCase{
        Dir:  reflect.SelectRecv,
        Chan: reflect.ValueOf(ch),
        // Send: reflect.Value{},
      })
    }

    for len(reflectSelectCases) > 0 {
      idx, val, recvOk := reflect.Select(reflectSelectCases)
      if !recvOk {
        reflectSelectCases = append(reflectSelectCases[:idx], reflectSelectCases[idx+1:]...)
        continue
      }
      mergeChan <- val.Interface()
    }
  }()

  return mergeChan
}
