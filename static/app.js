'use strict';

console.log('hi js');

main();

async function main() {
  let tmp = await getPosts();

  tmp = await addPost();

  tmp = await getPosts();

  tmp = await deletePost();

  tmp = await getPosts();

  tmp = await changePost();

  tmp = await getPosts();
}

async function getPosts() {
  let response = await fetch('http://localhost:8080/api/posts');
  let responseJSON = await response.json();

  if (!responseJSON instanceof Array) {
    console.log(`Expected array, got: ${responseJSON}`);
  }
  console.log(responseJSON);
}

async function addPost() {
  let newPost = {
    Title: "new post",
    Text: "post test text",
  }

  let response = await fetch('http://localhost:8080/api/post', {
    method: 'POST',
    body: JSON.stringify(newPost),
  });

  let result = await response.json();
  console.log(result);
}

async function deletePost() {
  let postID = 2
  let response = await fetch(`http://localhost:8080/api/post/${postID}`, {
    method: 'DELETE',
  });
}

async function changePost() {
  let postID = 1;
  let changedTitle = {
    Title: 'Title_2_changed',
    Text: 'text_2',
    IsDone: true,
  }

  let response = await fetch(`http://localhost:8080/api/post/${postID}`, {
    method: 'POST',
    body: JSON.stringify(changedTitle)
  })


  postID = 0;
  changedTitle = {
    Title: 'Title_1',
    Text: 'text_1_changed',
    IsDone: false,
  }

  response = await fetch(`http://localhost:8080/api/post/${postID}`, {
    method: 'POST',
    body: JSON.stringify(changedTitle)
  })

}