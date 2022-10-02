package flex

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type view struct {
	w int
	h int
}

func (v *view) SetSize(width int, height int) {
	v.w = width
	v.h = height
}

func (v *view) View() string {
	return ""
}

func TestSetSize(t *testing.T) {
	t.Run("should grow", func(t *testing.T) {
		type testData struct {
			w         int
			h         int
			basis     []int
			grow      []float64
			expectedW []int
		}

		tests := []testData{
			{100, 80, []int{0, 0, 0}, []float64{1, 1}, []int{50, 50}},
			{100, 80, []int{0, 0, 0}, []float64{1, 2, 4}, []int{14, 28, 58}},
			{100, 80, []int{0, 0, 0}, []float64{0, 2, 4}, []int{0, 33, 67}},
			{100, 80, []int{20, 30, 10}, []float64{2, 1, 1}, []int{40, 40, 20}},
		}

		execute := func(td testData) {
			layout := NewContainer(Row)
			views := make([]view, len(td.grow))
			for i, g := range td.grow {
				views[i] = view{}
				layout.AddBox(&views[i], NewStyle().FlexCustom(g, 1, td.basis[i]))
			}
			layout.SetSize(td.w, td.h)
			for i, v := range views {
				if v.w != td.expectedW[i] {
					t.Errorf("Expected width is %d, found %d", td.expectedW[i], v.w)
				}
				if v.h != td.h {
					t.Errorf("Expected height is %d, found %d", td.h, v.h)
				}
			}
		}

		for _, td := range tests {
			testName := fmt.Sprintf(
				"||Container:%dx%d||Basis:%s||Grow:%s||Expected:%s",
				td.w, td.h, iToS(td.basis), fToS(td.grow), iToS(td.expectedW))

			t.Run(testName, func(t *testing.T) {
				execute(td)
			})
		}
	})

	t.Run("should shrink", func(t *testing.T) {
		type testData struct {
			w         int
			h         int
			basis     []int
			shrink    []float64
			expectedW []int
		}

		tests := []testData{
			{100, 80, []int{999, 999}, []float64{1, 1}, []int{50, 50}},
			{120, 80, []int{999, 999 + 30, 999}, []float64{1, 1, 1}, []int{30, 30 + 30, 30}},
			{100, 80, []int{None, 90, 40}, []float64{1, 1, 1}, []int{0, 75, 25}},
		}

		execute := func(td testData) {
			layout := NewContainer(Row)
			views := make([]view, len(td.basis))
			for i, g := range td.basis {
				views[i] = view{}
				layout.AddBox(&views[i], NewStyle().FlexCustom(1, td.shrink[i], g))
			}
			layout.SetSize(td.w, td.h)
			for i, v := range views {
				if v.w != td.expectedW[i] {
					t.Errorf("Expected width is %d, found %d", td.expectedW[i], v.w)
				}
				if v.h != td.h {
					t.Errorf("Expected height is %d, found %d", td.h, v.h)
				}
			}
		}

		for _, td := range tests {
			testName := fmt.Sprintf(
				"||Container:%dx%d||Basis:%s||Grow:%s||Expected:%s",
				td.w, td.h, iToS(td.basis), fToS(td.shrink), iToS(td.expectedW))

			t.Run(testName, func(t *testing.T) {
				execute(td)
			})
		}
	})

	t.Run("max size", func(t *testing.T) {
		type testData struct {
			w         int
			h         int
			maxSize   []int
			expectedW []int
		}

		tests := []testData{
			{100, 80, []int{30, 120}, []int{30, 70}},
			{100, 80, []int{30, 60}, []int{30, 60}},
		}

		execute := func(td testData) {
			layout := NewContainer(Row)
			views := make([]view, len(td.maxSize))
			for i, m := range td.maxSize {
				views[i] = view{}
				layout.AddBox(
					&views[i],
					NewStyle().MaxSize(m).Flex(1),
				)
			}
			layout.SetSize(td.w, td.h)
			for i, v := range views {
				if v.w != td.expectedW[i] {
					t.Errorf("Expected width is %d, found %d", td.expectedW[i], v.w)
				}
				if v.h != td.h {
					t.Errorf("Expected height is %d, found %d", td.h, v.h)
				}
			}
		}

		for _, td := range tests {
			testName := fmt.Sprintf(
				"||Container:%dx%d||MaxSize:%s||Expected:%s",
				td.w, td.h, iToS(td.maxSize), iToS(td.expectedW))

			t.Run(testName, func(t *testing.T) {
				execute(td)
			})
		}
	})

	t.Run("min size", func(t *testing.T) {
		type testData struct {
			w         int
			h         int
			basis     []int
			minSize   []int
			expectedW []int
		}

		tests := []testData{
			{100, 80, []int{999, 999}, []int{70, 0}, []int{70, 30}},
		}

		execute := func(td testData) {
			layout := NewContainer(Row)
			views := make([]view, len(td.minSize))
			for i, m := range td.minSize {
				views[i] = view{}
				layout.AddBox(
					&views[i],
					NewStyle().MinSize(m).FlexCustom(1, 1, td.basis[i]),
				)
			}
			layout.SetSize(td.w, td.h)
			for i, v := range views {
				if v.w != td.expectedW[i] {
					t.Errorf("Expected width is %d, found %d", td.expectedW[i], v.w)
				}
				if v.h != td.h {
					t.Errorf("Expected height is %d, found %d", td.h, v.h)
				}
			}
		}

		for _, td := range tests {
			testName := fmt.Sprintf(
				"||Container:%dx%d||MinSize:%s||Expected:%s",
				td.w, td.h, iToS(td.minSize), iToS(td.expectedW))

			t.Run(testName, func(t *testing.T) {
				execute(td)
			})
		}
	})
}

func fToS(values []float64) string {
	s := make([]string, len(values))
	for i, v := range values {
		s[i] = strconv.FormatFloat(v, 'f', 1, 64)
	}
	return strings.Join(s, ",")
}

func iToS(values []int) string {
	s := make([]string, len(values))
	for i, v := range values {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}
