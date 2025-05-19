package calendar

var eventJson = `{
    "uid": "",
    "dtstamp": "20231001T000000Z",
    "dtstart": "20231001T000000-0400",
    "dtend": "20231001T000000-0400",
    "summary": "Test Event",
    "description": "Test Event",
    "location": "Test Location",
    "status": "CONFIRMED",
    "sequence": 0,
    "valarm": {
        "trigger": "-PT15M",
        "description": "Test Event",
        "action": "DISPLAY"
    }
}`

func Calendar(path string) string {
    return eventJson
}

func Render(path string) string {
    return "```calendar\n\n" + eventJson + "\n```"
}
