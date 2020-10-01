systemctl stop pihole-geoip
rm -f -r /etc/pihole/geoip
mkdir -p /etc/pihole/geoip
cd /etc/pihole/geoip
wget https://github.com/efrenbg1/pihole-geoip/releases/download/0.1/pihole-geoip.tar.gz
tar -xzvf pihole-geoip.tar.gz
rm pihole-geoip.tar.gz
nano config.json
cp pihole-geoip.service /lib/systemd/system/pihole-geoip.service
systemctl daemon-reload
systemctl start pihole-geoip