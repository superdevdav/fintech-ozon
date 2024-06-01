## Задача
Реализовать систему для добавления и чтения постов и комментариев с использованием GraphQL

## Структура
1. Файл server.go главный
2. В /graph/... реализация graphql запросов
3. В /storage/... находится функционал, связанный с постами, комментариями и таблицами

## /storage/storePosts
В postRepository.go имеется структура PostRepository с методами AddPost, GetAllPosts, getCommentsForPost, getUserByID.
В postRepository_test.go находятся тесты для PostRepository.

## /storage/storeComments
В commentsRepository.go имеется структура CommentsRepository с методом AddComment.
В commentsRepository_test.go находятся тесты для CommentsRepository, такие как TestAddComment, TestAddCommentMore2000Symbols, TestAddCommentEmpty

## Про базу данных
Хранение данных в PostgreSQL. Имеется база данных productdb с тремя таблицами (entities.sql)
```
-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);

-- Таблица постов
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    author_id INTEGER REFERENCES users(id),
    url VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    permission_to_comment BOOLEAN DEFAULT TRUE
);

-- Таблица комментариев
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(id),
    user_id INTEGER REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
Т.к. в ТЗ было прописано только об функционале добавления постов и комментариев, то пользователей я создавал вручную просто в таблице.

## Примеры GraphQL запросов
**1. Создание поста**<br/>
**Запрос:**
```
   mutation
  createPost(input: {
    title: "Название поста",
    description: "Описание поста",
    authorId: "123",
    url: "https://example.com",
  	permissionToComment: true
  }) {
    id
    title
    description
    author {
      id
    }
    url
    createdAt
  }
}
```
**Ответ от сервера:**
```
{
  "data": {
    "createPost": {
      "id": "3",
      "title": "Название поста",
      "description": "Описание поста",
      "author": {
        "id": "123"
      },
      "url": "https://example.com",
      "createdAt": "2024-06-01T18:06:54+03:00"
    }
  }
}
```
**2. Создание комментария**<br/>
P.S. Значение permissionToComment у запроса должно совпадать с permissionToComment у поста
**Запрос:**
```
mutation {
  createComment(input: {
    description: "Это комментарий к посту",
    authorId: "2",
    postId: "2"
  }, permissionToComment: true) {
    id
    description
    author {
      id
      name
      email
    }
    createdAt
  }
}
```
**Ответ от сервера при permissionToComment = true**
```
{
  "data": {
    "createComment": {
      "id": "2",
      "description": "Это комментарий к посту",
      "author": {
        "id": "2",
        "name": "",
        "email": ""
      },
      "createdAt": "2024-06-01T17:48:59+03:00"
    }
  }
}
```
P.S. значения name и email заполнены при выводе постов с комментариями

**Ответ сервера при permissionToComment = false**
```
{
  "errors": [
    {
      "message": "the author has prohibited commenting on this post",
      "path": [
        "createComment"
      ]
    }
  ],
  "data": null
}
```
**3. Получение постов**<br/>
**Запрос:**
```
query {
  posts {
    id
    title
    description
    author {
      id
    }
    url
    comments {
      id
      description
      author {
        id
        name
        email
      }
      createdAt
    }
    createdAt
  }
}
````
**Ответ от сервера:**
```
{
  "data": {
    "posts": [
      {
        "id": "1",
        "title": "Название поста1",
        "description": "Описание поста1",
        "author": {
          "id": "1"
        },
        "url": "https://example.com1",
        "comments": [
          {
            "id": "",
            "description": "Это комментарий к посту",
            "author": {
              "id": "2",
              "name": "user2",
              "email": "email2"
            },
            "createdAt": "2024-05-30T15:41:40Z"
          }
        ],
        "createdAt": "2024-05-30T15:32:59Z"
      },
      {
        "id": "2",
        "title": "Название поста2",
        "description": "Описание поста2",
        "author": {
          "id": "1"
        },
        "url": "https://example.com2",
        "comments": [
          {
            "id": "",
            "description": "Это комментарий к посту какой то",
            "author": {
              "id": "2",
              "name": "user2",
              "email": "email2"
            },
            "createdAt": "2024-06-01T17:48:59Z"
          }
        ],
        "createdAt": "2024-05-30T15:33:17Z"
      }
    ]
  }
}
```
