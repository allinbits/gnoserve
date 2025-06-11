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

func Register(elementTag string, frame gnomark.Frame) {
	assertAccess()
	register(elementTag, webHost, frame)
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
		if strings.HasPrefix(prefix, reserved+"-") {
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

func register(elementTag string, webHost gnomark.WebHost, frame interface{}) {
	if !IsValidTagName(elementTag) {
		panic("invalid tag name: " + elementTag)
	}

	if !IsFrameType(frame) {
		panic("frame must implement Frame interface")
	}

	registry.Set(elementTag, gnomark.Record{
		TagName:      elementTag,
		WebHost:      webHost, // REVIEW consider validation or white-list by github username
		Frame:        frame,
		RegisteredBy: std.CurrentRealm(),
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	})
}

func RenderFrame(element, path string, props ...interface{}) string {
	el, ok := registry.Get(element)
	if !ok {
		panic("element not registered: " + element)
	}
	record, ok := el.(gnomark.Record)
	if !ok {
		panic("element is not a valid Record type: " + element)
	}
	if frameFunc, ok := record.Frame.(gnomark.Frame); ok {
		return frameFunc(path, props...) // call with props
	}
	panic("frame is not a valid Frame type: " + element)
}

func SvgWrapper(xhtml string) string {
	var sb strings.Builder
	sb.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" >\n")
	sb.WriteString("<rect x=\"0\" y=\"0\" width=\"100%\" height=\"100%\" fill=\"none\" stroke=\"black\" stroke-width=\"1\"/>\n")
	sb.WriteString("<foreignObject width=\"100%\" height=\"100%\">\n")
	sb.WriteString("<div xmlns=\"http://www.w3.org/1999/xhtml\" style=\"font-family: Arial, sans-serif; font-size: 14px;\">\n")
	sb.WriteString(xhtml)
	sb.WriteString("</foreignObject>\n")
	sb.WriteString("</svg>\n")
	return sb.String()
}

func Render(path string) string {
	var sb strings.Builder

	registry.IterateByOffset(0, registry.Size(), func(key string, value interface{}) bool {
		if record, ok := value.(gnomark.Record); ok {
			if key != record.TagName {
				panic("registry key mismatch: " + key + " != " + record.TagName)
			}
			sb.WriteString("### &lt;" + record.TagName + "&gt;\n")
			if frameFunc, ok := record.Frame.(gnomark.Frame); ok {

				sb.WriteString(SvgWrapper(frameFunc(path))) // call without props for testing
			} else {
				panic("frame is not a valid Frame type")
			}
		}
		return false // continue iterating
	})
	return sb.String()
}
