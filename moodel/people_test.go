package model

import (
	"fmt"
	"testing"
)

// testmain函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	//通过m.Run()来执行子测试函数
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	// 通过t.run()来执行子测试函数
	t.Run("测试获取用户：", testGetUserByname)
	t.Run("测试获取用户：", testGetUsers)
}

func testAddUser(t *testing.T) {

	fmt.Println("子测试添加用户：")
	//people := &People{}
	// 调用添加用户的方法
	//people.AddPeople()
	//people.AddPeople2()
}
func testGetUserByname(t *testing.T) {
	fmt.Println("测试查询一个记录：")
	user := People{
		Name: "bob",
	}
	// 调用获取User方法
	u, _ := user.GetUserByname()
	fmt.Println("得到的信息是：", u)
}

// c测试获取所有user
func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录：")
	user := &People{}
	// 调用获取User方法
	us, _ := user.GetUsers()
	// 遍历切片
	for k, v := range us {
		fmt.Printf("第%v个用户是%v：\n", k+1, v)

	}

}
