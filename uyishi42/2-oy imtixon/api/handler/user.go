package handler

import (
	"net/http"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
)
//Bu yerda User yaratiladi 
func (h *Handler) CreateUser(c *gin.Context) {
	var user models.Users
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.User.CreateUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, user)
}
// Bu yerda bitta User ID si bo'yich keladi
func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}
	
	user, err := h.User.ReadUser(id)
	if err != nil {
  
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, user)
	
}
// Bu yerda malumotni update qilinadi
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.Users
	
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.User.UpdateUser(&user,c.Param("user_id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(200, "SUCCESS")
	
}
// Bu yerda ma'lumotni ID si bo'yicha o'chiriladi
func (h *Handler) DeleteUsers(c *gin.Context) {
	id := c.Param("user_id")
	if id == "" {
		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
	}
	
	err := h.User.DeleteUser(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(200, "SUCCESS")
	
}
// Bu yerda barcha Userlarni ko'rish mumkin
func (h *Handler) GetFilterAllUsers(c *gin.Context) {
	
	userFilter := models.UserF{}
	userFilter.ID = c.Param("user_id")
	userFilter.Name = c.Param("name")
	userFilter.Email = c.Param("email")
	userFilter.Birthday = c.Param("birthday")
	userFilter.Password = c.Param("password")
	res, err := h.User.FilterAllUser(userFilter)
	if err != nil{
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}


// 1-TASK
// Bu yerda Userga tegisli barcha course ma'lumotlari keladi
func (h *Handler) GetCourseByUsers(c *gin.Context) {
	userId, course, err := h.User.GetCourseByUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(course); i++ {
			c.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"status":  http.StatusOK,
				"time":    time.Now(),
				"user_id": userId,
				"Course": gin.H{
					"course_id":   course[i].ID,
					"title":       course[i].Title,
					"description": course[i].Description,
				},
			})
		}
	}

}
// 4-TASK
// Bu yerda Userni ismi yoki emaili bo'yicha qidib topish mumkin
func (h *Handler) SearchUsers(c *gin.Context) {
	users, err := h.User.SearchUser(c.Param("name"), c.Param("email"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
		} else {
			for i := 0; i < len(users); i++ {

				c.JSON(http.StatusOK, gin.H{
				"message":    "Success",
				"status":     http.StatusOK,
				"time":       time.Now(),
				"user_id":    users[i].ID,
				"name":       users[i].Name,
				"email":      users[i].Email,
				"birthday":   users[i].Birthday,
				"password":   users[i].Password,
				"created_at": users[i].CreatedAt,
				"updated_at": users[i].UpdateAt,
				"deleted_at": users[i].DeleteAt,
			})
		}
	}
}