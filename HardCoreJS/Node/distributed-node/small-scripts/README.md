* Testing Compression :- 
    - You can run *server-gzip.js* to test encoding.
    - You can make a curl request like 
    ```shell
    curl -H 'Accept-Encoding: gzip' http://localhost:3002 --output -
    ```
    - This will show you the binary version without decompressed.
    - Run the folloqing if you want to see the binary propoerly 
    ```shell
    curl -H 'Accept-Encoding: gzip' http://localhost:3002 | xxd
    ```
    - Run the following if you want to check the decompressed version:
    ```shell
    curl -H 'Accept-Encoding: gzip' http://localhost:3002 | gunzip
    ```
    - **.pipe() **

