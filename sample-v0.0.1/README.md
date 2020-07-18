## 状態

- とりあえずginでAPIサーバを立ち上げた状態


## 動かし方

```bash
docker build -t goapp .
docker run -p 28080:8080 -d --name goapp goapp
curl localhost:28080/ping
```