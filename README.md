# DynDNS Agent

A simple agent to update DynDNS records on the Lyratris DNS platform.

## Usage

### Debian

We provide prebuilt packages through our repository for Debian 12 and Ubuntu 22.04 & 23.10.

**Install our PGP Key**
```bash
curl https://forge.lyratris.com/api/packages/public/debian/repository.key -o /etc/apt/trusted.gpg.d/lyratris.asc
```

**Add our repository**
```bash
echo "deb https://forge.lyratris.com/api/packages/public/debian bookworm main" | sudo tee -a /etc/apt/sources.list.d/lyratris.list
```

If you are using Ubuntu, please make sure to replace the distro and version based on your OS.
\
\
**Install the agent**
```bash
apt update && apt install dyndns-agent
```

You can find all required configuration files at `/etc/dyndns-agent/`.


### Docker Container

Using docker compose `docker-compose-yaml`:

``` yaml
version: "3"
services:
    ironforge:
        image: forge.lyratris.com/public/dyndns-agent:dev
        environment:
            - LTRS_DDNS_API_ID=00000000-0000-0000-0000-000000000000
            - LTRS_DDNS_API_KEY=XXXXXXXXXXXXXX
            - LTRS_DDNS_ENDPOINT=00000000-0000-0000-0000-000000000000
            - LTRS_DDNS_INTERVAL=600
            - LTRS_DDNS_PROTOCOL=any
```


## Configuration

The agent is configured either through environment variables or through a config.ini file. The following configuration options are available:


| Environment Variable | config.ini | Description |
|-----------------------|------------|-------------|
| LTRS_DDNS_API_ID      | apiID      | Your API ID. Needs to have at least the DNS Write access level. |
| LTRS_DDNS_API_KEY     | apiKey     | Corresponding key to your API ID. |
| LTRS_DDNS_ENDPOINT    | endpoint   | The ID of your DynDNS endpoint. |
| LTRS_DDNS_INTERVAL    | interval   | The interval, how often the agent should check if your IP address has changed. Needs to be a value between 120 and 86400 seconds. |
| LTRS_DDNS_PROTOCOL    | protocol   | Defines which IP protocol should be used. Accepts `any` for IPv4 + IPv6, `ipv4` for IPv4 only, and `ipv6` for IPv6 only. |