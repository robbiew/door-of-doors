package input

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	gd "github.com/robbiew/door-of-doors/common"
)

// Select asks the user to select a item from the given list by the number.
// It shows the given query and list to user. The response is returned as string
// from the list. By default, it checks the input is the number and is not
// out of range of the list and if not returns error. If Loop is true, it continue to
// ask until it receives valid input.
//
// If the user sends SIGINT (Ctrl+C) while reading input, it catches
// it and return it as a error.
func (i *UI) Select(query string, list []string, opts *Options) (string, error) {
	// Set default val
	i.once.Do(i.setDefault)

	// Input must not be empty if no default is specified.
	// Because Select ask user to input by number.
	// If empty, can not transform it to int.
	opts.Required = true

	// Find default index which opts.Default indicates
	defaultIndex := -1
	defaultVal := opts.Default
	if defaultVal != "" {
		for i, item := range list {
			if item == defaultVal {
				defaultIndex = i
			}
		}

		// DefaultVal is set but doesn't exist in list
		if defaultIndex == -1 {
			// This error message is not for user
			// Should be found while development
			return "", fmt.Errorf("opt.Default is specified but item does not exist in list")
		}
	}

	// Construct the query & display it to user
	// var buf bytes.Buffer
	gd.MoveCursor(2, 6)
	fmt.Println(fmt.Sprintf(gd.WhiteHi+"%s\n\n"+gd.Reset, query))

	count := 1
	yLoc1 := 8
	yLoc2 := 8
	for i, item := range list {
		if count < 13 {
			gd.MoveCursor(2, yLoc1)
			fmt.Fprintf(os.Stdout, gd.Cyan+" %d. %s\n"+gd.Reset, i+1, item)
			yLoc1++
		}
		if count > 13 {
			gd.MoveCursor(40, yLoc2)
			fmt.Fprintf(os.Stdout, gd.Cyan+" %d. %s\n"+gd.Reset, i+1, item)
			yLoc2++
		}
		count++

	}

	// buf.WriteString("\n")
	// fmt.Fprintf(i.Writer, buf.String())

	// resultStr and resultErr are return val of this function
	var resultStr string
	var resultErr error
	for {

		// Construct the asking line to input
		var buf bytes.Buffer
		gd.MoveCursor(2, 22)
		buf.WriteString(gd.CyanHi + "Enter a number" + gd.Reset)

		// Add default val if provided
		if defaultIndex >= 0 && !opts.HideDefault {
			buf.WriteString(fmt.Sprintf(" (Default is %d)", defaultIndex+1))
		}

		buf.WriteString(": ")
		fmt.Fprintf(i.Writer, buf.String())

		// Read user input from reader.
		line, err := i.read(opts.readOpts())
		if err != nil {
			resultErr = err
			break
		}

		// line is empty but default is provided returns it
		if line == "" && defaultIndex >= 0 {
			resultStr = list[defaultIndex]
			break
		}

		if line == "" && opts.Required {
			if !opts.Loop {
				resultErr = ErrEmpty
				break
			}

			fmt.Fprintf(i.Writer, "Input must not be empty. Answer by a number.")

			continue
		}

		if line == "q" || line == "Q" {
			os.Exit(0)
		}

		// Convert user input string to int val
		n, err := strconv.Atoi(line)
		if err != nil {
			if !opts.Loop {
				resultErr = ErrNotNumber
				break
			}

			fmt.Fprintf(i.Writer,
				"%q is not a valid input. Answer by a number.\n\n", line)
			continue
		}

		// Check answer is in range of list
		if n < 1 || len(list) < n {
			if !opts.Loop {
				resultErr = ErrOutOfRange
				break
			}

			fmt.Fprintf(i.Writer,
				"%q is not a valid choice. Choose a number from 1 to %d.\n\n",
				line, len(list))
			continue
		}

		// validate input by custom function
		validate := opts.validateFunc()
		if err := validate(line); err != nil {
			if !opts.Loop {
				resultErr = err
				break
			}

			fmt.Fprintf(i.Writer, "Failed to validate input string: %s\n\n", err)
			continue
		}

		// Reach here means it gets ideal input.
		resultStr = list[n-1]
		break
	}

	// Insert the new line for next output
	fmt.Fprintf(i.Writer, "\n")

	return resultStr, resultErr
}
