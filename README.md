# ds18b20-science-fair

My kid is doing a science fair project and wants to log temperature changes in her experiment.  We got some inexpensive DS18B20 sensors online and hooked them up to a raspberry pi.   Most examples online do some string manipulation to extract the temperature from the sensor.   I wanted to extract data from multiple sensors and export it to CSV or JSON so it could be graphed out later.

This is a very basic executable that gets the data from all attached sensors and appends it to a file.   We call the executable via cron as needed.   Worked great.

We used <https://github.com/yryz/ds18b20> to access the sensors and temperatures.   Thank you for the excellent library!
