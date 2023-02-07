# Commutee-Server

Golang server for Commuttee Project 

## Api Points 
### "/" 
  provides a `Helloworld` string 
### "/api/recommendation" 
  Provides a 4 random recomended location (Vikalp)
### "/api/matrix" 
  Provides the distance matrix for the data given in 
  start="" - for starting point 
  end="" for ending point 
  mode = "" to define the mode of transport  `transit` mode dont work (not accurate data)

Example :  `http://127.0.0.1:8080/api/matrix?start=Kelambakam%20Tamil%20Nadu&end=Tambaram%20TamilNadu&mode=walk`
