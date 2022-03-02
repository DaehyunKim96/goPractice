//package main
//
//import (
//	"bufio"
//	"fmt"
//	"io/fs"
//	"io/ioutil"
//	"mapreduceApp/mappers"
//	"mapreduceApp/reducers"
//	"os"
//	"sync"
//)
//
//func split(inputDir string) []fs.FileInfo {
//	files, _ := ioutil.ReadDir(inputDir)
//	return files
//}
//
//func shuffle(log mappers.Log, num int, rdrs *reducers.Reducers) {
//	reducer := rdrs.GetReducer(num)
//	reducer.Aggregate(log)
//}
//func write(resultNum int, rdrs *reducers.Reducers, outputDirName string) {
//	//for i := 0; i < resultNum; i++ {
//	//	doc, _ := json.Marshal(rdrs.GetReducer(i))
//	//	//doc := rdrs.GetReducer(i).GetByteArray()
//	//	//rdrs.GetReducer(i).PrintByteArray()
//	//	//rdrs.PrintBucket(1)
//	//
//	//	fileName := "part-0000" + fmt.Sprint(i) + "-r.json"
//	//	dirName := outputDirName + "/" + fileName
//	//	err := ioutil.WriteFile(dirName, doc, os.FileMode(0644))
//	//	if err != nil {
//	//		fmt.Println(err)
//	//		return
//	//	}
//	//}
//}
//
//func main() {
//	inputDirName := "./log_temp"
//	//outputDirName := "./log_r"
//	resultNum := 3
//	rdrs := reducers.NewReducers(resultNum)
//	mutex := new(sync.Mutex)
//	c1 := make(chan mappers.Log)
//	c2 := make(chan mappers.Log)
//	c3 := make(chan mappers.Log)
//	files := split(inputDirName)
//	for _, file := range files {
//		fileName := inputDirName + "/" + file.Name()
//		f, _ := os.Open(fileName)
//		defer f.Close()
//		go func(f *os.File) {
//			fileScanner := bufio.NewScanner(f)
//			for fileScanner.Scan() {
//				byteLog := []byte(fileScanner.Text())
//				mapper := mappers.NewMapper(byteLog)
//				mapper.Map()
//				adGroupId := mapper.GetLog().GetAdGroupId()
//				if adGroupId%3 == 1 {
//					c1 <- mapper.GetLog()
//				} else if adGroupId%3 == 2 {
//					c2 <- mapper.GetLog()
//				} else {
//					c3 <- mapper.GetLog()
//				}
//				//count += 1
//			}
//		}(f)
//	}
//	for {
//		select {
//		case log1 := <-c1:
//			mutex.Lock()
//			shuffle(log1, 1, rdrs)
//			mutex.Unlock()
//			fmt.Println(1, log1)
//		case log2 := <-c2:
//			mutex.Lock()
//			shuffle(log2, 2, rdrs)
//			mutex.Unlock()
//			fmt.Println(2, log2)
//		case log3 := <-c3:
//			mutex.Lock()
//			shuffle(log3, 0, rdrs)
//			mutex.Unlock()
//			fmt.Println(3, log3)
//		}
//	}
//	//for log := range shuffleC {
//	//	fmt.Println(log)
//	//	reduce(log, resultNum, bkts)
//	//}
//	//todo: Write Json
//	//write(resultNum, rdrs, outputDirName)
//}
