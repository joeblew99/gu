package app

import (
	"github.com/influx6/gu/css"
)

// RootCSS defines a css component which defines the page rendering styles.
var RootCSS = css.New(`
  html, body {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
  }

  body {
    max-wdith: 800px;
  }

  svg path, svg rect {
    fill: inherit;
    stroke: inherit;
    stroke-wdith: inherit;
  }

  $, $ *{
    box-sizing: border-box;
    -o-box-sizing: border-box;
    -moz-box-sizing: border-box;
    -webkit-box-sizing: border-box;
  }

  $ {
    text-align: center;
    margin: 100px 0px;
  }

  $ h1 {
    color: rgba(0,0,0,0.7);
  }
`)

// SubscribeCSS defines the css component which defines the rendering for a
// subscriber form.
var SubscribeCSS = css.New(`

  $, $ *{
    box-sizing: border-box;
    -o-box-sizing: border-box;
    -moz-box-sizing: border-box;
    -webkit-box-sizing: border-box;
  }

  $ {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
  }

  $.subscription .form{
    width: 85%;
    height: 70px;
    margin: 0 auto;
    border: 1px solid rgba(0,0,0,0.3);
  }

  $.subscription .form .email, $.subscription .form .buttons {
    vertical-align: top;
    display: block;
  }

  $.subscription .form .email {
    width: 100%;
    height: 100%;
  }

  $.subscription .form .email input{
    width: 100%;
    height: 100%;
    outline: none;
    border: none;
    display: block;
    font-size: 1.3em;
    padding: 5px 10px;
  }

  $.subscription .form .buttons {
    display: block;
    width: 100%;
    height: 100%;
  }

  $.subscription .form .buttons button {
    outline: none;
    border: none;
    display: block;
    cursor: pointer;
    border: 1px solid rgba(26, 143, 187, 1); 
  }

  $.subscription .form .buttons .button.named {
    width: 100%;
    height: 100%;
    font-size: 1.0em;
    font-weight: bold;
    {{ if ne .SubmitBtnColor "" }}
    background: {{ .SubmitBtnColor }};
    {{ else }}
    background: #42b0da;
    {{ end }}
  }

  @media (min-width:640px){

    $.subscription .form .email {
      width: 75%;
      display: inline-block;
    }

    $.subscription .form .buttons {
      width: 25%;
      display: inline-block;
    }

  }
`)
