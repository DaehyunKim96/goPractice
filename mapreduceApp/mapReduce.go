package main

import (
	"flag"
	"mapreduceApp/buckets"
	"mapreduceApp/logs"
	reader2 "mapreduceApp/reader"
	"mapreduceApp/reducers"
	writer2 "mapreduceApp/writer"
	"sync"
)

func shuffle(log logs.Log, num int, bkts *buckets.Buckets) *buckets.Bucket {
	id := log.AdGroupId()
	idx := id % num
	currentBucket := bkts.GetCurrentBucket(idx)
	return currentBucket
}
func main() {
	//Todo main Args
	inputDirName := "log"
	outputDirName := "log_r"
	resultNum := flag.Int("num", 10, "number of outputfiles")
	//Todo make channel, watigroup
	channel := make(chan logs.Log)
	wg := sync.WaitGroup{}
	//Todo read
	reader := reader2.NewReader()
	reader2.ReadFile(reader, inputDirName, &wg, channel)
	//Todo shuffle and reduce
	bkts := buckets.NewBuckets(*resultNum)
	reducer := reducers.NewReducer()
	for log := range channel {
		currentBucket := shuffle(log, *resultNum, bkts)
		reducer.Reduce(log, currentBucket)
	}
	//Todo write
	writer := writer2.NewWriter()
	writer.Write(*resultNum, bkts, outputDirName)
}
