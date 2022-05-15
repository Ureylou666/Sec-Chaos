package v1

import (
	"Backend/Model"
	"Backend/Utils/ErrMsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 增
func CategoryAdd(c *gin.Context) {
	var data Model.Category
	_ = c.ShouldBindJSON(&data)
	code := Model.CheckCategory(data.Name)
	if code > 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_CATEGORY_EXIST,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_EXIST),
		})
	} else {
		status := Model.AddCategory(&data)
		if status == ErrMsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"Status":  ErrMsg.ERROR_CATEGORY_ADD_FAIL,
				"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_ADD_FAIL),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status":  ErrMsg.SUCCESS,
				"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
			})
		}
	}
}

// 删
func CategoryDelete(c *gin.Context) {
	var data Model.Category
	_ = c.ShouldBindJSON(&data)
	code := Model.DeleteCategory(data.Name)
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_CATEGORY_DELETE_FAIL,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_DELETE_FAIL),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.SUCCESS,
			"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
		})
	}
}

// 改
func CategoryEdit(c *gin.Context) {
	type UpdateCategory struct {
		Name       string `gorm:"type:varchar(50);not null" json:"name"`
		UpdateName string `gorm:"type:varchar(50);not null" json:"updateName"`
	}
	var data UpdateCategory
	var code int64
	_ = c.ShouldBindJSON(&data)
	// 判断待修改分类是否存在
	code = Model.CheckCategory(data.Name)
	if code == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_CATEGORY_NOT_EXIST,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_NOT_EXIST),
		})
		return
	}
	// 判断updated分类是否存在
	code = Model.CheckCategory(data.UpdateName)
	if code > 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_CATEGORY_EXIST,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_EXIST),
		})
		return
	}
	// 修改分类
	code = Model.EditCategory(data.Name, data.UpdateName)
	if code == ErrMsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_CATEGORY_EDIT_FAIL,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_EDIT_FAIL),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.SUCCESS,
			"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
		})
	}
}

// 查
func CategoryList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	data := Model.ListCategory(pageSize, pageNum)
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrMsg.ERROR_CATEGORY_LIST_FAIL,
			"Message": ErrMsg.GetErrMsg(ErrMsg.ERROR_CATEGORY_LIST_FAIL),
		})
	}
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  200,
			"Message": ErrMsg.GetErrMsg(ErrMsg.SUCCESS),
			"Data":    data,
		})
	}
}
