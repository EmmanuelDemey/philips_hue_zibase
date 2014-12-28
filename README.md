philips_hue_zibase
==================

Bridge between the Zibase box and Philips Hue lights

This script is the first implementation of a bridge between the Zibase Home automation box, and Philips Hue lights

As it is not possible to send PUT request from the Zibase box, only GET requests are possible, we need a bridge that will convert a request to the 
corresponding Philips Hue API

How To
==================
```
go run hue_zibase.go
```

REST API
==================
* GET ip:3000/light?id=<id of the light>&switch=on : Will switch on the corresponding light
* GET ip:3000/light?id=<id of the light>&switch=off : Will switch off the corresponding light
* GET ip:3000/light?id=<id of the light : According to the current state of the light, will switch on or off. 
