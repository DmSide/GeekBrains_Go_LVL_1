package lesson4

import "fmt"

func f1() {
	defer fmt.Println("f1") // (["f1"])
	// defer add to the beginning(LIFO), because we already have defer
	defer f2() // (["f1" f2 ]) = (["f1" ("f2" "f3" "f4" "f6" "f5")])
} // Simplify: (["f1" ("f2" "f3" "f4" "f6" "f5")]) = (("f2" "f3" "f4" "f6" "f5") "f1") = ("f2" "f3" "f4" "f6" "f5" "f1")

func f2() {
	fmt.Println("f2") // ("f2")
	defer f3()        // ("f2" [f3]) = ("f2" [("f3" "f4" "f6" "f5")])
} // Simplify: ("f2" [("f3" "f4" "f6" "f5")]) = ("f2" ("f3" "f4" "f6" "f5")) = ("f2" "f3" "f4" "f6" "f5")

func f3() {
	fmt.Println("f3") // ("f3")
	f4()              // ("f3" f4) = ("f3" ("f4" "f6" "f5"))
} // Simplify: ("f3" ("f4" "f6" "f5")) = ("f3" "f4" "f6" "f5")

func f4() {
	fmt.Println("f4")       // ("f4")
	defer fmt.Println("f5") // ("f4" ["f5"])
	defer fmt.Println("f6") // ("f4" ["f5", "f6"])
} // Simplify: ("f4" ["f5" "f6"]) = ("f4" "f6" "f5")

// Defer works on LIFO
func DeferMagic() {
	fmt.Println("Defer started")
	f1()
}
