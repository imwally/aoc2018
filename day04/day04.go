package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	TimeStamp time.Time
	Event     []string
}

type Guard struct {
	ID      string
	Sleep   int
	Minutes []int
}

func processGuard(entries []Entry) *Guard {
	if len(entries) == 0 {
		return nil
	}

	total := 0.0
	id := entries[0].Event[1]
	var minutes []int
	var t1, t2 time.Time

	for i := 0; i < len(entries); i++ {
		status := entries[i].Event[0]
		if status == "falls" {
			t1 = entries[i].TimeStamp
		}
		if status == "wakes" {
			t2 = entries[i].TimeStamp
			for i := t1.Minute(); i < t2.Minute(); i++ {
				minutes = append(minutes, i)
			}
			total += t2.Sub(t1).Minutes()
		}
	}

	return &Guard{id, int(total), minutes}
}

func duplicateMinutes(m []int) (int, int) {
	seen := make(map[int]int)
	max_seen := 0
	max_minute := 0
	for _, minute := range m {
		seen[minute] += 1
		if seen[minute] > max_seen {
			max_seen = seen[minute]
			max_minute = minute
		}
	}

	return max_minute, max_seen
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("specify a file, ya dummy")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var log []Entry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		logEvent := fields[2:]
		logDate := strings.Trim(fields[0], "[")
		logTime := strings.Trim(fields[1], "]")
		timeStamp, _ := time.Parse(time.RFC3339, logDate+"T"+logTime+":00Z")

		log = append(log, Entry{timeStamp, logEvent})
	}

	sort.Slice(log, func(i, j int) bool {
		return log[i].TimeStamp.Before(log[j].TimeStamp)
	})

	totalSleep := make(map[string]int)
	minutesSleep := make(map[string][]int)

	var lines []Entry
	for _, entry := range log {
		if entry.Event[0] == "Guard" {
			g := processGuard(lines)
			if g != nil {
				totalSleep[g.ID] += g.Sleep
				for _, m := range g.Minutes {
					minutesSleep[g.ID] = append(minutesSleep[g.ID], m)
				}
			}
			lines = nil
		}
		lines = append(lines, entry)
	}

	for k, v := range totalSleep {
		fmt.Printf("Guard: %s\tTotal: %d\t", k, v)
		minute, seen := duplicateMinutes(minutesSleep[k])
		fmt.Printf("Minute: %d\t Seen: %d\n", minute, seen)
	}
}
