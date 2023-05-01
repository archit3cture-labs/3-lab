package lang

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/archit3cture-labs/3-lab/painter"
)

// Parser �쳺 ��������� ��� � �������� io.Reader �� ��������� ������ �������� ����������� ������� ��������.
type Parser struct {
}

// Parse reads and parses input from the provided io.Reader and returns the corresponding list of painter.Operation.
func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	var result []painter.Operation
	for scanner.Scan() { // loop through the input stream using the scanner
		commandLine := scanner.Text()

		oprtn := parse(commandLine) // parse the command line into an operation
		if oprtn == nil {
			return nil, fmt.Errorf("Failed to parse this command: %s", commandLine)
		}

		// Replace any previous BgRectangle operation with the new one
		if bgRect, ok := oprtn.(*painter.BgRectangle); ok {
			for i, oldOp := range result {
				if _, ok := oldOp.(*painter.BgRectangle); ok {
					result[i] = bgRect
					break
				}
			}
		}
		result = append(result, oprtn)
	}
	return result, nil
}

func parse(commandLine string) painter.Operation {
	parts := strings.Split(commandLine, " ")
	instruction := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	var iArgs []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err == nil {
			iArgs = append(iArgs, i)
		}
	}

	var figureOps []painter.Figure

	switch instruction {
	case "white":
		return painter.OperationFunc(painter.WhiteFill)
	case "green":
		return painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		return &painter.BgRectangle{X1: iArgs[0], Y1: iArgs[1], X2: iArgs[2], Y2: iArgs[3]}
	case "figure":
		clr := color.RGBA{R: 255, G: 255, B: 0, A: 1}
		figure := painter.Figure{X: iArgs[0], Y: iArgs[1], C: clr}
		figureOps = append(figureOps, figure)
		return &figure
	case "move":
		return &painter.Move{X: iArgs[0], Y: iArgs[1], Figures: figureOps}
	case "reset":
		figureOps = figureOps[0:0]
		return painter.OperationFunc(painter.ResetScreen)
	case "update":
		return painter.UpdateOp
	}
	return nil
}