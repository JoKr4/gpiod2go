# gpiod2go

Golang bindings for libgpiod v2.1.x  
Link-time dependency is used

Incomplete surely, as I'm only interested in the very simple usecase to switch relais with a RaspberryPi  

## Building
Assuming you have Linux+go+gcc+pkgconf+libgpiod just do
```
cd cmd/example
CC=gcc go build
./example
```
Note that you need according rights to use /dev/gpiochip*

## More Notes

Here are some of my notes I did while coding:


```gpiod_line_request_set_value``` takes a 'request' and 'offset' and 'value', but 'request' has 'line_config' which has 'line_setting' which has 'offset' and 'value' again  
--> why is this repeated?!

TODO consider using the plural version ```gpiod_line_request_set_values``` everytime, even for single lines

There is a mixture of 'free' and 'release' suffixes or missing '_new' prefix, e.g.  
```gpiod_line_settings_new``` -> ```gpiod_line_settings_free```  
```gpiod_line_config_new``` -> ```gpiod_line_config_free```  
but  
```gpiod_chip_request_lines``` --> ```gpiod_line_request_release```

---

references
- https://github.com/brgl/libgpiod/blob/v2.1.x/examples/toggle_line_value.c
- https://github.com/brgl/libgpiod/blob/v2.1.x/examples/toggle_multiple_line_values.c