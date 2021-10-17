package utils

import (
	"github.com/wcharczuk/go-chart"
	"os"
	"path/filepath"
)

func DrawTimeChart(randSlice, incSlice, decSlice []Resources, algorithm string) {
	var xRSliceSize, xISliceSize, xDSliceSize []float64
	var yRSliceTime, yISliceTime, yDSliceTime []float64

	for i := 0; i < len(randSlice); i++ {
		xRSliceSize = append(xRSliceSize, float64(randSlice[i].Size))
		xISliceSize = append(xISliceSize, float64(incSlice[i].Size))
		xDSliceSize = append(xDSliceSize, float64(decSlice[i].Size))
	}
	for i := 0; i < len(randSlice); i++ {
		yRSliceTime = append(yRSliceTime, float64(randSlice[i].Time))
		yISliceTime = append(yISliceTime, float64(incSlice[i].Time))
		yDSliceTime = append(yDSliceTime, float64(decSlice[i].Time))
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 10,
			},
		},
		XAxis: chart.XAxis{
			Name: "Size of Slice (N)",
		},
		YAxis: chart.YAxis{
			Name: "Time (ms)",
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

func DrawSwapsChart(randSlice, incSlice, decSlice []Resources, algorithm string) {
	var xRSliceSize, xISliceSize, xDSliceSize []float64
	var yRSliceSwaps, yISliceSwaps, yDSliceSwaps []float64

	for i := 0; i < len(randSlice); i++ {
		xRSliceSize = append(xRSliceSize, float64(randSlice[i].Size))
		xISliceSize = append(xISliceSize, float64(incSlice[i].Size))
		xDSliceSize = append(xDSliceSize, float64(decSlice[i].Size))
	}
	for i := 0; i < len(randSlice); i++ {
		yRSliceSwaps = append(yRSliceSwaps, float64(randSlice[i].Swaps))
		yISliceSwaps = append(yISliceSwaps, float64(incSlice[i].Swaps))
		yDSliceSwaps = append(yDSliceSwaps, float64(decSlice[i].Swaps))
	}

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 0,
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

func DrawComparisonsChart(randSlice, incSlice, decSlice []Resources, algorithm string) {
	var xRSliceSize, xISliceSize, xDSliceSize []float64
	var yRSliceComparisons, yISliceComparisons, yDSliceComparisons []float64

	for i := 0; i < len(randSlice); i++ {
		xRSliceSize = append(xRSliceSize, float64(randSlice[i].Size))
		xISliceSize = append(xISliceSize, float64(incSlice[i].Size))
		xDSliceSize = append(xDSliceSize, float64(decSlice[i].Size))
	}
	for i := 0; i < len(randSlice); i++ {
		yRSliceComparisons = append(yRSliceComparisons, float64(randSlice[i].Comparisons))
		yISliceComparisons = append(yISliceComparisons, float64(incSlice[i].Comparisons))
		yDSliceComparisons = append(yDSliceComparisons, float64(decSlice[i].Comparisons))
	}
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  0,
				Left: 0,
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
				YValues: yRSliceComparisons,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + incSlice[0].SortingOrder + " slice order",
				XValues: xISliceSize,
				YValues: yISliceComparisons,
			},
			chart.ContinuousSeries{
				Name:    algorithm + " - " + decSlice[0].SortingOrder + " slice order",
				XValues: xDSliceSize,
				YValues: yDSliceComparisons,
			},
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
