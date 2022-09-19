'use strict';

console.log('hi js');

getPosts();

async function getPosts() {
  let response = await fetch('http://localhost:8080/GetAllPosts');
  let responseJSON = await response.json();

  if (!responseJSON instanceof Array) {
    console.log(`Expected array, got: ${responseJSON}`);
  }

  for (let item of responseJSON) {
    createElement(obj);
  }
}