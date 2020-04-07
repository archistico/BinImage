![Console](https://github.com/archistico/ShadeOfColor/raw/master/screenshot/console1.png)

# SHADE OF COLOR
Encrypt/decrypt file to image

### UTILIZZO
#### Eseguire da go  
`go run main.go NOMEFILE`

#### Sample esportazione
Tutta la "Divina commedia" di Dante sta in mezza immagine 640x480  
![Immagine di un export](https://github.com/archistico/ShadeOfColor/raw/master/testi/divinacommedia-640x480-000.png)  

#### Compilazione  
`go build -o shadeofcolor.exe main.go`  
`shadeofcolor NOMEFILE`

#### Opzioni  
Per scegliere il formato desiderato dell'immagine  
`shadeofcolor NOMEFILE -f`

### TODO
 - Encrypt data
 - Poter gestire un solo canale (R, G, B, A)
 
### LICENSE
GNU GPL 3.0