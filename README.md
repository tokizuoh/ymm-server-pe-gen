# ymm-server-pe-gen
  
[【新卒】サーバーサイドエンジニア応募者向けの模試 | ゆめみ](https://www.yumemi.co.jp/serverside_recruit) > コーディング試験の例  
にて使用する、CSV形式のファイルを生成するスクリプト。  
  
```bash
> head -6 game_score_log.csv
player_id,score
11,9100
8,8100
14,4700
11,7600
9,1600
```
  
## Docker
  
### Version

```bash
> docker --version
Docker version 19.03.12, build 48a66213fe

> docker-compose --version
docker-compose version 1.27.2, build 18f557f9
```
  
### Build
  
```bash
> docker-compose up --build -d
```
  
### Run
  
```bash
> docker-compose exec app go run main.go
```