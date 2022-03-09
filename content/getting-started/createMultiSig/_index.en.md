---
title: "Creating a multi-signature (multisig) wallet"
date: 2022-02-22T16:03:11.470Z
lastmod: 2022-02-22T16:03:11.470Z
weight: 4
draft: false
# search related keywords
keywords: ["mac, macos, windows, linux, ios, android"]
---

{{< notice note >}}
 You should have at least *two* keys in your account. If you need help creating keys click <u>[here](/getting-started/createsoftwarekey/).</u>
{{< /notice >}}

1. Click the + symbol next to "Wallets" to add a key.

<img src = addNewWallet.png
      alt = "Wallet + button (Screenshot)"/>

2. Name your wallet, give it a description (optional), and then proceed to the next screen.

<img src = multisigWalletConfig.png
    alt = "Add wallet name and description (Screenshot)"/>

3. Click the checkbox of the keys you want to associate with your wallet.

<img src = chooseKeys.png
    alt = "Choose the keys to associate with wallet (Screenshot)"/>

4. Assign the keys you selected to the wallet by clicking the "Assign To Wallet" button. You should see the keys listed on the right side now under "Key Configuration."

<img src = assignToWallet.png
    alt = "Assign to Wallet button (Screenshot)"
    style = "float: left"/>

<img src = assignedKeys.png 
    alt = "Assigned keys (Screenshot)"
    style = "float: center"/>


5. Add the number of signatures (keys) required to sign/send a transaction. You can only require as many keys that you have associated with the wallet. Each key required is a signer for the transaction.

<img src = requiredSignatures.png
    alt = "Required signature (Screenshot)"/>

6. Proceed to the next screen by clicking "Next: Review Wallet" and confirm the details of the wallet. Finalize your wallet by clicking the "Create Wallet" button.

<img src = confirmWalletDetails.png 
    alt = "Confirm wallet details (Screenshot)"/>

{{< notice warning >}}
  You will be prompted to back up your wallet by creating a BSMS (Bitcoin Secure Multisig Setup) file. This is highly recommended for a multi-signature wallet. There are 2 options for recovery. You will need either (a) the BSMS file and enough keys to satisfy the number of signatures required or (b) all keys and "guess" what the original BSMS file was. The latter might require specialized tools and can fail for highly-customized wallets. If you do not have a backup or a working wallet, you risk losing access to your funds forever.
{{< /notice >}}

<img src = downloadBSMSFile.png 
    alt = "Download wallet backup (Screenshot)"/>

{{< notice note >}}
This backup file strictly contains information about the wallet. This means access to this file could be used to <u>track</u> the wallet's funds but not to <u>spend</u> from the wallet.
{{< /notice >}}

