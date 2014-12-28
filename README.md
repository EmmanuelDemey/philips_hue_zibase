philips_hue_zibase
==================

Bridge between the Zibase box and Philips Hue lights

This script is the first implementation of a bridge between the Zibase Home automation box, and Philips Hue lights

As it is not possible to send PUT request from the Zibase box, only GET requests are possible, we need a bridge that will convert a request to the 
corresponding Philips Hue API

How To
==================
In the GO script, you need first to update the baseUrl property with the IP address of your Philips Hue module and with the available Hue user. More information on the Getting started doc : http://www.developers.meethue.com/documentation/getting-started

```
go run hue_zibase.go
```

REST API
==================
* GET ip:3000/light?id=<id of the light>&switch=on : Will switch on the corresponding light
* GET ip:3000/light?id=<id of the light>&switch=off : Will switch off the corresponding light
* GET ip:3000/light?id=<id of the light : According to the current state of the light, will switch on or off. 
