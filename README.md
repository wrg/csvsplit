csvsplit - lightweight csv string splitter.

Version
=======

    csvsplit 1.0

Synopsis
========

Package csvsplit is a simple csv line parser that properly handles quotes
and nested quotes (i.e. 'text, "quoted", with commas').
This does not trim spaces or newline characters, and only returns string
values.  Conversion or manipulation is not assumed and left to the developer.
The field delimiter is a comma (,) and the quote characters are single and double
quotes.

```go
package main

import (
   "io"
   "bufio"
   "fmt"
   "os"
   "github.com/x86pgmer/csvsplit"
)

func main() {
   file, err := os.Open(os.Args[1])
   if err != nil { panic(err) }
   defer file.Close()
   r := bufio.NewReader(file)
   for {
       line, err := r.ReadString('\n')
       if err != nil && err != io.EOF { panic(err) }
       if line != "" {
         // remove newline chars yourself!
         if line[len(line)-1] == '\n' {
            line = line[:len(line)-1]
         }
         fields, err := csvsplit.Split(line)
         if err != nil {
            fmt.Println(err)
         }
         fmt.Println(fields)
       }
       if err == io.EOF { break }
   }
}
```

Given a CSV file 'example.csv' with contents

```
Bob,28,'text, "quoted", with commas',3/13/2013
Rick,44,"text, 'inverse' commas",3/12/2013
John,68,No Comment,1/1/1970
Jake,18,,2/2/2000
```

When the above program is called and given the path 'example.csv', the following
is printed to standard output.

```
[Bob 28 'text, "quoted", with commas' 3/13/2013]
[Rick 44 "text, 'inverse' commas" 3/12/2013]
[John 68 No Comment 1/1/1970]
[Jake 18  2/2/2000]
```

About
=====

    csvsplit is a lightweight string parser that will split the input text into
    an array of strings.  It is not intended to be a full blown file parser, and
    will does not do any type conversion.

Features
========

    * Properly handles quotes and nested quotes

    * Lightweight - no unwanted code imported into you application
Todo
====

   * Cannot change the field separator, would like to make that an option in a future release.

Install
=======

The easiest installation of csvsplit is done through goinstall.

    goinstall github.com/x86pgmer/csvsplit

Documentation
=============

The best way to read the current csvsplit documentation is using
godoc.

    godoc github.com/x86pgmer/csvsplit

Or better yet, you can run a godoc http server.

    godoc -http=":6060"

Then go to the url http://localhost:6060/pkg/github.com/x86pgmer/csvsplit/

Copyright & License
===================

Copyright (c) 2013, Rick Gibson.
All rights reserved.

Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.
