package main

import "os"
import "io"
import "bufio"
import "strings"
import "bytes"
import "fmt"

type WinterWonderland struct {
  input *bufio.Reader
  current *bytes.Buffer
}

func NewWinterWonderland(r io.Reader) (*WinterWonderland) {
  this := new(WinterWonderland)
  this.input = bufio.NewReader(r)
  this.current = new(bytes.Buffer)
  return this
}

func (this WinterWonderland) Read(p []byte) (int, error) {
  written := 0 
  for written < len(p) {
    if this.current.Len() == 0 {
      err := this.readAndTransformOneLine()
      if err != nil {
        return written, err
      }
    }

    n, err := this.current.Read(p[written:])
    written += n 
    if err != nil {
      return written, err
    }
  }

  return written, nil
}

func (this WinterWonderland) readAndTransformOneLine() error {
  line, err := this.input.ReadString('\n')
  if err != nil && len(line) == 0 {
    return err 
  }
  line_parts := strings.Split(line, " ")
  
  var matched bool
  for i, line_part := range line_parts {
    if i == len(line_parts) -1 && line_part[len(line_part) - 1] == '\n' {
      line_part = line_part[:len(line_part) - 1]
    }

    if !matched && line_part == "snowman" { 
      matched = true
      line_part = "\u2603"
    } 

    _, err := fmt.Fprint(this.current, line_part)
    if err != nil {
      return err
    }

    if i != len(line_parts) - 1 {
        _, err := fmt.Fprint(this.current, " ")
      if err != nil {
        return err
      }
    }
  }

  _, err = fmt.Fprint(this.current, "\n")
  if err != nil {
    return err
  }
  return nil
}

func main() {
  var err error
  // Read input from standard in, write output to standard out
  _, err = io.Copy(os.Stdout, NewWinterWonderland(os.Stdin))
  if err != nil {
    panic(err)
  }

}
