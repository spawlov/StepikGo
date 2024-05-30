package main

func swapNeighborElements(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i += 2 {
			arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

func main(){

}