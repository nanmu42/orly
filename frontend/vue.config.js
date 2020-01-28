/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

const merge = require('deepmerge')
const CompressionWebpackPlugin = require('compression-webpack-plugin')
const fileToCompressRegex = /\.(js|css|html|svg)$/

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

  configureWebpack: () => {
    let plugins = []
    if (process.env.NODE_ENV === 'production') {
      plugins = [
        new CompressionWebpackPlugin({
          filename: '[path].br[query]',
          algorithm: 'brotliCompress',
          test: fileToCompressRegex,
          compressionOptions: {level: 11},
          threshold: 512,
          minRatio: 0.9,
          deleteOriginalAssets: false,
        }),
        new CompressionWebpackPlugin({
          filename: '[path].gz[query]',
          algorithm: 'gzip',
          test: fileToCompressRegex,
          compressionOptions: {level: 9},
          threshold: 512,
          minRatio: 0.9,
          deleteOriginalAssets: false,
        }),
      ]
    }
    return {plugins}
  },
}

