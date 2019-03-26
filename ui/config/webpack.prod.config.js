const merge = require('webpack-merge');
const path = require('path');
const webpack = require('webpack')
const webpackBaseConfig = require('./webpack.base.config');

const webpackProdConfig = merge(webpackBaseConfig, {
    output: {
        publicPath: '/',
        path: path.join(__dirname, '../static/'),
        filename: '[name]-[hash:8].min.js',
        chunkFilename: '[name]-[chunkhash:8].chunk.min.js',
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env.NODE_ENV': '"production"'
        }),
    ]
})
module.exports = webpackProdConfig;