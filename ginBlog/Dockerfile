FROM golang:latest

WORKDIR /opt/ginBlog
COPY ginBlog /opt/ginBlog/
COPY conf /opt/ginBlog/conf/

EXPOSE 8000
CMD ["./ginBlog"]