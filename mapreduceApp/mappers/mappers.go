package mappers

import (
	"mapreduceApp/logs"
	"mapreduceApp/parser"
)

type Mapper struct {
	log logs.Log
}

func NewMapper() *Mapper {
	log := logs.NewLog()
	mapper := Mapper{log: *log}
	return &mapper
}

func (mapper *Mapper) Map(byteLog []uint8) {
	adGroupId, bidAmount, pctr := parser.Parsing(byteLog)
	mapper.log.SetAdGroupId(adGroupId.(int))
	mapper.log.SetBidAmount(bidAmount.(float64))
	mapper.log.SetPctr(pctr.(float64))
}
func (mapper Mapper) GetLog() logs.Log {
	return mapper.log
}
