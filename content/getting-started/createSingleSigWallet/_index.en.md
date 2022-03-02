---
title: "Creating a SingleSig Wallet"
date: 2022-02-22T16:03:11.470Z
lastmod: 2022-02-22T16:03:11.470Z
weight: 3
draft: false
# search related keywords
keywords: ["mac, macos, windows, linux, ios, android"]
---

{{< notice note >}}
 You should have at least *one* key in your account.
{{< /notice >}}

1. Click the + symbol next to "Wallets" to add a key.

<img src = addNewWallet.png
      alt = "Wallet + button (Screenshot)"/>

2. Name your wallet, give it a description (optional), and then proceed to the next screen.

<img src = walletConfig.png
    alt = "Add wallet name and description (Screenshot)"/>

3. Click the boxes of the keys you want to associate with your wallet. 

<img src = chooseKeys.png
    alt = "Choose the keys to associate with wallet (Screenshot)"/>

4. Assign the key(s) you selected to the wallet by clicking the "Assign To Wallet" button. You should see the key(s) listed on the right side now under "Key Configuration"

<img src = assignToWallet.png
    alt = "Assign to Wallet button (Screenshot)"
    style = "float: left"/>

<img src = assignedKeys.png 
    alt = "Assigned keys (Screenshot)"
    style = "float: center"/>


5. Add the number of signatures (keys) required to sign/send a transaction. You can only require as many keys that you have associated with the wallet.

<img src = requiredSignatures.png
    alt = "Required signature (Screenshot)"/>

6. Proceed to the next screen by clicking "Next: Review Wallet" and confirm the details of the wallet. Finalize your wallet by clicking the "Create Wallet" button.

<img src = confirmWalletDetails.png 
    alt = "Confirm wallet details (Screenshot)"/>

{{< notice warning >}}
  You will be prompted to backup your wallet by creating a BSMS (Bitcoin Secure Multisig Setup) File. If you created a 1 of 1 wallet this is not necessary to do. If you created a multisignature wallet this is highly recommended. In order to recreate a multisig wallet you are required to have either the BSMS file or all the keys that are associated with it. Put simply, if you create a 2 of 3 and lose one of the keys, you would need all 3 keys to recreate the wallet and therefore risk losing access to your funds forever.
{{< /notice >}}

8. Proceed past the backup screen. You should now see your new wallet in the "Wallets" section of the app.

<img src = firstWalletAdded.png
    alt = "My first wallet in Nunchuk (Screenshot)"/>