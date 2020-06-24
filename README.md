# A Vault Controller for the Pimoroni Blinkt! #

A simple way to physically/visually display the seal status of a [HashiCorp](https://www.hashicorp.com) [Vault](https://github.com/hashicorp/vault) server by using a [Pimoroni Blinkt!](https://shop.pimoroni.com/products/blinkt).

The Blinkt is a low-profile strip of eight super-bright, color LED indicators that plugs directly onto the Raspberry Pi's GPIO header. Several available software libraries make it easy to control the color and brightness of each LED independently.

## How It Works ##

This little tool will fetch the health and seal status from the Vault server and will display the status using the Blinkt! LEDs.
When sealed the LEDs are turned red, when unsealed they will show a green color.
During the unsealing process, a portion of the LEDs will turn green for each entered key, until the required treshold is reached and the Vault is unsealed.

## Acknowledgements ##

Based on the [blinkt-k8s-controller](https://github.com/apprenda/blinkt-k8s-controller) of @apprenda, this project draws inspiration and borrows heavily from the work done by @alexellis on [Docker on Raspberry Pis](http://blog.alexellis.io/visiting-pimoroni/) and his [Blinkt Go libraries](https://github.com/alexellis/blinkt_go), themselves based on work by @gamaral for using the `/sys/` fs interface [instead of special libraries or elevated privileges](https://guillermoamaral.com/read/rpi-gpio-c-sysfs/) to `/dev/mem` on the Raspberry Pi.

## Requirements ##

Physically install a [Pimoroni Blinkt](https://shop.pimoroni.com/products/blinkt) on a Raspberry Pi you want to use for display. **No additional sofware or setup is required for the Blinkt**.

