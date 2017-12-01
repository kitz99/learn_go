package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16
	break_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}
// value receiver => Does not modify any value
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

// pointer receiver => Does modify values of struct
func (c *car) new_top_speed(new_speed float64) {
	c.top_speed_kmh = new_speed
}

func main() {
	a_car := car{ 
		gas_pedal: 65000, 
		break_pedal: 0, 
		steering_wheel: 12560, 
		top_speed_kmh: 150.55 }
	
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	
	fmt.Println("Changed top speed")
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}