# CSS
CSS provides a library which greatly simplify how we write css styles in a more 
flexible way by using the power of Go templates and css parsers.


## Install

```bash
go get -u github.com/influx6/gu/css
```

## Example

```go
csr := css.New(`

    $:hover {
      color: red;
    }

    $::before {
      content: "bugger";
    }

    $ div a {
      color: black;
      font-family: {{ .Font }}
    }

    @media (max-width: 400px){

      $:hover {
        color: blue;
        font-family: {{ .Font }}
      }

    }
`)

  sheet, err := csr.Stylesheet(struct {
    Font string
  }{Font: "Helvetica"}, "#galatica")

  sheet.String() // => "#galatica:hover {\n  color: red;\n}\n#galatica::before {\n  content: \"bugger\";\n}\n#galatica div a {\n  color: black;\n  font-family: Helvetica;\n}\n@media (max-width: 400px) {\n  #galatica:hover {\n    color: blue;\n    font-family: Helvetica;\n  }\n}"

```

## Gratitude
Thanks to the awesome work of the [CSS tokenizer by the Gorilla team](https://github.com/gorilla/css)  
and [Aymerick's css parser](https://github.com/aymerick/douceur) through all whom by God's grace 
made this library possible.
