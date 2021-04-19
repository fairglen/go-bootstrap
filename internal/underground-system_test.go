package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckInInvalidEntry(t *testing.T) {
	scenarios := []struct{
		name string
		id int
		stationName string
		timestamp int
		entryMap map[int]map[int]Entry
		expError error
	}{
		{
			name: "duplicate check in",
			id: 1,
			stationName: "A",
			timestamp: 100,
			entryMap: map[int]map[int]Entry{1: {100: Entry{"A", IN}}},
			expError: InvalidCheckIn,
		},
		{
			name: "no check in until checked out",
			id: 1,
			stationName: "A",
			timestamp: 101,
			entryMap: map[int]map[int]Entry{1: {100: Entry{"A", IN}}},
			expError: InvalidCheckIn,
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T){
			u := &UndergroundSystem{scenario.entryMap}
			err := u.CheckIn(scenario.id, scenario.stationName, scenario.timestamp)
			require.ErrorIs(t, scenario.expError, err)
		})
	}
}

func TestCheckIn(t *testing.T) {
	scenarios := []struct {
		name     string
		id int
		ts int
		entry    Entry
		entryMap map[int]map[int]Entry
	}{
		{
			name:     "valid check-in",
			id: 1,
			ts: 100,
			entry:    Entry{"A", IN},
			entryMap: map[int]map[int]Entry{},
		},
		{
			name:     "valid check-in",
			id: 2,
			ts: 100,
			entry:    Entry{"A", IN},
			entryMap: map[int]map[int]Entry{1: {101: Entry{"A", IN}}},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			us := &UndergroundSystem{scenario.entryMap}
			err := us.CheckIn(
				scenario.id,
				scenario.entry.stationName,
				scenario.ts,
			)
			require.NoError(t, err)
			require.Equal(t, us.entryMap[scenario.id][scenario.ts], scenario.entry)
		})
	}
}

func TestCheckOut(t *testing.T) {
	scenarios := []struct {
		name     string
		id int
		ts int
		entry    Entry
		entryMap map[int]map[int]Entry
	}{
		{
			name:     "valid check-out",
			id: 1,
			ts: 101,
			entry:    Entry{"A", OUT},
			entryMap: map[int]map[int]Entry{1: {100: Entry{"A", IN}}},
		},
	}
	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			us := &UndergroundSystem{scenario.entryMap}
			err := us.CheckOut(
				scenario.id,
				scenario.entry.stationName,
				scenario.ts,
			)
			require.NoError(t, err)
			require.Equal(t, us.entryMap[scenario.id][scenario.ts], scenario.entry)
		})
	}
}
