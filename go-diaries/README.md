# go-diaries

## API一覧

| Method | URL                                         | Interface |
| :----- | :------------------------------------------ | :-------- |
| POST   | http://localhost:8080/api/diary             | Store     |
| PUT    | http://localhost:8080/api/diary/{id}        | Update    | 
| DELETE | http://localhost:8080/api/diary/{id}/delete | Delete    | 
| GET    | http://localhost:8080/api/diary/{id}        | FindByID  | 
| GET    | http://localhost:8080/api/diaries           | FindAll   | 

### API呼出の例
{id}にはデータのidを指定する 例えば 1

#### 登録
```
curl -i -X POST -H "Content-Type=application/json;Charset=UTF-8" \
-d '{"title":"英語", "description":"英語30分勉強したよ"}' http://localhost:8080/api/diary
```

#### 更新
```
curl -i -X PUT -H "Content-Type=application/json;Charset=UTF-8" \
-d '{"title":"英語", "description":"英語30分勉強したよ"}' http://localhost:8080/api/diary/{id}
```

#### 削除
```
curl -i -X DELETE -H "Content-Type=application/json;Charset=UTF-8" \
http://localhost:8080/api/diary/{id}/delete
```

#### 取得
```
curl -i -X GET -H "Content-Type=application/json;Charset=UTF-8" \
http://localhost:8080/api/diary/{id}
```

#### 一覧取得
```
curl -i -X GET -H "Content-Type=application/json;Charset=UTF-8" \
http://localhost:8080/api/diaries
```
