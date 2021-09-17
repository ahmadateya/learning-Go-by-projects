# Exercise: URL Shortener

The goal of this exercise is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.


#### this project is part of exercising and exploring the GoLang features, I did this project practicing after URL Shortener exercise in [calhoun.io](https://www.calhoun.io/).

#### Packages I Used:
- net/http
- gopkg.in/yaml.v2