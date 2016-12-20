package app

import (
	"github.com/gu-io/gu/css"
)

// IndexCSS defines the css component which defines the rendering for
// a notification page.
var IndexCSS = css.New(`
  *{
    margin: 0;
    padding: 0;
    font-family: "Lato", "Open Sans",sans-serif;
  }

  html, body {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
  }

  body {
    width: 100%;
    margin: 0 auto;
    background: #efefef;
  }

  .roboto {
    font-family: "Roboto", "Lato", "helvetica";
  }
`)

// RootCSS defines a css component which defines the page rendering styles.
var RootCSS = css.New(`
  $, $ *{
    box-sizing: border-box;
    -o-box-sizing: border-box;
    -moz-box-sizing: border-box;
    -webkit-box-sizing: border-box;
  }

  $ {
    text-align: center;
    margin: 150px 0px;
  }

  $ h1 {
    color: rgba(0,0,0,0.7);
    font-size: 3em;
  }
`)

// NotificationCSS defines the css component which defines the rendering for
// a notification page.
var NotificationCSS = css.New(`
  $, $ *{
    box-sizing: border-box;
    -o-box-sizing: border-box;
    -moz-box-sizing: border-box;
    -webkit-box-sizing: border-box;
  }

  $ {
    width: 800px;
    height: 100px;
    margin:0px auto;
  }

  $ .submission.title {
    margin:90px auto;
    text-align: center;
  }

  $ .submission.content .header {
    margin:0px auto;
  }

  $ .submission.content h2{
    font-size: 10em;
    text-align: center;
    color: rgba(0,0,0,0.2);
  }

  $ .submission.content h2.failed{
    color: #ab3809;
  }

  $ .submission.content h2.passed{
    color: #8caf29;
  }

  $ .submission.content .desc{
    margin:50px auto auto auto;
    font-size: 2em;
    text-align: center;
  }

  $ .submission.content .desc .email{
    color: rgba(255,255,255,1);
    background: rgba(67, 164, 189, 0.87);
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
    height: auto;
    margin: 0;
    padding: 0;
  }

  $.subscription {
    height: 70px;
    margin: 60px 0px 0px 0px;
  }

  $.subscription .form{
    width: 85%;
    height: 100%;
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

  $.subscription .form .buttons .button.named {
    font-weight: bold;
    font-size: 1.7em;
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
