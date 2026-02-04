package main

import (
	"encoding/csv"
	"os"
)

var (
	columnTitles = []string{"date", "recorded_time"}
)

func writeToCSV(entries []TimeEntry) error {
	defer logTime("Write entries to csv")()
	out("Writing entries to csv")
	defer out("Written entries to csv")
	// https://stackoverflow.com/a/53840463
	file, err := os.Create("export.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	var exportArray [][]string
	exportArray = append(exportArray, columnTitles)
	for _, entry := range entries {
		recordedTime := MyTime{seconds: entry.time}
		entryArray := []string{entry.date, recordedTime.toString()}
		exportArray = append(exportArray, entryArray)
	}

	w.WriteAll(exportArray)

	return nil
}
