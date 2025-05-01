package www


var body = `
{
"gnoMark": "frame",
"frame": "gno-frame",
"version": "<div> <h1>Stack Dump Example</h1> </div>"
}
`

var htmlBody = `
<div>
    <h1>Stack Dump Example</h1>

    <form action="/stackdump" method="post">
        <label for="input">Enter some text:</label><br>
        <input type="text" id="input" name="input"><br><br>
        <input type="submit" value="Submit">
    </form>
</div>
`

func Render(_ string) string {
    return "<gno-mark>\n\n" + htmlBody + "\n\n</gno-mark>"
}