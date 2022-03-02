package writer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mapreduceApp/buckets"
	"os"
)

type Writer struct{}
type WriterAggLog struct {
	AdGroupId    int     `json:"adGroupId"`
	AvgBidAmount float64 `json:"avgBidAmount"`
	AvgPctr      float64 `json:"avgPctr"`
	SumImp       int     `json:"sumImp"`
}

func NewWriter() *Writer {
	writer := Writer{}
	return &writer
}
func NewWriterAggLog() *WriterAggLog {
	return &WriterAggLog{}
}
func (writer Writer) DeleteFiles(outputDirName string) {
	path, _ := os.Getwd()
	outputDir := path + "/" + outputDirName
	fileList, _ := ioutil.ReadDir(outputDir)
	for _, file := range fileList {
		fileName := outputDir + "/" + file.Name()
		e := os.Remove(fileName)
		if e != nil {
			log.Fatal(e)
		}
	}
}
func (writer Writer) Write(resultNum int, bkts *buckets.Buckets, outputDirName string) {
	newWriterLog := NewWriterAggLog()
	writer.DeleteFiles(outputDirName)
	for i := 0; i < resultNum; i++ {
		aggLogSlice := []WriterAggLog{}
		for _, val := range *bkts.GetCurrentBucket(i).GetBucket() {
			newWriterLog.AdGroupId = val.GetAdGroupId()
			newWriterLog.AvgBidAmount = val.GetAvgBidAmount()
			newWriterLog.AvgPctr = val.GetAvgPctr()
			newWriterLog.SumImp = val.GetSumImp()
			aggLogSlice = append(aggLogSlice, *newWriterLog)
		}
		doc, _ := json.MarshalIndent(aggLogSlice, "", "\t")
		path, _ := os.Getwd()
		fileName := "part-0000" + fmt.Sprint(i) + "-r.json"
		dirName := path + "/" + outputDirName + "/" + fileName
		err := ioutil.WriteFile(dirName, doc, os.FileMode(0644))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
