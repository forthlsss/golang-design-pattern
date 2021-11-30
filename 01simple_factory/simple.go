/**
* @Author: Seanhuang
* @Date: 2021/11/30 3:37 下午
 */

package _1simple_factory

import "fmt"

type API interface {
 	Say(name string) string
 }

 type hiAPI struct {}

 type helloAPI struct {}

func ( *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func (receiver *helloAPI ) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func NewAPI(t int) API{
	if t==1{
		return &hiAPI{}
	}
	return &helloAPI{}
}
