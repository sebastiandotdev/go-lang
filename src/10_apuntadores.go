package src

var i *int  // apuntador a ints
var x *bool // apuntador a bools
var pi *int

pi = nil
i := 10
p := &i

a := *p
*p = 21


if pi == nil {
	fmt.Println("No puedo hacer nada con este apuntador")
	fmt.Println("Porque no apunta a nada!")
}