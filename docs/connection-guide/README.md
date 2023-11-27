# Nunchuk Connection Guide

- [Start9](start9.md)
- [Umbrel](umbrel.md)
- [MyNode](mynode.md)
- [RaspiBlitz](raspiblitz.md)


<b>Troubleshouting</b>

1. Ensure you're using the latest Nunchuk app.
2. Enter the correct address:
   - For SSL-enabled nodes (usually port `50002`): `ssl://host:port`
   - For non-SSL-enabled nodes (usually port `50001`): `host:port` or `tcp://host:port`
   - If your copied address ends with `:t` or `:s`, please remove that.
3. If you encounter a `403 Forbidden` error with `Orbot`, changing the country in Orbot should help. 
4. When using `electrs`, make sure to use version `0.9.14` or newer.