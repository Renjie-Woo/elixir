package progressBar

import (
	"elixir/utils"
	"errors"
	"fmt"
	"strconv"
)

type ProgressBar struct {
	title                             string  // 进度条标题
	rateGraph                         string  // 进度条图标
	totalCount                        float64 // 数据总量
	currentCount                      float64 // 当前完成量
	percent                           float64 // 完成比例
	currentBar                        string  // 当前条形图
	unit                              string  // 单位
	hasTotalOrCurrentCountInitialized bool
}

func (pb *ProgressBar) SetTitle(title string) {
	if utils.IsStringEmpty(title) {
		title = defaultTitle
	}
	pb.title = title
}

func (pb *ProgressBar) GetTitle() string {
	return pb.title
}

func (pb *ProgressBar) SetGraph(graph string) {
	if utils.IsStringEmpty(graph) {
		graph = defaultGraph
	}
	pb.rateGraph = graph
}

func (pb *ProgressBar) GetGraph() string {
	return pb.rateGraph
}

// SetTotalCount
// support all number format such as int,int32,int64,uint,uint32,uint64,float,float32,float64, and their string ones
func (pb *ProgressBar) SetTotalCount(total interface{}) error {
	totalStr := fmt.Sprintf("%v", total)
	totalFloat, err := strconv.ParseFloat(totalStr, float64Base)
	if err != nil {
		return err
	}
	pb.totalCount = totalFloat
	if pb.hasTotalOrCurrentCountInitialized && pb.totalCount < pb.currentCount {
		return errors.New(fmt.Sprintf(currentCountGreaterThanTotalOneError, pb.currentCount, pb.totalCount))
	}
	pb.hasTotalOrCurrentCountInitialized = true
	return nil
}

// SetCurrentCount
// support all number format such as int,int32,int64,uint,uint32,uint64,float,float32,float64, and their string ones
func (pb *ProgressBar) SetCurrentCount(current interface{}) error {
	if current == nil {
		current = 0
	}
	currentStr := fmt.Sprintf("%v", current)
	currentFloat, err := strconv.ParseFloat(currentStr, float64Base)
	if err != nil {
		return err
	}
	pb.currentCount = currentFloat
	if pb.hasTotalOrCurrentCountInitialized && pb.totalCount < pb.currentCount {
		return errors.New(fmt.Sprintf(currentCountGreaterThanTotalOneError, pb.currentCount, pb.totalCount))
	}
	pb.hasTotalOrCurrentCountInitialized = true
	return nil
}

func (pb *ProgressBar) setPercent(percent float64) {
	pb.percent = percent
}

func (pb *ProgressBar) getPercent() float64 {
	return pb.percent
}

func (pb *ProgressBar) setCurrentBar() string {
	var currentPercent = pb.getPercent()
	var newestPercent = pb.currentCount / pb.totalCount
	increase := int(newestPercent*100) - int(currentPercent*100)
	pb.setPercent(newestPercent)
	for i := 0; i < increase; i++ {
		pb.currentBar += pb.GetGraph()
	}
	return pb.currentBar
}

func (pb *ProgressBar) SetUnit(uint string) {
	pb.unit = uint
}

func NewProgressBar(title string, current, total interface{}) ProgressBar {
	var bar = ProgressBar{}
	bar.SetTitle(title)
	bar.SetCurrentCount(current)
	bar.SetTotalCount(total)
	bar.SetGraph(defaultGraph)
	return bar
}

func (pb *ProgressBar) Run(current interface{}) {
	pb.SetCurrentCount(current)
	currentBar := pb.setCurrentBar()
	rate := pb.getPercent() * 100
	strRate := utils.ParseFloatToStringWithAccuracy(rate, 2)
	strCurrent := utils.ParseFloatToStringWithAccuracy(pb.currentCount, 1)
	strTotal := utils.ParseFloatToStringWithAccuracy(pb.totalCount, 1)
	fmt.Printf("\r%s: [%-100s]%8s%%  %8s %s/%s %s", pb.title, currentBar, strRate, strCurrent, pb.unit, strTotal, pb.unit)
	if pb.currentCount == pb.totalCount {
		fmt.Println()
	}
}
