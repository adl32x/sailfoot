![Alt logo](./logo.svg)

# Sailfoot

A software for controlling browsers for the purpose of end-to-end testing, scraping and other automation.

Design philosophies:
* Easy to use (straightforward scripting, small learning curve)
* Easy to install, just a single binary
* Making things easier for testers

## Etymology

Pleopod:
* a forked swimming limb of a crustacean, five pairs of which are typically attached to the abdomen.
* from Greek plein ‘swim, **sail**’ + pous, pod- ‘**foot**’.

## Documentation

### Forewords

Too many enterprise apps have no e2e tests at all. One can't blame the users. Many testing frameworks are cumbersome to install, use and are just 1-to-1 mappings to the WebDriver API, providing nothing more than pretty printing.

Hence sailfoot, a highly opinionated testing application/framework that focuses on making testing as simple as possible, in one portable binary.

### Installation

Grab the binary from the release page. You will also need either a selenium server, or a webdriver like chromedriver.

### Running

Run the binary. `sf`

### Hello world
When you run `sf` it expects a start.txt in you current directory. You can also use the `-file` flag to change the start script.

start.txt:
```
navigate https://url_to_hello_world

has_text 'Hello world!'
```

Run `sf`.

### Built-in keywords

Here's a quick list of the built-in keywords. More detailed explanations here.

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



#### Directory structure

```
├── keywords
│   ├── ... .txt
│   ├── ... .txt
│   └── ... .txt
└── start.txt
```