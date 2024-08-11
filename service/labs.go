package service

import (
	"devbeginner-doc-api/database"
	"devbeginner-doc-api/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Labs *labsSrvMethod

type labsSrvMethod struct{}

func (p *labsSrvMethod) Create(c *gin.Context) {
	recv := model.Lab{}
	err := c.ShouldBindJSON(&recv)
	if err != nil {
		fmt.Println("[Service.Lab] 解析JSON数据失败! -> ", err.Error())
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "传递数据无效!"})
		return
	}
	res := database.Labs.Create(&recv)
	if res != nil {
		fmt.Println("[Service.Lab] 创建数据失败 -> ", res.Error())
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "创建数据失败!"})
		return
	}
	fmt.Println("[Service.Lab] 创建数据成功!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "创建数据成功!"})
}

func (p *labsSrvMethod) Get(c *gin.Context) {
	var res []model.Lab
	var err error
	var isRelease bool
	param := c.Query("release")
	if param == "" {
		isRelease = true
	} else if param == "true" {
		isRelease = true
	} else {
		isRelease = false
	}
	res, err = database.Labs.Query(isRelease)
	if err != nil {
		fmt.Println("[Service.Lab] 查询数据失败 -> ", err.Error())
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "查询数据失败!"})
		return
	}
	c.JSON(http.StatusOK, res)
	fmt.Println("[Service.Lab] 查询数据成功!")
}

func (p *labsSrvMethod) Delete(c *gin.Context) {
	param := c.Query("uid")
	if param == "" {
		fmt.Println("[Service.Lab] 删除数据失败 -> 参数为空!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "删除数据失败,参数不能为空!"})
		return
	}
	uid, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("[Service.Lab] 删除数据失败 -> 参数错误!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "删除数据失败,参数错误!"})
		return
	}
	err = database.Labs.Delete(uid)
	if err != nil {
		fmt.Println("[Service.Lab] 删除数据失败 -> 删除数据失败!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
		return
	}
	fmt.Println("[Service.Lab] 删除数据成功!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "删除数据成功!"})
}

func (p *labsSrvMethod) Update(c *gin.Context) {
	paraUid := c.Query("uid")
	col := c.Query("column")
	if paraUid == "" || col == "" {
		fmt.Println("[Service.Lab] 修改数据失败 -> 参数为空或不完整!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "修改数据失败,参数为空或不完整!"})
		return
	}
	if !model.IsJsonInclude(model.Lab{}, col) {
		fmt.Println("[Service.Lab] 修改数据失败 -> 字段名不存在!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "修改数据失败,字段名不存在!"})
		return
	}
	uid, err := strconv.Atoi(paraUid)
	if err != nil {
		fmt.Println("[Service.Lab] 修改数据失败 -> uid参数错误!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "修改数据失败,uid参数错误!"})
		return
	}
	body := struct {
		Content any `json:"content"`
	}{}
	err = c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println("[Service.Lab] 解析JSON数据失败! -> ", err.Error())
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "传递数据无效!"})
		return
	}
	err = database.Labs.Update(uid, col, body.Content)
	if err != nil {
		fmt.Println("[Service.Lab] 修改数据失败 -> 修改数据失败!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
		return
	}
	fmt.Println("[Service.Lab] 修改数据成功!")
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "修改数据成功!"})
}
