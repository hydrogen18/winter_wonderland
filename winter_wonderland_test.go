package main 

import "errors"
import "bytes"
import "strings"
import "testing"
import "io"
import . "gopkg.in/check.v1" 

type TestSuite struct{
}

func Test(t *testing.T){ TestingT(t) }

var _ = Suite(&TestSuite{})

//Expected cases

func (s *TestSuite) TestDocumentWithSnowman(c *C) {
  var err error 
  buf := new(bytes.Buffer)

  _, err = buf.WriteString("Hello world\n")
  c.Assert(err,IsNil)
  _, err = buf.WriteString("Do you know what a snowman is?\n")
  c.Assert(err,IsNil) 
  _, err = buf.WriteString("Hello bob\n")
  c.Assert(err,IsNil)

  reader := NewWinterWonderland(buf)

  output := new(bytes.Buffer)
  _, err = io.Copy(output, reader)
  c.Assert(err, IsNil)
  
  c.Assert(strings.Count(output.String(), "\u2603"), Equals, 1)
}

func (s *TestSuite) TestDocumentWithSnowmanTwice(c *C) {
  var err error 
  buf := new(bytes.Buffer)

  _, err = buf.WriteString("Hello world\n")
  c.Assert(err,IsNil)
  _, err = buf.WriteString("Do you know what a snowman snowman is?\n")
  c.Assert(err,IsNil) 
  _, err = buf.WriteString("Hello bob\n")
  c.Assert(err,IsNil)

  reader := NewWinterWonderland(buf)

  output := new(bytes.Buffer)
  _, err = io.Copy(output, reader)
  c.Assert(err, IsNil)
  
  c.Assert(strings.Count(output.String(), "\u2603"), Equals, 1)
}

func (s *TestSuite) TestDocumentWithoutSnowman(c *C) {
  var err error 
  buf := new(bytes.Buffer)

  _, err = buf.WriteString("Hello world\n")
  c.Assert(err,IsNil)
  _, err = buf.WriteString("Do you know what a yeti is?\n")
  c.Assert(err,IsNil) 
  _, err = buf.WriteString("Hello bob\n")
  c.Assert(err,IsNil)

  reader := NewWinterWonderland(buf)

  output := new(bytes.Buffer)
  _, err = io.Copy(output, reader)
  c.Assert(err, IsNil)
  
  c.Assert(strings.Count(output.String(), "\u2603"), Equals, 0)
}

//Edge cases
func (s *TestSuite) TestZeroLengthDocument(c *C){
 var err error 
  buf := new(bytes.Buffer)

  reader := NewWinterWonderland(buf)

  output := new(bytes.Buffer)
  _, err = io.Copy(output, reader)
  c.Assert(err, IsNil)
  
  c.Assert(buf.Len(), Equals, 0)
}

func (s *TestSuite) TestDocumentWithoutNewline(c *C){
 var err error 
  buf := new(bytes.Buffer)

  _, err = buf.WriteString("Hello snowman")
  c.Assert(err,IsNil)

  reader := NewWinterWonderland(buf)

  output := new(bytes.Buffer)
  _, err = io.Copy(output, reader)
  c.Assert(err, IsNil)
  
  c.Assert(strings.Count(output.String(), "\u2603"), Equals, 1)
}

func (s *TestSuite) TestDocumentWithWindowsNewlines(c *C){
 var err error 
  buf := new(bytes.Buffer)

  _, err = buf.WriteString("Hello snowman\r\nHello bob")
  c.Assert(err,IsNil)

  reader := NewWinterWonderland(buf)

  output := new(bytes.Buffer)
  _, err = io.Copy(output, reader)
  c.Assert(err, IsNil)
  
  // Windows newline throws off the program, don't care about supporting this
  c.Assert(strings.Count(output.String(), "\u2603"), Equals, 0)

}

// Failure cases

type failureReader int //Type does not matter here

var failureReaderSentinel = errors.New("Sentinel from failure reader")

func (failureReader) Read([]byte) (int, error) {
  return 0, failureReaderSentinel
}

func (s *TestSuite) TestReaderFailure(c *C) {
  reader := NewWinterWonderland(failureReader(0))
  var output [1]byte
  _, err := reader.Read(output[:])
  c.Assert(err, Equals, failureReaderSentinel)
}

