function turnServo()
    pwm.setup(4, 50, 20)
    print("start servo")
    pwm.start(4)
    tmr.alarm(2, 2000, 0, function()
        print("stop servo")
        pwm.stop(4)
    end)
end
