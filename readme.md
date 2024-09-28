Сервер запускается на порте 8001 по стандарту.
Поменять параметры сервера и базы данных можно в файле по пути config/local.env

Запрос во внешний API:
  ссылка на внешний API находится в файле local.env,  ExtAddress="http://localhost:8002"
  туда отправляются два query параметра (group и song)
  оттуда принимаются три параметра (release_date, link, text)

Сваггер:
  http://localhost:8001/swagger/index.html

Пример Post запроса на роут /song
{
    "group": "aria",
    "song": "angel",
    "release_date":" ",
    "text":" ",
    "link":" "
}

Пример Put запроса на роут /song
id = 1
{
    "group": "newGroup",
    "song": "newSong",
    "text" : [{
        "verse_id" : 1,
        "verse" : "newText for 1 verse"
    }]
}

Пример Get запроса на роут /song
id = 1
limit = 1
verse = 1

Пример Delete запроса на роут /song
id = 1
