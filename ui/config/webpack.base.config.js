const HtmlWebPackPlugin = require("html-webpack-plugin");
const fs = require('fs');
const path = require('path');
const pkgPath = path.join(__dirname, '../package.json');
const pkg = fs.existsSync(pkgPath) ? require(pkgPath) : {};
let theme = {};
if (pkg.theme && typeof (pkg.theme) === 'string') {
    let cfgPath = pkg.theme;
    // relative path
    if (cfgPath.charAt(0) === '.') {
        cfgPath = path.resolve(__dirname, cfgPath);
    }
    const getThemeConfig = require(cfgPath);
    theme = getThemeConfig();
} else if (pkg.theme && typeof (pkg.theme) === 'object') {
    theme = pkg.theme;
}
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
                test: /\.less|css$/,
                exclude:/node_modules/,
                use: [{
                    loader: "style-loader" // creates style nodes from JS strings
                }, {
                    loader: "css-loader",
                    options: {
                        modules: true,
                        camelCase: true,
                        localIdentName: '[name]_[local]__[hash:base64:5]',
                        importLoaders: 2,
                        sourceMap: false,
                    }// translates CSS into CommonJS
                }, {
                    loader: "less-loader",
                    options: {
                        javascriptEnabled: true,
                        sourceMap: true,
                        modifyVars: theme
                    } // compiles Less to CSS
                }]
            },
            {
                test: /\.less|css$/,
                include:/node_modules/,
                use: [{
                    loader: "style-loader" // creates style nodes from JS strings
                }, {
                    loader: "css-loader",
                }, {
                    loader: "less-loader",
                    options: {
                        javascriptEnabled: true,
                        sourceMap: true,
                        modifyVars: theme
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