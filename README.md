# BRouter Documentation

BRouter is light weight, fast and strict HTTP router specialfically design to build REST API in Go language. It has less features and flexiblity as compared to other HTTP router, which helps to reduce the resource requirements and boost the performance. It use redis tree to efficiently compare URL paths. It also implements the Handler interface of net/http module, which makes it suitable to use with the http package.

## structure

It has two different kind of connection between each router node.

- Router to Router
- Router to Handler

This structure will helps us to create multi layer routing easily.

It supports only four kind of methods - `GET`, `POST`, `PUT`, `DELETE`. These four methods are enough to build any complex API, that's why all other methos are eliminated.

## Installation

```bash
go get github.com/SubhoBasak/brouter
```

## Import

```go
import "github.com/SubhoBasak/brouter"
```

## Create a router - BRouter

BRouter is a structure. Under the hood it is used as the node of redis tree.

```go
// create a router
router := brouter.Router{}
```

## GET

Add GET handler to the router

```go
// add GET handler to the rouer
router.GET("/authors", getAuthors)
router.GET("/blogs", getBlogs)
router.GET("/products", getProducts)
```

## POST

Add POST handler to the router

```go
// add POST handler to the router
router.POST("/author", newAuthor)
router.POST("/blog", createBlog)
router.POST("/product", addProduct)
```

## PUT

Add PUT handler to the router

```go
// add PUT handler to the router
router.PUT("/author", updateAuthor)
router.PUT("/blog", editBlog)
router.PUT("/product", editProduct)
```

## DELETE

Add DELETE handler to the router

```go
// add DELETE handler to the router
router.DELETE("/author", removeAuthor)
router.DELETE("/blog", deleteBlog)
router.DELETE("/product", deleteProduct)
```

## ROUTER & SUB-ROUTER

Add sub-routers to a parent router

```go
// add another router to the current router
schoolRouter = brouter.Router{}

// student router
studentRouter = brouter.Router{}

studentRouter.GET("/homework", allHomeWorks)
studentRouter.POST("/homework", submitHomeWork)
studentRouter.GET("/notice", allNotice)
studentRouter.PUT("/profile", updateProfile)

// teacher router
teacherRouter = brouter.Router{}

teacherRouter.GET("/homework", viewHomeWorks)
teacherRouter.POST("/homework", postNewHomeWork)
teacherRouter.PUT("/homework", editHomeWork)
teacherRouter.DELETE("/homework", deleteHomeWork)
teacherRouter.POST("/report", postStudentReport)

// parent router
parentRouter = Router{}

parentRouter.GET("/report", viewStudentReport)
parentRouter.GET("/fees", getFees)

// add sub routers to the main router
router.Router("/student", &studentRouter)
router.Router("/teacher", &teacherRouter)
router.Router("/parent", &parentRouter)
```

## Add router to the net/http server

```go
router := brouter.Router{}

// add handlers and sub routers
// ...

// add the root router to the http server
http.ListenAndServe("127.0.0.1:5000", &router)
```
