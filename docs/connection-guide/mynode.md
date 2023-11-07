# Connect Nunchuk to MyNode


### Connect via Tor

1. Make sure Tor is up and running on your system [https://mynodebtc.github.io/tor/setup.html](https://mynodebtc.github.io/tor/setup.html)
2. Make sure you have Electrum Server running on MyNode [https://mynodebtc.github.io/electrum/electrum.html](https://mynodebtc.github.io/electrum/electrum.html)
3. Copy the Electrum Tor Service hostname
4. Open Nunchuk and go to Profile -> Settings -> Network Settings
5. Enter the IP you just copied and the Electrum Server port (default is 50001), separated by `:` into the Mainnet server. For example: `y6j7bn4rogpz3pk5jcvwhlsboizofbn7dvk7i6y6wbqjcid.onion:50001`
 Note that if you use the secure port `50002`, you need to prefix the address with `ssl://`.  For example:
 `ssl://y6j7bn4rogpz3pk5jcvwhlsboizofbn7dvk7i6y6wbqjcid.onion:50002`
1. Save network settings then restart Nunchuk

##### Mobile
<p float="left">
  <img src="./assets/mobile-1.png" width="408" />
  <img src="./assets/mobile-2-tor.png" width="400" />
</p>

##### Desktop
![Nunchuk desktop](./assets/desktop-tor.png)
