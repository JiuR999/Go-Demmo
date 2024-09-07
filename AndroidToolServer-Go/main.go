package main

import (
	"AndroidToolServer-Go/roof/env"
	"AndroidToolServer-Go/routers"
	"math/rand"
	"time"
)

//	@title			My GoDemoAPP API Doc
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost:8080

//	@securityDefinitions.basic	BasicAuth
func main() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	/*	hardwareType := []string{"Linux", "centos", "Windows"}
		hardwareGroup := []string{"测试分组", "计算机处理逻辑", "测试分组2"}
		mainPeople := []string{"科长", "组长", "主管"}*/

	//name := "Test1"
	/*	hd := &model.HardWare{
			Name: &name,
		}
		ret := hardware.Create(&hd)*/

	/*	var hws []models.HardWare
		for i := 0; i < 10; i++ {
			monitorStatus := int32(i % 3)

			randomNumber1 := rand.Intn(len(hardwareType))

			randomNumber2 := rand.Intn(len(hardwareGroup))
			groupId := strconv.Itoa(randomNumber2)
			hws = append(hws, models.HardWare{
				Name:          fmt.Sprintf("Test%d", i),
				Code:          fmt.Sprintf("%4d", i),
				CreateTime:    time.Now(),
				UpdateTime:    time.Now(),
				MonitorStatus: &monitorStatus,
				HardWareType:  &hardwareType[randomNumber1],
				UnitDepartID:  &groupId,
				UnitGroupName: &hardwareGroup[randomNumber2],
				UserID:        &groupId,
				UserName:      &mainPeople[randomNumber1],
			})
		}
		err := hardware.Create(hws).Error
		fmt.Println(err.Error())*/
	//hardware.AutoMigrate(&hw)

	routers.Router.Init().Run(":" + env.Config.ServerPort)
}
