package igdb

import (
	"strings"

	IGDBService "steam-list-api.com/pkg/igdb/igdbService"
)

type ImageSize string

const (
	CoverSmall     ImageSize = "cover_small"
	ScreenshotMed  ImageSize = "screenshot_med"
	CoverBig       ImageSize = "cover_big"
	LogoMed        ImageSize = "logo_med"
	ScreenshotBig  ImageSize = "screenshot_big"
	ScreenshotHuge ImageSize = "screenshot_huge"
	Thumb          ImageSize = "thumb"
	Micro          ImageSize = "micro"
	HD             ImageSize = "720p"
	FHD            ImageSize = "1080p"
)

func findCurrentImageSize(url string) ImageSize {
	if strings.Index(url, string(CoverSmall)) > 0 {
		return CoverSmall
	}
	if strings.Index(url, string(ScreenshotMed)) > 0 {
		return ScreenshotMed
	}
	if strings.Index(url, string(CoverBig)) > 0 {
		return CoverBig
	}
	if strings.Index(url, string(LogoMed)) > 0 {
		return LogoMed
	}
	if strings.Index(url, string(ScreenshotBig)) > 0 {
		return ScreenshotBig
	}
	if strings.Index(url, string(ScreenshotHuge)) > 0 {
		return ScreenshotHuge
	}
	if strings.Index(url, string(ScreenshotMed)) > 0 {
		return ScreenshotMed
	}
	if strings.Index(url, string(Thumb)) > 0 {
		return Thumb
	}
	if strings.Index(url, string(Micro)) > 0 {
		return Micro
	}
	if strings.Index(url, string(HD)) > 0 {
		return HD
	}
	if strings.Index(url, string(FHD)) > 0 {
		return FHD
	}
	return ""
}

func SetImageSize(url string, size ImageSize) string {
	currentSize := findCurrentImageSize(url)
	return strings.Replace(url, string(currentSize), string(size), -1)
}

func Login(clientid string, clientsecret string) *IGDBService.Client {
	client := IGDBService.NewClient(clientid, clientsecret)
	return client
}

func Limit(limit int) string {
	return "limit " + string(rune(limit)) + ";"
}

func Search(name string) string {
	return "search \"" + name + "\";"
}

func Where(conditions string) string {
	return "where " + conditions + ";"
}

func Fields(fields string) string {
	return "fields " + fields + ";"
}

func Sort(fields string, descSort bool) string {
	sort := "sort " + fields
	if descSort {
		sort += " desc"
	}
	return sort + ";"
}

func Page(page int) string {
	return " "
}
