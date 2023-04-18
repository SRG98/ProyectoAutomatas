package main

import(
"fmt"
"github.com/SRG98/ProyectoAutomatas/contrladores"
"github.com/SRG98/ProyectoAutomatas/vistas"
)

func main()  {
	controlador := contrladores.NewControlador()
	uiinterfaz := vistas.NewUI(controlador)
	err:= uiinterfaz.RunUI()
	if err:=nil {
		fmt.Println("Error:",err)
	} 
	
}