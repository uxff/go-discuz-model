discuz model for go

use xorm model

at first you must install xorm cmd tool:
```
go get github.com/go-xorm/cmd/xorm
```

then you must have a discuz database, a available datasource

generate your discuz models:
```
xorm reverse mysql 'root:123456@tcp(192.168.38.134:3306)/discuz_co7?charset=utf8' x3.4/templates/goxorm/ x3.4/models

```




