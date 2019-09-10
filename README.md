# Draft: Upload/FFmpeg/Amazon S3  
Backend, system for uploading, trimming video, creating video previews with scaling, creating photo previews, uploading to Amazon s3 using the generated presigned URL.   

# Install FFmpeg
1. apt install ffmpeg  

# Install GO

1. apt install snapd  
  
Restart your  session  
  
2. snap install go --classic  
3. go version  

# Deploy bd_video

1. git clone ...   
2. echo "GOPATH=$HOME/.example_video" >> ~/.bashrc  
3. cd ~/example_video  
4. go get -d ./  

# Build & run

1. change IP address & port  
2. go build  
3. ./example_video &  

# Check port listening

netstat -ntlp | grep LISTEN  

# Stop GO script by his name

pkill main  

