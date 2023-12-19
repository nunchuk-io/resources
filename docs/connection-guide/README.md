# Nunchuk Connection Guide

- [Start9](start9.md)
- [Umbrel](umbrel.md)
- [MyNode](mynode.md)
- [RaspiBlitz](raspiblitz.md)


<b>Troubleshooting</b>

1. Ensure that you are using the latest Nunchuk app.
2. Enter the correct address:
   - For SSL-enabled nodes (usually port `50002`): `ssl://host:port`
   - For non-SSL-enabled nodes (usually port `50001`): `host:port` or `tcp://host:port`
   - If your address ends with `:t` or `:s`, please remove those characters.
3. If you encounter a `403 Forbidden error` in Orbot, try changing the country in Orbot settings.
4. When using `electrs`, please make sure to use version `0.9.14` or above.
