---
title: "Creating a self-managed, multi-device multisig wallet"
date: 2022-02-22T16:03:11.470Z
lastmod: 2022-02-22T16:03:11.470Z
weight: 5
draft: false
# search related keywords
keywords: ["mac, macos, windows, linux, ios, android"]
---

{{< notice note >}}
This tutorial is for someone that would like to require the signing of transactions from multiple devices. In this workflow, we will be using a laptop, desktop, and mobile phone to create a 3 of 3 wallet. You can also use this workflow as a template to easily create a 2 of 2 wallet where you're required to sign a transaction from your desktop and your phone.
{{< /notice >}}

{{< notice info >}}
You must be signed in to your account in order to access the messaging system. <br/>
You will need a key on each of the devices you wish to use in this wallet. The keys cannot be the same. If you need help creating keys click <u>[here](/getting-started/createsoftwarekey/)</u>
{{</ notice >}}

<!-- {{< tabs >}}

  {{< tab "Desktop" >}} -->
  
1. Open the chat feature located on the left side bar.
    
<img src = chatButton.png
    alt = "Wallet + button (Screenshot)"/>

2. Click the "Start new message" button and type in your user email. You should see your name and email pop up below, click it!

<img src = startNewMessage.png
    alt = "Wallet + button (Screenshot)"/>
<img src = nameSelected.png
    alt = "Wallet + button (Screenshot)"/>

3. Send a friendly message to yourself to initiate the chat room. After sending the message, you should see a new chat room with yourself!

<img src = joinedChatRoom.png
    alt = "Wallet + button (Screenshot)"/>

4. Click the "Create a collaborative wallet" button to initiate the multisig configuration setup.

<img src = createCollabWalletBttn.png
    alt = "Wallet + button (Screenshot)"/>

5. Name your wallet and click "Continue".
    
<img src = walletName.png
    alt = "Wallet + button (Screenshot)"/>

6. Add the number of keys you will be using in your wallet. Then, add the required keys to sign a transaction.

<img src = walletConfig.png
    alt = "">

7. Be sure to check your configuration before clicking "Continue".

<img src = currentConfig.png
    alt = "">

8. You can now review other properties of your wallet configuration, then continue.

<img src = reviewWallet.png
    alt = "">

9. Assign the keys that are required for signing wallet transactions, and then continue.

<img src = assignKeys.png
    alt = "">

{{< notice tip >}}
The key in green (Software) is an eligible key to sign with for your current device. Foreign keys are not eligible for signing locally and will need to be signed on the device they were imported or created on.
{{< /notice >}}

10. 