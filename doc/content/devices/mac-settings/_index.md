---
title: "MAC Settings"
description: ""
---

This section provides guidelines for configuring MAC settings for end devices from the CLI.

<!--more-->

{{< cli-only >}}

MAC settings on {{% tts %}} are configurable per end device.

Updates to the `mac-state` have an effect during the current session, and are lost on reset.

Updates to `mac-settings` take effect after device creation or after MAC state reset (e.g. after OTAA join or ABP FCnt reset, or ResetInd MAC command)

If no settings are provided on device creation or unset, defaults are taken from [Network Server Configuration]({{< ref src="/reference/configuration/network-server" >}}).

### Available MAC settings

Run the following command to get a list of all available MAC settings:

```bash
$ ttn-lw-cli end-devices update --help
```

You can also refer to the [End Device API Reference page]({{< ref "reference/api/end_device#message:MACSettings" >}}) for documentation on the available MAC settings and MAC state parameters.

### Setting Duty Cycle

To change the duty cycle, set the `max-duty-cycle` parameter. For example, to set the duty cycle to 0.098%, use the following command:

```bash
$ ttn-lw-cli end-devices update "app-id" "device-id" --mac-settings.max-duty-cycle DUTY_CYCLE_1024
```

>Note: See the [End Device API Reference]({{< ref "reference/api/end_device#message:MACSettings" >}}) for available fields and definitions of constants. DUTY_CYCLE_1024 represents 1/1024 ≈ 0.98%.

### Setting Uplink and Downlink Dwell Time

To enable uplink and downlink dwell time, set the `uplink-dwell-time` or `downlink-dwell-time` parameters of `mac-state`. For example, to enable downlink dwell time in the current session, use the command:

```bash
$ ttn-lw-cli end-devices update "app-id" "device-id" --mac-state.current-parameters.downlink_dwell_time true 
```

### Setting RX1 Delay

The RX1 delay of end devices is set to 1 second by default. For some end devices, this may lead to downlink messages not being scheduled in time. Therefore, it is recommended that the RX1 delay be increased to 5 seconds:

```bash
$ ttn-lw-cli end-devices update "app-id" "device-id" --mac-settings.rx1-delay RX_DELAY_5
```

### Unsetting MAC settings

The CLI can also be used to unset MAC settings (so that the default ones are used):

```bash
$ ttn-lw-cli end-devices update "app-id" "device-id" --unset mac-settings.rx1-delay
```
