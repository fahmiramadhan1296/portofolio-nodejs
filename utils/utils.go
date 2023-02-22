package utils

import (
	"strconv"
	"time"
)


func UniqueId() string {
    uniqueIdWithTimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return "POR"+uniqueIdWithTimestamp
}