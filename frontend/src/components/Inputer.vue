<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <form class="inputer" autocomplete="off">
    <label for="title">Title</label>
    <input id="title" name="title" placeholder="required">

    <label for="guide_text">Guide Text</label>
    <input id="guide_text" name="guide_text" placeholder="such as 'The Definitive Guide'">

    <label for="guide_text_placement">Guide Text Placement</label>
    <select id="guide_text_placement" name="guide_text_placement">
      <option selected value="BR">Bottom Right</option>
      <option value="BL">Bottom Left</option>
      <option value="TR">Top Right</option>
      <option value="TL">Top Left</option>
    </select>

    <label for="author">Author</label>
    <input id="author" name="author" placeholder="required">

    <label for="top_text">Top Text</label>
    <input id="top_text" name="top_text" placeholder="required">

    <label for="animal_code">Animal Code</label>
    <input id="animal_code" name="animal_code" placeholder="0-40 (listed below, defaults to random)">

    <label for="color_code">Color Code</label>
    <input id="color_code" name="color_code" placeholder="0-16 (listed below, defaults to random)">

    <button>Generate O'RLY</button>
  </form>
</template>

<script>
  export default {
    name: "Inputer"
  }
</script>

<style scoped>
  form {
    max-width: 500px;
    min-width: 400px;
  }
  @media screen and (max-width: 500px) {
    form {
      max-width: 96%;
      min-width: auto;
    }
  }
  label, input, select, button {
    -webkit-appearance: none;
    outline: none;
    display: block;
    width: 100%;
  }
  label {
    text-align: left;
    font-size: 0.9em;
    color: black;
  }
  label:not(:first-child) {
    margin: 18px 0 0 0;
  }
  input, select {
    font-size: 1em;
    margin: 0;
    padding: 8px 10px;
    border: solid 2px lightgray;
  }
  select {
    background: url('../assets/options.png') no-repeat right white;
  }
  button {
    white-space: nowrap;
    margin: 16px 0 0 0;
    font-size: 1.6em;
    padding: 12px;
    background-color: lightseagreen;
    color: white;
    border: none;
    font-weight: bold;
  }

  button:active {
    background-color: white;
    box-shadow: 0 0 0 2px lightseagreen inset;
    color: lightseagreen;
  }
</style>