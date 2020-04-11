/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

const merge = require('deepmerge')

module.exports = {
  productionSourceMap: false,

  devServer: {
    proxy: {
      '/generate': {
        target: 'http://localhost:3000/',
        changeOrigin: true
      }
    }
  },

  chainWebpack: config => {
    config.module
    .rule('vue')
    .use('vue-loader')
    .tap(options =>
    merge(options, {
      loaders: {
        i18n: '@kazupon/vue-i18n-loader'
      }
    })
    )
  },

  pluginOptions: {
    i18n: {
      locale: 'en',
      fallbackLocale: 'en',
      localeDir: 'locales',
      enableInSFC: true
    }
  },

  pwa: {
    workboxOptions: {
      skipWaiting: true
    },
  },
}

