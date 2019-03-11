const HtmlWebPackPlugin = require("html-webpack-plugin");

module.exports = {
    entry: './src/App.jsx',
    resolve: {
        extensions: ['.js', '.jsx'],
    },
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        babelrc: false,
                        presets: [
                            require('babel-preset-react'),
                            [require('babel-preset-env'), { modules: false }],
                        ],
                        cacheDirectory: true,
                        plugins: [
                            ["import", { "libraryName": "antd", "style": true }],
                            // [
                            //     "@babel/plugin-transform-runtime",
                            //     {
                            //       "corejs": false,
                            //       "helpers": true,
                            //       "regenerator": true,
                            //       "useESModules": false
                            //     }
                            //   ]
                        ]
                    },
                },
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader']
            },
            {
                test: /\.less$/,
                use: [{
                    loader: "style-loader" // creates style nodes from JS strings
                }, {
                    loader: "css-loader" // translates CSS into CommonJS
                }, {
                    loader: "less-loader",
                    options: {
                        javascriptEnabled: true
                    } // compiles Less to CSS
                }]
            }
        ]
    },
    plugins: [
        new HtmlWebPackPlugin({
            template: "./src/index.html",
            filename: "index.html"
        })
    ]
};