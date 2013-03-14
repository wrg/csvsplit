// Copyright 2013, Rick Gibson.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package csvsplit is a simple csv line parser that properly handles quotes
// and nested quotes (i.e. 'text, "quoted", with commas').
// This does not trim spaces or newline characters, and only returns string
// values.  Conversion or manipulation is not assumed and left to the developer.
// The field delimiter is a comma (,) and the quote characters are single and double
// quotes (' and ").
package csvsplit

import "errors"

var ErrQuote = errors.New("unmatched quote")
var ErrNull = errors.New("string value is nil")

type SplitError struct {
     Where string // where failed
     Text string // input text
     Err error // reason for failure
}

func (e *SplitError) Error() string {
     return "error: csvsplit." + e.Where + ": " + e.Text + ": " + e.Err.Error()
}

// error function for an unmatched quote character
func quoteError(w, text string) *SplitError {
   return &SplitError{w, text, ErrQuote}
}

// error function for an empty string
func nullError(w, text string) *SplitError {
   return &SplitError{w, text, ErrNull}
}

const startSize = 10
const doubleQuote = '"'
const singleQuote = '\''

// inQuote struct is used to determine if the parser is currently inside a quote, and remembers which
// quote char we are looking for to end the quote.
type inQuote struct {
    True bool
    Char uint8
}

const delim = ','

// Split parses a csv formatted string and returns a string array containing each field
// of the string along with any error information.
func Split(text string) (s []string, err error) {
   if len(text) == 0 { return nil, nullError("Split",text) }
   startpos :=0
   in_quote := inQuote{false,doubleQuote}
   result := make([]string, 0, startSize)
   for p := 0; p < len(text); p++ {
      if in_quote.True {
         if text[p] == in_quote.Char {
            in_quote.True = false
         }
      } else {
         if text[p] == delim {
            if text[p-1] == delim {
               result = append(result, "")
            } else {
               result = append(result, text[startpos:p])
            }
            startpos=p+1
         } else {
            if text[p] == doubleQuote || text[p] == singleQuote {
               in_quote.True = true
               in_quote.Char = text[p]
            }
         }
      }
   }
   if in_quote.True {
      return nil, quoteError("Split",text)
   }
   result = append(result, text[startpos:len(text)])
   return result, nil
}
