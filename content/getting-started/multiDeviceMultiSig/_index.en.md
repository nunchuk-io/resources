---
title: "Creating a self-managed, multi-device multisig wallet"
date: 2022-02-22T16:03:11.470Z
lastmod: 2022-02-22T16:03:11.470Z
weight: 6
draft: false
# search related keywords
keywords: ["mac, macos, windows, linux, ios, android"]
---

{{< notice note >}}
This tutorial is for someone that would like their transactions to be signed from multiple devices. In this workflow, we will be using a desktop, laptop, and a mobile phone to create a 2-of-3 multisig wallet. You can also use this workflow as a template to easily create another version of a multisig wallet.
{{< /notice >}}

{{< notice info >}}
You must be signed in to your account in order to access the messaging system. <br/>
You will need to import or create a key for each of the devices you wish to use in the wallet. The keys cannot be the same. If you need help creating keys click <u>[here](/getting-started/createsoftwarekey/)</u>
{{</ notice >}}

{{< tabs >}}

  {{< tab "Desktop" >}}
  
1. Open the chat feature located on the left side bar. <br/><br/>
    
<img src = desktopChatButton.png
    alt = "Chat Button (Screenshot)"/>

2. Click the "Start new message" button and type in your user email. You should see your name and email pop up below, click it!<br/><br/>

<img src = desktopStartNewMessage.png
    alt = "start new message button (Screenshot)"/>
<img src = desktopNameSelected.png
    alt = "Name selected (Screenshot)"/>

3. Send a friendly message to yourself to initiate the chat room. After sending the message, you should see a new chat room with yourself!<br/><br/>

<img src = desktopJoinedChatRoom.png
    alt = "Wallet + button (Screenshot)"/>

4. Click the "Create a collaborative wallet" button to initiate the multisig configuration setup.<br/>

<img src = desktopCreateCollabWalletBttn.png
    alt = "collaborative wallet question (Screenshot)"/>

5. Name your wallet and click "Continue".<br/><br/>
    
<img src = desktopWalletName.png
    alt = "Create collaborative wallet view (Screenshot)"/>

6. Add the number of keys you will be using in your wallet. Then, add the required keys to sign a transaction.

{{< notice tip >}}
NOTE: The total number of keys refers to how many keys are valid for this wallet. The required keys refers to how many keys are needed to sign and send a transaction.
{{< /notice >}}

<img src = desktopWalletConfig.png
    alt = "Configure collaborative wallet (Screenshot)">


7. Be sure to check your configuration before clicking "Continue".<br/><br/>

<img src = desktopCurrentConfig.png
    alt = "your current configuration (Screenshot)">

8. You can now review other parameters of your wallet configuration, then continue.<br/><br/>

<img src = desktopReviewWallet.png
    alt = "Review collaborative wallet (Screenshot)">


9. Ensure your local key is selected (for me, that is my Desktop Key) and click "Continue".

{{< notice tip >}}
The key in green (Software) is an eligible key to sign with for your current device. Foreign keys are not eligible for signing locally and will need to be signed on the device they were imported or created on.
{{< /notice >}}

<img src = desktopAssignKeys.png
    alt = "Assign key (Screenshot)">

<img src = desktopSelectKey.png
    alt = "Assign Key Desktop (Screenshot)">

10. You should now see the original chat screen you created in step 2. This is where you can find the pending wallet configuration when you begin adding your keys on the next device.<br/><br/>

<img src = desktopConfigDone.png
    alt = "Chat screen pending keys (Screenshot)">

{{< notice note >}}
You do not need to continue to the other tabs if you were able to add all the keys to your wallet (number 9). If you want to add a key created from another device, continue to the next tab.
{{< /notice >}}

{{< /tab >}}

{{< tab "Laptop" >}}

{{< notice note >}}
If you did not have all your keys available for the initial setup, you can continue from another device to continue the setup process. It is assumed that you have started to create a wallet. If you still have a pending configuration notification in your chat screen you can continue on your laptop here. If you do not, you should refer to the desktop tab to start a configuration.
{{< /notice >}}

1. Sign into your laptop and open Nunchuk. If you need help installing, refer to the install documentation found <u> [here](/getting-started/createsoftwarekey/) </u>
<br/><br/>
2. Navigate to the chat screen. There you should see the pending wallet notification. Click the "View config" button.<br/><br/>

<img src = laptopConfigNote.png
    alt = "pending keys config note (Screenshot)">

3. Here you see the number of current keys assigned among other information about the wallet. Click the "Continue to assign keys" button.<br/><br/>

<img src = laptopWalletConfig.png
    alt = "wallet configuration (Screenshot)">

4. You can now assign the key stored locally on your laptop. Notice that the other two keys are considered foreign and your local key is recognized as a "SOFTWARE" key.<br/><br/>

<img src = laptopAssignKey.png
    alt = "assign key configuration (Screenshot)">

5. After assigning your key, click the "Continue" button. You should now be returned to the chat screen and see that your config notification is now only pending 1 key.<br/><br/>

<img src = laptopPendingKeys.png
    alt = "pending keys (Screenshot)">

6. Continue to your next device!

{{< /tab >}}

{{< tab "Mobile" >}}

{{< notice note >}}
If you did not have all your keys available for the initial setup, you can continue from another device to continue the setup process. It is assumed that you have started to create a wallet. If you still have a pending configuration notification in your chat screen you can continue on your laptop here. If you do not, you should refer to the desktop tab to start a configuration.
{{< /notice >}}

1. Open your Nunchuk mobile app and navigate to the chat screen. You should see your wallet config notification here waiting for the last key!<br/><br/>

<img src = mobilePendingKey.png
    alt = "mobile app chat screen (Screenshot)">

2. Click on the "View config" button.<br/><br/>

3. Click the "Continue to assign keys" button at the bottom of the screen.<br/><br/>

<img src = mobileContinueAssignKeys.png
    alt = "Wallet configuration (Screenshot)">

4. You can now assign the last key (in my case, it is the Mobile Key).<br/><br/>

5. Click the "Assign keys and send invite" to finalize your wallet with the last key.<br/><br/>

<img src = mobileAssignKeysSendInvite.png
    alt = "assign keys and send invite (Screenshot)">

6. You should be returned to the chat screen. You should now see the wallet config notification says "Pending finalization". On any of your devices you can open the config again by clicking the "View Config" and then the "Finalize wallet" button!<br/><br/>

<img src = mobileAllKeysAssigned.png
    alt = "mobile phone chat screen (Screenshot)">

7. You will now be able to see the finalize button. When you click that button you will see all of the keys associated with the wallet and other important information. Click the "Finalize wallet" button once you're sure everything is correct.<br/><br/>

<img src = desktopFinalizeWalletConfig.png
    alt = "desktop wallet config (Screenshot)">

8. You're all done! You should see your wallet on the main screen.<br/><br/>

{{< notice tip >}}
It is highly recommended to back up your wallet (noted in red, above). There are 2 options for recovery. Either (a) the BSMS file AND enough keys to satisfy the number of signatures requirement or (b) all keys and brute-forcing to "guess" what the original BSMS file was (this can fail with highly-customized and more advanced wallet configurations).If you do not have a backup or a working wallet, you risk losing access to your funds forever.
{{< /notice>}}


{{< /tab >}}
{{</ tabs >}}

Typically for collaborative wallets, the rule is that the creator of the wallet has to be the one to finalize it. However, for self-managed multisig, the wallet can be finalized from any device (since it's the same user). Once you have added all the required keys to your wallet, you will see a "Finalize wallet" button in the chat screen. You can click that on any device and finalize the wallet! Be sure to look for the backup option in the chat screen and <u>make a backup!</u>

<img src = desktopCompletedMultiSig.png
    alt = "desktop chat screen (Screenshot)">