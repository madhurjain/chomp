-- Blink OnBoard Blue LED (Active LOW)
function blinkBlueLED()
    pin=0
    gpio.mode(pin,gpio.OUTPUT)
    gpio.write(pin,gpio.LOW)
    lighton=1
    tmr.alarm(0,1000,1,function()
        if lighton==0 then 
            lighton=1
            gpio.write(pin,gpio.LOW)
        else
            lighton=0 
            gpio.write(pin,gpio.HIGH) 
        end 
    end)
end