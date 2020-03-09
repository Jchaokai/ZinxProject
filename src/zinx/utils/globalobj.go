package utils

import (
	"ZinxProject/src/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

/*
	读取全局配置文件
*/

type GlobalObj struct {
	//关于server的配置
	TcpServer ziface.IServer
	Host      string `json:"Host"`
	Port      int    `json:"Port"`
	Name      string `json:"Name"`
	//关于zinx配置
	Version           string `json:"Version"`
	MaxConn           uint64 `json:"MaxConn"`
	MaxPackageSize    uint64 `json:"MaxPackageSize"`
	WorkPoolSize      uint64 `json:"WorkPoolSize"`      //Zinx 框架允许用户最多开辟多少个worker
	MaxWorkerTaskSize uint64 `json:"MaxWorkerTaskSize"` //每个worker 对应的消息队列的最大任务数
}

//定义一个全局对外 的GlobalObj
var GlobalObject *GlobalObj

//导入该包是自动执行
func init() {
	GlobalObject = &GlobalObj{
		Host:              "127.0.0.1",
		Port:              8999,
		Name:              "ZinxDemo",
		Version:           "V0.8",
		MaxConn:           10000,
		MaxPackageSize:    4096,
		WorkPoolSize:      4,
		MaxWorkerTaskSize: 1024,
	}
	//从 conf/zinx.json 加载用户自定义的参数
	GlobalObject.reload()
}

func (GlobalObject *GlobalObj) reload() {
	bytes, e := ioutil.ReadFile("F:\\GoProgram\\src\\ZinxProject\\src\\zinx\\zinx.json")
	if e != nil {
		panic(e)
	}
	if e := json.Unmarshal(bytes, &GlobalObject); e != nil {
		panic(e)
	}

}
