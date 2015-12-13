-- WiFi

function connectToWiFi()
  wifi.setmode(wifi.STATION)

  -- Your WiFi Credentials
  wifi.sta.config("ssid","password")

  wifi.sta.connect()
  tmr.delay(4000000)   -- wait 4,000,000 us = 4 second
  print("WiFi Connection Status")
  print(wifi.sta.status())
  print(wifi.sta.getip())
end

function listWiFiAps()
  function listap(t)
    for k,v in pairs(t) do
      print(k.." : "..v)
    end
  end
  wifi.sta.getap(listap)
end
