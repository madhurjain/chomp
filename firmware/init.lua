-- PetFeeder Proof-Of-Concept
-- Madhur Jain (msj25@njit.edu)
-- 11/24/2015


print("\n")
print("\n Initializing Pet Feeder System")

-- Compile Include files
local includeFiles = {"wifi", "led", "mqtt", "servo", "proximity"}
for i, f in ipairs(includeFiles) do
    if file.open(f..".lua") then
        file.close()
        print("Compile File:"..f)
        node.compile(f..".lua")
        print("Remove File:"..f)
        file.remove(f..".lua")
    end
end

-- Include Compiled files
for i, f in ipairs(includeFiles) do
    if file.open(f..".lc") then
        dofile(f..".lc")
    else
        print(f..".lc not exist")
    end
end
includeFilesile = nil
collectgarbage()

blinkBlueLED()
connectToWiFi()

-- proximity check every 5seconds
sensor = proximity.init()
tmr.alarm(3, 2000, 1, function()
    -- check if someone near feeder
    distance = sensor.measure()    
    if distance > 2 and distance < 80 then
        print("proximity: ")
        print(distance)
        if (wifi.sta.status() == 5) then
            sendLog(distance)
        else
            print("WiFi Not Connected. Distance Not Sent.")
        end
    end
end)

-- Control Command received from MQTT
function controlOp(data)    
    if data == "feed" then
        turnServo()
    end
end
connectMQTT(controlOp)

