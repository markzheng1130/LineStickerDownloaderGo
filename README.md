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


- 類型1 (web service)
    - (1) 先做成local command line tool
        - input, hard code url字串
        - output, gzip落地儲存

    - (2) 接口改成API
        - API request要夾帶input url
        - API response要回傳gzip檔
    
    - (3) 再增加web UI
        - 透過web UI，發起API request
        - 透過web UI，展示API response 收回來的貼圖


- 類型2 (streaming service)
    - (1) 先做成local command line tool
        - input, hard code url字串
        - 把貼圖的binary資料跟metadata封裝起來

    - (2-1)塞進mongo db
        - 練習mongo db infra設置 & 資料結構定義 (應用端)

    - (2-2) 塞進kafka 
        - 練習kafka infra設置，以及message system的應用
    
    - (2-3) 塞進rabbit mq
        - 練習message queue的應用
        - rabbit mq 架構較單純，啟動就可以用了，不像mongo或kafka要讀通基礎infra架構

    - (3) 承(2)，設計1個handler (consumer)，把貼圖資料讀出來
     - 再看是要落地，或是秀在GUI上
     - 貼圖怎麼使用沒那麼重要，這關賣點在練習event consumer的設計
