package test

import (
	"image/color"
	"strings"
	"testing"

	"github.com/archit3cture-labs/3-lab/painter"
	"github.com/archit3cture-labs/3-lab/painter/lang"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parse_struct(t *testing.T) {
	tests := []struct {
		op      painter.Operation
		command string
		name    string
	}{
		{
			op:      &painter.Move{X: 100, Y: 100},
			command: "move 100 100",
			name:    "move",
		},
		{
			op:      painter.UpdateOp,
			command: "update",
			name:    "update painter",
		},
		{
			op:      &painter.Figure{X: 200, Y: 200, C: color.RGBA{R: 255, G: 255, B: 0, A: 1}},
			command: "figure 200 200",
			name:    "figure",
		},
		{
			op:      &painter.BgRectangle{X1: 0, Y1: 0, X2: 100, Y2: 100},
			command: "bgrect 0 0 100 100",
			name:    "background of rect",
		},
		{
			op:      nil,
			command: "invalidcommand",
			name:    "command is invalid",
		},
	}

	parser := &lang.Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))
			if tc.op == nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Len(t, ops, 1)
				assert.IsType(t, tc.op, ops[0])
				assert.Equal(t, tc.op, ops[0])
			}
		})
	}
}

func Test_parse_func(t *testing.T) {
	tests := []struct {
		op      painter.Operation
		command string
		name    string
	}{
		{
			op:      painter.OperationFunc(painter.GreenFill),
			command: "green",
			name:    "filling with green",
		},
		{
			op:      painter.OperationFunc(painter.WhiteFill),
			command: "white",
			name:    "filling with white",
		},
		{
			op:      painter.OperationFunc(painter.ResetScreen),
			command: "reset",
			name:    "resetting screen",
		},
	}

	parser := &lang.Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))
			require.NoError(t, err)
			require.Len(t, ops, 1)
			assert.IsType(t, tc.op, ops[0])

		})
	}
}
