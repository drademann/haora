package cmd

import "fmt"

const version = "1.0.0"

func ExecVersion() error {
	_, err := fmt.Fprintf(Config.Out, "haora version %s\n", version)
	if err != nil {
		return err
	}
	return nil
}
