# pihole-geoip
My Pi-hole server was receiving too many requests (yes I know, public Pi-hole is not recommended in any way), but as I need a public server to provida fallback to some networks I decided to implement this tool.

<!-- markdownlint-disable MD033 -->
<p align="center">
    <pre>
                     +---------+         +---------+
US:80634  +--------->+         +-------->+         |
                     |         |         |         |
CH:64789  +--------->+  GeoIP  +-->x     | Pi-hole |
                     |         |         |         |
ES:28010  +--------->+         +-------->+         |
                     +---------+         +---------+
    </pre>
</p>
<!-- markdownlint-enable MD033 -->

## One-Step Automated Install

Similar to Pi-hole installer, you can easily install it with this command:

##### `curl -sSL https://raw.githubusercontent.com/efrenbg1/pihole-geoip/master/install.sh | bash`

##
## License
Copyright © 2020 Efrén Boyarizo <efren@boyarizo.es><br>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See http://www.wtfpl.net/ for more details.