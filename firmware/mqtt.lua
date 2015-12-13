-- initiate the mqtt client and set keepalive timer to 120sec
m = mqtt.Client("feeder", 120, "sreudmfs", "ivMAiXmR-v-1")

m:on("connect", function(con) print ("mqtt connected") end)
m:on("offline", function(con) print ("mqtt offline") end)

function connectMQTT(cbControl)
    -- on receive message
    m:on("message", function(conn, topic, data)
      print(topic .. ":" )
      if data ~= nil then
        print(data)
        if topic == "/feeder/control" then
            cbControl(data)
        end
      end
    end)
    -- connect to CloudMQTT Service
    m:connect("m11.cloudmqtt.com", 18959, 0, function(conn) 
      print("connected")      
      -- subscribe topic with qos = 0
      m:subscribe("/feeder/#",0)
    end)
end

function sendLog(log)
    m:publish("/feeder/status/log",log,0,0, function(conn) 
        print("sent log status "..log) 
    end)
end
