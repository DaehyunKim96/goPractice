package reader

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"mapreduceApp/logs"
	"mapreduceApp/mappers"
	"os"
	"sync"
)

type Reader struct{}
type Files struct {
	files []fs.FileInfo
	dir   string
}

func NewReader() *Reader {
	return &Reader{}
}
func NewFiles() *Files {
	return &Files{}
}
func (reader Reader) FileSplit(inputDir string) *Files { //ListFiles, 직관적인 변수명
	path, _ := os.Getwd()
	dirName := path + "/" + inputDir
	fs := NewFiles()
	fileList, _ := ioutil.ReadDir(dirName)
	fs.files = fileList
	fs.dir = inputDir
	return fs
}
func (reader Reader) Read(f *os.File, wg *sync.WaitGroup, channel chan<- logs.Log) { //struct channel, wg
	defer wg.Done()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		byteLog := []byte(fileScanner.Text())
		mapper := mappers.NewMapper()
		mapper.Map(byteLog)
		channel <- mapper.GetLog()
	}
}
func (reader Reader) ReadFiles(fs Files, wg *sync.WaitGroup, channel chan<- logs.Log) {
	for _, file := range fs.files {
		fileName := fs.dir + "/" + file.Name()
		f, _ := os.Open(fileName)
		defer f.Close()
		wg.Add(1)
		go reader.Read(f, wg, channel)
	}
	wg.Wait()
	close(channel)
}
func ReadFile(reader *Reader, inputDir string, wg *sync.WaitGroup, channel chan<- logs.Log) {
	fs := reader.FileSplit(inputDir)
	go reader.ReadFiles(*fs, wg, channel)
}
