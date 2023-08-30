package main
import "fmt"
func main(){

	var array [5]int

	fmt.Println(array)

	var array1 [5] float32
	  fmt.Println(array1)

	  array2 :=  [3]float64{1.2,2.3}
fmt.Println(array2)

array2[2]=3.4
fmt.Println(array2)


var matrix [5][4] int

for i := 0 ; i <=3 ; i++ {
	for j := 0 ; j <= 3 ; j++ {
		matrix[i][j] = matrix[i][j]+10
	}
}
fmt.Println(matrix)

fmt.Println ("Lenght of a row ",len(matrix))
fmt.Println("Length of a col ", len(matrix[0]))

  }
