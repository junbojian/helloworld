package main

import (
	"fmt"
	"time"
	"reflect"
)

func main()  {
	rt := time.Now()

	fmt.Println("Hello")
	fmt.Println(reflect.TypeOf(rt))

	joel := user{1,"joel","3125","j",time.Now()}
	fmt.Println(joel)
	fmt.Println(joel.regtime.Format("2006-01-02 15:04:05.0000000"))
}

//用户类型
type user struct {
	userid		int			//系统内部id
	username	string		//用户可见id
	password	string		//用户密码
	nickname	string		//用户昵称（无昵称时显示username）
	regtime		time.Time	//注册时间
}
