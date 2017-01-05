package main

import "fmt"

func iteration(num int, k int, output[] int)([]int){
	var s[] int
	var array[] int
	var sum int
	var i int
 	var j int
 	var num2 int
	 if k < num{
	 	fmt.Scanf("%d\n", &num2) 
	 	newarray := arraylist(num2 , s)
 		list := dopositive(newarray , i, array);
 		nsum:= dosum(list, j, sum )
 		k++
 		output = append(output, nsum)
 		return iteration(num, k, output)
 	}
 	return output
}

func arraylist(num2 int, s[] int)([]int){
	var i int
	fmt.Scanf("%d", &i)
	s = append(s, i);
	if len(s) == num2{
		return s	
	}
	return arraylist(num2, s)
}

func dopositive(newarray[] int, i int, array[] int)([]int){
	if i < len(newarray) {
		if newarray[i] < 0{
			i ++
			return dopositive(newarray, i, array)
		}
		if newarray[i] >= 0 {
			array = append(array, newarray[i])
			i ++
			return dopositive(newarray, i, array)
		}
	}
	return array 
}

func dosum(list[] int , i int, sum int)(int){
	if i < len(list){
		sum = sum + list[i]*list[i] 
		i++
		return dosum(list, i , sum)
	}
	return sum
}

func main() {
 var num int
 var k int
 var output[] int
 fmt.Scanf("%d\n", &num) 
 final := iteration(num, k, output);
 fmt.Printf("%d\n", final[0])
 fmt.Printf("%d\n", final[1]) 
 }
 

 




 
