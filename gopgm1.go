package main
import ("fmt")

func main(){
//integer data type
intNum1 := 10
intNum2 := 20
//float data type
floatNum1 := 4.5
floatNum2 := 5.0
//string data type
stringVal1 := "Hey,"
stringVal2 := "How are you?"
//boolean data type
boolVal1 := true 
boolVal2 := false
//integer operations
sum := intNum1 + intNum2
average := (intNum1 + intNum2)/2 
//float operations
product := floatNum1 * floatNum2
quotient := floatNum1 / floatNum2 
//string opeartions
stringConcatenate := stringVal1 + stringVal2
stringLength := len(stringConcatenate)
//boolean operations 
andFunc := boolVal1 && boolVal2
orFunc := boolVal1 || boolVal2

fmt.Printf("Sum: %d\n", sum)
fmt.Printf("Average: %d\n\n", average)

fmt.Printf("Product: %.2f\n", product)
fmt.Printf("Quotient: %.2f\n\n", quotient)

fmt.Printf("Concatenated String: %s\n", stringConcatenate)
fmt.Printf("String Length: %d\n\n", stringLength)

fmt.Printf("AND Operation: %v\n", andFunc)
fmt.Printf("OR Operation:  %v\n", orFunc)
}