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
		p := client.NewPersonalization()
		p.AddTos(client.NewEmail("", to)).
			AddCCs(client.NewEmail("", cc))

		m := client.NewMessage(subject)
		m.SetFrom(client.NewEmail("From User", "test@example.com")).
			AddContents(client.NewContent(client.ContentTypePlain, plainContent)).
			AddContents(client.NewContent(client.ContentTypeHTML, htmlContent)).
			AddPersonalizations(p)

		err := cl.Send(m)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Done.")
}
