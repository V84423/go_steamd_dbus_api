# go_steamd_dbus_api

Arguments:
1. reboot = reboots the device, please see systemctl reboot
2. shutdown = power downs the device, please see systemctl poweroff
3. net = shows only "ether" type IP and MAC addresses, please see networkctl list
4. net-renew [if-name] = renews dynamic configurations for [if-name], please see networkctl renew
5. restart [service-name] = restarts service, please see systemctl restart

The project is a showcase how to integrate Golang app with Systemd.
As you may know Systemd exposes some API via filesystem endpoints in /run/systemd, and while we can use those endpoints too, 
they just don't provide all necessary info (for example types of network interfaces). So I see no other options rather than using D-Bus to query Systemd.

I used systemd dbus API for (reboot/shutdown/restart service)
and netlink linux API for (net/renew)

console:
If you're not familiar with go,

to build:

go build and then you can see available commands:

./gosystem
