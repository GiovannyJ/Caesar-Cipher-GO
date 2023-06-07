package caesarciphergo

import (
	"log"
	"os"
	"github.com/urfave/cli"
	"fmt"
)


func CLI(){
	app := cli.NewApp()
	app.Name = "Caesar Cipher"
	app.Usage = "Enter a string and a shift to encrypt or decrypt the text"

	startFlags := []cli.Flag{
		cli.StringFlag{
			Name: "t",
			Usage: "Text to be encrypted/decrypted",
		},
		cli.StringFlag{
			Name: "s",
			Usage: "The shift of the text",

		},
	}
	
	app.Commands = []cli.Command{
		{
			Name:  "enc",
			Usage: "Encrypt the text",
			Flags: startFlags,
			Action: func(c *cli.Context) error {
				text := c.String("t")
				shift := c.Int("s")

				encrypted := Enc(text, shift)
				fmt.Println("Encrypted:", encrypted)

				return nil
			},
		},
		{
			Name:  "dec",
			Usage: "Decrypt the text",
			Flags: startFlags,
			Action: func(c *cli.Context) error {
				text := c.String("t")
				shift := c.Int("s")

				decrypted := Dec(text, shift)
				fmt.Println("Decrypted:", decrypted)

				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil{
		log.Fatal(err)
	}
}
