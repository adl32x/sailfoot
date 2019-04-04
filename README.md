[![Go Report Card](https://goreportcard.com/badge/github.com/adl32x/sailfoot)](https://goreportcard.com/report/github.com/adl32x/sailfoot) [![CircleCI](https://circleci.com/gh/adl32x/sailfoot.svg?style=svg)](https://circleci.com/gh/adl32x/sailfoot)

![Alt logo](./logo.svg)

# Sailfoot

A software for controlling browsers for the purpose of end-to-end testing, scraping and other automation.

Design philosophies:
* Easy to use (straightforward scripting, small learning curve)
* Easy to install, just a single binary

## Etymology

Pleopod:
* a forked swimming limb of a crustacean, five pairs of which are typically attached to the abdomen.
* from Greek plein ‘swim, **sail**’ + pous, pod- ‘**foot**’.

## Documentation

Status: Still a very early work in progress!

### Forewords

Too many enterprise apps have no e2e tests at all. One can't blame the users. Many testing frameworks are cumbersome to install and use, and are just 1-to-1 mappings to the WebDriver API, providing nothing more than pretty printing.

In the sea of untested enterprise web apps, one is allowed to do it quick and dirty. Hence sailfoot, a highly opinionated testing application/framework that focuses on making testing as simple as possible, in one portable binary.

### Installation

Grab the binary from the [release page](https://github.com/adl32x/sailfoot/releases). You will also need either a selenium server, or a webdriver like chromedriver.

### Running

Run the binary. `sf`

### Hello world
When you run `sf` it expects a start.txt in you current directory. You can use the `-file` flag to change the start script.

start.txt:
```
navigate https://url_to_hello_world

has_text 'Hello world!'
```

Run `sf`.

### Built-in keywords

Here's a shortened list of the built-in keywords. Full list of keywords explained [in wiki](https://github.com/adl32x/sailfoot/wiki/Keywords).

|Keyword|Arguments|Example|
|---|---|---|
| has_text | text | has_text 'Hello world!' |
| has_text | selector, text | has_text '.class' 'Hello world!' |
| click | selector | click '.button' |
| click_on_text | text | click 'Log in' |
| click_closest_to | selector, selector | click_closest_to '.username' '.button' |
| input | text | input 'andy' |
| input | selector, text | input '.username' 'andy' |
| read | selector, variable | read '.username' 'USERNAME' |
| execute | text | execute 'ls -l' |

### Writing your own keywords

Making your keyword is as simple as putting it under a directory named `keywords/` and saving it into a .txt file. The filename will be the name of the keyword.

```
├── keywords
│   ├── keyword1.txt
│   ├── keyword2.txt
│   └── ... .txt
└── start.txt
```

#### Arguments

Arguments end up in variables 0..9:
```
log 'Sample keyword with arguments'
input '.long-boring-selector > div > ... $$0$$' $$1$$
```

If you save this as `keyword.txt` you can then use it like this:

`keyword '.username' 'admin'`

### Need to know more?

[Full documentation in Wiki.](https://github.com/adl32x/sailfoot/wiki/Home)

