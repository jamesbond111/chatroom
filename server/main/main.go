package main
import (
	"fmt"
	"net"
	"time"
	"chatroom/server/model"
)


func init() {
	//当服务器启动时，我们就去初始化我们的redis的连接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
}

func initUserDao() {
	//这里需要注意一个初始化顺序问题
	//先initPool, 再initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return 
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来链接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=" ,err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯。。
		go process(conn)
	}
}

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	processor := &Processor{
		Conn : conn,
	}
	//调用总控
	err := processor.coreProcess()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}