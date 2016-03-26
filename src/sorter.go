package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	stm "sorterMethod"
	"strconv"
)

func main() {
	//获取输入参数
	infile := flag.String("in", "no place", "input file path")
	outfile := flag.String("out", "no place", "output file path")
	method := flag.String("method", "no place", "sort method")
	flag.Parse()
	//解析输入文件数据
	inputInt, err := parseInputFile(*infile)
	//调用排序算法
	if err != nil || inputInt == nil {
		fmt.Println("read input file failed")
		return
	}
	switch *method {
	case "ks":
		stm.QuickSort(inputInt)
	case "mp":
		stm.BubbleSort(inputInt)
	default:
		fmt.Println("find no sort method for you input,please input 'mp' for bubble,'ks' for quick ")
		return
	}
	//将排序结果写出到输出文件
	writeOutputFile(inputInt, *outfile)
}

func writeOutputFile(outputInt []int, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("write output file error", filepath)
		return err
	}
	defer file.Close()
	bfw := bufio.NewWriter(file)
	for _, v := range outputInt {
		vstr := strconv.Itoa(v)
		bfw.WriteString(vstr + "\n")
	}
	bfw.Flush()
	return nil
}

func parseInputFile(filepath string) ([]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("open file %v has error", filepath)
		return nil, err
	}
	defer file.Close()
	bf := bufio.NewReader(file)
	inputInt := make([]int, 0)
	for {
		line, isPrefix, err := bf.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Println("read finished")
				break
			}
			fmt.Println("read file line has err:%v", line)
			return nil, err
		}
		if isPrefix {
			fmt.Println("is prefix true,maybe line to long")
			return nil, nil
		}
		//将字节数组转为字符串
		linestr := string(line)
		fmt.Println("linestr:", linestr)
		ivalue, conErr := strconv.Atoi(linestr)
		if conErr != nil {
			fmt.Println("Value must be int,default it to zero:", ivalue)
		}
		inputInt = append(inputInt, ivalue)
	}
	fmt.Println(inputInt)
	return inputInt, nil
}
