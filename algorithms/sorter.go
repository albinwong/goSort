/*package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting") 
var outfile *string = flag.String("o", "outfile", "File to receive sorted values") 
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	} 
}*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string)(values []int, err error) {
	file, err := os.Open(file)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	value := make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexcepted.")
			return
		}

		str := string(line) // 转换字符数组为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		value = append(values, value)
	}
	return

}


func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithms =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}