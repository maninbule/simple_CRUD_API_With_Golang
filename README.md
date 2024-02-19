# simple_CRUD_API_With_Golang
一个简单的golang增删改查练习

功能：对电影数据进行了增删改查
## api
```
get /movies 获取全部书籍

get /movies/id 获取指定id的书

post /movies 创建一本书

put /movies/id 修改指定id的书

delete /movies/id 删除指定id的书
```
## 实现
```
本项目没有使用数据库，图书的数据存储在内存，以切片形式存在

api返回和接受的都是json数据

路由选择的是"github.com/gorilla/mux"
```
