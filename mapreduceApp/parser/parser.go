package parser

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Parser interface {
	Parse(win map[string]interface{})
	Getter() interface{}
}

type AdGroupIdParser struct {
	adGroupId int
}

type BidAmountParser struct {
	bidAmount float64
}

type PctrParser struct {
	pctr float64
}

func ParseWinner(byteLog []uint8) map[string]interface{} {
	var rawLog map[string]interface{}
	json.Unmarshal(byteLog, &rawLog)
	rawM := rawLog["rawMessage"].(map[string]interface{})
	win := rawM["winner"].(map[string]interface{})
	return win
}

func NewAdGroupIdParser() *AdGroupIdParser {
	parser := AdGroupIdParser{}
	return &parser
}

func (parser *AdGroupIdParser) Parse(win map[string]interface{}) {
	adGroupId := int(win["adGroupId"].(float64))
	parser.adGroupId = adGroupId
}

func (parser AdGroupIdParser) Getter() interface{} {
	return parser.adGroupId
}

func NewBidAmountParser() *BidAmountParser {
	parser := BidAmountParser{}
	return &parser
}

func (parser *BidAmountParser) Parse(win map[string]interface{}) {
	bidAmount := win["bidAmount"].(float64)
	parser.bidAmount = bidAmount
}

func (parser BidAmountParser) Getter() interface{} {
	return parser.bidAmount
}

func NewPctrParser() *PctrParser {
	parser := PctrParser{}
	return &parser
}

func (parser *PctrParser) Parse(win map[string]interface{}) {
	rankJson, _ := json.Marshal(win["ranking"])
	var ranking map[string]interface{}
	json.Unmarshal(rankJson, &ranking)
	rankStr := string(rankJson)
	if strings.Contains(rankStr, "ad") {
		rankSlice := strings.Split(rankStr, ":")
		var pctrStr string
		for idx, str := range rankSlice {
			if strings.Contains(str, "pctr") {
				pctrStr = rankSlice[idx+1]
				break
			}
		}
		if pctrStr != "" {
			pctrSlice := strings.Split(pctrStr, ",")
			pctr, _ := strconv.ParseFloat(pctrSlice[0], 64)
			parser.pctr = pctr
		} else {
			parser.pctr = 0
		}
	}
}

func (parser PctrParser) Getter() interface{} {
	return parser.pctr
}
func ParseLog(parser Parser, byteLog []uint8) {
	winLog := ParseWinner(byteLog)
	parser.Parse(winLog)
}
func Parsing(byteLog []uint8) (interface{}, interface{}, interface{}) {
	var adGroupIdParser, bidAmountParser, pctrParser Parser
	adGroupIdParser = NewAdGroupIdParser()
	bidAmountParser = NewBidAmountParser()
	pctrParser = NewPctrParser()
	ParseLog(adGroupIdParser, byteLog)
	ParseLog(bidAmountParser, byteLog)
	ParseLog(pctrParser, byteLog)
	return adGroupIdParser.Getter(), bidAmountParser.Getter(), pctrParser.Getter()
}
