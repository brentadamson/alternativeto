package alternativeto

var (
	// categories
	calendars          = "Calendars"
	gitHostingServices = "Git Hosting Services"
	importExportData   = "Import and Export Data"
	messagingServices  = "Messaging Services"
	paymentProcessors  = "Payment Processors"
	searchEngines      = "Search Engines"
	spreadsheets       = "Spreadsheets"
	streamingServices  = "Streaming Services"
	twoFactorAuth      = "Two-factor authentication"
	webBrowsers        = "Web Browsers"
	surveySoftware     = "Survey Software"

	// platforms
	platformIOS               = "iOS"
	platformGoogleSheetsAddOn = "Google Sheets Add-on"
	platformLinux             = "Linux"
	platformMac               = "Mac"
	platformWeb               = "Web"
	platformWindows           = "Windows"
)

var allPlatforms = []string{
	platformIOS,
	platformGoogleSheetsAddOn,
	platformLinux,
	platformMac,
	platformWeb,
	platformWindows,
}

// Company is an individual company
type Company struct {
	Name       string   `json:"name"`
	Website    string   `json:"website"`
	Category   string   `json:"category"`
	Platforms  []string `json:"platforms"`
	OpenSource bool     `json:"open_source"`
}

// allCompanies holds all of the companies
// TODO: if this gets unwieldy we can move each category to their own separate file
var allCompanies = []Company{
	// calendars
	{
		Name:       "Google Calendar",
		Website:    "https://calendar.google.com/",
		Category:   calendars,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformMac, platformWeb, platformWindows,
		},
	},
	{
		Name:       "Thunderbird",
		Website:    "https://www.thunderbird.net/",
		Category:   calendars,
		OpenSource: true,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWeb, platformWindows,
		},
	},
	// gitHostingServices
	{
		Name:       "Github",
		Website:    "https://github.com/",
		Category:   gitHostingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformMac, platformWeb, platformWindows,
		},
	},
	{
		Name:       "GitLab",
		Website:    "https://gitlab.com/",
		Category:   gitHostingServices,
		OpenSource: true,
		Platforms: []string{
			platformIOS, platformMac, platformWeb, platformWindows,
		},
	},
	// importExportData
	{
		Name:       "API Connector",
		Website:    "https://mixedanalytics.com/api-connector/",
		Category:   importExportData,
		OpenSource: false,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	{
		Name:       "Apipheny",
		Website:    "https://apipheny.io/",
		Category:   importExportData,
		OpenSource: false,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	{
		Name:       "Data Connector",
		Website:    "https://dataconnector.app/",
		Category:   importExportData,
		OpenSource: true,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	{
		Name:       "KPIBees",
		Website:    "https://kpibees.com/",
		Category:   importExportData,
		OpenSource: false,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	{
		Name:       "Flex.io",
		Website:    "https://www.flex.io/",
		Category:   importExportData,
		OpenSource: false,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	{
		Name:       "ImportJSON / ImportFromWeb",
		Website:    "https://nodatanobusiness.com/",
		Category:   importExportData,
		OpenSource: false,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	{
		Name:       "Supermetrics",
		Website:    "https://supermetrics.com/",
		Category:   importExportData,
		OpenSource: false,
		Platforms: []string{
			platformGoogleSheetsAddOn,
		},
	},
	// messagingServices
	{
		Name:       "Mattermost",
		Website:    "https://mattermost.com/",
		Category:   messagingServices,
		OpenSource: true,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWeb, platformWindows,
		},
	},
	{
		Name:       "Microsoft Teams",
		Website:    "https://www.microsoft.com/en-us/microsoft-teams/group-chat-software",
		Category:   messagingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWeb, platformWindows,
		},
	},
	{
		Name:       "Rocket.Chat",
		Website:    "https://rocket.chat/",
		Category:   messagingServices,
		OpenSource: true,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWeb, platformWindows,
		},
	},
	{
		Name:       "Slack",
		Website:    "https://slack.com/",
		Category:   messagingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWeb, platformWindows,
		},
	},
	// paymentProcessors
	{
		Name:       "Paypal",
		Website:    "https://www.paypal.com/",
		Category:   paymentProcessors,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "Stripe",
		Website:    "https://stripe.com/",
		Category:   paymentProcessors,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	// searchEngines
	{
		Name:       "Bing",
		Website:    "https://www.bing.com/",
		Category:   searchEngines,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "DuckDuckGo",
		Website:    "https://duckduckgo.com/",
		Category:   searchEngines,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "Ecosia",
		Website:    "https://www.ecosia.org/",
		Category:   searchEngines,
		OpenSource: false,
		Platforms: []string{
			platformWeb,
		},
	},
	{
		Name:       "Google",
		Website:    "https://www.google.com/",
		Category:   searchEngines,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "Searx",
		Website:    "https://searx.space/",
		Category:   searchEngines,
		OpenSource: true,
		Platforms: []string{
			platformWeb,
		},
	},
	{
		Name:       "Yahoo",
		Website:    "https://www.yahoo.com/",
		Category:   searchEngines,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	// spreadsheets
	{
		Name:       "Airtable",
		Website:    "https://airtable.com/",
		Category:   spreadsheets,
		OpenSource: false,
		Platforms: []string{
			platformWeb,
		},
	},
	{
		Name:       "Excel",
		Website:    "https://www.microsoft.com/en-us/microsoft-365/excel",
		Category:   spreadsheets,
		OpenSource: false,
		Platforms: []string{
			platformMac, platformWeb, platformWindows,
		},
	},
	{
		Name:       "Google Sheets",
		Website:    "https://sheets.google.com/",
		Category:   spreadsheets,
		OpenSource: false,
		Platforms: []string{
			platformWeb,
		},
	},
	{
		Name:       "LibreOffice Calc",
		Website:    "https://www.libreoffice.org/",
		Category:   spreadsheets,
		OpenSource: true,
		Platforms: []string{
			platformLinux, platformMac, platformWindows,
		},
	},
	{
		Name:       "ONLYOFFICE",
		Website:    "https://www.onlyoffice.com/",
		Category:   spreadsheets,
		OpenSource: true,
		Platforms: []string{
			platformLinux, platformMac, platformWindows,
		},
	},
	// streamingServices
	{
		Name:       "Netflix",
		Website:    "https://www.netflix.com/",
		Category:   streamingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb, platformWindows,
		},
	},
	{
		Name:       "Amazon Prime Video",
		Website:    "https://www.amazon.com/",
		Category:   streamingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "HBO Max",
		Website:    "https://www.hbomax.com/",
		Category:   streamingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "Hulu",
		Website:    "https://www.hulu.com/",
		Category:   streamingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	{
		Name:       "MUBI",
		Website:    "https://mubi.com/",
		Category:   streamingServices,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformWeb,
		},
	},
	// surveySoftware
	{
		Name:       "Onva",
		Website:    "https://www.onva.io/",
		Category:   surveySoftware,
		OpenSource: false,
		Platforms: []string{
			platformWeb,
		},
	},
	{
		Name:       "Survey Monkey",
		Website:    "https://www.surveymonkey.co.uk/",
		Category:   surveySoftware,
		OpenSource: false,
		Platforms: []string{
			platformWeb,
		},
	},
	{
		Name:       "Typeform",
		Website:    "https://www.typeform.com/",
		Category:   surveySoftware,
		OpenSource: false,
		Platforms: []string{
			platformWeb,
		},
	},
	// twoFactorAuth
	{
		Name:       "Google Authenticator",
		Website:    "https://apps.apple.com/us/app/google-authenticator/id388497605",
		Category:   twoFactorAuth,
		OpenSource: false,
		Platforms: []string{
			platformIOS,
		},
	},
	{
		Name:       "Authy",
		Website:    "https://authy.com/",
		Category:   twoFactorAuth,
		OpenSource: false,
		Platforms: []string{
			platformIOS,
		},
	},
	{
		Name:       "1Password",
		Website:    "https://1password.com/",
		Category:   twoFactorAuth,
		OpenSource: false,
		Platforms: []string{
			platformIOS,
		},
	},
	// webBrowsers
	{
		Name:       "Google Chrome",
		Website:    "https://www.google.com/chrome/",
		Category:   webBrowsers,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWindows,
		},
	},
	{
		Name:       "Firefox",
		Website:    "https://www.mozilla.org/en-US/firefox/",
		Category:   webBrowsers,
		OpenSource: true,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWindows,
		},
	},
	{
		Name:       "Safari",
		Website:    "https://www.apple.com/safari/",
		Category:   webBrowsers,
		OpenSource: false,
		Platforms: []string{
			platformIOS, platformLinux, platformMac, platformWindows,
		},
	},
	{
		Name:       "LibreWolf",
		Website:    "https://librewolf-community.gitlab.io/",
		Category:   webBrowsers,
		OpenSource: true,
		Platforms: []string{
			platformLinux,
		},
	},
	{
		Name:       "GNOME Web",
		Website:    "https://wiki.gnome.org/Apps/Web",
		Category:   webBrowsers,
		OpenSource: true,
		Platforms: []string{
			platformLinux,
		},
	},
	{
		Name:       "Pale Moon",
		Website:    "https://www.palemoon.org/",
		Category:   webBrowsers,
		OpenSource: false,
		Platforms: []string{
			platformLinux,
		},
	},
	{
		Name:       "Chromium",
		Website:    "https://www.chromium.org/Home",
		Category:   webBrowsers,
		OpenSource: true,
		Platforms: []string{
			platformLinux,
		},
	},
}
