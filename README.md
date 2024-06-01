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
3. Создание комментария
```
   mutation {
  createComment(input: {
    description: "Это комментарий к посту",
    authorId: "124",
    postId: "1"
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
