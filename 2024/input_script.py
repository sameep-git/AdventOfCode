import os
import datetime

done = False
while True:
    if done:
        break
    curr_time = datetime.datetime.now()

    target_time = datetime.datetime(2024, 12, 1, 23, 0, 0)

    if curr_time < target_time:
        continue
    else:
        os.system("aocd 2024 2 > day2/input.txt")
        done = True
