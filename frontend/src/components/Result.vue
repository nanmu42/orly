<!--
  - Copyright (c) 2018 LI Zhennan
  -
  - Use of this work is governed by an MIT License.
  - You may find a license copy in project root.
  -
  -->

<template>
  <div class="result">
    <progressive-img
      class="img"
      :blur="3"
      :src="imgSrc"
      :placeholder="baseURL + 'example.gif'"
      :fallback="baseURL + 'error.gif'"
      alt="Your O'RLY Book Cover"
    ></progressive-img>
    <a :href="imgSrc" target="_blank" download>Download</a>
  </div>
</template>

<script>
  export default {
    name: "Result",
    props: {
      inputSrc: String,
    },
    beforeMount: function() {
      this.imgSrc = this.inputSrc
    },
    data: function () {
      return {
        baseURL: process.env.BASE_URL,
        imgSrc: "",
      }
    },
    watch: {
      inputSrc: function (newVal) {
        if (newVal === this.imgSrc) {
          return
        }
        this.imgSrc = newVal
      }
    },
  }
</script>

<style scoped>
  .result {
    width: 400px;
  }
  @media screen and (min-width: 970px) {
    .result {
      margin: 0 0 0 50px;
    }
  }
  @media screen and (max-width: 970px) {
    .result {
      margin: 30px 50px 0 50px;
    }
  }
  @media screen and (max-width: 401px) {
    .result {
      max-width: 96%;
      margin: 30px 0 0 0;
    }
  }
  .img, img, a {
    -webkit-appearance: none;
    outline: none;
    display: block;
    width: 100%;
  }
  .img div {
    width: 100%;
  }
  a {
    white-space: nowrap;
    margin: 20px 0 0 0;
    font-size: 1.2em;
    padding: 12px;
    background: black;
    color: white;
    border: none;
    font-weight: bold;
    text-decoration: none;
  }
  a:active {
    background: white;
    color: black;
    box-shadow: 0 0 0 2px black inset;
  }
  .img {
    box-shadow: 8px 8px 30px black;
  }

</style>