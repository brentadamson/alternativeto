package alternativeto

var (
	// categories
	calendars          = "Calendars"
	gitHostingServices = "Git Hosting Services"
	googleSheetsAddons = "Google Sheets Add-ons"
	messagingServices  = "Messaging Services"
	paymentProcessors  = "Payment Processors"
	searchEngines      = "Search Engines"
	spreadsheets       = "Spreadsheets"
	streamingServices  = "Streaming Services"
	webBrowsers        = "Web Browsers"

	// subcategories
	noSubCategory    = "none"
	importExportData = "Import and Export Data"
)

// Company is an individual company
type Company struct {
	Name        string `json:"name"`
	Website     string `json:"website"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	OpenSource  bool   `json:"open_source"`
}

// allCompanies holds all of the companies
// TODO: if this gets unwieldy we can move each category to their own separate file
var allCompanies = map[string]map[string][]Company{
	calendars: {
		noSubCategory: {
			{
				Name:       "Google Calendar",
				Website:    "https://calendar.google.com/",
				OpenSource: false,
			},
			{
				Name:       "Thunderbird",
				Website:    "https://www.thunderbird.net/",
				OpenSource: true,
			},
		},
	},
	gitHostingServices: {
		noSubCategory: {
			{
				Name:       "Github",
				Website:    "https://github.com/",
				OpenSource: false,
			},
			{
				Name:       "GitLab",
				Website:    "https://gitlab.com/",
				OpenSource: true,
			},
		},
	},
	googleSheetsAddons: {
		importExportData: {
			{
				Name:        "API Connector",
				Website:     "https://mixedanalytics.com/api-connector/",
				OpenSource:  false,
				SubCategory: importExportData,
			},
			{
				Name:        "Apipheny",
				Website:     "https://apipheny.io/",
				OpenSource:  false,
				SubCategory: importExportData,
			},
			{
				Name:        "Data Connector",
				Website:     "https://dataconnector.app/",
				OpenSource:  true,
				SubCategory: importExportData,
			},
			{
				Name:        "KPIBees",
				Website:     "https://kpibees.com/",
				OpenSource:  false,
				SubCategory: importExportData,
			},
			{
				Name:        "Flex.io",
				Website:     "https://www.flex.io/",
				OpenSource:  false,
				SubCategory: importExportData,
			},
			{
				Name:        "ImportJSON / ImportFromWeb",
				Website:     "https://nodatanobusiness.com/",
				OpenSource:  false,
				SubCategory: importExportData,
			},
			{
				Name:        "Supermetrics",
				Website:     "https://supermetrics.com/",
				OpenSource:  false,
				SubCategory: importExportData,
			},
		},
	},
	messagingServices: {
		noSubCategory: {
			{
				Name:       "Mattermost",
				Website:    "https://mattermost.com/",
				OpenSource: true,
			},
			{
				Name:       "Microsoft Teams",
				Website:    "https://www.microsoft.com/en-us/microsoft-teams/group-chat-software",
				OpenSource: false,
			},
			{
				Name:       "Rocket.Chat",
				Website:    "https://rocket.chat/",
				OpenSource: true,
			},
			{
				Name:       "Slack",
				Website:    "https://slack.com/",
				OpenSource: false,
			},
		},
	},
	paymentProcessors: {
		noSubCategory: {
			{
				Name:       "Paypal",
				Website:    "https://www.paypal.com/",
				OpenSource: false,
			},
			{
				Name:       "Stripe",
				Website:    "https://stripe.com/",
				OpenSource: false,
			},
		},
	},
	searchEngines: {
		noSubCategory: {
			{
				Name:       "Bing",
				Website:    "https://www.bing.com/",
				OpenSource: false,
			},
			{
				Name:       "DuckDuckGo",
				Website:    "https://duckduckgo.com/",
				OpenSource: false,
			},
			{
				Name:       "Ecosia",
				Website:    "https://www.ecosia.org/",
				OpenSource: false,
			},
			{
				Name:       "Google",
				Website:    "https://www.google.com/",
				OpenSource: false,
			},
			{
				Name:       "Searx",
				Website:    "https://searx.space/",
				OpenSource: true,
			},
			{
				Name:       "Yahoo",
				Website:    "https://www.yahoo.com/",
				OpenSource: false,
			},
		},
	},
	spreadsheets: {
		noSubCategory: {
			{
				Name:       "Airtable",
				Website:    "https://airtable.com/",
				OpenSource: false,
			},
			{
				Name:       "Excel",
				Website:    "https://www.microsoft.com/en-us/microsoft-365/excel",
				OpenSource: false,
			},
			{
				Name:       "Google Sheets",
				Website:    "https://sheets.google.com/",
				OpenSource: false,
			},
			{
				Name:       "LibreOffice Calc",
				Website:    "https://www.libreoffice.org/",
				OpenSource: true,
			},
			{
				Name:       "ONLYOFFICE",
				Website:    "https://www.onlyoffice.com/",
				OpenSource: true,
			},
		},
	},
	streamingServices: {
		noSubCategory: {
			{
				Name:       "Netflix",
				Website:    "https://www.netflix.com/",
				OpenSource: false,
			},
			{
				Name:       "Amazon Prime Video",
				Website:    "https://www.amazon.com/",
				OpenSource: false,
			},
			{
				Name:       "HBO Max",
				Website:    "https://www.hbomax.com/",
				OpenSource: false,
			},
			{
				Name:       "Hulu",
				Website:    "https://www.hulu.com/",
				OpenSource: false,
			},
			{
				Name:       "MUBI",
				Website:    "https://mubi.com/",
				OpenSource: false,
			},
		},
	},
	webBrowsers: {
		noSubCategory: {
			{
				Name:       "Google Chrome",
				Website:    "https://www.google.com/chrome/",
				OpenSource: false,
			},
			{
				Name:       "Firefox",
				Website:    "https://www.mozilla.org/en-US/firefox/",
				OpenSource: true,
			},
			{
				Name:       "Safari",
				Website:    "https://www.apple.com/safari/",
				OpenSource: false,
			},
			{
				Name:       "LibreWolf",
				Website:    "https://librewolf-community.gitlab.io/",
				OpenSource: true,
			},
			{
				Name:       "GNOME Web",
				Website:    "https://wiki.gnome.org/Apps/Web",
				OpenSource: true,
			},
			{
				Name:       "Pale Moon",
				Website:    "https://www.palemoon.org/",
				OpenSource: false,
			},
			{
				Name:       "Chromium",
				Website:    "https://www.chromium.org/Home",
				OpenSource: true,
			},
		},
	},
}
