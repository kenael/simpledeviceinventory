INSERT INTO system (machineid, attrs) VALUES ('46a8f9ab071f421a965ab3caeee6c917', '{
  "sysinfo": {
    "version": "0.9.3",
    "timestamp": "2021-06-01T09:51:21.251358237+02:00"
  },
  "node": {
    "hostname": "ets-20160689.eos.lcl",
    "machineid": "46a8f9ab071f421a965ab3caeee6c917",
    "timezone": "Europe/Berlin"
  },
  "os": {
    "name": "Ubuntu 18.04.5 LTS",
    "vendor": "ubuntu",
    "version": "18.04",
    "release": "18.04.5",
    "architecture": "amd64"
  },
  "kernel": {
    "release": "4.15.0-143-generic",
    "version": "#147-Ubuntu SMP Wed Apr 14 16:10:11 UTC 2021",
    "architecture": "x86_64"
  },
  "product": {
    "name": "20FAS3TN00",
    "vendor": "LENOVO",
    "version": "ThinkPad T460s"
  },
  "board": {
    "name": "20FAS3TN00",
    "vendor": "LENOVO",
    "version": "SDK0J40697 WIN",
    "assettag": "Not Available"
  },
  "chassis": {
    "type": 10,
    "vendor": "LENOVO",
    "version": "None",
    "assettag": "No Asset Information"
  },
  "bios": {
    "vendor": "LENOVO",
    "version": "N1CET71W (1.39 )",
    "date": "09/04/2018"
  },
  "cpu": {
    "vendor": "GenuineIntel",
    "model": "Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz",
    "cache": 4096,
    "cpus": 1,
    "cores": 2,
    "threads": 4
  },
  "memory": {},
  "storage": [
    {
      "name": "sda",
      "driver": "sd",
      "vendor": "ATA",
      "model": "SAMSUNG MZNLN512",
      "serial": "S2XANX0H537287",
      "size": 512
    }
  ],
  "network": [
    {
      "name": "enp0s31f6",
      "driver": "e1000e",
      "macaddress": "50:7b:9d:f5:8d:fa",
      "port": "tp",
      "speed": 1000
    },
    {
      "name": "wlp4s0",
      "driver": "iwlwifi",
      "macaddress": "44:85:00:b1:ae:7c"
    },
    {
      "name": "wwp0s20f0u2c2",
      "driver": "cdc_ether",
      "macaddress": "02:1e:10:1f:00:00"
    }
  ]
}') 
;

-- packages

INSERT INTO packages (machineid, attrs) 
VALUES ('46a8f9ab071f421a965ab3caeee6c917', '[{"package":"fonts-sil-abyssinica","version":"1.500-1","status":"installed"},{"package":"printer-driver-cups-pdf","version":"3.0.1-5","status":"installed"},{"package":"libatk-adaptor","version":"2.26.2-1","status":"installed"},{"package":"libvorbisfile3","version":"1.3.5-4.2","status":"installed"},{"package":"binutils-x86-64-linux-gnu","version":"2.30-21ubuntu1~18.04.5","status":"installed"},{"package":"libconfuse-common","version":"3.2.1+dfsg-4ubuntu0.1","status":"installed"},{"package":"libjsoncpp1","version":"1.7.4-3","status":"installed"},{"package":"software-properties-common","version":"0.96.24.32.14","status":"installed"},{"package":"python-ldb","version":"2:1.2.3-1ubuntu0.2","status":"installed"},{"package":"libkeybinder-3.0-0","version":"0.3.2-1","status":"installed"},{"package":"libxcb-xkb1","version":"1.13-2~ubuntu18.04","status":"installed"},{"package":"libavahi-core7","version":"0.7-3.1ubuntu1.2","status":"installed"}]'
;

