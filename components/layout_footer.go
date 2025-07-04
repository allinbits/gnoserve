package components

type FooterData struct {
	Analytics  bool
	AssetsPath string
	Sections   []FooterSection
}

type FooterLink struct {
	Label string
	URL   string
}

type FooterSection struct {
	Title string
	Links []FooterLink
}

func EnrichFooterData(data FooterData) FooterData {
	data.Sections = []FooterSection{
		{
			Title: "Footer navigation",
			Links: []FooterLink{
				//{Label: "About", URL: "/about"},
				//{Label: "Docs", URL: "https://docs.gno.land/"},
				//{Label: "Faucet", URL: "https://faucet.gno.land/"},
				//{Label: "Blog", URL: "https://gno.land/r/gnoland/blog"},
				//{Label: "Status", URL: "#"},
			},
		},
		{
			Title: "Social media",
			Links: []FooterLink{
				{Label: "GitHub", URL: "https://github.com/allinbits/gnoserve"},
				//{Label: "Twitter", URL: "#"},
				//{Label: "Discord", URL: "#"},
				//{Label: "Youtube", URL: "#"},
			},
		},
		{
			Title: "Legal",
			Links: []FooterLink{
				{Label: "Terms", URL: "https://github.com/gnolang/gno/blob/master/LICENSE.md"},
				{Label: "Privacy", URL: "https://github.com/gnolang/gno/blob/master/LICENSE.md"},
			},
		},
	}

	return data
}
