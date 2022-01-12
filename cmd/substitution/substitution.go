package main

import (
	"flag"
	"fmt"
	"gosendgridexample/client"
	"gosendgridexample/di"
	"os"
)

type items []string

func (i *items) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *items) Set(v string) error {
	*i = append(*i, v)
	return nil
}

func main() {
	var subject, plainContent, htmlContent string
	var tos, ccs items

	{
		flag.Var(&tos, "to", "To Address")
		flag.Var(&ccs, "cc", "CC Address")
		flag.StringVar(&subject, "subject", "Hello SendGrid;)", "Subject")
		flag.StringVar(&plainContent, "plain", "Hello, world! %name%", "Plain content")
		flag.StringVar(&htmlContent, "html", "<strong>Hello</strong>, <small>world!</small>! <i>%name%</i>", "HTML content")
		flag.Parse()
	}

	{
		cl := di.InjectMailClient()

		m := client.NewMessage(subject)

		for _, to := range tos {
			p := client.NewPersonalization()
			p.AddTos(client.NewEmail("", to)).
				SetSubstitution("%name%", to)

			for _, cc := range ccs {
				p.AddCCs(client.NewEmail("", cc))
			}
			m.AddPersonalizations(p)
		}

		m.SetFrom(client.NewEmail("From User", "test@example.com")).
			AddContents(client.NewContent(client.ContentTypePlain, plainContent)).
			AddContents(client.NewContent(client.ContentTypeHTML, htmlContent))

		err := cl.Send(m)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Done.")
}
