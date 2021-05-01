# Vibely
Learn song lyrics quickly so you can vibe

![vibely](https://user-images.githubusercontent.com/7995105/116646135-5511ce80-a945-11eb-9eea-8cc27a296779.png)


## Details
Vibley is written in [Torus](https://github.com/thesephist/torus) on the frontend and Go on the backend, deployed as a systemd file with nginx on my Digital Ocean server.

To run this locally, run `go run cmd/main.go` and navigate to the `127.0.0.1:8996`



## API
Vibely takes advantage of the "public" Genius API to search for songs. It then uses a [Go implementation of beautiful soup](https://github.com/anaskhan96/soup) to get the lyrics of the song and display it on the page
