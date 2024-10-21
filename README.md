# LineStickerDownloaderGo

- main
    - new a handler for each API request
        - retrieve the input url (from API body)
        - do the following steps


- a handler
    - download page source (from input url)

    - sticker url parser
        - parse sticker urls

    - sticker downloader
        - download stickers

    - post handling
        - compress stickers to gzip
        - save to disk or API response
