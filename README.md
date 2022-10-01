# SING ME A MUSIC
Find a music buy a term.


# REQUIRED/OPTION
- SHAZAM_RAPIDAPI_KEY=xxx (required) 
   - https://rapidapi.com/apidojo/api/shazam
- TERM=DOG (optional)

# BUILD
```bash
$ docker build -t diegofonseca/sing-me-a-music .
```

# RUN
```bash
$ docker run --rm -e SHAZAM_RAPIDAPI_KEY=xxx -e TERM=cat -it diegofonseca/sing-me-a-music
```