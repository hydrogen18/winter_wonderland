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
  // Track the number of bytes written into p
  written := 0 
  // Process until p is full
  for written < len(p) {
    // If no data is in the buffer, populate the buffer
    if this.current.Len() == 0 {
      err := this.readAndTransformOneLine()
      if err != nil {
        return written, err
      }
    }

    //Copy the buffer into p
    n, err := this.current.Read(p[written:])
    written += n 
    if err != nil {
      return written, err
    }
  }

  return written, nil
}

func (this WinterWonderland) readAndTransformOneLine() error {
  // Read a single line
  line, err := this.input.ReadString('\n')
  // If err is not nil, but data was returned, process the line
  if err != nil && len(line) == 0 {
    return err 
  }

  // Split the line up into words
  line_parts := strings.Split(line, " ")
  
  // Track if we matched "snowmna" this line already
  var matched bool
  var strippedNewline bool
  for i, line_part := range line_parts {
    // If this is the last element in the line and it ends with the newline, strip the newline
    if i == len(line_parts) -1 && line_part[len(line_part) - 1] == '\n' {
      line_part = line_part[:len(line_part) - 1]
      strippedNewline = true
    }

    // Convert snowman to the unicode glyph if it has not been done already in this line
    if !matched && line_part == "snowman" { 
      matched = true
      line_part = "\u2603"
    } 

    // Write this part of the line out.
    _, err := fmt.Fprint(this.current, line_part)
    if err != nil {
      return err
    }

    // If this is not the last part of the line, write a space character
    if i != len(line_parts) - 1 {
        _, err := fmt.Fprint(this.current, " ")
      if err != nil {
        return err
      }
    }
  }

  // Write a newline character, if it was stripped above
  if  strippedNewline {
    _, err = fmt.Fprint(this.current, "\n")
    if err != nil {
      return err
    }
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
