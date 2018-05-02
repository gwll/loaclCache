### 实际项目中很多时候缓存数据不用持久化和跨系统通信 所以我把这些功能封装了一下，多线程是安全的
>package main

>import (
	
>	"github.com/gwll/loaclCache"

>	"fmt"
>)
>
>func main() {
	
>	loaclCache.Set("keyx", 11, 10)

>	vx, _ := loaclCache.Get("keyx")

>	vxs, ok := ii.(string)

>	if ok {
	
>		fmt.Println(vxs)

>	}

>
>}
>
>
>
>
