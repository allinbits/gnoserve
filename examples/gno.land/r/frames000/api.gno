package frame

import (
	"regexp"
	"std"
	"strings"
	"time"

	"gno.land/p/demo/avl"
	"gno.land/p/gnomark"
)

var (
	registry         = avl.NewTree()
	realmAllowPrefix = []string{}
	reservedPrefixes = []string{}

	webHost = gnomark.WebHost{
		Base: "", // KLUDGE: use empty base for local development
		Tag:  "",
		Path: "/static/",
	}
)

func init() {
	realmAllowPrefix = append(realmAllowPrefix, std.CurrentRealm().PkgPath()+"/") // must be in realm sub-path
	reservedPrefixes = append(reservedPrefixes, "gno", "gnoland", "gnomark")

	register("rss-feed", webHost, "renders an rss widget") // REVIEW: maybe try a more useful example
}

func hasAllowedPrefix() bool {
	prevRealm := std.PreviousRealm()
	for _, callerPath := range realmAllowPrefix {
		if strings.HasPrefix(prevRealm.PkgPath(), callerPath) {
			return true
		}
	}
	return false
}

func assertAccess() {
	if !hasAllowedPrefix() {
		panic("access denied: " + std.PreviousRealm().PkgPath() +
			" realm must match an allowed prefix:[" + strings.Join(realmAllowPrefix, ",") + "]")
	}
}

// REVIEW: let user update or change webhost
func Register(elementTag string, description string) {
	assertAccess()
	register(elementTag, webHost, description)
}

func ReservePrefix(prefix string) {
	assertAccess()
	if prefix == "" || len(prefix) < 3 {
		panic("reserved prefix must be at least 3 characters long")
	}
	if strings.Contains(prefix, "-") {
		panic("reserved prefix cannot contain hyphens")
	}
	reservedPrefixes = append(reservedPrefixes, prefix)
}

func IsReservedPrefix(prefix string) bool {
	for _, reserved := range reservedPrefixes {
		if strings.HasPrefix(prefix, reserved + "-") {
			return true
		}
	}
	return false
}

func IsValidTagName(tagName string) bool {
	if tagName == "" || len(tagName) < 8 || IsReservedPrefix(tagName) {
		return false
	}
	match, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_-]*-[a-zA-Z0-9_-]+$`, tagName)
	return match
}

func IsFrameType(frame interface{}) bool {
	if frame == nil {
		return false
	}
	_, ok := frame.(gnomark.Frame)
	return ok
}

func register(elementTag string, webHost gnomark.WebHost, description string) {
	if !IsValidTagName(elementTag) {
		panic("invalid tag name: " + elementTag)
	}

	registry.Set(elementTag, gnomark.Record{
		TagName:      elementTag,
		WebHost:      webHost, // REVIEW consider validation or white-list by github username
		RegisteredBy: std.CurrentRealm(),
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	})
}

func RenderTag(element, path string, obj interface{}) string {
	el, ok := registry.Get(element)
	if !ok {
		panic("html element not registered: " + element)
	}
	record, ok := el.(gnomark.Record)
	if !ok {
        panic("element is not a valid record: " + element)
    }
    var sb strings.Builder
    sb.WriteString(record.WebHost.Frame(element, obj.(gnomark.Frame).ToGnoMark(path)))
    return sb.String()
}

// TODO: return web hosts and tags as json or md
func Render(path string) string {
	var sb strings.Builder

	registry.IterateByOffset(0, registry.Size(), func(key string, value interface{}) bool {
		if record, ok := value.(gnomark.Record); ok {
			if key != record.TagName {
				panic("registry key mismatch: " + key + " != " + record.TagName)
			}
			sb.WriteString("### &lt;" + record.TagName + "&gt;\n")
		}
		return false // continue iterating
	})
	return sb.String()
}
