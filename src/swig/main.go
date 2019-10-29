package main

import (
	"fmt"

	"demo/src/swig/example"
)

func main() {
	// Call our gcd() function
	x := 42
	y := 105
	g := example.Gcd(x, y)
	fmt.Println("The gcd of", x, "and", y, "is", g)

	// Manipulate the Dvar global variable

	// Output its current value
	fmt.Println("Dvar =", example.GetDvar())

	// Change its value
	example.SetDvar(3.1415926)

	// See if the change took effect
	fmt.Println("Dvar =", example.GetDvar())

	//	constant
	fmt.Println("-----------------constant-----------------")
	fmt.Println("Iconst  = ", example.Iconst)
	fmt.Println("Fconst  = ", example.Fconst)

	fmt.Println("ICONST  = ", example.ICONST, " (should be 42)")
	fmt.Println("PI  = ", example.PI, " (should be 3.1415926)")
	fmt.Println("SCONST  = ", example.SCONST, " (should be 'Hello World')")
	fmt.Println("AREA    = ", example.AREA)
	fmt.Println("LENGTH  = ", example.LENGTH)

	fmt.Println("-----------------variable-----------------")
	fmt.Println("read-only variable, status = ", example.GetStatus())
	fmt.Println("read-only variable, path = ", example.GetPath())
	example.Print_vars()

	fmt.Println("former variable of pointer, iptrvar = ", *example.GetIptrvar())
	example.SetIptrvar(example.New_int(37))
	fmt.Println("later variable address of pointer, iptrvar = ", example.GetIptrvar())
	fmt.Println("later variable value of pointer, iptrvar = ", *example.GetIptrvar())

	example.SetPtptr(example.New_Point(37, 42))

	fmt.Println("struct variable of pointer, ptptr = ", "0x", example.GetPtptr(), example.Point_Print(example.GetPtptr()))

	fmt.Println("-----------------enumuration-----------------")
	fmt.Println("*** color enum  ***")
	fmt.Println("    RED = ", example.RED)
	fmt.Println("    BLUE = ", example.BLUE)
	fmt.Println("    GREEN = ", example.GREEN)

	fmt.Println("-----------------enumuration in class(C++)-----------------")
	fmt.Println("\n*** Foo::speed enum ***")
	fmt.Println("    Foo::IMPULSE = ", example.FooIMPULSE)
	fmt.Println("    Foo::WARP = ", example.FooWARP)
	fmt.Println("    Foo::LUDICROUS = ", example.FooLUDICROUS)

	fmt.Println("\nTesting use of enums with functions\n")

	example.Enum_test(example.RED, example.FooIMPULSE)

	fmt.Println("\nTesting use of enum with class method")
	f := example.NewFoo()

	f.Enum_test(example.FooIMPULSE)

	fmt.Println("-----------------class(C++)-----------------")
	fmt.Println("Creating some objects:")
	c := example.NewCircle(10)
	fmt.Println("   Created circle", c)
	s := example.NewSquare(10)
	fmt.Println("   Created square", s)

	// ----- Access a static member -----
	fmt.Println("A total of", example.GetShapeNshapes(), "shapes were created")

	// member data access
	c.SetX(20)
	c.SetY(30)

	var shape example.Shape = s
	shape.SetX(-10)
	shape.SetY(5)
	fmt.Println("Here is their current position:")
	fmt.Println("    Circle = (", c.GetX(), " ", c.GetY(), ")")
	fmt.Println("    Square = (", s.GetX(), " ", s.GetY(), ")")

	fmt.Println("\nHere are some properties of the shapes:")
	shapes := []example.Shape{c, s}
	for i := 0; i < len(shapes); i++ {
		fmt.Println("   ", shapes[i])
		fmt.Println("        area      = ", shapes[i].Area())
		fmt.Println("        perimeter = ", shapes[i].Perimeter())
	}

	fmt.Println("	Guess I'll clean up now")
	example.DeleteCircle(c)
	example.DeleteSquare(s)

	fmt.Println(example.GetShapeNshapes(), " shapes remain")

}
