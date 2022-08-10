# farsi reshaper
Farsi reshaper written in go 

i just used http://bobardo.com/reshaper/ opensource javascript and extracted the core code and also used https://github.com/robertkrimen/otto 
package to import js into go.

the way it works is so simple just run the following in the terminal to get the package:
```
go get github.com/javad-majidi/farsi-reshaper@v1.0.1
```
then write this sample code:
```go
package main

import (
	"fmt"
	"github.com/javad-majidi/farsi-reshaper"
)

func main() {
	fmt.Println(FarsiReshaper.Reshape("سلام احوال شما؟"))
}

```
output(reshaped farsi):
```
؟ﺎﻤﺷ ﻝﺍﻮﺣﺍ ﻡﻼﺳ
```

it is extremely useful for some packages like [Fyne](https://github.com/fyne-io/fyne) or other desktop packages in `arabic` and `farsi` languages.

It also has some unit tests and code coverage.
