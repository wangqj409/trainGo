## Category
### 1.category list
url: localhost:8080/v1/category

method: GET

return 
```json
{
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2021-06-25T12:47:10.603+08:00",
            "UpdatedAt": "2021-06-25T12:47:10.603+08:00",
            "DeletedAt": null,
            "id": 1,
            "pid": 0,
            "cat_name": "cartoon"
        }
    ]
}
```

### 2.create category 
url: localhost:8080/v1/category

method: POST

|name|type|value|remark|
|---|---|---|---|
|pid|int||default:0|
|cat_name|string|||

### 3.category detail
url: localhost:8080/v1/category/{cat_id}

method: GET

return demo
```json
{
    "data": {
        "ID": 0,
        "CreatedAt": "2021-06-25T12:47:10.603+08:00",
        "UpdatedAt": "2021-06-25T12:47:10.603+08:00",
        "DeletedAt": null,
        "id": 1,
        "pid": 0,
        "cat_name": "cartoon"
    }
}
```

## Article
### 4.list article
url: localhost:8080/v1/article

method: GET

return demo
```json
{
    "data": [
        {
            "id": 1,
            "cat_id": 3,
            "title": "昨天",
            "sub_title": "昨天",
            "content": "昨天是高考第二天了，内心的激动无以言语"
        },
        {
            "id": 2,
            "cat_id": 2,
            "title": "昨天是高考第一天",
            "sub_title": "昨天是?",
            "content": "昨天是高考第二天了，内心的激动无以言语"
        }
    ]
}
```

### 5.create article
url: localhost:8080/v1/article

method: POST

|name|type|value|remark|
|---|---|---|---|
|cat_id|int||default:0|
|title|string||title content|
|content|string|||

return 
```json
{
    "data": {
        "id": 0,
        "cat_id": 2,
        "title": "昨天是高考第一天",
        "content": "昨天是高考第二天了，内心的尴尬无以言语"
    }
}
```