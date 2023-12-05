# Dokumentasi API Filan Project 2

## Link Github = https://github.com/Bobby-P-dev/FinalProject2_kel7.git

***

Domains = https://finalproject2kel7.adaptable.app

Domains Database = postgres://kesnvrjv:IXjhC8WBdPHCOW_9EKZOEENRvTiDX8iZ@flora.db.elephantsql.com/kesnvrjv


## account admin
    {
    "email":    "admin@gmail.com",
    "password": "admin1"
    }

***
# user
## regist account

Method = POST

Domain = https://finalproject2kel7.adaptable.app/users/register


request body 
```
{
    "username": string,
    "email":    string,
    "password": string,
    "age":  int
}
```

## Login

Method = POST

Domain = https://finalproject2kel7.adaptable.app/users/login

request body

```
{
    "email":    string,
    "password": string
}
```
## put account

Method = PUT

Domain = https://finalproject2kel7.adaptable.app/users/edit/id

request
`
bearer token authorization & param id
`

request body
```
{
    "username": string,
    "email":    string
}
```
## delete account

Method = DELETE

Domain = https://finalproject2kel7.adaptable.app/users/delete

request
`
bearer token authorizaiton
`
***

# Photo
## Post photo

Method = POST

Domain = https://finalproject2kel7.adaptable.app/users/register

request
`
bearer token authorization
`
request body 
```
{
    "title": string,
    "caption": string,
    "photo_url": string
}
```

## Get Photo

Method = GET

Domain = https://finalproject2kel7.adaptable.app/photo/get

request
`
bearer token authorization
`

## Put Photo

Method = PUT

Domain = https://finalproject2kel7.adaptable.app/users/edit/id

request
`
bearer token authorization & param id
`
request body 
```
{
    "title": string,
    "caption": string,
    "photo_url": string
}
```

## Delete Photo

Method = DELETE

Domain = https://finalproject2kel7.adaptable.app/photo/del/id

request
`
bearer token authorization & param id
`

***

# Comment
## Post Comment

Method = POST

Domain = https://finalproject2kel7.adaptable.app/comment/upload

request
`
bearer token authorization
`
request body 
```
{
    "message": string,
    "photo_id": int
}
```

## Get Comment

Method = GET

Domain = https://finalproject2kel7.adaptable.app/comment/get

request
`
bearer token authorization
`

## Put Comment

Method = PUT

Domain = https://finalproject2kel7.adaptable.app/comment/put/id

request
`
bearer token authorization & param id
`
request body 
```
{
     "message":  string
}
```

## Delete Comment

Method = DELETE

Domain = https://finalproject2kel7.adaptable.app/comment/del/id

request
`
bearer token authorization & param id
`

***

# Social Media URL
## Post SocialMediaURL

Method = POST

Domain = https://finalproject2kel7.adaptable.app/socialm/upload

request
`
bearer token authorization
`
request body 
```
{
    "name": string,
    "social_media_url": string
}
```

## Get SocialMediaURL

Method = GET

Domain = https://finalproject2kel7.adaptable.app/socialm/get

request
`
bearer token authorization
`

## Put SocialMediaURL

Method = PUT

Domain = https://finalproject2kel7.adaptable.app/socialm/put/id

request
`
bearer token authorization & param id
`
request body 
```
{
    "name": string,
    "social_media_url": string
}
```

## Delete SocialMediaURL

Method = DELETE

Domain = https://finalproject2kel7.adaptable.app/socialm/del/id

request
`
bearer token authorization & param id