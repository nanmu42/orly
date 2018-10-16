<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div class="result">
    <img src="../assets/example.gif" alt="O'RLY? Book Cover Example">
    <button>Download</button>
  </div>
</template>

<script>
  export default {
    name: "Result"
  }
</script>

<style scoped>
  .result {
    max-width: 400px;
  }
  @media screen and (max-width: 400px) {
    .result {
      max-width: 96%;
    }
  }
  @media screen and (min-width: 1000px) {
    .result {
      margin: 0
    }
  }
  @media screen and (min-width: 900px) {
    .result {
      margin: 0 0 0 50px;
    }
  }
  @media screen and (max-width: 835px) {
    .result {
      margin: 30px 0 0 0;
    }
  }
  img, button {
    -webkit-appearance: none;
    outline: none;
    display: block;
    width: 100%;
  }
  button {
    white-space: nowrap;
    margin: 20px 0 0 0;
    font-size: 1.2em;
    padding: 12px;
    background: black;
    color: white;
    border: none;
    font-weight: bold;
  }
  button:active {
    background: white;
    color: black;
    box-shadow: 0 0 0 2px black inset;
  }
  img {
    box-shadow: 8px 8px 30px black;
  }

</style>