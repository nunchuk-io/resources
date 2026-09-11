# Connect Nunchuk to RaspiBlitz


1. Follow the guide to install Electrum Rust Server to your RaspiBlitz node: [https://github.com/raspiblitz/raspiblitz#electrum-rust-server](https://github.com/raspiblitz/raspiblitz#electrum-rust-server)
2. Make sure Tor is up and running on your system.
3. Copy the Tor Hidden Service address for electrs.
4. Open Nunchuk and go to Profile -> Settings -> Network Settings
5. Enter the address you just copied, prefixed with `ssl://` and followed by the port `:50002`. For example:
`ssl://y6j7bn4rogpz3pk5jcvwhlsboizofbn7dvk7i6y6wbqjcid.onion:50002`
6. On desktop, enable "Use TOR Proxy" and enter `127.0.0.1` as the proxy address. Use port `9050` if you are running the Tor service, or `9150` if you are using Tor Browser. Keep the Tor service or Tor Browser running while using Nunchuk.
7. Save network settings, then restart Nunchuk.