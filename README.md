# xk6-dhcp

A k6 extension DHCP

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  xk6 build master \
    --with github.com/mgazza/xk6-dhcp
  ```

## Example
## TODO
```javascript

import dhcp from 'k6/x/dhcp';

const client = dhcp.new('eth0')

export default function () {
    dhcp.discoverOffer(client,)
}
```