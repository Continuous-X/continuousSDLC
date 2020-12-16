# IT-System
Für die Betrachtung stelle ich ein IT-System auf einem Raspberry PI bereit, 
welcher in einem eigenen Netzwerk hinter einem Nano Router steht.
## Bereitstellung
Für den Raspberry PI gibt es eine hohe Anzahl von [Betriebssystemen](https://www.elektronikpraxis.vogel.de/45-betriebssysteme-fuer-den-raspberry-pi-a-488934/).
Als Betriebssystem wähle ich ``Raspberry Pi OS``.

1. [Image ZIP](https://www.raspberrypi.org/software/operating-systems/) laden und entpacken
1. Image auf eine SD-Karte mit [Raspberry Pi Imager](https://www.raspberrypi.org/software/) übertragen
1. In der Boot-Partition die leere Datei ``ssh`` im root-Verzeichnis ablegen. 
1. In der Boot-Partition die Datei ``wpa_supplicant.conf`` im root-Verzeichnis ablegen.\
````editorconfig
country=DE 
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
network={
       ssid="WLAN-NAME"
       psk="Passwort"
       key_mgmt=WPA3-PSK
}
````
1. SD-Karte einsetzen und Raspberry Pi starten
1. IP Adresse des Raspberry Pi am Router ausfindig machen
1. Über eine Konsole am Raspberry Pi anmelden: ``ssh pi@192.168.0.xxx`` => Default Passwort `raspberry`
1. ``sudo apt-get update && sudo apt-get upgrade -y``
1. ``sudo raspi-config``
    1. vollen speicherplatz der SD-Karte nutzen
    1. Locale
    1. Zeitzone
    1. Passwort ändern
    1. Hostname ändern, wenn mehrere Raspberry Pi's im Einsatz sind
    

## Metriken
Installation von InfluxDB und Grafana: https://simonhearne.com/2020/pi-influx-grafana/
Installation von Prometheus mit Node Exporter: https://www.mytinydc.com/en/blog/prometheus-grafana-installation/

### Node-Exporter
- aktuelle Version ermitteln: https://github.com/prometheus/node_exporter/releases
- ``wget https://github.com/prometheus/node_exporter/releases/download/v1.0.1/node_exporter-1.0.1.linux-armv7.tar.gz``
- ``tar xfz node_exporter-1.0.1.linux-armv7.tar.gz``
- Ablage unter``/usr/local/bin`` => `sudo cp -R ... /usr/local/bin`
- ``cd /usr/local/bin && sudo ln -s node_exporter-1.0.1.linux-armv7 node_exporter``
- ``cd node_exporter``
- ``nohup ./node_exporter &``
- http://192.168.0.xxx:9100/metrics
