# Object-Storage-Golang

Object-Storage golang is build for easy file serving and object storage file upload (Min.IO)

## Installation

Git Clone This Project

```bash
docker build -t golang-minio-file-server -f Dockerfile .
docker-compose up -d
```

## Usage

```golang
http://localhost:9000 for accessing the object storage
http://localhost:8081/upload to upload file
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
