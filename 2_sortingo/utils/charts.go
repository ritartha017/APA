package utils

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"path/filepath"
)

func DrawTimeChart(randSlice, incSlice, decSlice, halfSortSlice, almostSortSlice []Resources, algorithm string) {
	var xRSliceSize, xISliceSize, xDSliceSize, xHSortSliceSize, xASortSliceSize []float64
	var yRSliceTime, yISliceTime, yDSliceTime, yHSortSliceTime, yASortSliceTime []float64

	for i := 0; i < len(randSlice); i++ {
		xRSliceSize = append(xRSliceSize, float64(randSlice[i].Size))
		xISliceSize = append(xISliceSize, float64(incSlice[i].Size))
		xDSliceSize = append(xDSliceSize, float64(decSlice[i].Size))
		xHSortSliceSize = append(xHSortSliceSize, float64(halfSortSlice[i].Size))
		xASortSliceSize = append(xASortSliceSize, float64(almostSortSlice[i].Size))
	}
	for i := 0; i < len(randSlice); i++ {
		yRSliceTime = append(yRSliceTime, float64(randSlice[i].Time.Microseconds()))
		yISliceTime = append(yISliceTime, float64(incSlice[i].Time.Microseconds()))
		yDSliceTime = append(yDSliceTime, float64(decSlice[i].Time.Microseconds()))
		yHSortSliceTime = append(yHSortSliceTime, float64(halfSortSlice[i].Time.Microseconds()))
		yASortSliceTime = append(yASortSliceTime, float64(almostSortSlice[i].Time.Microseconds()))
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:   0,
				Left:  0,
				Right: 0,
			},
		},
		XAxis: chart.XAxis{
			Name: "Size of Slice (N)",
		},
		YAxis: chart.YAxis{
			Name: "Time (µs)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    algorithm + " - " + randSlice[0].SortingOrder + " slice order",
				XValues: xRSliceSize,
				YValues: yRSliceTime,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + incSlice[0].SortingOrder + " slice order",
				XValues: xISliceSize,
				YValues: yISliceTime,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + decSlice[0].SortingOrder + " slice order",
				XValues: xDSliceSize,
				YValues: yDSliceTime,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + halfSortSlice[0].SortingOrder + " slice order",
				XValues: xHSortSliceSize,
				YValues: yHSortSliceTime,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + almostSortSlice[0].SortingOrder + " slice order",
				XValues: xASortSliceSize,
				YValues: yASortSliceTime,
			},
			chart.AnnotationSeries{
				Annotations: []chart.Value2{
					{XValue: 10, YValue: 0, Label: "10"},
					{XValue: 1000, YValue: 0, Label: "1k"},
					{XValue: 10000, YValue: 0, Label: "10k"},
					{XValue: 30000, YValue: 0, Label: "30k"},
					{XValue: 50000, YValue: 0, Label: "50k"},
					{XValue: 70000, YValue: 0, Label: "70k"},
					{XValue: 100000, YValue: 0, Label: "100k"},
				},
			},
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	newpath := filepath.Join(".", algorithm+"charts")
	os.MkdirAll(newpath, os.ModePerm)

	f, _ := os.Create(algorithm + "charts/" + algorithm + "TimeChart.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func DrawSwapsChart(randSlice, incSlice, decSlice, halfSortSlice, almostSortSlice []Resources, algorithm string) {
	var xRSliceSize, xISliceSize, xDSliceSize, xHSortSliceSize, xASortSliceSize []float64
	var yRSliceSwaps, yISliceSwaps, yDSliceSwaps, yHSortSliceSwaps, yASortSliceSwaps []float64

	for i := 0; i < len(randSlice); i++ {
		xRSliceSize = append(xRSliceSize, float64(randSlice[i].Size))
		xISliceSize = append(xISliceSize, float64(incSlice[i].Size))
		xDSliceSize = append(xDSliceSize, float64(decSlice[i].Size))
		xHSortSliceSize = append(xHSortSliceSize, float64(halfSortSlice[i].Size))
		xASortSliceSize = append(xASortSliceSize, float64(almostSortSlice[i].Size))
	}
	for i := 0; i < len(randSlice); i++ {
		yRSliceSwaps = append(yRSliceSwaps, float64(randSlice[i].Swaps))
		yISliceSwaps = append(yISliceSwaps, float64(incSlice[i].Swaps))
		yDSliceSwaps = append(yDSliceSwaps, float64(decSlice[i].Swaps))
		yHSortSliceSwaps = append(yHSortSliceSwaps, float64(halfSortSlice[i].Swaps))
		yASortSliceSwaps = append(yASortSliceSwaps, float64(almostSortSlice[i].Swaps))
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:   0,
				Left:  0,
				Right: 0,
			},
		},
		XAxis: chart.XAxis{
			Name: "Size of Slice (N)",
		},
		YAxis: chart.YAxis{
			Name: "Swaps (N)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    algorithm + " - " + randSlice[0].SortingOrder + " slice order",
				XValues: xRSliceSize,
				YValues: yRSliceSwaps,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + incSlice[0].SortingOrder + " slice order",
				XValues: xISliceSize,
				YValues: yISliceSwaps,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + decSlice[0].SortingOrder + " slice order",
				XValues: xDSliceSize,
				YValues: yDSliceSwaps,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + halfSortSlice[0].SortingOrder + " slice order",
				XValues: xHSortSliceSize,
				YValues: yHSortSliceSwaps,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + almostSortSlice[0].SortingOrder + " slice order",
				XValues: xASortSliceSize,
				YValues: yASortSliceSwaps,
			},
			// chart.AnnotationSeries{
			// 	Annotations: []chart.Value2{
			// 		{XValue: 10, YValue: 0, Label: "10"},
			// 		{XValue: 10000, YValue: 0, Label: "10k"},
			// 		{XValue: 100000, YValue: 0, Label: "100k"},
			// 		{XValue: 1000000, YValue: 0, Label: "1mln"},
			// 		{XValue: 5000000, YValue: 0, Label: "5mln"},
			// 	},
			// },
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	newpath := filepath.Join(".", algorithm+"charts")
	os.MkdirAll(newpath, os.ModePerm)

	f, _ := os.Create(algorithm + "charts/" + algorithm + "SwapsChart.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func DrawComparisonsChart(randSlice, incSlice, decSlice, halfSortSlice, almostSortSlice []Resources, algorithm string) {
	var xRSliceSize, xISliceSize, xDSliceSize, xHSortSliceSize, xASortSliceSize []float64
	var yRSliceComp, yISliceComp, yDSliceComp, yHSortSliceComp, yASortSliceComp []float64

	for i := 0; i < len(randSlice); i++ {
		xRSliceSize = append(xRSliceSize, float64(randSlice[i].Size))
		xISliceSize = append(xISliceSize, float64(incSlice[i].Size))
		xDSliceSize = append(xDSliceSize, float64(decSlice[i].Size))
		xHSortSliceSize = append(xHSortSliceSize, float64(halfSortSlice[i].Size))
		xASortSliceSize = append(xASortSliceSize, float64(almostSortSlice[i].Size))
	}
	for i := 0; i < len(randSlice); i++ {
		yRSliceComp = append(yRSliceComp, float64(randSlice[i].Comparisons))
		yISliceComp = append(yISliceComp, float64(incSlice[i].Comparisons))
		yDSliceComp = append(yDSliceComp, float64(decSlice[i].Comparisons))
		yHSortSliceComp = append(yHSortSliceComp, float64(halfSortSlice[i].Comparisons))
		yASortSliceComp = append(yASortSliceComp, float64(almostSortSlice[i].Comparisons))
	}
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:   0,
				Left:  0,
				Right: 0,
			},
		},
		XAxis: chart.XAxis{
			Name: "Size of Slice (N)",
		},
		YAxis: chart.YAxis{
			Name: "Comparisons (N)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    algorithm + " - " + randSlice[0].SortingOrder + " slice order",
				XValues: xRSliceSize,
				YValues: yRSliceComp,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + incSlice[0].SortingOrder + " slice order",
				XValues: xISliceSize,
				YValues: yISliceComp,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + decSlice[0].SortingOrder + " slice order",
				XValues: xDSliceSize,
				YValues: yDSliceComp,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + halfSortSlice[0].SortingOrder + " slice order",
				XValues: xHSortSliceSize,
				YValues: yHSortSliceComp,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + almostSortSlice[0].SortingOrder + " slice order",
				XValues: xASortSliceSize,
				YValues: yASortSliceComp,
			},
			// chart.AnnotationSeries{
			// 	Annotations: []chart.Value2{
			// 		{XValue: 10, YValue: 0, Label: "10"},
			// 		{XValue: 10000, YValue: 0, Label: "10k"},
			// 		{XValue: 100000, YValue: 0, Label: "100k"},
			// 		{XValue: 1000000, YValue: 0, Label: "1mln"},
			// 		{XValue: 5000000, YValue: 0, Label: "5mln"},
			// 	},
			// },
		},
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	newpath := filepath.Join(".", algorithm+"charts")
	os.MkdirAll(newpath, os.ModePerm)

	f, _ := os.Create(algorithm + "charts/" + algorithm + "ComparisonsChart.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
