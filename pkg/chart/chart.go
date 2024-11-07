package chart

import (
	"fmt"

	"github.com/AaravSibbal/SquashWebsite/pkg/elo"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func genarateXAxisArr() []int {
	arr := make([]int, 0, 20)
	
	for i:=1; i<21; i++{
		arr = append(arr, i)
	}

	return arr
}

func eloToLineData(matches []*elo.MatchJson, playerName string) []opts.LineData {
	items := make([]opts.LineData, 0, len(matches))

	for _, match := range matches {
		items = append(items, opts.LineData{Value: getElo(match, getElo(match, playerName))})
	}

	return items
}

func getElo(match *elo.MatchJson, playerName string) string {
	if match.PlayerA == playerName {
		return string(match.PlayerARating)
	}

	return string(match.PlayerBRating)
}

func GetEloChart(matches []*elo.MatchJson, playerName string) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("%s-Elo Progression of last 20 games", playerName),
		}),
	)

	line.SetXAxis(genarateXAxisArr()).
		AddSeries(playerName, eloToLineData(matches, playerName))

	return line
}