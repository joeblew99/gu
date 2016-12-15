package css

import "github.com/gu-io/gu/css"

// Index defines the css component which defines the rendering for
// a notification page.
var Index = css.New(`
  *{
    margin: 0;
    padding: 0;
  }

  html, body {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
  }

  body {
    margin: 0 auto;
    padding: 250px;
    background: #efefef;
  }
`)
