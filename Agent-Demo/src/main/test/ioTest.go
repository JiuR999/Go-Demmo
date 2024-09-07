package main

import "Agent-Demo/src/main/server/systemPlugins"

func main() {

	/*	pid := systemPlugins.GetPid()
		fmt.Println("进程数：", len(pid))

		for _, v := range pid {
			fmt.Println(v)
		}*/

	/*	systemPlugins.ExecuteMemory(models.CollectReq{ParamNum: "200"})
		systemPlugins.ExecuteMemory(models.CollectReq{ParamNum: "202"})
		systemPlugins.ExecuteMemory(models.CollectReq{ParamNum: "203"})*/
	systemPlugins.TotalMemory()
	systemPlugins.UsedMemory()

	systemPlugins.UsedPersent()
}
