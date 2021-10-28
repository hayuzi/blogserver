package v1

import "github.com/gin-gonic/gin"

type Comment struct{}

func NewComment() Comment {
	return Comment{}
}

func (t Comment) Get(c *gin.Context)    {}
func (t Comment) List(c *gin.Context)   {}
func (t Comment) All(c *gin.Context)    {}
func (t Comment) Create(c *gin.Context) {}
func (t Comment) Update(c *gin.Context) {}
func (t Comment) Delete(c *gin.Context) {}
