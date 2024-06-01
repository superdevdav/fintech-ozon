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
   mutation {<br/>
  createPost(input: {<br/>
    title: "Название поста",<br/>
    description: "Описание поста",<br/>
    authorId: "123",<br/>
    url: "https://example.com",<br/>
  	permissionToComment: true<br/>
  }) {<br/>
    id<br/>
    title<br/>
    description<br/>
    author {<br/>
      id<br/>
    }<br/>
    url<br/>
    createdAt<br/>
  }<br/>
}<br/>
2. Создание комментария
   mutation {<br/>
  createComment(input: {<br/>
    description: "Это комментарий к посту",<br/>
    authorId: "124",<br/>
    postId: "1"<br/>
  }, permissionToComment: true) {<br/>
    id<br/>
    description<br/>
    author {<br/>
      id<br/>
      name<br/>
      email<br/>
    }<br/>
    createdAt<br/>
  }<br/>
}<br/>
