package internal

import (
	"errors"
	"sort"
)

var InvalidCheckIn = errors.New("A customer can only be checked into one place at a time.")

type CheckType bool

const (
	IN  = true
	OUT = false
)

type Entry struct {
	stationName string
	checkType   CheckType
}

type UndergroundSystem struct {
	// {id: {ts: {sn: "<stationName>"}}}
	entryMap map[int]map[int]Entry
}

func (u *UndergroundSystem) CheckIn(id int, sn string, ts int) error {
	if entries := u.entryMap[id]; entries != nil {
		keys := make([]int, 0, len(entries))
		for k := range entries {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			if entries[k].checkType == IN && k <= ts {
				return InvalidCheckIn
			}
		}
		// new and valid entry for a given customer
		entries[ts] = Entry{sn, IN}
		return nil
	}
	// new customer
	u.entryMap[id] = map[int]Entry{ts: {sn, IN}}
	return nil
}
