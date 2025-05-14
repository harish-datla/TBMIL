package main
import( 
	"fmt"
	"classes_and_objects_demo/products"
)

func main() {
	fmt.Println("Demo for Classes and Objects")
	barOneChocolate := product.NewProduct("BarOne Chocolate", 5)
	fmt.Println("BarOne Chocolate Price is", barOneChocolate.GetPrice())
}
