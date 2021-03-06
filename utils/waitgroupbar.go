package utils

import (
	"sync"

	"github.com/vbauerster/mpb/v6"
	"github.com/vbauerster/mpb/v6/decor"
)

var (
	ScreenshotBar Bar
	CrawlerBar    Bar
)

type Bar struct {
	waitGroup *sync.WaitGroup
	bar       *mpb.Bar
	name      string
	total     int
	mutex     sync.Mutex
}

type WaitGroupBar struct {
	waitGroup sync.WaitGroup
	progress  *mpb.Progress
	mutex     sync.Mutex
	bars      []*Bar
}

func InitWaitGroupBar() *WaitGroupBar {
	var groupBar WaitGroupBar
	groupBar.progress = mpb.New(mpb.WithWidth(1))
	return &groupBar
}

func (groupBar *WaitGroupBar) AddBar(name string, main bool) (newBar Bar) {
	newBar.name = name
	newBar.total = 0
	newBar.bar = groupBar.progress.Add(int64(1),
		mpb.NewSpinnerFiller([]string{}, mpb.SpinnerOnLeft),
		mpb.PrependDecorators(decor.Name("[")),
		mpb.AppendDecorators(
			decor.Name("] ["),
			decor.Name(name),
			decor.Name("] ["),
			decor.Counters(0, "%d / %d"),
			decor.OnComplete(decor.Name("] [Running]"), "] [Finished]"),
		),
	)
	newBar.waitGroup = &groupBar.waitGroup
	groupBar.bars = append(groupBar.bars, &newBar)
	return newBar
}

func (bar *Bar) Add(delta int) {
	bar.mutex.Lock()
	bar.total += delta
	bar.bar.SetTotal(int64(bar.total), false)
	bar.waitGroup.Add(delta)
	bar.mutex.Unlock()
}

func (bar *Bar) Done() {
	bar.mutex.Lock()
	bar.waitGroup.Done()
	bar.bar.IncrBy(1)
	bar.mutex.Unlock()
}

func (groupBar *WaitGroupBar) Wait() {
	groupBar.waitGroup.Wait()
	for _, item := range groupBar.bars {
		item.bar.SetTotal(int64(item.total), true)
	}
	groupBar.progress.Wait()
}
