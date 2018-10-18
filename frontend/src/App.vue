<template>
  <div id="app">
    <TitleBar class="w100" title-msg="O'RLY Cover Generator" :init-lang="this.$i18n.locale" :lang-set="langSet"/>
    <Inputer class="w50" v-on:input-submit="handleSubmit"/>
    <Result class="w50 gap" v-bind:input-src="imgSrc"/>
    <Thumbnails class="w100"/>
    <Colors class="w100" v-bind:colors="colors"/>
    <paperwork class="w100"/>
    <Footer class="w100"/>
  </div>
</template>

<script>
  import TitleBar from "./components/TitleBar"
  import Inputer from "./components/Inputer"
  import Result from "./components/Result"
  import Colors from "./components/Colors"
  import Thumbnails from "./components/thumbnails"
  import Footer from "./components/Footer"
  import Paperwork from "./components/Paperwork"

  export default {
    name: 'app',
    components: {
      Paperwork,
      Footer,
      Thumbnails,
      Colors,
      Result,
      Inputer,
      TitleBar,
    },
    beforeMount: function() {
      let nativeLang = navigator.language ? navigator.language.substring(0,2) : "en"
      let langSet = {
        en: "English",
        zh: "中文",
      }
      this.langSet = langSet
      if (nativeLang in langSet) {
        this.$i18n.locale = nativeLang
      } else {
        this.$i18n.locale = "en"
      }
    },
    data: function () {
      return {
        colors: [
          "#61005e",
          "#70706d",
          "#890029",
          "#c4000e",
          "#6d001d",
          "#6a00bd",
          "#f10000",
          "#0071b1",
          "#f9bc00",
          "#2c0077",
          "#ba009a",
          "#009047",
          "#009d9e",
          "#222e85",
          "#bd002e",
          "#009d1a",
          "#75a500",
        ],
        imgSrc: process.env.BASE_URL + "example.gif",
        langSet : {},
      }
    },
    methods: {
      handleSubmit: function (input) {
        let color, coverID
        if (input.colorCode !== "" && input.colorCode >= 0 && input.colorCode < this.colors.length) {
          color = this.colors[input.colorCode].substring(1)
        } else {
          color = this.colors[Math.floor(Math.random() * (this.colors.length - 1))].substring(1)
        }
        if (input.animalCode !== "" && input.animalCode >= 0 && input.animalCode <= 40) {
          coverID = input.animalCode
        } else {
          coverID = Math.floor(Math.random() * 40)
        }

        let rawRequest = "/generate?g_loc=" + input.guideTextPlacement +
          "&g_text=" + input.guideText +
          "&color=" + color +
          "&img_id=" + coverID +
          "&author=" + input.author +
          "&top_text=" + input.topText +
          "&title=" + input.title
        this.imgSrc = encodeURI(rawRequest)
      }
    }
  }
</script>

<style>
  body {
    margin: 0;
    background-color: #fefefe;
  }

  h1 {
    display: block;
    font-size: 2.2em;
    margin: 2em 0 1em 0;
    font-weight: bold;
  }

  h2 {
    display: block;
    font-size: 1.6em;
    margin: 1.6em 0 1em 0;
    font-weight: bold;
  }

  #app {
    box-sizing: border-box;
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    text-align: center;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }

  #app * {
    box-sizing: inherit;
    font-family: inherit;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  .w100 {
    flex: 0 0 100%;
  }

  .w50 {
    flex: 0 0 auto;
  }
</style>
