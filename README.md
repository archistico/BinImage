![Console](https://github.com/archistico/ShadeOfColor/raw/master/screenshot/console1.png)

# SHADE OF COLOR
Encrypt/decrypt file to image

### UTILIZZO
#### Eseguire da go  
`go run crypt.go NOMEFILE`

#### Sample esportazione

##### Image
Tutta la "Divina commedia" di Dante sta in mezza immagine 640x480  
![Immagine di un export](https://github.com/archistico/ShadeOfColor/raw/master/screenshot/divinacommedia-640x480-000.png)  

##### Yaml
`NomeFile: divinacommedia.txt`  
`DataLength: 576609`  
`NomeImmagine: testi/divinacommedia-640x480`  
`EstensioneImmagine: png`  
`LarghezzaImmagine: 640`  
`AltezzaImmagine: 480`  
`NumeroImmagini: 1`  
`Sha1: 40bbbb5f38a74037dbfa72d6a6e818bc57f537f0`  

#### Compilazione  
`go build -o crypt.exe crypt.go`  
`go build -o decrypt.exe decrypt.go`    
`crypt NOMEFILE`  
`decrypt NOMEFILE.yaml`  

#### Opzioni  
Per scegliere il formato desiderato dell'immagine  
`crypt NOMEFILE -f`

### TODO
 - Encrypt data
 
### LICENSE
GNU GPL 3.0