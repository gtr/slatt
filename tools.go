package main

import (
	"time"

	"github.com/fatih/color"
)

func termLog(text string) {
	start := time.Now()
	now := start.Format("2006-01-02 15:04:05")
	color.Cyan(now)
}
