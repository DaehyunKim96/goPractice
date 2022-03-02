package reducers

import (
	"mapreduceApp/buckets"
	"mapreduceApp/logs"
)

type Reducer struct{}

func NewReducer() *Reducer {
	reducer := Reducer{}
	return &reducer
}
func (reducer Reducer) Reduce(log logs.Log, bucket *buckets.Bucket) {
	reducer.Aggregate(log, bucket)
}
func (reducer *Reducer) Aggregate(log logs.Log, bucket *buckets.Bucket) {
	find := false
	index := -1
	for idx, elem := range *bucket.GetBucket() {
		if elem.GetAdGroupId() == log.AdGroupId() {
			find = true
			index = idx
		}
	}
	if find {
		bucket.GetCurrentAggLog(index).AvgBA(log.BidAmount())
		bucket.GetCurrentAggLog(index).AvgPctr(log.Pctr())
		bucket.GetCurrentAggLog(index).SumImp()
	} else {
		newLog := buckets.NewAggLog(log.AdGroupId(), log.BidAmount(), log.Pctr())
		bucket.Append(newLog)
	}
}
