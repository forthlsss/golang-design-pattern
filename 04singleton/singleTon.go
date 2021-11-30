/**
* @Author: Seanhuang
* @Date: 2021/11/30 5:52 下午
 */

package _4singleton

import "sync"

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton{
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}


