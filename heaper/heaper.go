package heaper

import (
	"container/heap"
	"fmt"
)

// Constant map
func GetAlertGroupMap() map[string]string {
	return map[string]string{
		"RedAlert":    "RED",
		"YellowAlert": "YELLOW",
	}
}

const (
	highAlertSeverity   = 1
	mediumAlertSeverity = 2
	lowAlertSeverity    = 3
)

type alert struct {
	Message  string
	Group    string
	Severity int
}

type alertHeap []alert

func (ah alertHeap) Less(i, j int) bool {
	// agm := GetAlertGroupMap()
	if ah[i].Group == ah[j].Group {
		if ah[i].Severity < ah[j].Severity {
			return true
		} else {
			return false
		}
	}

	if ah[i].Group == "RED" && ah[j].Group == "YELLOW" {
		return true
	}

	return false
}

func (ah alertHeap) Len() int {
	return len(ah)
}

func (ah alertHeap) Swap(i, j int) {
	ah[i], ah[j] = ah[j], ah[i]
}

func (ah *alertHeap) Push(x interface{}) {
	*ah = append(*ah, x.(alert))
}

func (ah *alertHeap) Pop() interface{} {
	oldh := *ah
	n := len(oldh)
	e := oldh[n-1]
	*ah = oldh[0 : n-1]
	return e
}

func Run() {
	agm := GetAlertGroupMap()
	h := &alertHeap{
		alert{
			Group:    agm["RedAlert"],
			Severity: mediumAlertSeverity,
			Message:  "Rack plonie",
		},
		alert{
			Group:    agm["RedAlert"],
			Severity: lowAlertSeverity,
			Message:  "Serwer plonie",
		},
		alert{
			Group:    agm["YellowAlert"],
			Severity: highAlertSeverity,
			Message:  "BGP session down",
		},
		alert{
			Group:    agm["YellowAlert"],
			Severity: lowAlertSeverity,
			Message:  "Service down",
		},
		alert{
			Group:    agm["RedAlert"],
			Severity: highAlertSeverity,
			Message:  "Serwerownia plonie",
		},
		alert{
			Group:    agm["RedAlert"],
			Severity: highAlertSeverity,
			Message:  "ZOnk",
		},
	}

	heap.Init(h)
	ne := alert{
		Group:    agm["YellowAlert"],
		Severity: mediumAlertSeverity,
		Message:  "UPS sie zalaczyl",
	}
	heap.Push(h, ne)

	for h.Len() > 0 {
		e := heap.Pop(h)
		fmt.Printf("%s: %s %d\n", e.(alert).Group, e.(alert).Message, e.(alert).Severity)

	}
}
