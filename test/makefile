all:
	xk6 build --with github.com/mgazza/xk6-dhcp="$(shell pwd)/.."
	sudo ./k6 run dhcp.js