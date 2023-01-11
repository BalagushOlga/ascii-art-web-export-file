# ASCII-art-web-export-file project

Written in **GO**,

By **[Olha Balahush](https://github.com/BalagushOlga)**
___
## Description:
___

**ASCII-art-web-export-file** is based on the ([ascii-art-web](https://01.kood.tech/git/Olya/ascii-art-web)) with the user-friendly interface ([ascii-art-web-stylize](https://01.kood.tech/git/Olya/ascii-art-web-stylize)) and with the possibility of *downloading the result to a file* (in *txt* format).

___
Webpage allow to use different banners:

- *shadow*;
- *standard*;
- *thinkertoy*.

The 'main' page have:

- buttons to start use the program;
- the result of program on the page;
- ***button download*** for this task. 

The 'write' page have:

- text input
- buttons to switch between banners
- button, which sends a POST request to '/ascii-art'

This project is handle an input with *numbers*, *letters*, *spaces*, *special characters* and ```\n```.
___
## How to use:
___

To run use this commad:
```
go run .
```
To check the work of app open [localhost:8080](http://localhost:8080/) in a browser.
