user=`id -u`
if [ $user != "0" ]; then
    echo "\e[31mThis script must be run with privileges!\e[0m"
    exit
fi
systemctl stop pihole-geoip
rm -f -r /etc/pihole/geoip 2> /dev/null
mkdir -p /etc/pihole/geoip
cd /etc/pihole/geoip
echo "\e[34mDownloading files...\e[0m"
wget -q -nv https://github.com/efrenbg1/pihole-geoip/releases/download/0.1/pihole-geoip.tar.gz
echo "\e[34mExtracting files...\e[0m"
tar -xzf pihole-geoip.tar.gz
echo "\e[34mRemoving archive...\e[0m"
rm pihole-geoip.tar.gz
echo "\e[34mInstalling systemd service...\e[0m"
cp pihole-geoip.service /lib/systemd/system/pihole-geoip.service
systemctl daemon-reload
systemctl enable pihole-geoip
echo " \e[31m→\e[0m You can edit the settings in '\e[35m/etc/pihole/geoip/config.json\e[0m'"
echo " \e[31m→\e[0m To start the service use: '\e[35msystemctl start pihole-geoip\e[0m'"
echo "\e[92mAll done. Have a great day!\e[0m"
cd /root/go/src/github.com/efrenbg1/pihole-geoip