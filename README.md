![Console](https://github.com/archistico/BinImage/raw/master/screenshot/console1.png)

# SHADE OF COLOR
Software che converte un file in un'immagine

### UTILIZZO
#### Eseguire da go  
`go run main.go NOMEFILE`

#### Sample esportazione
![Immagine di un export](https://github.com/archistico/BinImage/raw/master/testi/divinacommedia-640x480-000.png)
Tutta la "Divina commedia" di Dante sta in mezza immagine 640x480 

#### Compilazione  
`go build -o shadeofcolor.exe main.go`  
`shadeofcolor NOMEFILE`

#### Opzioni  
Per scegliere il formato desiderato dell'immagine  
`shadeofcolor NOMEFILE -f`

### TODO
 - MD5/SHA1 hash file originale
 - Encrypt data
 - Poter gestire un solo canale (R, G, B, A)
 
### LICENSE
GNU GPL 3.0