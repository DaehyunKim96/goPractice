package logs

type Log struct {
	adGroupId int
	bidAmount float64
	pctr      float64
	imp       int
}

func NewLog() *Log {
	log := Log{imp: 1}
	return &log
}

func (log Log) AdGroupId() int {
	return log.adGroupId
}
func (log *Log) SetAdGroupId(adGroupId int) {
	log.adGroupId = adGroupId
}

func (log Log) BidAmount() float64 {
	return log.bidAmount
}
func (log *Log) SetBidAmount(bidAmount float64) {
	log.bidAmount = bidAmount
}

func (log Log) Pctr() float64 {
	return log.pctr
}
func (log *Log) SetPctr(pctr float64) {
	log.pctr = pctr
}

func (log Log) Imp() int {
	return log.imp
}
func (log *Log) SetImp(imp int) {
	log.imp = imp
}
