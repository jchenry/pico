package pastes

import (
	"fmt"

	"git.sr.ht/~erock/pico/shared"
	"git.sr.ht/~erock/pico/wish/cms/config"
)

func NewConfigSite() *shared.ConfigSite {
	domain := shared.GetEnv("PASTES_DOMAIN", "pastes.sh")
	email := shared.GetEnv("PASTES_EMAIL", "hello@pastes.sh")
	subdomains := shared.GetEnv("PASTES_SUBDOMAINS", "0")
	port := shared.GetEnv("PASTES_WEB_PORT", "3000")
	dbURL := shared.GetEnv("DATABASE_URL", "")
	protocol := shared.GetEnv("PASTES_PROTOCOL", "https")
	subdomainsEnabled := false
	if subdomains == "1" {
		subdomainsEnabled = true
	}

	intro := "To get started, enter a username.\n"
	intro += "Then create a folder locally (e.g. ~/pastes).\n"
	intro += "Then write your paste post (e.g. feature.patch).\n"
	intro += "Finally, send your files to us:\n\n"
	intro += fmt.Sprintf("scp ~/pastes/* %s:/", domain)

	return &shared.ConfigSite{
		SubdomainsEnabled: subdomainsEnabled,
		ConfigCms: config.ConfigCms{
			Domain:      domain,
			Port:        port,
			Protocol:    protocol,
			Email:       email,
			DbURL:       dbURL,
			Description: "a pastebin for hackers.",
			IntroText:   intro,
			Space:       "pastes",
			Logger:      shared.CreateLogger(),
		},
	}
}