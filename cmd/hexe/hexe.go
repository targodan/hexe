package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/targodan/go-errors"
	"github.com/targodan/hexe/decoder"
	"github.com/targodan/hexe/endianess"
	"github.com/targodan/hexe/interpreter"
	"github.com/urfave/cli/v2"
)

var defaultDecoders = []string{"raw", "hex", "base64"}
var defaultInterpreters = []string{"int"}

func decode(c *cli.Context) error {
	if c.NArg() != 1 {
		return errors.New("Expected exactly one argument. Try --help if you need help.")
	}

	decoderNames := c.StringSlice("decoder")
	decoders := make([]decoder.Decoder, len(decoderNames))
	for i, dec := range decoderNames {
		var ok bool
		decoders[i], ok = decoder.Get(dec)
		if !ok {
			return errors.Newf("Unknown decoder \"%s\".", dec)
		}
	}

	interpreterNames := c.StringSlice("interpreter")
	interpreters := make([]interpreter.Interpreter, len(interpreterNames))
	for i, inter := range interpreterNames {
		var ok bool
		interpreters[i], ok = interpreter.Get(inter)
		if !ok {
			return errors.Newf("Unknown interpreter \"%s\".", inter)
		}
	}

	var inputBuff *bytes.Buffer
	arg := c.Args().Get(0)
	if arg == "-" {
		inputBuff = &bytes.Buffer{}
		inputBuff.ReadFrom(os.Stdin)
	} else {
		inputBuff = bytes.NewBufferString(arg)
	}
	input := inputBuff.Bytes()

	for _, decoder := range decoders {
		fmt.Printf("===== decoding input as %s =====\n", decoder.Name())

		data, err := decoder.Decode(bytes.NewReader(input))
		if err != nil {
			fmt.Printf("Failed to decode: %v\n", err)
			fmt.Println()
			continue
		}

		for _, interpreter := range interpreters {
			interpretations, err := interpreter.Interpret(data)
			for _, interpretation := range interpretations {
				fmt.Printf("Interpreted as %s", interpretation.Description)
				if interpretation.Endianess != endianess.Independent {
					fmt.Printf(" (%s)", interpretation.Endianess)
				}
				fmt.Printf(": %s\n", interpretation.Value)
			}
			if len(interpretations) == 0 && err != nil {
				fmt.Printf("Error interpreting: %v\n", err)
			}
		}

		fmt.Println()
	}

	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			&cli.Command{
				Name:  "decode",
				Usage: "Decodes a given value and interprets it.",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "decoder,d",
						Usage: "List of decoders to try.",
						Value: cli.NewStringSlice(defaultDecoders...),
					},
					&cli.StringSliceFlag{
						Name:  "interpreter,i",
						Usage: "List of interpreters to try.",
						Value: cli.NewStringSlice(defaultInterpreters...),
					},
				},
				Action: decode,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
