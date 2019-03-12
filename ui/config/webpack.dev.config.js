const merge = require('webpack-merge');
const path = require('path');
const webpack = require('webpack')
const webpackBaseConfig = require('./webpack.base.config');

const webpackDevConfig = merge(webpackBaseConfig, {
    output: {
        publicPath: '/',
        path: path.join(__dirname, '../'),
        filename: 'build/[name].js',
        chunkFilename: 'build/chunk-[name].js',
    },
    devtool: 'cheap-module-source-map',
    devServer: {
        headers: {
            "Access-Control-Allow-Origin": "*"
        },
        stats: {
            version: false,
            hash: false,
            maxModules: 0
        },
        compress: false,
        clientLogLevel: "none",
        port: 8082,
        hot: true,
        hotOnly: true,
        host: "0.0.0.0",
        disableHostCheck: true,
        historyApiFallback: true,
        overlay: {
            warnings: true,
            errors: true
        },
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env.NODE_ENV': '"development"'
        }),
    ]
})
module.exports = webpackDevConfig;