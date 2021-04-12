package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestHello(t *testing.T){
	//given
	exp :="Hello World!"
	//when
	act := Hello()
	//then
	require.Equal(t, exp, act, "Messages don't match")
}

func TestPersonalHello(t *testing.T){
	scenarios := []struct {
		name string
	}{
		{name: "Leo" },
		{name: "Estrela" },
		{name: "Batatas" },
	}
	for _, scenario := range scenarios{
		t.Run(fmt.Sprintf("returns personal message to %s.", scenario.name), func(t *testing.T){
			//given
			exp :=fmt.Sprintf("Hello %s!", scenario.name)
			//when
			act := Hello(scenario.name)
			//then
			require.Equal(t, exp, act, "Messages don't match")
		})
	}
}
