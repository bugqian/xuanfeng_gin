package request

type UserCreateReq struct {
	Name string `json:"name" binding:"required"` // 姓名
	Age  int64  `json:"age" binding:"required"`  // 年龄
}
