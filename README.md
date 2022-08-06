# farsi reshaper
Farsi reshaper written in go 

i just used http://bobardo.com/reshaper/ opensource javascript and extracted the core code and also used https://github.com/robertkrimen/otto 
package to import js into go.

the way it works is so simple just call:

```go
package main

import (
	"fmt"
	"github.com/javad-majidi/farsi-reshaper"
)

func main() {
	fmt.Println(farsi_reshaper.ReshapeToFarsi("سلام احوال شما؟"))
}

```
output(reshaped farsi):
```
؟ﺎﻤﺷ ﻝﺍﻮﺣﺍ ﻡﻼﺳ
```

it is extremely useful for some packages like [Fyne](https://github.com/fyne-io/fyne) or other desktop packages in `arabic` and `farsi` languages.
