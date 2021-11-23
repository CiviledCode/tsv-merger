# tsv-merger
A quick and simple Text Seperated Value (TSV) file merger and parser written in Go. 

## Disclaimer
**This code is experimental and prone to bugs/performance issues in it's current state. I am NOT responsible for any possible issues or damages that come from
using this software**

## What it includes
In it's current state it's fairly limited, but more features are planned for the future.

- [x] Parsing TSV files.
- [x] Merging TSV files with addition.
- [ ] Merging TSV files with appending.
- [ ] Merging multiple columns from configuration.
- [ ] Individual merging operations for columns (Addition, Appending, Subtraction, Multiplication, Division, Floating point numbers, Constants). These should be represented as individual characters.
- [ ] Formula interpretation for merging. Some sort of basic arithmetic parsing library.

## How to use it
1. Download the latest config.json and prebuilt binary from the releases tab.
2. Configure the seperator and files within the config file. Make sure the config is within the same directory as said files.
3. Feed the config into the executable with the following command `tsv-merger config.json output.tsv` where config.json is your config file and output.tsv is the name of the file you want to output to.
4. If the console prints OK, then your files have been merged.
