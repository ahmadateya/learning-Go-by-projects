# Exercise: HTML Link Parser

the goal of this exercise is to create a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link the package should return a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Links will be nested in different HTML elements, and it is very possible that we will have to deal with HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

In situations like these we want to get output that looks roughly like:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

#### this project is part of exercising and exploring the GoLang features, I did this project practicing after HTML Link Parser exercise in [calhoun.io](https://www.calhoun.io/).

#### Packages I Used:
- io
- golang.org/x/net/html
- 
