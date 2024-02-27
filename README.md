TODO beautify

api findings:

```gpiod_line_request_set_value``` takes a 'request' and 'offset' and 'value', but 'request' has 'line_config' which has 'line_setting' which has 'offset' and 'value' again  
--> why is this repeated?!

TODO consider using the plural version ```gpiod_line_request_set_values``` everytime, even for single offsets

There is a mixture of 'free' and 'release' suffixes or missing '_new' prefix, e.g.  
```gpiod_line_settings_new``` -> ```gpiod_line_settings_free```  
```gpiod_line_config_new``` -> ```gpiod_line_config_free```  
but  
```gpiod_chip_request_lines``` --> ```gpiod_line_request_release```

---

references
- https://github.com/brgl/libgpiod/blob/v2.1.x/examples/toggle_line_value.c
- https://github.com/brgl/libgpiod/blob/v2.1.x/examples/toggle_multiple_line_values.c