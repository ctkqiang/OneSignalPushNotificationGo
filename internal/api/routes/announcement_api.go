package routes

import (
	"net/http"
	"pushnotification_services/internal/config"
	"pushnotification_services/internal/repositories"
	"pushnotification_services/internal/structure"
	"pushnotification_services/internal/utilities"

	"github.com/gin-gonic/gin"
)

func Announcement(router *gin.Engine) {
	public := router.Group(config.ANNOUNCEMENT)
	{
		public.POST(config.ANNOUNCEMENT_CREATE, CreateAnnouncement())
		public.DELETE(config.ANNOUNCEMENT_DELETE, DeleteAnnouncement())
		public.GET(config.ANNOUNCEMENT_LATEST, GetLatestAnnouncement())
		public.PUT(config.ANNOUNCEMENT_UPDATE, UpdateAnnouncement())
		public.GET(config.ANNOUNCEMENT_LIST_ALL, GetAllAnnouncements())
	}
}

// CreateAnnouncement 创建新公告
// @Summary 创建新公告
// @Description 创建一个新的公告并保存到数据库
// @Tags 公告管理
// @Accept json
// @Produce json
// @Param announcement body structure.Announcement true "公告信息"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /announcement/create [post]
func CreateAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		var announcement structure.Announcement
		if err := c.ShouldBindJSON(&announcement); err != nil {
			utilities.Log(utilities.ERROR, "解析请求参数失败: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request parameters",
			})
			return
		}

		err := repositories.WriteAnnouncement(announcement)
		if err != nil {
			utilities.Log(utilities.ERROR, "创建公告失败: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "创建公告失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "公告创建成功",
			"data":    announcement,
		})
	}
}

// DeleteAnnouncement 删除公告
// @Summary 删除公告
// @Description 根据 ID 删除公告
// @Tags 公告管理
// @Accept json
// @Produce json
// @Param id query string true "公告 ID"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /announcement/delete [delete]
func DeleteAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "ID 不能为空",
			})
			return
		}

		err := repositories.DeleteAnnouncement(id)
		if err != nil {
			utilities.Log(utilities.ERROR, "删除公告失败: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "删除公告失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "公告删除成功",
		})
	}
}

// GetLatestAnnouncement 获取最新公告
// @Summary 获取最新公告
// @Description 获取最新的公告信息
// @Tags 公告管理
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 404 {object} map[string]interface{} "没有找到公告"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /announcement/latest [get]
func GetLatestAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		announcement, err := repositories.GetLatestAnnouncement()
		if err != nil {
			utilities.Log(utilities.ERROR, "获取最新公告失败: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "获取最新公告失败",
			})
			return
		}

		if announcement == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "没有找到公告",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "获取最新公告成功",
			"data":    announcement,
		})
	}
}

// UpdateAnnouncement 更新公告
// @Summary 更新公告
// @Description 根据 ID 更新公告信息
// @Tags 公告管理
// @Accept json
// @Produce json
// @Param id query string true "公告 ID"
// @Param announcement body structure.Announcement true "公告信息"
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /announcement/update [put]
func UpdateAnnouncement() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "ID 不能为空",
			})
			return
		}

		var announcement structure.Announcement
		if err := c.ShouldBindJSON(&announcement); err != nil {
			utilities.Log(utilities.ERROR, "解析请求参数失败: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Invalid request parameters",
			})
			return
		}

		err := repositories.UpdateAnnouncement(id, announcement)
		if err != nil {
			utilities.Log(utilities.ERROR, "更新公告失败: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "更新公告失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "公告更新成功",
			"data":    announcement,
		})
	}
}

// GetAllAnnouncements 获取所有公告
// @Summary 获取所有公告
// @Description 获取所有公告信息，按创建时间降序排序
// @Tags 公告管理
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "成功响应"
// @Failure 500 {object} map[string]interface{} "服务器内部错误"
// @Router /announcement/all [get]
func GetAllAnnouncements() gin.HandlerFunc {
	return func(c *gin.Context) {
		announcements, err := repositories.GetAllAnnouncements()
		if err != nil {
			utilities.Log(utilities.ERROR, "获取所有公告失败: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "获取所有公告失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "获取所有公告成功",
			"data":    announcements,
		})
	}
}