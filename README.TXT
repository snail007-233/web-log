# intro
record http post content to log file.   
# usage
1.parameters  
-h show help   
-d dir of log file to store,default current dir ./   
-p port to listen , default 8010  
2.access  
http post "content" to url http://IP:port/xxxx  
"content" will be store in {logDir}/xxxx-{day}.log  
3."xxxx" must match [0-9a-zA-Z_-]+ and max length is 30  


