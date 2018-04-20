package main

import (
	"html/template"
	"net/http"
)

var tmpl = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>INNOQ Blockchain Go</title>
</head>

<style>
    .box {
        margin: 50px 0px 0px 100px
    }
    .mbottom25 {
        margin-bottom: 25px
    }
    .mbottom50 {
        margin-bottom: 50px
    }
    .mtop10 {
        margin-top: 15px
    }
    a {
        text-decoration: none
    }
    input {
        width: 300px
    }
</style>
<body>
    <div class="box">
        <div class="mbottom25">
            <a href="/" target="_blank">
                Get Details</a>
        </div>

        <div class="mbottom25">
            <a href="/blocks" target="_blank">
                Get Blocks</a>
        </div>

        <div class="mbottom25">
            <a href="/transactions" target="_blank">
                Get Transactions</a>
        </div>

        <div class="mbottom50">
            <a href="/mine" target="_blank">
                Mine A Block</a>
        </div>

        <span><b>Create Transaction</b></span><br>
        <form class="mbottom25 mtop10" method="post"
            action="/transcations">

            <span>Payload</span><br>
            <input type="text" name="payload">
            <button type="submit">Submit</button>
        </form>
    </div>
</body>

</html>

`

func GetIndex(w http.ResponseWriter, r *http.Request) {
	t := template.New("main") //name of the template is main
	t, _ = t.Parse(tmpl)      // parsing of template string
	t.Execute(w, "")
}