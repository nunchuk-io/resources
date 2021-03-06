---
title: "Personal wallet recovery"
date: 2022-02-22T16:03:11.470Z
lastmod: 2022-02-22T16:03:11.470Z
weight: 1
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

**The number of seed phrases you will need is equal to the number of signatures required to spend from the wallet.** For examples: for a singlesig wallet, you will need exactly 1 seed phrase; for a 2-of-3 multisig wallet, you will need 2 seed phrases.

{{< notice info >}}
  Seed phrases are backups for your keys. They are defined in [BIP-0039](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki).
{{< /notice >}}

{{< notice info >}}
  The wallet configuration file contains important information on how your wallet is constructed. It is defined in [BIP-0129](https://github.com/bitcoin/bips/blob/master/bip-0129.mediawiki).
{{< /notice >}}

{{< notice note >}}
  For simplicity, this is a "clean slate" recovery. It assumes that the user has nothing other than the seed phrases and the wallet configuration file. It uses Guest mode, which does not require an account. It also assumes that there are no hardware keys that survive from the original setup. If there are hardware keys, simply pair the hardware keys with Nunchuk in step 3 (instead of recovering).
{{< /notice >}}

{{< notice note >}}
  Any key that has a passphrase must be recovered using both the seed phrase and the passphrase.
{{< /notice >}}

{{< notice note >}}
  This recovery guide uses the Nunchuk desktop app. The same process can be done using the Nunchuk mobile apps.
{{< /notice >}}

### Step 1: Download the Nunchuk app and log in using Guest mode <a name="step1"></a>
Download the app from our website: https://nunchuk.io/#Download.

Log in as Guest.

![Use guest mode](guest_mode.jpg)

### Step 2: Make sure that you are on mainnet <a name="step2"></a>

![Open settings](settings.jpg)
![Make sure your network is mainnet](network_settings.jpg)

### Step 3: Recover your key(s) <a name="step3"></a>

Recover your keys one-by-one by entering the seed phrase for each key. If you have a singlesig wallet, you will only need to recover one key.

![Add a key](add_key.jpg)
![Add a software key](add_software_key.jpg)
![Recover using seed](recover_with_seed.jpg)

### Step 4: Recover your wallet <a name="step4"></a>

Recover your wallet by importing the wallet configuration file (.bsms). Remember to name the wallet before selecting a file to import.

After the import, you should see the wallet show up in the list of wallets. Select the wallet and wait for it to finish syncing.

{{< notice note >}}
  If the wallet's transaction history and balance are still not up-to-date after a while, try restarting the app.
{{< /notice >}}

![Add a wallet](add_wallet.jpg)
![Import BSMS](import_bsms.jpg)

### Step 5: Spend from the wallet <a name="step5"></a>

Hurrah, you have fully recovered the wallet! You can now proceed to withdraw your bitcoin.

Create a Send transaction, sign it with your recovered key(s), and wait for the transaction to confirm on the blockchain.

{{< notice warning >}}
  When signing, please make sure that all transaction data, such as destination address and transaction fee, are accurate. If you are unsure, try sending a small amount first.
{{< /notice >}}

{{< notice note >}}
  To drain the entire wallet, tick the "Send all" box when creating the Send transaction.
{{< /notice >}}

![Sweep wallet](sweep_wallet.jpg)

### Step 6: Clean up <a name="step6"></a>

To wipe all data used during the recovery process, remove the wallet first, then all of the keys.

![Delete wallet](delete_wallet.jpg)
![Select key](select_key.jpg)
![Delete key](delete_key.jpg)
