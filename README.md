# countdown-timer
A small webapp to create a countdown to an event, using URL params and defaults to countdown instead of environment variables.

Forked from https://github.com/shaneutt/countdown-timer but I wanted to use URL params instead of environment variables. I wanted to use URL parameters so I can display multiple countdowns without running multiple containers and didn't mind longer URL paths to support this. 

Default running port 8080 - you can change it in main.go file.

Use `date -d "YYYY-MM-DD HH:MM:SS" +%s` in Linux or (https://www.epochconverter.com/)[https://www.epochconverter.com/] to get the Unix timestamp of the countdown datetime.
Use `perl -MURI::Escape -e 'print uri_escape($ARGV[0]);' "Your string here"` or (https://www.utilities-online.info/urlencode)[https://www.utilities-online.info/urlencode] to convert URL escaped strings.

URL Parameters
- largestunit: largest unit visible in countdown. Either days (D), hours (H), minutes (M), or seconds (S).
- time: unix timestamp for countdown to count to eg. 2524626001.
- title: title displayed for a countdown.
- fontcolor: font color for title and countdown to accomodate for different background images used.
- backgroundimage: URL for background image to use.

# Direct usage

## Prereqs
This program is for personal use so I'm leveraging the existing support for Linux.
- Go, follow steps from `sudo apt install golang-go`
- Make, `sudo apt install make`

## Steps

Create the app with `make`

On Linux, call the countdown binary with `./countdown`

Now view the webpage locally:
```
http://127.0.0.1:8080/?largestunit=D&time=2524626001&title=2050&fontcolor=white&backgroundimage=https://images.unsplash.com/photo-1569470451072-68314f596aec?q=80&w=1931
```


# Docker usage:

## Prereqs
- Docker installation, follow steps from (https://docs.docker.com/engine/install/ubuntu/)[https://docs.docker.com/engine/install/ubuntu/]

## Steps

If you want to edit web layout
```
vi countdown-timer.html
```

Create docker image
```
make container
```

Start image
```
docker run -d -p 8080:8080 --name countdown countdown-timer
docker ps
```

Access the countdown with URL params
```
http://127.0.0.1:8080/?largestunit=D&time=2524626001&title=2050%20docker&fontcolor=white&backgroundimage=https://images.unsplash.com/photo-1569470451072-68314f596aec?q=80&w=1931
```

## Pull the container from Docker hub
```
docker pull saleh99er/countdown-timer:latest
```

## Upload new version of the container 

Remember to revise the image name to your user if you want to upload your own. 
```
docker image tag countdown-timer:latest saleh99er/countdown-timer
docker push saleh99er/countdown-timer
```


