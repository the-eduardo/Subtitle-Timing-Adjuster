package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseTimestamp(timestamp string) (time.Duration, error) {
    // Split timestamp into parts (HH:MM:SS,mmm)
    parts := strings.Split(timestamp, ":")

    if len(parts) != 3 {
      return 0, fmt.Errorf("invalid timestamp format")
    }

    hours, err := strconv.Atoi(parts[0])
    if err != nil {
    return 0, err
    }
    minutes, err := strconv.Atoi(parts[1])
    if err != nil {
      return 0, err
    }


    // Split seconds and milliseconds
    secondsParts := strings.Split(parts[2], ",")
    if len(secondsParts) != 2 {
        return 0, fmt.Errorf("invalid seconds format")
    }
    seconds, err := strconv.Atoi(secondsParts[0])
    if err != nil {
        return 0, err
    }
    milliseconds, err := strconv.Atoi(secondsParts[1])
    if err != nil {
        return 0, err
    }
    duration := time.Duration(hours)*time.Hour +
                time.Duration(minutes)*time.Minute +
                time.Duration(seconds)*time.Second +
                time.Duration(milliseconds)*time.Millisecond
                
    return duration, nil
}


func formatTimestamp(duration time.Duration) string {
    hours := int(duration.Hours())
    minutes := int(duration.Minutes()) % 60
    seconds := int(duration.Seconds()) % 60
    milliseconds := int(duration.Milliseconds()) % 1000
    
    return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, milliseconds)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: fixsub <filename> [offset_ms]")
        os.Exit(1)
    }
    
    filename := os.Args[1]
    offsetMs := -500 // default offset
    
    if len(os.Args) > 2 {
        var err error
        offsetMs, err = strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Printf("Invalid offset value: %v\n", err)
            os.Exit(1)
        }
    }
    
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()
    
    // Create a new file for output
    outputFile, err := os.Create(filename + ".adjusted.srt")
    if err != nil {
        fmt.Printf("Error creating output file: %v\n", err)
        os.Exit(1)
    }
    defer outputFile.Close()
    
    scanner := bufio.NewScanner(file)
    writer := bufio.NewWriter(outputFile)
    
    for scanner.Scan() {
        line := scanner.Text()
        
        // Check if line contains the timestamp default format (00:00:01,543 --> 00:00:06,381) 
        if strings.Contains(line, " --> ") {
            timestamps := strings.Split(line, " --> ")
            if len(timestamps) != 2 {
                continue
            }
            
            // Parse and adjust start time
            startTime, err := parseTimestamp(timestamps[0])
            if err != nil {
                continue
            }
            
            // Parse and adjust end time
            endTime, err := parseTimestamp(timestamps[1])
            if err != nil {
                continue
            }
            
            // Apply offset
            offset := time.Duration(offsetMs) * time.Millisecond
            startTime += offset
            endTime += offset
            
            // Write adjusted timestamp
            fmt.Fprintf(writer, "%s --> %s\n", formatTimestamp(startTime), formatTimestamp(endTime))
        } else {
            fmt.Fprintln(writer, line)
        }
    }
    
    writer.Flush()
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("Adjusted subtitles saved to %s\n", filename+".adjusted.srt")
}

