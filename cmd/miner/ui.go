package main

import (
	"html/template"
	"net/http"
)

type Ui struct {
	overview *Overview
}

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
			<span>Node ID: {{.NodeId}}</span><br>
			<span>Current block height: {{.CurrentBlockHeight}}</span>
		</div>
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

func NewUi(overview *Overview) *Ui {
	return &Ui{
		overview: overview,
	}
}

func (u *Ui) GetIndex(w http.ResponseWriter, r *http.Request) {
	u.overview.CurrentBlockHeight = u.overview.chain.Height()
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	t.Execute(w, u.overview)
}
