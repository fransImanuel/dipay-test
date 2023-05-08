
# Project Title

Dipay Test



## Acknowledgements

 - [Export to Thunder Client/Postman/etc...](https://drive.google.com/file/d/1w-enO1s63jnZAUGD73GXqIjN2MbDK6KQ/view?usp=sharing)

## Run Locally

If you want to run the program locally, you need to specify MongoDB URI 
```bash
  docker build --tag=dipay-test:1.0 .
  docker run --name dipay-test -e MONGO_URI="mongodb://localhost:27017/" -p 8080:8080 dipay-test:1.0
```

or if you want to run without docker you need to download this repo to your local and do: 

```bash
  go mod tidy
  go run .\main.go -MONGO_URI="mongodb://localhost:27017/"
```

and dont forget to use exported thunder client api to do the test ( link above )

## Appendix

I think there is a bit ambiguity for the task description, therefore I use my own perception to did the task.

The ambiguity such as
- Address is nullable but there is min/max length to be specify
- on the docs ( no.3 ) there is table spec for object_id but in the postman it was integer data type
-  etc..



## Lessons Learned

Lessons Learned: This project made me feel familiar with Docker again, although it may not be the best as I believe there are still improvements that can be made to the program in the future.
