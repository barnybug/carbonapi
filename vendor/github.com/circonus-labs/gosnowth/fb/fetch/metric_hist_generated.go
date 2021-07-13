// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fetch

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HistogramBucketT struct {
	Val   int8
	Exp   int8
	Count uint64
}

func HistogramBucketPack(builder *flatbuffers.Builder, t *HistogramBucketT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	HistogramBucketStart(builder)
	HistogramBucketAddVal(builder, t.Val)
	HistogramBucketAddExp(builder, t.Exp)
	HistogramBucketAddCount(builder, t.Count)
	return HistogramBucketEnd(builder)
}

func (rcv *HistogramBucket) UnPack() *HistogramBucketT {
	if rcv == nil {
		return nil
	}
	t := &HistogramBucketT{}
	t.Val = rcv.Val()
	t.Exp = rcv.Exp()
	t.Count = rcv.Count()
	return t
}

type HistogramBucket struct {
	_tab flatbuffers.Table
}

func GetRootAsHistogramBucket(buf []byte, offset flatbuffers.UOffsetT) *HistogramBucket {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HistogramBucket{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *HistogramBucket) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HistogramBucket) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *HistogramBucket) Val() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *HistogramBucket) MutateVal(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *HistogramBucket) Exp() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *HistogramBucket) MutateExp(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func (rcv *HistogramBucket) Count() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *HistogramBucket) MutateCount(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func HistogramBucketStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func HistogramBucketAddVal(builder *flatbuffers.Builder, val int8) {
	builder.PrependInt8Slot(0, val, 0)
}
func HistogramBucketAddExp(builder *flatbuffers.Builder, exp int8) {
	builder.PrependInt8Slot(1, exp, 0)
}
func HistogramBucketAddCount(builder *flatbuffers.Builder, count uint64) {
	builder.PrependUint64Slot(2, count, 0)
}
func HistogramBucketEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type HistogramT struct {
	Buckets []*HistogramBucketT
}

func HistogramPack(builder *flatbuffers.Builder, t *HistogramT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	bucketsOffset := flatbuffers.UOffsetT(0)
	if t.Buckets != nil {
		bucketsLength := len(t.Buckets)
		bucketsOffsets := make([]flatbuffers.UOffsetT, bucketsLength)
		for j := 0; j < bucketsLength; j++ {
			bucketsOffsets[j] = HistogramBucketPack(builder, t.Buckets[j])
		}
		HistogramStartBucketsVector(builder, bucketsLength)
		for j := bucketsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(bucketsOffsets[j])
		}
		bucketsOffset = builder.EndVector(bucketsLength)
	}
	HistogramStart(builder)
	HistogramAddBuckets(builder, bucketsOffset)
	return HistogramEnd(builder)
}

func (rcv *Histogram) UnPack() *HistogramT {
	if rcv == nil {
		return nil
	}
	t := &HistogramT{}
	bucketsLength := rcv.BucketsLength()
	t.Buckets = make([]*HistogramBucketT, bucketsLength)
	for j := 0; j < bucketsLength; j++ {
		x := HistogramBucket{}
		rcv.Buckets(&x, j)
		t.Buckets[j] = x.UnPack()
	}
	return t
}

type Histogram struct {
	_tab flatbuffers.Table
}

func GetRootAsHistogram(buf []byte, offset flatbuffers.UOffsetT) *Histogram {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Histogram{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Histogram) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Histogram) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Histogram) Buckets(obj *HistogramBucket, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Histogram) BucketsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func HistogramStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func HistogramAddBuckets(builder *flatbuffers.Builder, buckets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(buckets), 0)
}
func HistogramStartBucketsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func HistogramEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type MetricHistogramResultT struct {
	Timestamp uint64
	Period    int32
	Histogram *HistogramT
}

func MetricHistogramResultPack(builder *flatbuffers.Builder, t *MetricHistogramResultT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	histogramOffset := HistogramPack(builder, t.Histogram)
	MetricHistogramResultStart(builder)
	MetricHistogramResultAddTimestamp(builder, t.Timestamp)
	MetricHistogramResultAddPeriod(builder, t.Period)
	MetricHistogramResultAddHistogram(builder, histogramOffset)
	return MetricHistogramResultEnd(builder)
}

func (rcv *MetricHistogramResult) UnPack() *MetricHistogramResultT {
	if rcv == nil {
		return nil
	}
	t := &MetricHistogramResultT{}
	t.Timestamp = rcv.Timestamp()
	t.Period = rcv.Period()
	t.Histogram = rcv.Histogram(nil).UnPack()
	return t
}

type MetricHistogramResult struct {
	_tab flatbuffers.Table
}

func GetRootAsMetricHistogramResult(buf []byte, offset flatbuffers.UOffsetT) *MetricHistogramResult {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MetricHistogramResult{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MetricHistogramResult) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MetricHistogramResult) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MetricHistogramResult) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MetricHistogramResult) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *MetricHistogramResult) Period() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MetricHistogramResult) MutatePeriod(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *MetricHistogramResult) Histogram(obj *Histogram) *Histogram {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Histogram)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MetricHistogramResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func MetricHistogramResultAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(0, timestamp, 0)
}
func MetricHistogramResultAddPeriod(builder *flatbuffers.Builder, period int32) {
	builder.PrependInt32Slot(1, period, 0)
}
func MetricHistogramResultAddHistogram(builder *flatbuffers.Builder, histogram flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(histogram), 0)
}
func MetricHistogramResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type MetricHistogramResultListT struct {
	Results []*MetricHistogramResultT
}

func MetricHistogramResultListPack(builder *flatbuffers.Builder, t *MetricHistogramResultListT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	resultsOffset := flatbuffers.UOffsetT(0)
	if t.Results != nil {
		resultsLength := len(t.Results)
		resultsOffsets := make([]flatbuffers.UOffsetT, resultsLength)
		for j := 0; j < resultsLength; j++ {
			resultsOffsets[j] = MetricHistogramResultPack(builder, t.Results[j])
		}
		MetricHistogramResultListStartResultsVector(builder, resultsLength)
		for j := resultsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(resultsOffsets[j])
		}
		resultsOffset = builder.EndVector(resultsLength)
	}
	MetricHistogramResultListStart(builder)
	MetricHistogramResultListAddResults(builder, resultsOffset)
	return MetricHistogramResultListEnd(builder)
}

func (rcv *MetricHistogramResultList) UnPack() *MetricHistogramResultListT {
	if rcv == nil {
		return nil
	}
	t := &MetricHistogramResultListT{}
	resultsLength := rcv.ResultsLength()
	t.Results = make([]*MetricHistogramResultT, resultsLength)
	for j := 0; j < resultsLength; j++ {
		x := MetricHistogramResult{}
		rcv.Results(&x, j)
		t.Results[j] = x.UnPack()
	}
	return t
}

type MetricHistogramResultList struct {
	_tab flatbuffers.Table
}

func GetRootAsMetricHistogramResultList(buf []byte, offset flatbuffers.UOffsetT) *MetricHistogramResultList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MetricHistogramResultList{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MetricHistogramResultList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MetricHistogramResultList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MetricHistogramResultList) Results(obj *MetricHistogramResult, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MetricHistogramResultList) ResultsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MetricHistogramResultListStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MetricHistogramResultListAddResults(builder *flatbuffers.Builder, results flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(results), 0)
}
func MetricHistogramResultListStartResultsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MetricHistogramResultListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
