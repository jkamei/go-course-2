package main
import ("bufio"
"os"
"strings"
"strconv"
"fmt")


func main() {
	intArray := getUserInput()
	var pos int
	sorted := true
	for !(pos == len(intArray)-1 && sorted) {
		fmt.Println(intArray, pos)
		if pos < len(intArray)-1 && intArray[pos] > intArray[pos+1] {
			swap(intArray, pos)
			sorted = false
		}
		pos = pos + 1
		if pos == len(intArray) {
			reset(&sorted, &pos)
		}
	}
}

func swap (intArray []int, pos int) {
	intArray[pos], intArray[pos+1] = intArray[pos+1], intArray[pos]
}

func reset(sorted *bool, pos *int) {
	*sorted = true
	*pos = 0
}

func getUserInput() []int {
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("enter a sequence of integers: ")
	userInput, _ := in.ReadString('\n')
	intArray := []int{}
	strArray := strings.Split(strings.TrimSpace(userInput), " ")
	for _, i := range strArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intArray = append(intArray, j)
	}
	return intArray
}