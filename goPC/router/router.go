package router

import (
	"github.com/gin-gonic/gin"
	"goPC/controller"
	"goPC/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//跨域
	r.Use(middleware.Cors())
	v1 := r.Group("CanteenPC")
	{
		v1.POST("/login", controller.PCLogin) //登录界面

		v1.GET("/Reserve", controller.Reserve)                  //显示预约界面
		v1.PUT("/AgreeReserve", controller.AgreeReserve)        //同意预约
		v1.POST("/DisagreeReserve", controller.DisagreeReserve) //拒绝预约

		v1.GET("/AdminApply", controller.AdminApply)        //显示管理员申请界面
		v1.PUT("/AgreeApply", controller.AgreeApply)        //同意申请
		v1.POST("/DisagreeApply", controller.DisagreeApply) //拒绝申请

		v1.GET("/MeetingRoom", controller.MeetingRoom) //显示会议室界面
		v1.POST("/DeleteRoom", controller.DeleteRoom)  //删除会议室
		v1.PUT("/EditRoom", controller.EditRoom)       //修改会议室
		v1.POST("/AddRoom", controller.AddRoom)        //添加会议室

		v1.GET("/AppointmentRecord", controller.AppointmentRecord) //显示预约历史记录
		v1.POST("/RecordSearch", controller.RecordSearch)          //搜索预约历史记录

		v1.GET("/Opinion", controller.Opinion)              //建议显示界面
		v1.POST("/DeleteOpinion", controller.DeleteOpinion) //删除建议
	}
	//小程序
	v2 := r.Group("Applets")
	{
		v2.POST("/AddAdvice", controller.AddAdvice) //添加意见

		v2.POST("/GetRoom", controller.GetRoom) //显示预约界面
	}
	return r
}
