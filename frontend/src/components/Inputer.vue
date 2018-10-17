<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <form class="inputer" autocomplete="off" v-on:submit.stop.prevent="onSubmit">
    <label for="title">Title</label>
    <textarea id="title" name="title" placeholder="required, at most one linebreak" cols="2" required v-model="input.title"
              maxlength="43"></textarea>

    <label for="guide_text">Guide Text</label>
    <input id="guide_text" name="guide_text" placeholder="such as 'The Definitive Guide'" v-model.trim="input.guideText"
           maxlength="40">

    <label for="guide_text_placement">Guide Text Placement</label>
    <select id="guide_text_placement" name="guide_text_placement" v-model="input.guideTextPlacement">
      <option selected value="BR">Bottom Right</option>
      <option value="BL">Bottom Left</option>
      <option value="TR">Top Right</option>
      <option value="TL">Top Left</option>
    </select>

    <label for="author">Author</label>
    <input id="author" name="author" placeholder="required" required v-model.trim="input.author" maxlength="36">

    <label for="top_text">Top Text</label>
    <input id="top_text" name="top_text" placeholder="required" required v-model.trim="input.topText" maxlength="60">

    <label for="animal_code">Animal Code</label>
    <input id="animal_code" name="animal_code" type="number" placeholder="0-40 (listed below, defaults to random)"
           v-model.number="input.animalCode" min="0" max="40">

    <label for="color_code">Color Code</label>
    <input id="color_code" name="color_code" type="number" placeholder="0-16 (listed below, defaults to random)"
           v-model.number="input.colorCode" min="0" max="16">

    <button type="submit" v-bind:disabled="isSubmitDisabled">{{submitWord}}</button>
  </form>
</template>

<script>
  export default {
    name: "Inputer",
    data: function () {
      return {
        input: {
          title: "",
          guideText: "",
          guideTextPlacement: "BR",
          author: "",
          topText: "",
          animalCode: "",
          colorCode: "",
        },

        submitWord: "Generate O'RLY",
        isSubmitDisabled: false,
      }
    },
    methods: {
      onSubmit: function (e) {
        let self = this
        this.disableSubmit()
        setTimeout(self.enableSubmit, 2000)

        this.$emit("input-submit", this.input)
      },
      disableSubmit: function () {
        this.isSubmitDisabled = true
        this.submitWord = "Generating..."
      },
      enableSubmit: function () {
        this.isSubmitDisabled = false
        this.submitWord = "Generate O'RLY"
      }
    },
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

  label, input, select, button, textarea {
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
    margin: 14px 0 0 0;
  }

  input, select, textarea {
    font-size: 1em;
    margin: 0;
    padding: 8px 10px;
    border: solid 2px lightgray;
  }

  input:focus, select:focus, textarea:focus {
    border-color: lightseagreen;
  }

  textarea {
    resize: none;
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
    transition: background-color ease-in-out 500ms;
  }

  button:active {
    background-color: white;
    box-shadow: 0 0 0 2px lightseagreen inset;
    color: lightseagreen;
    transition: none;
  }

  button:focus {
    box-shadow: 0 0 0 2px springgreen inset;
  }

  button:disabled {
    background-color: darkseagreen;
  }
</style>