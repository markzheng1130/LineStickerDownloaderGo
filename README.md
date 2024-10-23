# LineStickerDownloaderGo

- main
    - new a handler for each API request
        - retrieve the input url (from API body)
        - do the following steps


- handler
    - download page source (from input url)

    - sticker url parser
        - parse sticker urls

    - sticker downloader
        - download stickers

    - post handling
        - compress stickers to gzip
        - save to disk or API response


- 預計分成3個階段
    - (1) 先做成local command line tool
        - input, hard code url字串
        - output, gzip落地儲存

    - (2) 接口改成API
        - API request要夾帶input url
        - API response要回傳gzip檔
    
    - (3) 再增加web UI
        - 透過web UI，發起API request
        - 透過web UI，展示API response 收回來的貼圖
