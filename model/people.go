package model

import (
	"cloud/utils"
	"fmt"
)

//  结构体
type People struct {
	Name     string
	Num      string
	Picture  string
	Keywords string
}

// AddUser 添加User
func (people *People) AddPeople(p *People) error {

	//1.写sql语句
	sqlStr := "insert into quiz1(name,num,picture,keywords) values(?,?,?,?)"

	// 2.执行
	_, err := utils.Db.Exec(sqlStr, p.Name, p.Num, p.Picture, p.Keywords)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

// 删除User
func (people *People) DeletePeople(p *People) error {

	//1.写sql语句
	sqlStr := "delete from quiz1 where name =?"

	// 2.执行
	_, err := utils.Db.Exec(sqlStr, p.Name)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//  根据name查询一条记录
func (people *People) GetUserByname() (*People, error) {
	//sql语句
	sqlSter := "select name,num,picture,keywords from quiz1 where name = ?"
	//执行
	row := utils.Db.QueryRow(sqlSter, people.Name)
	//声明
	var name string
	var num string
	var picture string
	var keywords string
	err := row.Scan(&name, &num, &picture, &keywords)
	if err != nil {
		return nil, err
	}
	u := &People{
		Name:     name,
		Num:      num,
		Picture:  picture,
		Keywords: keywords,
	}
	return u, nil
}

// GetUsers 根据id查询所有记录
func (people *People) GetUsers() ([]*People, error) {
	//sql语句
	sqlSter := "select name,num,picture,keywords from quiz1"
	//执行
	rows, err := utils.Db.Query(sqlSter)
	if err != nil {
		return nil, err
	}
	//创建user切片
	var users []*People
	for rows.Next() {
		//声明
		var name string
		var num string
		var picture string
		var keywords string
		err := rows.Scan(&name, &num, &picture, &keywords)
		if err != nil {
			return nil, err
		}
		u := &People{
			Name:     name,
			Num:      num,
			Picture:  picture,
			Keywords: keywords,
		}
		users = append(users, u)
	}
	return users, nil
}
