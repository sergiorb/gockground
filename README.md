# Gockground v1.0
A simple golang script to set your linux background from Imgur galleries.

### Usage

Go to [Imgur][app-registry] and register your application for getting your CLIENT_ID code.

Execute:
```
gockground -clientId=CLIENT_ID -galleryId=GALLERYID
```

Example:

```
gockground -clientId=123456789 -galleryId=dbVN1
```

Params:
* **-clientId:** Your own Imgur Client Id.
* **-folder:** Folder to store downloaded images.
* **-galleryId:** Allows you to download and set an image from an Imgur gallery as your background".


[app-registry]: https://api.imgur.com/oauth2/addclient
