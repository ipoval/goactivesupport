package convert_test

import (
  "testing"

  . "github.com/ipoval/goactivesupport/convert"
)

func TestTypeBoolean_Cast(t1 *testing.T) {
  type fields struct {
    ValueAsStr string
  }
  tests := []struct {
    name   string
    fields fields
    want   bool
  }{
    {
      "value false",
      fields{
        ValueAsStr: "false",
      },
      false,
    },
    {
      "value f",
      fields{
        ValueAsStr: "false",
      },
      false,
    },
    {
      "value 0",
      fields{
        ValueAsStr: "false",
      },
      false,
    },
    {
      "value 1",
      fields{
        ValueAsStr: "1",
      },
      true,
    },
    {
      "value true",
      fields{
        ValueAsStr: "true",
      },
      true,
    },
    {
      "value -1",
      fields{
        ValueAsStr: "-1",
      },
      true,
    },
    {
      "value random string",
      fields{
        ValueAsStr: "abc",
      },
      true,
    },
  }
  for _, tt := range tests {
    t1.Run(tt.name, func(t1 *testing.T) {
      t := &TypeBoolean{
        ValueAsStr: tt.fields.ValueAsStr,
      }
      if got := t.Cast(); got != tt.want {
        t1.Errorf("Cast() = %v, want %v", got, tt.want)
      }
    })
  }
}
