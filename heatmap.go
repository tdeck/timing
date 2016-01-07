package main

import "os"
import "log"
import "bufio"
import "regexp"
import "strconv"
import "fmt"

var actionRegex, _ = regexp.Compile(`(\d\d):(\d\d) (.+)`)

func main() {
	path := os.Getenv("TIMING_LOGFILE")
	if path == "" {
		log.Panicln("TIMING_LOGFILE not set")
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// Print vertical numbers
	// e.g.	1 1
	//		2 3
	fmt.Printf("Time:       ")
	for i := 0; i < 24; i++ {
		fmt.Printf("%d ", i/10)
	}
	fmt.Printf("\n            ")
	for i := 0; i < 24; i++ {
		fmt.Printf("%d ", i%10)
	}
	fmt.Printf("\n")

	scanner := bufio.NewScanner(file)

	date := ""
	startmin := -1
	var bins [48]int
	for scanner.Scan() {
		line := scanner.Text()

		switch line[0] {
		case '#': // do nothing
		case '=':
			if date != "" {
				fmt.Printf("%s: ", date)
				printHeatmap(&bins)
			}

			for i := range bins {
				bins[i] = 0
			}
			date = line[2:]
		default:
			match := actionRegex.FindStringSubmatch(line)
			hour, _ := strconv.Atoi(match[1])
			hourmin, _ := strconv.Atoi(match[2])
			action := match[3]

			min := hour*60 + hourmin
			if action == "b" || action == "d" {
				if startmin == -1 {
					log.Panicf("End without start for line %s\n", line)
				}
				recordInterval(startmin, min, &bins)
				startmin = -1
			} else {
				if startmin == -1 {
					// We just ignore double starts, since daylog does the same
					startmin = min
				}
			}
		}
	}
}

func recordInterval(startmin int, endmin int, bins *[48]int) {
	mins := endmin - startmin
	for b := startmin / 30; mins > 0; b++ {
		space := 30 - bins[b]
		added := space
		if mins < space {
			added = mins
		}

		mins -= added
		bins[b] += added

		if bins[b] > 30 {
			log.Panicln("Overlapping interval")
		}
	}
}

func printHeatmap(bins *[48]int) {
	for i := range bins {
		switch bins[i] / 5 {
		case 0: // 0-5 minutes
			fmt.Printf(" ")
		case 1, 2: // 5-15 minutes
			fmt.Printf(".")
		case 3: // 15-20 minutes
			fmt.Printf("x")
		case 4, 5: // 20-30 minutes
			fmt.Printf("X")
		case 6: // All 30 minutes
			fmt.Printf("#")
		}
	}
	fmt.Printf("\n")
}
