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

## Примеры GraphQL запросов
1. Создание поста
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
2. Создание комментария
P.S. permissionToComment у запроса должно совпадать с permissionToComment у поста
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
Если permissionToComment = true, то ответ от сервера будет иметь вид
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

Иначе при permissionToComment = false
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
