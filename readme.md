# Plays youtube album in the background

Youtube suggests very good music album. It's better than spotify or kkbox interms of suggestions in my opinion. The major inconvenience using youtube app is that the music stops when app
is in the background. You'll have to purchase [youtube premium](https://www.youtube.com/premium) to enable this feature and it costs USD 179.99/ month. If we could retrieve suggested album from youtube and streams the music from our own server via our APP. We could eventually play the music in the background through out APP.


## Search music from youtube API

Youtube provides APIs for searching musics. We could utilize the APIs to retrieve the playlist info.


```go
call := service.Search.List("id,snippet").
  Q("hiphop").
  Type("paylist").
  MaxResults(*maxResults)

res, err := call.Do()
```

We are able to [retrieve list of playlists](https://developers.google.com/youtube/v3/docs/search/list) from the go SDK. We could render the list of playlist. When the user chooses one of the playlists, The server starts to download and streams the music in each music item in the playlist.

## Retrieve video details from the playlist

After user has selected the playlist, we will use the playlist ID to retrieve playlist item from [youtube API](https://developers.google.com/youtube/v3/docs/playlistItems/list?apix_params=%7B%22part%22%3A%22snippet%22%2C%22playlistId%22%3A%22PLAKw47hquUsLe6xcBQjOLBf69HFyPwhCZ%22%7D). 

```go
call := service.
    PlaylistItems("id,snippet").    
    PlaylistId("{{ __PLAYLIST_ID__ }}").
    MaxResults(*maxResults)
    
res, err := call.Do()
```

The playlist item contains video ID that we could retrieve the URL of the storage and download the video from there.

## HTTP server

We will be building server side application based on go native 'http' package. 

