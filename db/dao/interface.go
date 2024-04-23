package dao

// CounterInterface 计数器数据模型接口
type CounterInterface interface {
	GetQiguaData() (string, error)
}

// CounterInterfaceImp 计数器数据模型实现
type CounterInterfaceImp struct{}

// Imp 实现实例
var Imp CounterInterface = &CounterInterfaceImp{}
