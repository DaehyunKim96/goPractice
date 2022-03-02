package buckets

type Buckets struct {
	buckets []Bucket
}
type Bucket struct {
	index  int
	bucket []AggLog
}
type AggLog struct {
	adGroupId    int     `json:"adGroupId"`
	avgBidAmount float64 `json:"avgBidAmount"`
	avgPctr      float64 `json:"avgPctr"`
	sumImp       int     `json:"sumImp"`
}

func NewBuckets(resultNum int) *Buckets {
	b := Buckets{}
	for i := 0; i < resultNum; i++ {
		bucket := NewBucket(i)
		b.buckets = append(b.buckets, *bucket)
	}
	return &b
}
func (buckets Buckets) GetCurrentBucket(index int) *Bucket {
	return &buckets.buckets[index]
}
func (bucket Bucket) GetCurrentAggLog(index int) *AggLog {
	return &bucket.bucket[index]
}
func NewBucket(index int) *Bucket {
	bucket := Bucket{index: index}
	return &bucket
}

func (bucket Bucket) GetBucket() *[]AggLog {
	return &bucket.bucket
}
func NewAggLog(adGroupId int, avgBidAmount float64, avgPctr float64) *AggLog {
	return &AggLog{adGroupId: adGroupId, avgBidAmount: avgBidAmount, avgPctr: avgPctr, sumImp: 1}
}
func (bucket *Bucket) Append(aggLog *AggLog) {
	bucket.bucket = append(bucket.bucket, *aggLog)
}
func (aggLog AggLog) GetAdGroupId() int {
	return aggLog.adGroupId
}
func (aggLog AggLog) GetAvgBidAmount() float64 {
	return aggLog.avgBidAmount
}
func (aggLog AggLog) GetAvgPctr() float64 {
	return aggLog.avgPctr
}
func (aggLog AggLog) GetSumImp() int {
	return aggLog.sumImp
}

func (aggLog *AggLog) AvgBA(bidAmount float64) {
	sumBA := aggLog.avgBidAmount*float64(aggLog.sumImp) + bidAmount
	newAvgBA := sumBA / (float64(aggLog.sumImp) + 1)
	aggLog.avgBidAmount = newAvgBA
}
func (aggLog *AggLog) AvgPctr(pctr float64) {
	sumPctr := aggLog.avgPctr*float64(aggLog.sumImp) + pctr
	newAvgPctr := sumPctr / (float64(aggLog.sumImp) + 1)
	aggLog.avgPctr = newAvgPctr
}
func (aggLog *AggLog) SumImp() {
	aggLog.sumImp += 1
}
