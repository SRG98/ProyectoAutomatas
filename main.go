package main

import(
"fmt"
"github.com/SRG98/ProyectoAutomatas/controllers"
"github.com/SRG98/ProyectoAutomatas/views"
)

func main()  {
	controlador := contrladores.NewControlador()
	uiinterfaz := vistas.NewUI(controlador)
	err:= uiinterfaz.RunUI()
	if err:=nil {
		fmt.Println("Error:",err)
	} 
	
}