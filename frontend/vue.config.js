/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

module.exports = {
    devServer: {
        proxy: {
            '/generate': {
                target: 'http://localhost:3000/',
                changeOrigin: true
            }
        }
    }
}

