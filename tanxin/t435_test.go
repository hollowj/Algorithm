package tanxin

import (
	"fmt"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/stretchr/testify/assert"
)

type Intervals struct {
	slice [][]int
}

func NewIntervals() *Intervals {
	return &Intervals{slice: make([][]int, 0)}
}
func (this *Intervals) Add(t []int) {
	for _, ints := range this.slice {
		if ints[1] == t[0] {
			ints[1] = t[1]
			return
		}
		if ints[0] == t[1] {
			ints[0] = t[0]
			return
		}
	}

	this.slice = append(this.slice, t)

}
func (this *Intervals) checkOverlap(t []int) bool {
	minX, maxX := t[0], t[1]
	for _, ints := range this.slice {
		if !(ints[1] <= minX || maxX <= ints[0]) {
			return true
		}
	}
	return false
}

//	func eraseOverlapIntervals(intervals [][]int) int {
//		sort.Slice(intervals, func(i, j int) bool {
//			return math.Abs(float64(intervals[i][1]-intervals[i][0])) < math.Abs(float64(intervals[j][1]-intervals[j][0]))
//		})
//		t := NewIntervals()
//
//		index := 0
//		t.Add(intervals[0])
//		index++
//		delCount := 0
//		for index < len(intervals) {
//			if t.checkOverlap(intervals[index]) {
//				delCount++
//			} else {
//				t.Add(intervals[index])
//			}
//			index++
//		}
//		return delCount
//
// }
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	avaliableCount := 1
	maxX := intervals[0][1]
	for index := 1; index < len(intervals); index++ {
		if maxX <= intervals[index][0] {
			maxX = intervals[index][1]
			avaliableCount++
		}
	}
	return len(intervals) - avaliableCount
}

func TestT435(t *testing.T) {
	assert.Equal(t, 1, eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	assert.Equal(t, 2, eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	assert.Equal(t, 0, eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	assert.Equal(t, 2, eraseOverlapIntervals([][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}))
	assert.Equal(t, 2, eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
	assert.Equal(t, 7, eraseOverlapIntervals([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
	assert.Equal(t, 19, eraseOverlapIntervals([][]int{{-25322, -4602}, {-35630, -28832}, {-33802, 29009}, {13393, 24550}, {-10655, 16361}, {-2835, 10053}, {-2290, 17156}, {1236, 14847}, {-45022, -1296}, {-34574, -1993}, {-14129, 15626}, {3010, 14502}, {42403, 45946}, {-22117, 13380}, {7337, 33635}, {-38153, 27794}, {47640, 49108}, {40578, 46264}, {-38497, -13790}, {-7530, 4977}, {-29009, 43543}, {-49069, 32526}, {21409, 43622}, {-28569, 16493}, {-28301, 34058}}))
}

func TestDraw(t *testing.T) {
	c := [][]int{{-25322, -4602}, {-35630, -28832}, {-33802, 29009}, {13393, 24550}, {-10655, 16361}, {-2835, 10053}, {-2290, 17156}, {1236, 14847}, {-45022, -1296}, {-34574, -1993}, {-14129, 15626}, {3010, 14502}, {42403, 45946}, {-22117, 13380}, {7337, 33635}, {-38153, 27794}, {47640, 49108}, {40578, 46264}, {-38497, -13790}, {-7530, 4977}, {-29009, 43543}, {-49069, 32526}, {21409, 43622}, {-28569, 16493}, {-28301, 34058}}
	minX := c[0][0]
	maxX := c[0][1]
	for i := 1; i < len(c); i++ {
		if c[i][0] < minX {
			minX = c[i][0]
		}
		if c[i][1] > maxX {
			maxX = c[i][1]
		}
	}
	log.Print(minX, maxX)
	f := make([][]float64, 0)
	for _, t := range c {
		f = append(f, []float64{float64(t[0]), float64(t[1])})

	}
	draw2(f)

}

// func draw(values [][]float64) {
//
//	p, err := charts.LineRender(
//		values,
//		charts.TitleTextOptionFunc("Line"),
//		charts.XAxisDataOptionFunc([]string{
//			"min",
//			"max",
//		}),
//		charts.SVGTypeOption(),
//	)
//
//	if err != nil {
//		panic(err)
//	}
//
//	buf, err := p.Bytes()
//	if err != nil {
//		panic(err)
//	}
//
//	os.WriteFile("chart.svg", buf, 0666)
//
// }
func generateLineData(data []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}
func draw2(values [][]float64) {
	page := components.NewPage()
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
	)

	line.SetXAxis([]string{"min", "max"})
	for _, data := range values {
		line.AddSeries(fmt.Sprintf("%v-%v", data[0], data[1]), generateLineData(data))
	}
	f, _ := os.Create("chart.html")
	page.AddCharts(line)

	page.Render(f)

}
