/**
* @Author: Seanhuang
* @Date: 2021/11/30 4:56 下午
 */

//抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。
//如果抽象工厂退化成生成的对象无关联则成为工厂函数模式。
//比如本例子中使用RDB和XML存储订单信息，抽象工厂分别能生成相关的主订单信息和订单详情信息。
//如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换使用不同的存储方式了。

//means 主订单信息和订单详情信息是同一单的信息，例如都是电视的，都是电脑的



package _3abstract_factory

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

//DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO	//嵌套了子接口  创建订单-主要数据对象
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct {
}

func (receiver *RDBMainDAO) SaveOrderMain()  {
	fmt.Print("rdb main save\n")
}

type RDBDetailDAO struct {}

func (r *RDBDetailDAO) SaveOrderDetail()  {
	fmt.Print("rdb detail save\n")
}

type RDBDAOFactory struct {
}

func (r *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (r *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

//XMLMainDAO XML存储
type XMLMainDAO struct{}

//SaveOrderMain ...
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

//XMLDetailDAO XML存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

//XMLDAOFactory 是RDB 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
