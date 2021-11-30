/**
* @Author: Seanhuang
* @Date: 2021/11/30 4:08 下午
 */

package _2factoryMethod

//Operator是被封装的实际类接口
//BaseOperator 基类，实现了Operator接口需要的两个方法
//PlusOperator，MinusOperator 在BaseOperator基类上实现了Result方法，实现了Operator接口

//OperatorFactory是工厂接口
//PlusOperatorFactory，MinusOperatorFactory是实际工厂类，实现了Create方法

//PlusOperatorFactory，MinusOperatorFactory分别Create两种Operator

//可以创建一个factory，指定为PlusOperatorFactory还是MinusOperatorFactory，只需要知道工厂名即可




//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// BaseOperator  是Operator 接口实现的基类，封装公用方法
type BaseOperator struct {
	a, b int
}

func (receiver *BaseOperator) SetA(a int) {
	receiver.a = a
}
func (receiver *BaseOperator) SetB(b int) {
	receiver.b = b
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*BaseOperator
}

type MinusOperator struct {
	*BaseOperator
}

//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

//PlusOperatorFactory 是 PlusOperator 的工厂类, 该结构体实现了OperatorFactory接口
type PlusOperatorFactory struct {
}

func (p PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		BaseOperator: &BaseOperator{},
	}
}

func (p PlusOperator) Result() int {
	return p.b + p.a
}

type MinusOperatorFactory struct {
}

func (m MinusOperatorFactory) Create() Operator {
	return &MinusOperator{BaseOperator: &BaseOperator{}}
}

func (m MinusOperator) Result() int {
	return m.a - m.b
}
