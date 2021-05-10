# Date Utils

**_dateutils_** is a datetime function utility for golang

an example :

```go
package main

import(
   "github/derit/dateutils"
   "fmt"
)

func main(){ 
  now:= dateutils.Now()
  fmt.Println(now)
}

```

diff between two date

```go
package main

import(
   "github/derit/dateutils"
   "fmt"
   "date"
)

func main(){ 
  from:= dateutils.Now()
  to:=dateutils.AddNextDay(from)
  v := DateDiff(from, to)
  fmt.Println(now)
}

```
