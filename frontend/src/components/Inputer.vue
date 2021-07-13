<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <form class="inputer" autocomplete="off" v-on:submit.stop.prevent="onSubmit">
    <label for="title">{{$t("title")}}</label>
    <textarea id="title" name="title" :placeholder="$t('title_placeholder')" cols="2" required v-model="input.title" maxlength="43"></textarea>

    <label for="guide_text">{{$t("guide_text")}}</label>
    <input id="guide_text" name="guide_text" :placeholder="$t('guide_text_placeholder')" v-model.trim="input.guideText"
           maxlength="40">

    <label for="guide_text_placement">{{$t("guide_text_placement")}}</label>
    <select id="guide_text_placement" name="guide_text_placement" v-model="input.guideTextPlacement">
      <option selected value="BR">{{$t("bottom_right")}}</option>
      <option value="BL">{{$t("bottom_left")}}</option>
      <option value="TR">{{$t("top_right")}}</option>
      <option value="TL">{{$t("top_left")}}</option>
    </select>

    <label for="author">{{$t("author")}}</label>
    <input id="author" name="author" :placeholder="$t('required')" required v-model.trim="input.author" maxlength="36">

    <label for="top_text">{{$t("top_text")}}</label>
    <input id="top_text" name="top_text" :placeholder="$t('required')" required v-model.trim="input.topText" maxlength="60">

    <label for="animal_code">{{$t("animal_code")}}</label>
    <input id="animal_code" name="animal_code" type="number" :placeholder="$t('animal_code_placeholder')"
           v-model.number="input.animalCode" min="0" max="41">

    <label for="color_code">{{$t("color_code")}}</label>
    <input id="color_code" name="color_code" type="number" :placeholder="$t('color_code_placeholder')"
           v-model.number="input.colorCode" min="0" max="16">

    <button type="submit" v-bind:disabled="isSubmitDisabled">{{$t('submit_word')}}</button>
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

        submitWord: "Loading...",
        isSubmitDisabled: true,
      }
    },
    mounted: function() {
      this.isSubmitDisabled = false
      this.submitWord = this.$t("submit_word")
    },
    methods: {
      onSubmit: function () {
        let self = this
        this.disableSubmit()
        setTimeout(self.enableSubmit, 2000)

        this.$emit("input-submit", this.input)
      },
      disableSubmit: function () {
        this.isSubmitDisabled = true
        this.submitWord = this.$t("submitting_word")
      },
      enableSubmit: function () {
        this.isSubmitDisabled = false
        this.submitWord = this.$t("submit_word")
      }
    },
  }
</script>

<style scoped>
  form {
    width: 500px;
  }

  @media screen and (max-width: 500px) {
    form {
      max-width: 96%;
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
