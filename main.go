package main

import (
	"flag"
	"fmt"
	"os"

	"gosendgridexample/client"
	"gosendgridexample/di"
)

func main() {
	var to, cc, subject, plainContent, htmlContent string

	{
		flag.StringVar(&to, "to", "taro.yamada@example.com", "To Address")
		flag.StringVar(&cc, "cc", "hanako.sato@example.com", "CC Address")
		flag.StringVar(&subject, "subject", "Hello SendGrid;)", "Subject")
		flag.StringVar(&plainContent, "plain", "Hello, world!", "Plain content")
		flag.StringVar(&htmlContent, "html", "<strong>Hello</strong>, <small>world!</small>", "HTML content")
		flag.Parse()
	}

	{
		cl := di.InjectMailClient()
		m := client.NewMessage(subject)
		m.AddFrom(client.NewEmail("From User", "test@example.com")).
			AddTo(client.NewEmail("", to)).
			AddCC(client.NewEmail("", cc)).
			AddPlaintContent(plainContent).
			AddHTMLContent(htmlContent)

		err := cl.Send(m)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Done.")
}
