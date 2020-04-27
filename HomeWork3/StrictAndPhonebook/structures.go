package main

import "fmt"

//DasAuto Структура для всех автомобилей
type DasAuto struct {
	Brand      string
	Year       uint
	LoadSize   float32
	EngineOn   bool
	WindowOpen bool
}

//Lorry структура для грузового автомобиля
type Lorry struct {
	props       DasAuto
	WheelsCount uint
}

//Auto структура легкового автомобиля
type Auto struct {
	props    DasAuto
	SeatsNum uint
}

func main() {
	fmt.Println("Разные структуры")
	carFiat := Auto{props: DasAuto{Brand: "Fiat 310", Year: 1950, LoadSize: 0.5}, SeatsNum: 4}
	fmt.Println(carFiat)

	carFiat.props.WindowOpen = true
	fmt.Println(carFiat)

	lorryVolvo := Lorry{props: DasAuto{Brand: "Volvo", Year: 2005, LoadSize: 30}, WheelsCount: 8}
	fmt.Println(lorryVolvo)

	lorryVolvo.WheelsCount--
	fmt.Println(lorryVolvo)
}
