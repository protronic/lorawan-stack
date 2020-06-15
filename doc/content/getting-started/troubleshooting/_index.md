---
title: "Troubleshooting"
description: ""
---

This section provides help for common issues and frequently asked questions you may have when getting started with {{% tts %}}. If you encounter a bug, please [file an issue](https://github.com/TheThingsNetwork/lorawan-stack/issues/new/choose). If you have questions not covered here, please ask in [{{% tts %}} Forum](https://www.thethingsnetwork.org/forum/c/network-and-routing/v3).

<!--more-->

## My gateway won't connect. What do I do?

- Double check your gateway settings in {{% tts %}}. Ensure your Gateway EUI is correct
- Double check the configuration settings on your gateway. Is the network address the same address you use to connect to {{% tts %}}? Are your ports correct?
- If using CUPS, double check your API Keys and certificates

## My device won't join. What do I do?

- Double check your Device EUI, Join EUI, LoRaWAN Version, AppKey, and NwkKey
- Check gateway and application events for traffic from your device
- Double check frequency plan settings in your end device and gateways (they must be matching)
- Double check your network connection. If there is a slow connection from the server to the gateway, the join accept message may be sent too late (this can happen when a gateway uses 3G as a backhaul)
- Check for an overused App Nonce

## No downlinks are reaching my device. What do I do?

- Check gateway and application events for traffic from your device
- Check duty cycle restrictions
- Device clock drift often occurs when SF12 is used
- Check your antenna connections

## {{% tts %}} is no longer receiving uplinks from my device. What do I do?

- Check gateway and application events for traffic from your device
- Check for FCnt mismatch on ABP devices

## How do I get gateway logs?

Gateway logs can be found in the console in the gateway's general information page.

To view gateway logs with the CLI, use the following command:

```bash
$ ttn-lw-cli events subscribe --gateway-id <gateway ID>
```

## How do I get device logs?

Device logs can be found in the console in the device's general information page.

To view gateway logs with the CLI, use the following command:

```bash
$ ttn-lw-cli events subscribe --application-id <application ID> --device-id <device ID>
```

## How do I generate and authorize OAuth client keys?

See the [authentication reference]({{< ref src="reference/api/authentication" >}}).

## How do I create and authorize new users?

See [user and organization management]({{< ref src="getting-started/user-management" >}}).

## I'm getting entity not found errors in my gateway traffic

Your gateway receives traffic from all devices in range, and {{% tts %}} drops this traffic if no matching device is found. If you are expecting uplinks from a device but receiving this error, double check your Device EUI and Join EUI.
