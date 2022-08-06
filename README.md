# farsi reshaper
Farsi reshaper written in go 

i just used http://bobardo.com/reshaper/ opensource javascript and extracted the core code and also used https://github.com/robertkrimen/otto 
package to import js into go.

the way it works is so simple just call:

```go
fmt.Println(farsi_reshaper.ReshapeToFarsi("سلام احوال شما؟"))
```
output:
```
؟ﺎﻤﺷ ﻝﺍﻮﺣﺍ ﻡﻼﺳ
```
