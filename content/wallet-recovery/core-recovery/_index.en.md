---
title: "Wallet recovery using Bitcoin Core"
date: 2022-02-22T16:03:11.470Z
lastmod: 2022-02-22T16:03:11.470Z
weight: 3
draft: false
# search related keywords
keywords: ["recovery, personal"]
---

### Table of contents
* [Overview](#overview)
* [Step 1: Download the Nunchuk app and log in using Guest mode](#step1)
* [Step 2: Make sure that you are on mainnet](#step2)
* [Step 3: Recover your key(s)](#step3)
* [Step 4: Recover your wallet](#step4)
* [Step 5: Spend from the wallet](#step5)
* [Step 6: Clean up](#step6)

### Overview <a name="overview"></a>

A personal wallet is one where you possess all keys inside the wallet.

To recover a personal wallet, you will need:
* The **seed phrase(s)** for your key(s)
* The **wallet configuration file**

{{< notice info >}}
  Seed phrases are defined in [BIP-0039 specification](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki).
{{< /notice >}}

{{< notice info >}}
  The wallet configuration file has a .bsms extension and is defined in [BIP-0129 specification](https://github.com/bitcoin/bips/blob/master/bip-0129.mediawiki).
{{< /notice >}}

{{< notice note >}}
  For simplicity, this is a "clean slate" recovery. It assumes that the user has no prior data except for the seed phrases and the wallet configuration file. It uses Guest mode, which does not require an account.
{{< /notice >}}

{{< notice note >}}
  This recovery guide uses the Nunchuk desktop app. The same process can be done using the Nunchuk mobile apps.
{{< /notice >}}

### Step 1: Download the Nunchuk app and log in using Guest mode <a name="step1"></a>
Download the app from our website: https://nunchuk.io/#Download.

Log in as Guest.

